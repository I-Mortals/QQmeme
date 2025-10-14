<script setup lang="ts">
import { computed, nextTick } from 'vue'
import { memeStore, toastStore, contextStore } from '@/store'
import type { StarMemeItem } from '@/store/memeStore'
import { WriteFileToClipboard } from '@wailsjs/go/memeFile/MemeFile'
import LazyLoadImg from '@/components/LazyLoadImg.vue'
import { VueDraggable } from 'vue-draggable-plus'
import { joinPath, joinShowImgPath } from '@/utils/path'

const handleClick = async (starMeme: StarMemeItem) => {
  try {
    await WriteFileToClipboard(joinPath(memeStore.rootPath + '/' + starMeme.fromFolder, starMeme.fileName))
    toastStore.showToast('复制成功！')
  } catch (error) {
    toastStore.showToast(`复制失败：${error}`, 'error')
  }
}

const parseFolder = computed(() => {
  return (folder: string) => {
    const _firstIndex = folder.indexOf('_')
    return folder.substring(_firstIndex + 1)
  }
})

const handleContextMenu = (e: MouseEvent, starMeme: StarMemeItem) => {
  const menuItems = [
    {
      icon: '🗑️',
      label: '从收藏夹移除',
      action: () => {
        memeStore.removeFromStarMemes(starMeme.id)
        toastStore.showToast('已取消收藏！')
      }
    },
    {
      icon: '📌',
      label: '置顶',
      action: () => topStar(starMeme.id)
    },
    {
      icon: '🔗',
      label: `跳转到 ${parseFolder.value(starMeme.fromFolder)}`,
      action: () => goToParentTab(starMeme.fromFolder),
      separator: true
    },
    {
      icon: '🔄',
      label: '刷新缓存',
      action: () => {
        memeStore.refreshMemes()
        toastStore.showToast('缓存刷新成功！', 'success')
      }
    }
  ]

  contextStore.showContextMenu(e, menuItems)
}

const goToParentTab = (tabKey: string) => {
  const targetTab = memeStore.allMemesPath.find(tab => tab.code === tabKey)
  if (targetTab) {
    memeStore.handleTabClick(tabKey)

    nextTick(() => {
      const targetElement = document.querySelector(`[data-meme-tab="${tabKey}"]`) as HTMLElement
      if (targetElement) {
        targetElement.scrollIntoView({
          behavior: 'smooth',
          block: 'nearest',
          inline: 'center'
        })
      }
    })

    toastStore.showToast(`已跳转到 ${parseFolder.value(tabKey)} 标签页`)
  } else {
    toastStore.showToast(`找不到 ${parseFolder.value(tabKey)} 标签页`, 'error')
  }
};

const topStar = (id: string) => {
  const index = memeStore.starMemes.findIndex(item => item.id === id)
  if (index !== -1 && index !== 0) {
    const item = memeStore.starMemes.splice(index, 1)[0]
    memeStore.starMemes.unshift(item)
    toastStore.showToast('已置顶到收藏夹首位')
  }
}

const handleDragEnd = (event: any) => {
  const { newIndex, oldIndex } = event
  if (newIndex !== oldIndex) {
    toastStore.showToast('收藏夹表情包顺序已更新')
  }
}
</script>

<template>
  <div
    class="star-grid"
    :key="memeStore.forceRefreshKey">

    <div v-if="memeStore.starMemes.length === 0" class="star-meme-empty">
      <div class="empty-icon">⭐</div>
      <h3>收藏夹是空的</h3>
      <p>右键点击任意表情包选择"收藏"来添加到这里</p>
    </div>

    <VueDraggable
      v-else
      v-model="memeStore.starMemes"
      class="star-grid-inner"
      :ghostClass="'star-item-ghost'"
      @end="handleDragEnd">
      <TransitionGroup
        type="transition"
        name="star-list">
        <div
          v-for="star in memeStore.starMemes"
          :key="star.id"
          class="star-item"
          @contextmenu="handleContextMenu($event, star)">
          <LazyLoadImg
            :src="joinShowImgPath(memeStore.rootPath + '/' + star.fromFolder, star.fileName)"
            :alt="star.fileName"
            class="star-image"
            @click="handleClick(star)" />
          <div class="star-info">
<!--            <div class="star-name">{{ star.fileName }}</div>-->
            <div class="star-from">来自: {{ parseFolder(star.fromFolder) }}</div>
          </div>
        </div>
      </TransitionGroup>
    </VueDraggable>
  </div>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

.star-grid {
  background: @rgb-b1;
  width: 100%;
  padding: 1rem;
  position: relative;
}

.star-meme-empty {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 60vh;
  text-align: center;
  color: @rgb-bc;

  h3 {
    font-size: 1.5rem;
    margin: 0 0 0.5rem 0;
    color: @rgb-bc;
  }

  p {
    font-size: 1rem;
    margin: 0;
    opacity: 0.8;
  }
}

.empty-icon {
  font-size: 4rem;
  margin-bottom: 1rem;
  opacity: 0.5;
}

.star-grid-inner {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(160px, 1fr));
  justify-content: start;
  align-content: start;
  gap: 0.875rem;
}

.star-item {
  display: flex;
  flex-direction: column;
  border-radius: 1rem;
  position: relative;
  cursor: grab;
  background: @rgb-b1;
  padding: 0.75rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.05), 0 1px 3px rgba(0, 0, 0, 0.1);
  border: 1px solid @rgb-b3;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);

  &:active {
    cursor: grabbing;
  }

  &:hover {
    transform: translateY(-4px) scale(1.02);
    border-color: rgba(@p, 0.3);
  }
}

.star-image {
  width: 100%;
  height: 160px;
  object-fit: cover;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  z-index: 2;
  cursor: pointer;
  border-radius: 0.5rem;
  margin-bottom: 0.5rem;
}

.star-info {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.star-name {
  font-size: 0.875rem;
  font-weight: 600;
  color: @rgb-bc;
  word-break: break-all;
}

.star-from {
  font-size: 0.75rem;
  color: @rgb-bc;
  opacity: 0.8;
  word-break: break-all;
}

.star-list-move,
.star-list-enter-active,
.star-list-leave-active {
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
}

.star-list-enter-from,
.star-list-leave-to {
  opacity: 0;
  transform: scale(0.8) translateY(20px);
}

.star-list-leave-active {
  position: absolute;
  z-index: 0;
}

.star-item-ghost {
  opacity: 0.5;
  background: rgba(@p, 0.1);
  border: 2px dashed @rgb-p;
}
</style>
