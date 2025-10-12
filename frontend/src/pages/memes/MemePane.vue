<script setup lang="ts">
import { memeFile } from '../../../wailsjs/go/models'
import MemeInfo = memeFile.MemeInfo
import { joinShowImgPath, joinPath } from '@/utils/path'
import { store } from '@/store'
import { ref } from 'vue'
import { WriteFileToClipboard, DeleteMemeFile } from '../../../wailsjs/go/memeFile/MemeFile'
import LazyLoadImg from '@/components/LazyLoadImg.vue'
import RenameMemeModal from './modal/RenameMemeModal.vue'
import DeleteMemeModal from './modal/DeleteMemeModal.vue'
import { VueDraggable } from 'vue-draggable-plus'

interface MemePaneProps {
  memeInfo: MemeInfo
}
const { memeInfo } = defineProps<MemePaneProps>()

const selectedImage = ref('')

const isRenameModalVisible = ref(false)
const currentRenameFile = ref('')

const isDeleteModalVisible = ref(false)
const currentDeleteFile = ref('')

const handleClick = async (image: string) => {
  const realPath = joinPath(store.rootPath + '/' + memeInfo.code, image)
  if (realPath) {
    try {
      await WriteFileToClipboard(realPath)
      store.showToast('复制成功！')
    } catch (error) {
      store.showToast(`复制失败：${error}`, 'error')
    }
  }
}

const handleContextMenu = (e: MouseEvent, image: string) => {
  selectedImage.value = image

  const menuItems = [
    {
      icon: '📌',
      label: '置顶',
      action: topImage
    },
    {
      icon: '⭐',
      label: '收藏',
      action: () => {
        if (!selectedImage.value) return

        store.addToStarMemes(image, memeInfo.code)
      },
      separator: true
    },
    {
      icon: '✏️',
      label: '重命名',
      action: () => {
        if (!selectedImage.value) return
        showRenameModal(selectedImage.value)
      }
    },
    {
      icon: '🗑️',
      label: '删除',
      action: () => {
        if (!selectedImage.value) return
        showDeleteModal(selectedImage.value)
      }
    }
  ]

  store.showContextMenu(e, menuItems)
}

const topImage = () => {
  if (!selectedImage.value) return

  const index = memeInfo.memes.findIndex((item) => item === selectedImage.value)
  if (index !== -1 && index !== 0) {
    const item = memeInfo.memes.splice(index, 1)[0]
    memeInfo.memes.unshift(item)
    store.setMemeOrderChanged(memeInfo.code, true)
  }
}

const showRenameModal = (fileName: string) => {
  currentRenameFile.value = fileName
  isRenameModalVisible.value = true
}

const showDeleteModal = (fileName: string) => {
  currentDeleteFile.value = fileName
  isDeleteModalVisible.value = true
}

const deleteMemeFile = async (fileName: string) => {
  try {
    await DeleteMemeFile(store.rootPath, memeInfo.code, fileName)
    store.showToast(`已删除表情包: ${fileName}`)

    // 从列表中移除
    const index = memeInfo.memes.findIndex(item => item === fileName)
    if (index !== -1) {
      memeInfo.memes.splice(index, 1)
      store.setMemeOrderChanged(memeInfo.code, true)
    }
  } catch (error) {
    console.error('删除失败:', error)
    store.showToast(`删除失败: ${error}`, 'error')
  }
}

const hideRenameModal = () => {
  isRenameModalVisible.value = false
  currentRenameFile.value = ''
}

const hideDeleteModal = () => {
  isDeleteModalVisible.value = false
  currentDeleteFile.value = ''
}

const handleDeleteConfirm = async () => {
  if (currentDeleteFile.value) {
    await deleteMemeFile(currentDeleteFile.value)
    hideDeleteModal()
  }
}

const handleRenameSuccess = (oldName: string, newName: string) => {
  const index = memeInfo.memes.findIndex(m => m === oldName)
  if (index !== -1) {
    memeInfo.memes[index] = newName

    const currentMemeInfo = store.allMemesPath.find((meme) => meme.code === memeInfo.code)
    if (currentMemeInfo) {
      currentMemeInfo.memes = [...memeInfo.memes]
    }
  }
}

const handleDragEnd = (event: any) => {
  const { newIndex, oldIndex } = event
  if (newIndex !== oldIndex) {
    const currentMemeInfo = store.allMemesPath.find((meme) => meme.code === memeInfo.code)
    if (currentMemeInfo) {
      currentMemeInfo.memes = [...memeInfo.memes]
      store.setMemeOrderChanged(memeInfo.code, true)
      store.showToast('表情包顺序已更新')
    }
  }
}
</script>

<template>
  <div
    class="meme-grid"
    :key="store.forceRefreshKey">
    <VueDraggable
      v-model="memeInfo.memes"
      class="meme-grid-inner"
      :ghostClass="'meme-item-ghost'"
      @end="handleDragEnd">
      <TransitionGroup
        type="transition"
        name="meme-list">
        <div
          v-for="image in memeInfo.memes"
          :key="image"
          class="meme-item">
          <LazyLoadImg
            draggable="true"
            :src="joinShowImgPath(store.rootPath + '/' + memeInfo.code, image)"
            :alt="image"
            class="meme-image"
            @click="handleClick(image)"
            @contextmenu="handleContextMenu($event, image)" />
        </div>
      </TransitionGroup>
    </VueDraggable>
  </div>

  <RenameMemeModal
    v-model:visible="isRenameModalVisible"
    :current-file="currentRenameFile"
    :parent-path="memeInfo.parentPath"
    :on-success="handleRenameSuccess"
    @close="hideRenameModal"
  />

  <DeleteMemeModal
    v-model:visible="isDeleteModalVisible"
    :file-name="currentDeleteFile"
    @close="hideDeleteModal"
    @confirm="handleDeleteConfirm"
  />
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

.meme-grid {
  background: @rgb-b1;
  width: 100%;
  height: 100%;
  padding: 1rem;
  position: relative;
  overflow: auto;
}

.meme-grid-inner {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(140px, 1fr));
  justify-content: start;
  align-content: start;
  gap: 0.875rem;
}

.meme-item {
  display: flex;
  align-items: center;
  border-radius: 1rem;
  position: relative;
  cursor: grab;

  &:active {
    cursor: grabbing;

    .meme-image {
      transform: scale(0.98);
    }
  }
}

.meme-image {
  width: 100%;
  object-fit: cover;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  z-index: 2;
  cursor: pointer;
  border-radius: 0.75rem;
  background: @rgb-b1;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05), 0 1px 3px rgba(0, 0, 0, 0.1);
  border: 1px solid @rgb-b3;

  &:hover {
    transform: translateY(-4px) scale(1.1);
    border-color: rgba(@p, 0.3);
  }
}

.meme-list-move,
.meme-list-enter-active,
.meme-list-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.meme-list-enter-from,
.meme-list-leave-to {
  opacity: 0;
  transform: scale(0.8) translateY(20px);
}

.meme-list-leave-active {
  position: absolute;
  z-index: 0;
}

.meme-item-ghost {
  opacity: 0.5;
  background: rgba(@p, 0.1);
  border: 2px dashed @rgb-p;
}
</style>
