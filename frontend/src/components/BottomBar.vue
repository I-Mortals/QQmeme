<script lang="ts" setup>
import { ref } from 'vue'
import { store } from '../store'
import MeSetting from './Setting.vue'
import { GenerateAllMemePath, Hello } from '../../wailsjs/go/memeFile/memeFile'
import { joinShowImgPath } from '../utils'
import { clsx } from 'clsx'

const isShowSetting = ref(false)

const handleTabClick = (tab: string) => {
  store.handleTabClick(tab)
}

// setting
const openSetting = async () => {
  isShowSetting.value = true
  await Hello()
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
      <div
        v-for="meme in store.allMemesPath"
        :key="meme.Code"
        @click="handleTabClick(meme.Code)"
        :class="clsx(
           'meme-item',
           {'meme-item-action':meme.Code === store.tabCurrent}
         )"
      >
        <img :src="joinShowImgPath(meme.ParentPath,meme.Icon)" :alt="meme.Code"/>
      </div>
    </div>
    <MeSetting v-if="isShowSetting" :closeCall="closeSetting"/>
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
  background-color: #f3f4f7;
  z-index: 1000;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.tab-bar-inner {
  height: 100%;
  display: flex;
  gap: .5rem;
  padding: .5rem;
  overflow: auto;
}

.setting-btn,
.refresh-btn {
  white-space: nowrap;
  border: 1px solid #c6c6c6;
  border-radius: .25rem;
  cursor: pointer;
}

.meme-item {
  background-color: cadetblue;
  cursor: pointer;
  height: 100%;
  align-content: center;
  border-radius: .125rem;
  padding: .25rem;
}

.meme-item-action {
  background-color: aqua;
}

.meme-item img {
  height: 100%;
  /* 禁止拖拽 */
  user-drag: none;
  -webkit-user-drag: none;
  user-select: none;
  -moz-user-select: none;
  -webkit-user-select: none;
  -ms-user-select: none;
  /* 隐藏alt提示 */
  font-size: 0;
  color: transparent;
  object-fit: cover;
  border-radius: .125rem;
}
</style>
