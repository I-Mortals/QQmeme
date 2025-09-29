<script lang="ts" setup>
import { ref, onMounted, ImgHTMLAttributes } from 'vue'
import { store } from '../store'
import MeSetting from './Setting.vue'
import { GenerateAllMemePath } from '../../wailsjs/go/memeFile/memeFile'
import { joinShowImgPath } from '../utils'
import { clsx } from 'clsx'

const isShowSetting = ref(false)

const topTabActive = ref('')
const contextMenu = ref<HTMLElement>()


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

onMounted(async () => {
  document.addEventListener('contextmenu', (e) => {
    e.preventDefault(); // 阻止默认右键菜单
    const currentElement = e.target as ImgHTMLAttributes;
    if (currentElement?.id === 'MemeTabItem') {
      const code = currentElement?.alt
      topTabActive.value = code || ''
      showCustomMenu(e.clientX, e.clientY, e.target as HTMLElement)
    }
  });
  // 点击其他地方时隐藏菜单
  document.addEventListener('click', () => {
    hideCustomMenu();
  });

})

// 显示自定义菜单
const showCustomMenu = (x: number, y: number, element: HTMLElement) => {

  if (!contextMenu.value) return
  // 设置菜单位置


  contextMenu.value.style.left = x + 'px';
  contextMenu.value.style.top = 0 + 'px';

  // 显示菜单
  contextMenu.value.style.display = 'block';

  // 防止菜单超出屏幕边界
  const rect = contextMenu.value.getBoundingClientRect();
  const windowWidth = window.innerWidth;
  const windowHeight = window.innerHeight;

  if (rect.right > windowWidth) {
    contextMenu.value.style.left = (x - rect.width) + 'px';
  }

}

// 隐藏自定义菜单
const hideCustomMenu = () => {
  if (!contextMenu.value) return
  contextMenu.value.style.display = 'none';
}

// 置顶
const topTab = () => {
  console.log(topTabActive.value);
  if (!topTabActive.value) return
  // 将code移动到数组的第一位
  const index = store.allMemesPath.findIndex(item => item.code === topTabActive.value);
  if (index !== -1) {
    const item = store.allMemesPath.splice(index, 1)[0];
    store.allMemesPath.unshift(item);
  }
}

</script>


<template>
  <main class="tab-bar">
    <div class="tab-bar-inner">
      <button class="setting-btn" @click="openSetting()">设置</button>
      <button class="refresh-btn" @click="refreshMemes()">刷新</button>

      <div class="tab-list">
        <div v-for="meme in store.allMemesPath" :key="meme.code" @click="handleTabClick(meme.code)" :class="clsx(
          'tab-item',
          { 'tab-item-action': meme.code === store.tabCurrent }
        )">
          <img id="MemeTabItem" :src="joinShowImgPath(meme.parentPath, meme.icon)" :alt="meme.code" />
        </div>
      </div>
    </div>
    <MeSetting :visible="isShowSetting" :closeCall="closeSetting" />
    <!-- 自定义右键菜单 -->
    <div ref="contextMenu" class="context-menu" id="contextMenu">
      <div class="menu-item" @click="topTab">
        <i>📌</i> 置顶
      </div>
    </div>
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

/* 自定义右键 */

.context-menu {
  position: absolute;
  background: white;
  border-radius: 8px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
  /* width: 280px; */
  z-index: 1;
  display: none;
  overflow: hidden;
  animation: fadeIn 0.2s ease;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: scale(0.9);
  }

  to {
    opacity: 1;
    transform: scale(1);
  }
}

.menu-item {
  padding: 6px 6px;
  color: #333;
  cursor: pointer;
  display: flex;
  align-items: center;
  transition: background 0.2s;
  border-bottom: 1px solid #f0f0f0;
}

.menu-item:last-child {
  border-bottom: none;
}

.menu-item:hover {
  background: #f5f5f5;
}

.menu-item i {
  margin-right: 10px;
  width: 10px;
  text-align: center;
}
</style>
