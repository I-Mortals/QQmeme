<script lang="ts" setup>
import { defineProps, withDefaults } from 'vue'
import { GenerateAllMemePath, SelectRootDir } from '../../wailsjs/go/memeFile/memeFile'
import { store } from '../store'

interface SettingProps {
  visible?: boolean;
  closeCall?: () => void;
}

const props = withDefaults(defineProps<SettingProps>(), {
  title: '设置',
  visible: false,
  closeCall: () => { }
})

const selectRoot = async () => {
  const path = await SelectRootDir()
  if (path) {
    store.rootPath = path
    store.allMemesPath = await GenerateAllMemePath(store.rootPath)
  }
}

const closeSetting = () => {
  props.closeCall()
}

const clearCache = () => {
  store.clearCache()
  store.rootPath = ''
}
</script>

<template>
  <main class="setting">
    <div class="setting-inner">
      <h1 class="setting-title">设置</h1>
      <div class="setting-item">
        <button class="setting-label" @click="selectRoot">选择MEME主目录</button>
        <span class="setting-el">{{ store.rootPath || "请选择主目录" }}</span>
      </div>
      <div class="setting-item">
        <button @click="clearCache">清除图片内存缓存</button>
        <button @click="closeSetting">关闭</button>
      </div>
    </div>
  </main>
</template>


<style scoped>
.setting {
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  position: fixed;
  z-index: 100;
  top: 0;
  left: 0;
}

.setting-title {
  color: #ffffff;
}

.setting-inner {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  color: #333;
  position: fixed;
  top: 0;
  left: 0;
}

.setting-item {
  display: flex;
  align-items: center;
  height: 2rem;
}

.setting-item .setting-label {
  border: 1px solid #c6c6c6;
  padding: .25rem .5rem;
  border-top-left-radius: .25rem;
  border-bottom-left-radius: .25rem;
  cursor: pointer;
  background: #ffffff;
  align-content: center;
  height: 100%;
}

.setting-item .setting-el {
  border: 1px solid #c6c6c6;
  border-left: 0;
  margin-left: -1px;
  background-color: #a2a2a2;
  padding: .25rem .5rem;
  border-top-right-radius: .25rem;
  border-bottom-right-radius: .25rem;
  align-content: center;
  height: 100%;
}
</style>
