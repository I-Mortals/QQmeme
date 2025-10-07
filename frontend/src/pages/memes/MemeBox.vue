<script lang="ts" setup>
import { computed, ref } from 'vue'
import { store } from '@/store'
import { Tab, TabPanel } from '@/components/tab'
import MemePane from './MemePane.vue'
import StarMemePane from './StarMemePane.vue'
import TelegramStickerPane from './TelegramStickerPane.vue'
import type { TabItem } from '@/components/tab/types'
import { joinShowImgPath } from '@/utils/path'
import { VueDraggable } from 'vue-draggable-plus'
import SaveTabMemeOrderModal from './modal/SaveTabMemeOrderModal.vue'
import SaveTabOrderModal from './modal/SaveTabOrderModal.vue'

const tabsConfig = computed(() => {
  const configItems: Array<{
    key: string
    label: string
    icon: string
    component: any
    data?: any
    props?: any
  }> = [
    {
      key: 'starMemes',
      label: '收藏夹',
      icon: '⭐',
      component: StarMemePane
    },
    {
      key: 'telegramStickers',
      label: 'Telegram',
      icon: '📱',
      component: TelegramStickerPane
    }
  ]

  store.allMemesPath.forEach(memeInfo => {
    configItems.push({
      key: memeInfo.code,
      label: memeInfo.name,
      icon: memeInfo.icon,
      data: memeInfo,
      component: MemePane,
      props: { memeInfo }
    })
  })

  return configItems
})

const tabs = computed<TabItem[]>(() => {
  return tabsConfig.value.map(config => ({
    key: config.key,
    label: config.label,
    icon: config.icon,
    data: config.data
  }))
})

const topTabActive = ref('')
const isConfirmModalOpen = ref(false)
const isSaveTabOrderModalOpen = ref(false)

const handleTabChange = (newKey: string) => {
  store.handleTabClick(newKey)
}

const handleTabClick = (tab: TabItem) => {
  store.handleTabClick(tab.key)
}

const handleContextMenu = (e: MouseEvent, memeCode: string) => {
  topTabActive.value = memeCode

  const menuItems = [
    {
      icon: '📌',
      label: '置顶',
      action: topTab,
      separator: false
    },
    {
      icon: '💾',
      label: '保存表情包顺序',
      action: () => {
        topTabActive.value = memeCode
        isConfirmModalOpen.value = true
      },
      separator: true
    },
    {
      icon: '💾',
      label: '保存 tab 选项顺序',
      action: () => isSaveTabOrderModalOpen.value = true,
      separator: false
    }
  ]

  store.showContextMenu(e, menuItems)
}

const topTab = () => {
  if (!topTabActive.value) return

  const index = store.allMemesPath.findIndex(item => item.code === topTabActive.value)
  if (index !== -1) {
    const item = store.allMemesPath.splice(index, 1)[0]
    store.allMemesPath.unshift(item)
  }
}

const handleSaveMemeOrderSuccess = async () => {
  store.setMemeOrderChanged(topTabActive.value, false)

  await store.refreshMemes()

  store.forceRefreshCurrentTab()

  isConfirmModalOpen.value = false
}

const handleDragEnd = (event: any) => {
  const { newIndex, oldIndex } = event
  if (newIndex !== oldIndex) {
    store.setTabOrderChanged(true)
    store.showToast('tab 顺序已更新')
  }
}

const handleSaveTabOrderSuccess = () => {
  store.setTabOrderChanged(false)
  isSaveTabOrderModalOpen.value = false
}
</script>

<template>
  <main class="meme-box">
    <Tab
      v-model="store.tabCurrent"
      header-position="bottom"
      :tabs="tabs"
      :lazy="true"
      @tab-change="handleTabChange"
      @tab-click="handleTabClick"
    >
      <template #fixedLeft>
        <div class="tab-fixed-list">
          <div
            @click="store.handleTabClick('starMemes')"
            :class="[
              'tab-item',
              'star-tab-item',
              { 'tab-item-action': store.tabCurrent === 'starMemes' }
            ]">
            <div class="star-icon">⭐</div>
          </div>

          <div
            @click="store.handleTabClick('telegramStickers')"
            :class="[
              'tab-item',
              'telegram-tab-item',
              { 'tab-item-action': store.tabCurrent === 'telegramStickers' }
            ]">
            <div class="telegram-icon">📱</div>
          </div>
        </div>
      </template>

      <template #scrollable>
        <div class="tab-list" :key="store.forceRefreshKey">
          <VueDraggable
            v-model="store.allMemesPath"
            class="tab-list-inner"
            :ghostClass="'tab-item-ghost'"
            @end="handleDragEnd">
            <TransitionGroup name="tab-list">
              <div
                v-for="meme in store.allMemesPath"
                :key="meme.code"
                @click="store.handleTabClick(meme.code)"
                @contextmenu="handleContextMenu($event, meme.code)"
                :data-meme-tab="meme.code"
                :class="[
                  'tab-item',
                  { 'tab-item-action': meme.code === store.tabCurrent },
                  { 'tab-item-changed': meme.orderChanged }
                ]">
                <img
                  id="MemeTabItem"
                  :src="joinShowImgPath(meme.parentPath, meme.icon)"
                  :alt="meme.code"/>
              </div>
            </TransitionGroup>
          </VueDraggable>
        </div>
      </template>

      <TabPanel
        v-for="config in tabsConfig"
        :key="config.key"
        :tab-key="config.key"
        :active="store.tabCurrent === config.key"
      >
        <component
          :is="config.component"
          v-bind="config.props || {}"
        />
      </TabPanel>
    </Tab>

    <!-- 模态框 -->
    <SaveTabMemeOrderModal
      v-model:visible="isConfirmModalOpen"
      :tab-name="topTabActive"
      :parent-path="store.allMemesPath.find((meme: any) => meme.code === topTabActive)?.parentPath || ''"
      :memes="store.allMemesPath.find((meme: any) => meme.code === topTabActive)?.memes || []"
      :on-success="handleSaveMemeOrderSuccess"
      @close="() => isConfirmModalOpen = false"
    />

    <SaveTabOrderModal
      v-model:visible="isSaveTabOrderModalOpen"
      @close="() => isSaveTabOrderModalOpen = false"
      :on-success="handleSaveTabOrderSuccess"
    />
  </main>
</template>

<style lang="less" scoped>
.meme-box {
  width: 100%;
  height: 100%;
  overflow: auto;
  position: relative;
  scroll-behavior: smooth;

  &::before {
    content: '';
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: radial-gradient(circle at 20% 80%, rgba(120, 119, 198, 0.3) 0%, transparent 50%),
    radial-gradient(circle at 80% 20%, rgba(255, 119, 198, 0.3) 0%, transparent 50%),
    radial-gradient(circle at 40% 40%, rgba(120, 219, 255, 0.2) 0%, transparent 50%);
    pointer-events: none;
    z-index: 0;
  }
}

:deep(.tab-header) {
  background: var(--theme-primary);
  padding: 0 .25rem;
}

.tab-fixed-list {
  display: flex;
  gap: .25rem;
  align-items: center;
}

.tab-list {
  display: flex;
  gap: 0.25rem;
  align-items: center;
  padding: .125rem 0;

  &-inner {
    display: flex;
    gap: 0.25rem;
    align-items: center;
  }

  &-move,
  &-enter-active,
  &-leave-active {
    transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  }

  &-enter-from,
  &-leave-to {
    opacity: 0;
    transform: scale(0.8) translateX(-20px);
  }

  &-leave-active {
    position: absolute;
    z-index: 0;
  }
}

.tab-fixed-list .tab-item,
.tab-list .tab-item {
  background: rgba(255, 255, 255, 0.12);
  cursor: grab;
  height: 3rem;
  width: 3rem;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 0.5rem;
  padding: 0.25rem;
  flex: none;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border: 2px solid transparent;

  &:active {
    cursor: grabbing;
  }

  &:hover {
    background: rgba(255, 255, 255, 0.2);
    box-shadow: 0 12px 28px rgba(0, 0, 0, 0.15), 0 6px 16px rgba(255, 255, 255, 0.1);
    border-color: rgba(255, 255, 255, 0.3);

    img {
      transform: scale(1.08);
    }
  }

  &-action {
    background: rgba(255, 255, 255, 0.25);
    border-color: rgba(255, 255, 255, 0.6);
  }

  &-changed {
    position: relative;

    &::after {
      content: '';
      position: absolute;
      top: -2px;
      right: -2px;
      width: 12px;
      height: 12px;
      background: #ff6b6b;
      border-radius: 50%;
      font-size: 8px;
      color: white;
      display: flex;
      align-items: center;
      justify-content: center;
      border: 2px solid rgba(255, 255, 255, 0.8);
      box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
    }
  }

  img {
    height: 2rem;
    width: 2rem;
    max-width: 100%;
    color: transparent;
    object-fit: cover;
    border-radius: 0.375rem;
    transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  }
}

.tab-item-ghost {
  opacity: 0.5;
  transform: scale(0.95);
  background: rgba(255, 255, 255, 0.3);
  border-color: rgba(255, 255, 255, 0.5);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.2);
}

.tab-fixed-list {
  .star-tab-item,
  .telegram-tab-item {
    cursor: pointer;

    &:active {
      cursor: pointer;
    }
  }
}

.star-icon,
.telegram-icon {
  font-size: 1.25rem;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);

  &:hover {
    transform: scale(1.08);
  }
}
</style>

