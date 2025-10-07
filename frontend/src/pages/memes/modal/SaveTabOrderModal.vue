<script setup lang="ts">
import { ref } from 'vue'
import Modal from '@/components/Modal.vue'
import Button from '@/components/Button.vue'
import { RenameFoldersInOrder } from '../../../../wailsjs/go/memeFile/MemeFile'
import { store } from '@/store'

interface SaveTabOrderModalProps {
  visible: boolean
  onSuccess?: () => void
}

const props = withDefaults(defineProps<SaveTabOrderModalProps>(), {
  visible: false
})

const emit = defineEmits<{
  close: []
  'update:visible': [value: boolean]
}>()

const isSaving = ref(false)

const hideModal = () => {
  if (!isSaving.value) {
    emit('close')
    emit('update:visible', false)
  }
}

const confirmSaveTabOrder = async () => {
  if (!store.tabOrderChanged) {
    store.showToast('tab 顺序未改变，无需保存', 'info')
    hideModal()
    return
  }

  isSaving.value = true

  try {
    // 保存重命名前的文件夹名称映射
    const oldToNewFolderMap = new Map<string, string>()
    const folderNames = store.allMemesPath.map((meme, index) => {
      let cleanFolderName = meme.name
      if (meme.name.includes('_')) {
        const parts = meme.name.split('_')
        if (parts.length >= 2 && /^\d+$/.test(parts[0])) {
          cleanFolderName = parts.slice(1).join('_')
        }
      }
      const newFolderName = `${String(index + 1).padStart(2, '0')}_${cleanFolderName}`
      oldToNewFolderMap.set(meme.name, newFolderName)
      return meme.name
    })

    await RenameFoldersInOrder(store.rootPath, folderNames)

    store.starMemes.forEach(starMeme => {
      const newFolderName = oldToNewFolderMap.get(starMeme.fromFolder)
      if (newFolderName) {
        starMeme.fromFolder = newFolderName
      }
    })

    await store.refreshMemes()

    if (store.allMemesPath.length > 0) {
      store.tabCurrent = store.allMemesPath[0].code
    }

    store.showToast('tab 文件夹顺序已保存！', 'success')

    if (props.onSuccess) {
      props.onSuccess()
    }

    hideModal()
  } catch (error) {
    console.error('保存 tab 文件夹顺序失败:', error)
    store.showToast(`保存失败：${error}`, 'error')
  } finally {
    isSaving.value = false
  }
}
</script>

<template>
  <Modal
    :visible="props.visible"
    title="保存 tab 文件夹顺序"
    :closable="!isSaving"
    :mask-closable="!isSaving"
    @close="hideModal"
  >
    <div class="confirm-content">
      <div class="warning-icon">📋</div>
      <div class="confirm-text">
        <p class="main-text">确定要保存当前的 tab 文件夹顺序吗？</p>
        <p class="detail-text">
          文件夹将重命名为：<strong>01_文件夹名</strong>、<strong>02_文件夹名</strong>、<strong>03_文件夹名</strong>... 的格式
        </p>
        <p class="warning-text">⚠️ 此操作不可撤销，确定要继续吗？</p>
      </div>
    </div>

    <template #footer>
      <Button
        variant="secondary"
        @click="hideModal"
        :disabled="isSaving">
        取消
      </Button>
      <Button
        variant="primary"
        @click="confirmSaveTabOrder"
        :loading="isSaving"
        :disabled="isSaving">
        {{ isSaving ? '保存中...' : '确定保存' }}
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
