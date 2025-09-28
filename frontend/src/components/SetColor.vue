<script lang="ts" setup>
import { store } from '../store'
import ColorPicker from './ColorPicker.vue'
import { parse, Rgb } from 'culori'

const hexToRgba = (hex: string): string => {
  let rgb = parse(hex) as Rgb

  const { r, g, b, alpha = 1 } = rgb

  const _r = Math.round(r * 255),
    _g = Math.round(g * 255),
    _b = Math.round(b * 255)

  return `${_r},${_g},${_b},${alpha}`
}

const rgbaToRgba = (rgba: string): string => {
  const { r, g, b, alpha } = parse(rgba) as Rgb

  const parsedRgb = {
    r: Math.round(r * 255),
    g: Math.round(g * 255),
    b: Math.round(b * 255),
    alpha
  }

  return `${parsedRgb.r},${parsedRgb.g},${parsedRgb.b},${parsedRgb.alpha}`
}

const handleThemeColorChange = (colorValue: string) => {
  console.log(colorValue)
  const rgbaColor = colorValue.includes('rgba') ? rgbaToRgba(colorValue) : hexToRgba(colorValue)
  store.setThemeColor(rgbaColor)
}

const handleThemeBackgroundColorChange = (colorValue: string) => {
  const rgbaColor = colorValue.includes('rgba') ? rgbaToRgba(colorValue) : hexToRgba(colorValue)
  store.setThemeBackgroundColor(rgbaColor)
}
</script>

<template>
  <div class="set-color-body">
    <div class="color-picker-container">
      <h4>主题色：</h4>
      <ColorPicker
        :value="`rgba(${store.themeColor})`"
        @update:value="handleThemeColorChange" />
    </div>
    <div class="color-picker-container">
      <h4>主题背景色：</h4>
      <ColorPicker
        :value="`rgba(${store.themeBackgroundColor})`"
        @update:value="handleThemeBackgroundColorChange" />
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
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.color-picker-container h4 {
  margin: 0;
  font-size: 0.9rem;
  font-weight: 500;
  color: #333;
}
</style>
