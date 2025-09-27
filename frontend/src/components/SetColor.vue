<script lang="ts" setup>
import { ref } from 'vue'
import { store } from '../store'

// 主题色选择 - 实时变化
const handleThemeColorChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.value) {
    // 将十六进制颜色转换为RGB格式
    const hex = target.value
    const r = parseInt(hex.slice(1, 3), 16)
    const g = parseInt(hex.slice(3, 5), 16)
    const b = parseInt(hex.slice(5, 7), 16)
    const rgbColor = `${r},${g},${b}`
    store.setThemeColor(rgbColor)
  }
}

// 颜色选择完成时显示提示
const handleThemeColorInput = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.value) {
    // 将十六进制颜色转换为RGB格式
    const hex = target.value
    const r = parseInt(hex.slice(1, 3), 16)
    const g = parseInt(hex.slice(3, 5), 16)
    const b = parseInt(hex.slice(5, 7), 16)
    const rgbColor = `${r},${g},${b}`
    store.setThemeColor(rgbColor)
    store.showToast('主题色已更新', 'success')
  }
}

// 主题色选择 - 实时变化
const handleThemeBackgroundColorChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.value) {
    // 将十六进制颜色转换为RGB格式
    const hex = target.value
    const r = parseInt(hex.slice(1, 3), 16)
    const g = parseInt(hex.slice(3, 5), 16)
    const b = parseInt(hex.slice(5, 7), 16)
    const rgbColor = `${r},${g},${b}`
    store.setThemeBackgroundColor(rgbColor)
  }
}

// 颜色选择完成时显示提示
const handleThemeBackgroundColorInput = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.value) {
    // 将十六进制颜色转换为RGB格式
    const hex = target.value
    const r = parseInt(hex.slice(1, 3), 16)
    const g = parseInt(hex.slice(3, 5), 16)
    const b = parseInt(hex.slice(5, 7), 16)
    const rgbColor = `${r},${g},${b}`
    store.setThemeBackgroundColor(rgbColor)
    store.showToast('主题背景色已更新', 'success')
  }
}

</script>


<template>
      <div class="set-color-body">
        <div class="color-picker-container">
            <h4>主题色：</h4>
        <input
          type="color"
          class="color-picker"
          :value="`#${store.themeColor.split(',').map(c => parseInt(c).toString(16).padStart(2, '0')).join('')}`"
          @input="handleThemeColorChange"
          @change="handleThemeColorInput"
          title="选择主题色"
        />
      </div>
      <div class="color-picker-container">
        <h4>主题背景色：</h4>
        <input
          type="color"
          class="color-picker"
          :value="`#${store.themeBackgroundColor.split(',').map(c => parseInt(c).toString(16).padStart(2, '0')).join('')}`"
          @input="handleThemeBackgroundColorChange"
          @change="handleThemeBackgroundColorInput"
          title="选择主题背景色"
        />
      </div>
      </div>
</template>


<style scoped>
.set-color-body {
  display: flex;
  align-items: center;
  color: #333;
  gap: 15px;
}

.color-picker-container {
  position: relative;
  display: flex;
  align-items: center;
}

.color-picker {
  width: 2.5rem;
  height: 2.5rem;
  border-radius: 0.625rem;
  cursor: pointer;
  background: rgba(255, 255, 255, 0.15);
  backdrop-filter: blur(10px);
  border: 1px solid rgb(var(--theme-primary));
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  overflow: hidden;
}

.color-picker:hover {
  background: rgba(255, 255, 255, 0.25);
  transform: translateY(-3px) scale(1.02);
  box-shadow:
    0 8px 25px rgba(0, 0, 0, 0.15),
    0 4px 12px rgba(255, 255, 255, 0.1);
  border-color: rgba(255, 255, 255, 0.4);
}

.color-picker:active {
  transform: translateY(-1px) scale(0.98);
  transition: all 0.1s ease;
}

/* 隐藏默认的颜色选择器样式 */
.color-picker::-webkit-color-swatch-wrapper {
  padding: 0;
}

.color-picker::-webkit-color-swatch {
  border: none;
  border-radius: 0.5rem;
}

.color-picker::-moz-color-swatch {
  border: none;
  border-radius: 0.5rem;
}
</style>