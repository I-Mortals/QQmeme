<script setup lang="ts">
import { ref } from 'vue'
import Modal from '@/components/Modal.vue'
import { RenameFilesInOrder } from '@wailsjs/go/memeFile/MemeFile'
import { memeStore, toastStore } from '@/store'

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
  const currentMeme = memeStore.allMemesPath.find(meme => meme.code === props.tabName)
  if (!currentMeme?.orderChanged) {
    toastStore.showToast('当前表情包顺序未改变，无需保存', 'info')
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

    toastStore.showToast('表情包顺序已保存！')

    if (props.onSuccess) {
      props.onSuccess()
    }

    hideModal()
  } catch (error) {
    toastStore.showToast(`保存失败：${error}`, 'error')
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
    :cancelBtn="{
      text: '取消',
      disabled: isRenaming
    }"
    :confirmBtn="{
      text: '确定保存',
      loading: isRenaming,
      disabled: isRenaming
    }"
    @close="hideModal"
    @cancel="hideModal"
    @confirm="confirmSaveOrder"
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

  </Modal>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

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
  color: @rgb-bc;
  margin: 0 0 0.75rem 0;
}

.detail-text {
  font-size: 0.875rem;
  color: @rgb-bc;
  margin: 0 0 0.75rem 0;
  line-height: 1.5;
  opacity: 0.8;

  strong {
    color: @rgb-p;
    font-weight: 600;
  }
}

.warning-text {
  font-size: 0.875rem;
  color: @rgb-e;
  margin: 0;
  font-weight: 500;
}
</style>
