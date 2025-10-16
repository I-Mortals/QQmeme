<template>
  <div class="theme-switch">
    <Button
      ref="triggerRef"
      variant="primary"
      @click="toggleDropdown"
    >
      <div class="theme-preview-grid" :style="{ backgroundColor: getThemeColor(currentTheme, 'base1') }">
        <div class="preview-square" :style="{ backgroundColor: getThemeColor(currentTheme, 'baseContent') }"></div>
        <div class="preview-square" :style="{ backgroundColor: getThemeColor(currentTheme, 'primary') }"></div>
        <div class="preview-square" :style="{ backgroundColor: getThemeColor(currentTheme, 'warning') }"></div>
        <div class="preview-square" :style="{ backgroundColor: getThemeColor(currentTheme, 'success') }"></div>
      </div>
      <span style="text-transform: capitalize;">{{ currentTheme }}</span>
    </Button>

    <Popup
      :visible="showDropdown"
      :trigger="triggerRef"
      placement="bottom-end"
      :offset="8"
      :click-outside="true"
      @close="closeDropdown"
      class-name="theme-popup"
    >
      <div class="theme-dropdown">
        <div
          v-for="theme in availableThemes"
          :key="theme"
          class="theme-item"
          :class="{ active: theme === currentTheme }"
          @click="selectTheme(theme)"
        >
          <div class="theme-preview-grid" :style="{ backgroundColor: getThemeColor(theme, 'base1') }">
            <div class="preview-square" :style="{ backgroundColor: getThemeColor(theme, 'baseContent') }"></div>
            <div class="preview-square" :style="{ backgroundColor: getThemeColor(theme, 'primary') }"></div>
            <div class="preview-square" :style="{ backgroundColor: getThemeColor(theme, 'warning') }"></div>
            <div class="preview-square" :style="{ backgroundColor: getThemeColor(theme, 'success') }"></div>
          </div>
          <span class="theme-label">{{ theme }}</span>
        </div>
      </div>
    </Popup>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { themeStore } from '@/store'
import { Popup } from '@/components/popup'
import Button from '@/components/Button.vue'
import { getThemeColor } from '@/styles/themes'

const showDropdown = ref(false)
const triggerRef = ref<HTMLElement>()

const currentTheme = computed(() => themeStore.currentTheme)
const availableThemes = computed(() => themeStore.availableThemes)

const toggleDropdown = () => {
  showDropdown.value = !showDropdown.value
}

const closeDropdown = () => {
  showDropdown.value = false
}

const selectTheme = (theme: string) => {
  themeStore.setTheme(theme)
  closeDropdown()
}
</script>

<style lang="less" scoped>
@import '@/styles/variables.less';

.theme-switch {
  position: relative;
  display: inline-block;
}

.theme-dropdown {
  background: @rgb-b1;
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
  max-height: 20rem;
  overflow-y: auto;
}

.theme-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  cursor: pointer;
  transition: all 0.2s ease;
  color: @rgb-bc;
  border-radius: 0.25rem;

  &:hover {
    background: rgba(var(--bc), 0.06);
  }

  &.active {
    background: @rgb-pc;
    color: @rgb-p;
  }
}

.theme-preview-grid {
  width: 1.25rem;
  height: 1.25rem;
  border-radius: 0.25rem;
  border: 1px solid @rgb-b3;
  flex-shrink: 0;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  display: grid;
  gap: .15rem;
  padding: .15rem;
  grid-template-columns: 1fr 1fr;
  grid-template-rows: 1fr 1fr;
  overflow: hidden;
}

.preview-square {
  width: 100%;
  height: 100%;
  border: none;
  border-radius: 50%;
  margin: 0;
  padding: 0;
}

.theme-label {
  font-size: 0.8rem;
  font-weight: 500;
  text-transform: capitalize;
}
</style>
