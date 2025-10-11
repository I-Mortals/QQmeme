<template>
  <div class="theme-switch">
    <Button
      ref="triggerRef"
      variant="primary"
      @click="toggleDropdown"
      icon="tabler:category-filled"
    >
      {{ currentTheme }}
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
          <div class="theme-preview" :style="{ backgroundColor: getThemeConfig(theme).base1 }"></div>
          <span class="theme-label">{{ theme }}</span>
        </div>
      </div>
    </Popup>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { store } from '@/store'
import { Popup } from '@/components/popup'
import Button from '@/components/Button.vue'

const showDropdown = ref(false)
const triggerRef = ref<HTMLElement>()

const currentTheme = computed(() => store.currentTheme)
const availableThemes = computed(() => store.availableThemes)

const toggleDropdown = () => {
  showDropdown.value = !showDropdown.value
}

const closeDropdown = () => {
  showDropdown.value = false
}

const selectTheme = (theme: string) => {
  store.setTheme(theme)
  closeDropdown()
}

const getThemeConfig = (theme: string) => {
  return store.getThemeConfig(theme)
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

.theme-preview {
  width: 1.25rem;
  height: 1.25rem;
  border-radius: 0.25rem;
  border: 1px solid @rgb-b3;
  flex-shrink: 0;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}

.theme-label {
  font-size: 0.8rem;
  font-weight: 500;
  text-transform: capitalize;
}
</style>
