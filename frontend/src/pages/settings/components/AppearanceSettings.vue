<script lang="ts" setup>
import { store } from '@/store'
import ColorPicker from '@/components/ColorPicker.vue'
import { colorToRgb } from '@/utils/color'

const hexToRgbaStr = (hex: string): string => {
  const { r, g, b, alpha} = colorToRgb(hex)
  return `${r},${g},${b},${alpha}`
}

const setThemeColor = (type: string, rgba: string) => {
  switch (type){
    case 'theme':
      store.setThemeColor(hexToRgbaStr(rgba))
      break;
    case 'bg':
      store.setThemeBackgroundColor(hexToRgbaStr(rgba))
      break;
  }
}
</script>

<template>
  <div id="appearance-settings" class="content-section">
    <h3 class="section-title">外观设置</h3>
    <div class="setting-group">
      <div class="setting-item">
        <div class="setting-label">
          <span class="label-text">主题色</span>
          <span class="label-desc">自定义应用的主题颜色</span>
        </div>
        <div class="setting-control">
          <ColorPicker
            :value="`rgba(${store.themeColor})`"
            @update:value="(color) => {
              setThemeColor('theme', color)
            }">
            <template #trigger="{ currentColor }">
              <div class="color-picker-group">
                <div class="color-preview-container">
                  <div
                    class="color-preview"
                    :style="{ backgroundColor: currentColor }" />
                </div>
                <span class="color-hex">{{ currentColor }}</span>
              </div>
            </template>
          </ColorPicker>
        </div>
      </div>

      <div class="setting-item">
        <div class="setting-label">
          <span class="label-text">背景色</span>
          <span class="label-desc">自定义应用的背景颜色</span>
        </div>
        <div class="setting-control">
          <ColorPicker
            :value="`rgba(${store.themeBackgroundColor})`"
            @update:value="(color) => {
              setThemeColor('bg', color)
            }">
            <template #trigger="{ currentColor }">
              <div class="color-picker-group">
                <div class="color-preview-container">
                  <div
                    class="color-preview"
                    :style="{ backgroundColor: currentColor }" />
                </div>
                <span class="color-hex">{{ currentColor }}</span>
              </div>
            </template>
          </ColorPicker>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
.content-section {
  margin-bottom: 1.5rem;

  &:last-child {
    margin-bottom: 0;
  }
}

.section-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: #1f2937;
  margin: 0 0 1rem 0;
  padding-bottom: 0.5rem;
  border-bottom: 1px solid #e5e7eb;
}

.setting-group {
  background: #ffffff;
  border-radius: 0.5rem;
  border: 1px solid #e5e7eb;
  overflow: hidden;
}

.setting-item {
  display: flex;
  align-items: center;
  padding: 0.875rem 1.25rem;
  border-bottom: 1px solid #f3f4f6;

  &:last-child {
    border-bottom: none;
  }
}

.setting-label {
  flex: 1;
  min-width: 0;
}

.label-text {
  display: block;
  font-size: 0.875rem;
  font-weight: 500;
  color: #374151;
  margin-bottom: 0.25rem;
}

.label-desc {
  display: block;
  font-size: 0.75rem;
  color: #6b7280;
  line-height: 1.4;
}

.setting-control {
  flex-shrink: 0;
  margin-left: 1rem;
}

.color-picker-group {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.color-preview-container {
  position: relative;
  width: 1.5rem;
  height: 1.5rem;
  border-radius: 0.25rem;
  border: 1px solid rgba(0, 0, 0, 0.2);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  background-image: url("data:image/svg+xml,%3csvg width='100' height='100' xmlns='http://www.w3.org/2000/svg'%3e%3cdefs%3e%3cpattern id='a' patternUnits='userSpaceOnUse' width='10' height='10'%3e%3crect width='10' height='10' fill='%23fff'/%3e%3crect width='5' height='5' fill='%23ccc'/%3e%3crect x='5' y='5' width='5' height='5' fill='%23ccc'/%3e%3c/pattern%3e%3c/defs%3e%3crect width='100' height='100' fill='url(%23a)'/%3e%3c/svg%3e");
  overflow: hidden;
  cursor: pointer;
  transition: all 0.2s ease;

  &:hover {
    transform: scale(1.05);
  }
}

.color-preview {
  width: 100%;
  height: 100%;
  border-radius: 0.25rem;
  transition: transform 0.1s ease;
}

.color-hex {
  font-family: 'SF Mono', 'Monaco', 'Inconsolata', 'Roboto Mono', monospace;
  font-size: 0.875rem;
  font-weight: 500;
  color: #374151;
  background: #ffffff;
  padding: 0.375rem 0.625rem;
  border-radius: 0.375rem;
  border: 1px solid #d1d5db;
  min-width: 4.5rem;
  text-align: center;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}
</style>
