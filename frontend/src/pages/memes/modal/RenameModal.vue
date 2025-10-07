<script setup lang="ts">
import { ref, watch, nextTick } from 'vue'
import Modal from '@/components/Modal.vue'
import Button from '@/components/Button.vue'
import { store } from '@/store'
import { RenameFile } from '../../../../wailsjs/go/memeFile/MemeFile'

interface RenameModalProps {
  visible: boolean
  currentFile: string
  parentPath: string
  onSuccess?: (oldName: string, newName: string) => void
}

const props = withDefaults(defineProps<RenameModalProps>(), {
  visible: false,
  currentFile: '',
  parentPath: ''
})

const emit = defineEmits<{
  close: []
  'update:visible': [value: boolean]
}>()

const newFileName = ref('')
const inputRef = ref<HTMLInputElement>()
const isRenaming = ref(false)

watch(() => props.visible, (visible) => {
  if (visible && props.currentFile) {
    newFileName.value = props.currentFile.replace(/\.[^/.]+$/, '')

    nextTick(() => {
      inputRef.value?.focus()
      inputRef.value?.select()
    })
  } else {
    newFileName.value = ''
    isRenaming.value = false
  }
})

const hideRenameModal = () => {
  if (!isRenaming.value) {
    emit('close')
    emit('update:visible', false)
  }
}

const confirmRename = async () => {
  const trimmedName = newFileName.value.trim()

  if (!trimmedName) {
    store.showToast('文件名不能为空', 'error')
    return
  }

  const originalNameWithoutExt = props.currentFile.replace(/\.[^/.]+$/, '')
  if (trimmedName === originalNameWithoutExt) {
    store.showToast('文件名未改变', 'info')
    hideRenameModal()
    return
  }

  if (trimmedName) {
    const fileExt = props.currentFile.substring(props.currentFile.lastIndexOf('.'))
    const newFileNameWithExt = trimmedName + fileExt

    isRenaming.value = true

    try {
      await RenameFile(props.parentPath, props.currentFile, newFileNameWithExt)

      store.showToast('重命名成功！', 'success')

      props.onSuccess?.(props.currentFile, newFileNameWithExt)

      isRenaming.value = false
      hideRenameModal()
    } catch (error) {
      store.showToast(`重命名失败：${error}`, 'error')
      isRenaming.value = false
    }
  }
}

const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Enter') {
    event.preventDefault()
    confirmRename()
  }
}

const isValidFileName = (name: string) => {
  const invalidChars = /[<>:"/\\|?*]/
  return name.trim() !== '' && !invalidChars.test(name)
}
</script>

<template>
  <Modal
    :visible="props.visible"
    title="重命名文件"
    :closable="!isRenaming"
    :mask-closable="!isRenaming"
    @close="hideRenameModal"
  >
    <div class="rename-form">
      <div class="input-group">
        <label class="input-label">新文件名</label>
        <div class="input-wrapper">
          <input
            ref="inputRef"
            v-model="newFileName"
            type="text"
            class="rename-input"
            :class="{ 'error': !isValidFileName(newFileName) }"
            placeholder="请输入新的文件名"
            :disabled="isRenaming"
            @keydown="handleKeydown"
          />
          <div class="file-extension">
            {{ props.currentFile.substring(props.currentFile.lastIndexOf('.')) }}
          </div>
        </div>
        <div v-if="!isValidFileName(newFileName)" class="error-message">
          文件名不能包含以下字符：&lt; &gt; : " / \ | ? *
        </div>
      </div>

      <div class="current-name">
        <span class="label">当前文件名：</span>
        <span class="name">{{ props.currentFile }}</span>
      </div>
    </div>

    <template #footer>
      <div class="modal-actions">
        <Button
          variant="secondary"
          :disabled="isRenaming"
          @click="hideRenameModal"
        >
          取消
        </Button>
        <Button
          variant="primary"
          :loading="isRenaming"
          :disabled="!newFileName.trim() || !isValidFileName(newFileName) || newFileName.trim() === props.currentFile.replace(/\.[^/.]+$/, '')"
          @click="confirmRename"
        >
          {{ isRenaming ? '重命名中...' : '确认' }}
        </Button>
      </div>
    </template>
  </Modal>
</template>

<style lang="less" scoped>
.rename-form {
  padding: 0.5rem 0;
}

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
}

.rename-input {
  padding-right: 4rem;
}

.file-extension {
  position: absolute;
  right: 0.75rem;
  top: 50%;
  transform: translateY(-50%);
  font-size: 0.875rem;
  color: #6b7280;
  background: #f3f4f6;
  padding: 0.25rem 0.5rem;
  border-radius: 0.25rem;
  border: 1px solid #e5e7eb;
  font-weight: 500;
}

.error-message {
  margin-top: 0.5rem;
  font-size: 0.75rem;
  color: #ef4444;
  display: flex;
  align-items: center;
  gap: 0.25rem;
  font-weight: 500;
}

.current-name {
  padding: 1rem;
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  border-radius: 0.5rem;
  border: 1px solid #e2e8f0;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);

  .label {
    font-size: 0.75rem;
    color: #64748b;
    font-weight: 500;
    margin-bottom: 0.25rem;
  }

  .name {
    font-size: 0.875rem;
    color: #334155;
    font-weight: 600;
    word-break: break-all;
    font-family: 'Courier New', monospace;
    background: #ffffff;
    padding: 0.5rem;
    border-radius: 0.375rem;
    border: 1px solid #e5e7eb;
  }
}

.modal-actions {
  display: flex;
  gap: 0.75rem;
  justify-content: flex-end;
}
</style>
