<script setup lang="ts">
import { ref } from 'vue'
import Modal from '@/components/Modal.vue'
import Input from '@/components/Input.vue'
import { toastStore } from '@/store'
import { RenameFile } from '@wailsjs/go/memeFile/MemeFile'

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

const hideRenameModal = () => {
  if (!isRenaming.value) {
    emit('close')
    emit('update:visible', false)
  }
}

const confirmRename = async () => {
  if (!newFileName.value) {
    toastStore.showToast('文件名不能为空', 'error')
    return
  }

  const originalNameWithoutExt = props.currentFile.replace(/\.[^/.]+$/, '')
  if (newFileName.value === originalNameWithoutExt) {
    toastStore.showToast('文件名未改变', 'error')
    return
  }

  if (newFileName.value) {
    const fileExt = props.currentFile.substring(props.currentFile.lastIndexOf('.'))
    const newFileNameWithExt = newFileName.value + fileExt

    isRenaming.value = true

    try {
      await RenameFile(props.parentPath, props.currentFile, newFileNameWithExt)

      toastStore.showToast('重命名成功！', 'success')

      props.onSuccess?.(props.currentFile, newFileNameWithExt)

      isRenaming.value = false
      hideRenameModal()
    } catch (error) {
      toastStore.showToast(`重命名失败：${error}`, 'error')
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
</script>

<template>
  <Modal
    :visible="props.visible"
    title="重命名文件"
    :closable="!isRenaming"
    :mask-closable="!isRenaming"
    :cancelBtn="{
      text: '取消',
      disabled: isRenaming
    }"
    :confirmBtn="{
      text: '确认',
      loading: isRenaming,
      disabled: !newFileName
    }"
    @close="hideRenameModal"
    @confirm="confirmRename"
  >
    <div class="rename-form">
      <div class="form-field">
        <Input
          ref="inputRef"
          v-model="newFileName"
          type="text"
          prepend="新文件名"
          placeholder="请输入新的文件名"
          :disabled="isRenaming"
          :suffix="props.currentFile.substring(props.currentFile.lastIndexOf('.'))"
          @keydown="handleKeydown"
        />
      </div>

      <div class="current-name">
        <span class="label">当前文件名：</span>
        <span class="name">{{ props.currentFile }}</span>
      </div>
    </div>

  </Modal>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

.rename-form {
  padding: 0.5rem 0;
}

.form-field {
  margin-bottom: 1rem;
}

.current-name {
  padding: 1rem;
  background: rgba(var(--bc), 0.06);
  border-radius: 0.5rem;
  border: 1px solid @rgb-b3;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);

  .label {
    font-size: 0.75rem;
    color: @rgb-bc;
    font-weight: 500;
    margin-bottom: 0.25rem;
    opacity: 0.7;
  }

  .name {
    font-size: 0.875rem;
    color: @rgb-bc;
    font-weight: 600;
    word-break: break-all;
    background: @rgb-b1;
    padding: 0.5rem;
    border-radius: 0.375rem;
    border: 1px solid @rgb-b3;
  }
}

</style>
