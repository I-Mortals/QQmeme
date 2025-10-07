<script setup lang="ts">
import { ref } from 'vue'
import Modal from '@/components/Modal.vue'
import Button from '@/components/Button.vue'
import { RenameFilesInOrder } from '../../../../wailsjs/go/memeFile/MemeFile'
import { store } from '@/store'

interface SaveTabMemeOrderModalProps {
  visible: boolean
  tabName: string
  parentPath: string
  memes: string[]
  onSuccess?: () => void
}

const props = withDefaults(defineProps<SaveTabMemeOrderModalProps>(), {
  visible: false,
  tabName: '',
  parentPath: '',
  memes: () => []
})

const emit = defineEmits<{
  close: []
  'update:visible': [value: boolean]
}>()

const isRenaming = ref(false)

const hideModal = () => {
  if (!isRenaming.value) {
    emit('close')
    emit('update:visible', false)
  }
}

const confirmSaveOrder = async () => {
  const currentMeme = store.allMemesPath.find(meme => meme.code === props.tabName)
  if (!currentMeme?.orderChanged) {
    store.showToast('当前表情包顺序未改变，无需保存', 'info')
    hideModal()
    return
  }

  isRenaming.value = true

  try {
    await RenameFilesInOrder(
      props.tabName,
      props.parentPath,
      props.memes
    )

    store.showToast('表情包顺序已保存！')

    if (props.onSuccess) {
      props.onSuccess()
    }

    hideModal()
  } catch (error) {
    store.showToast(`保存失败：${error}`, 'error')
  } finally {
    isRenaming.value = false
  }
}
</script>

<template>
  <Modal
    :visible="props.visible"
    title="保存当前顺序"
    :closable="!isRenaming"
    :mask-closable="!isRenaming"
    @close="hideModal"
  >
    <div class="confirm-content">
      <div class="warning-icon">⚠️</div>
      <div class="confirm-text">
        <p class="main-text">为了保存排序，将会修改文件名字</p>
        <p class="detail-text">
          文件将重命名为：<strong>{{ props.tabName }}_01</strong>、<strong>{{ props.tabName }}_02</strong>、<strong>{{ props.tabName }}_03</strong>... 的格式
        </p>
        <p class="warning-text">此操作不可撤销，确定要继续吗？</p>
      </div>
    </div>

    <template #footer>
      <Button
        variant="secondary"
        @click="hideModal"
        :disabled="isRenaming">
        取消
      </Button>
      <Button
        variant="primary"
        @click="confirmSaveOrder"
        :loading="isRenaming"
        :disabled="isRenaming">
        {{ isRenaming ? '保存中...' : '确定保存' }}
      </Button>
    </template>
  </Modal>
</template>

<style lang="less" scoped>
.confirm-content {
  display: flex;
  align-items: flex-start;
  gap: 1rem;
  padding: 0.5rem 0;
}

.warning-icon {
  font-size: 2rem;
  flex-shrink: 0;
  margin-top: 0.25rem;
}

.confirm-text {
  flex: 1;
}

.main-text {
  font-size: 1.125rem;
  font-weight: 600;
  color: #374151;
  margin: 0 0 0.75rem 0;
}

.detail-text {
  font-size: 0.875rem;
  color: #6b7280;
  margin: 0 0 0.75rem 0;
  line-height: 1.5;

  strong {
    color: var(--theme-primary);
    font-weight: 600;
  }
}

.warning-text {
  font-size: 0.875rem;
  color: #dc2626;
  margin: 0;
  font-weight: 500;
}
</style>
