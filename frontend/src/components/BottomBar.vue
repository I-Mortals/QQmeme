<script lang="ts" setup>
import { ref } from 'vue'
import { store } from '../store'
import MeSetting from './Setting.vue'
import { GenerateAllMemePath } from '../../wailsjs/go/memeFile/memeFile'
import { joinShowImgPath } from '../utils'
import { clsx } from 'clsx'

const isShowSetting = ref(false)

const handleTabClick = (tab: string) => {
  store.handleTabClick(tab)
}

// setting
const openSetting = async () => {
  isShowSetting.value = true
}
const closeSetting = () => {
  isShowSetting.value = false
}

const refreshMemes = async () => {
  store.allMemesPath = await GenerateAllMemePath(store.rootPath)
}

</script>


<template>
  <main class="tab-bar">
    <div class="tab-bar-inner">
      <button class="setting-btn" @click="openSetting()">设置</button>
      <button class="refresh-btn" @click="refreshMemes()">刷新</button>

      <div class="tab-list">
        <div
          v-for="meme in store.allMemesPath"
          :key="meme.code"
          @click="handleTabClick(meme.code)"
          :class="clsx(
           'tab-item',
           {'tab-item-action':meme.code === store.tabCurrent}
         )"
        >
          <img :src="joinShowImgPath(meme.parentPath, meme.icon)" :alt="meme.code"/>
        </div>
      </div>
    </div>
    <MeSetting :visible="isShowSetting" :closeCall="closeSetting"/>
  </main>
</template>

<style scoped>
.tab-bar {
  /* 固定底部 */
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 5rem;
  background: var(--theme-primary);
  backdrop-filter: blur(20px);
  z-index: 1000;
  box-shadow:
    0 -8px 32px rgba(0, 0, 0, 0.12),
    0 -2px 8px rgba(0, 0, 0, 0.08);
  border-top: 1px solid rgba(255, 255, 255, 0.15);
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}


.tab-bar-inner {
  height: 100%;
  display: flex;
  gap: 0.75rem;
  padding: 0.75rem;
  align-items: center;
  position: relative;
}

.setting-btn,
.refresh-btn {
  white-space: nowrap;
  padding: 0.625rem 1rem;
  border-radius: 0.625rem;
  cursor: pointer;
  font-weight: 600;
  font-size: 0.875rem;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  background: rgba(255, 255, 255, 0.15);
  color: white;
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.2);
}


.setting-btn:hover,
.refresh-btn:hover {
  background: rgba(255, 255, 255, 0.25);
  transform: translateY(-3px) scale(1.02);
  box-shadow:
    0 8px 25px rgba(0, 0, 0, 0.15),
    0 4px 12px rgba(255, 255, 255, 0.1);
  border-color: rgba(255, 255, 255, 0.4);
}

.setting-btn:active,
.refresh-btn:active {
  transform: translateY(-1px) scale(0.98);
  transition: all 0.1s ease;
}

.tab-list {
  display: flex;
  gap: 0.5rem;
  overflow-x: auto;
  flex: 1;
  padding: 0.25rem 0;
  scroll-behavior: smooth;
}

.tab-list .tab-item {
  background: rgba(255, 255, 255, 0.12);
  cursor: pointer;
  height: 3.5rem;
  width: 3.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 0.75rem;
  padding: 0.375rem;
  flex: none;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border: 2px solid transparent;
}

.tab-list .tab-item:hover {
  background: rgba(255, 255, 255, 0.2);
  box-shadow:
    0 12px 28px rgba(0, 0, 0, 0.15),
    0 6px 16px rgba(255, 255, 255, 0.1);
  border-color: rgba(255, 255, 255, 0.3);
}

.tab-list .tab-item-action {
  background: rgba(255, 255, 255, 0.25);
  border-color: rgba(255, 255, 255, 0.6);
}

.tab-list .tab-item img {
  height: 2.5rem;
  width: 2.5rem;
  max-width: 100%;
  /* 禁止拖拽 */
  -webkit-user-drag: none;
  user-select: none;
  -moz-user-select: none;
  -webkit-user-select: none;
  -ms-user-select: none;
  /* 隐藏alt提示 */
  font-size: 0;
  color: transparent;
  object-fit: cover;
  border-radius: 0.625rem;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.tab-list .tab-item:hover img {
  transform: scale(1.08);
}

@media (max-width: 768px) {
  .tab-bar {
    height: 5rem;
  }

  .tab-bar-inner {
    gap: 0.75rem;
    padding: 0.75rem;
  }

  .setting-btn,
  .refresh-btn {
    padding: 0.625rem 1rem;
    font-size: 0.8125rem;
  }

  .tab-list .tab-item {
    height: 3.5rem;
    width: 3.5rem;
  }

  .tab-list .tab-item img {
    height: 2.5rem;
    width: 2.5rem;
  }
}

@media (max-width: 480px) {
  .tab-bar {
    height: 4.5rem;
  }

  .tab-bar-inner {
    gap: 0.5rem;
    padding: 0.5rem;
  }

  .setting-btn,
  .refresh-btn {
    padding: 0.5rem 0.75rem;
    font-size: 0.75rem;
  }

  .tab-list {
    gap: 0.5rem;
  }

  .tab-list .tab-item {
    height: 3rem;
    width: 3rem;
  }

  .tab-list .tab-item img {
    height: 2rem;
    width: 2rem;
  }
}
</style>
