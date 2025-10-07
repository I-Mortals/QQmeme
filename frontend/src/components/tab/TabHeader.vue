<script lang="ts" setup>
import { computed } from 'vue'
import type { TabItem, TabHeaderProps } from './types'

interface Props extends TabHeaderProps {
  fixedLeftSlot?: boolean
  fixedRightSlot?: boolean
  scrollableSlot?: boolean
  customSlot?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  position: 'top',
  show: true
})

const emit = defineEmits<{
  'tab-click': [tab: TabItem, event: MouseEvent]
}>()

const showDefaultHeader = computed(() => {
  return !props.customSlot && props.show
})

const handleTabClick = (tab: TabItem, event: MouseEvent) => {
  if (tab.disabled) return
  emit('tab-click', tab, event)
}

const getTabClass = (tab: TabItem) => {
  return {
    'tab-header-item': true,
    'tab-header-item--active': tab.key === props.activeKey,
    'tab-header-item--disabled': tab.disabled
  }
}
</script>

<template>
  <div
    v-if="show"
    :class="[
      'tab-header',
      `tab-header--${position}`,
      props.class
    ]"
  >
    <!-- header自定义插槽 -->
    <slot name="header" v-if="customSlot"/>

    <template v-else-if="showDefaultHeader">
      <!-- 左侧固定部分 -->
      <div v-if="fixedLeftSlot" class="tab-header__fixed tab-header__fixed--left">
        <slot name="fixedLeft"/>
      </div>

      <!-- 分隔符 -->
      <div v-if="fixedLeftSlot && scrollableSlot" class="tab-header__separator"></div>

      <!-- 可滚动部分 -->
      <div v-if="scrollableSlot" class="tab-header__scrollable">
        <slot name="scrollable"/>
      </div>

      <!-- 分隔符 -->
      <div v-if="scrollableSlot && fixedRightSlot" class="tab-header__separator"></div>

      <!-- 右侧固定部分 -->
      <div v-if="fixedRightSlot" class="tab-header__fixed tab-header__fixed--right">
        <slot name="fixedRight"/>
      </div>

      <!-- 默认tab列表 -->
      <div v-if="!fixedLeftSlot && !fixedRightSlot && !scrollableSlot" class="tab-header__default">
        <div
          v-for="tab in tabs"
          :key="tab.key"
          :class="getTabClass(tab)"
          @click="handleTabClick(tab, $event)"
        >
          <span v-if="tab.icon" class="tab-header-item__icon">{{ tab.icon }}</span>
          <span class="tab-header-item__label">{{ tab.label }}</span>
        </div>
      </div>
    </template>
  </div>
</template>

<style lang="less" scoped>
.tab-header {
  display: flex;
  gap: .25rem;
  align-items: center;
  background: white;
  border-bottom: 1px solid #e5e5e5;
  position: relative;
  z-index: 10;
  backdrop-filter: blur(20px);
  order: 0;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);

  &--bottom {
    order: -1;
    border-top: 1px solid #e5e5e5;
    border-bottom: none;

    .tab-header-item {
      border-bottom: none;
      border-top: 2px solid transparent;

      &--active {
        border-top-color: var(--theme-primary);
      }
    }
  }

  &__fixed {
    display: flex;
    align-items: center;
    flex-shrink: 0;
  }

  &__separator {
    width: 4px;
    height: 2rem;
    background: rgba(255, 255, 255, 0.3);
    border-radius: 4px;
    flex-shrink: 0;
  }

  &__scrollable {
    flex: 1;
    overflow-x: auto;
    display: flex;
    gap: 0.5rem;
    align-items: center;
  }

  &__default {
    display: flex;
    align-items: center;
    flex: 1;
    overflow-x: auto;
  }
}

.tab-header-item {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 12px 16px;
  cursor: pointer;
  transition: all 0.2s ease;
  white-space: nowrap;
  color: #666;
  font-weight: 400;
  position: relative;
  border-bottom: 2px solid transparent;
  background: transparent;

  &:hover {
    color: var(--theme-primary);
    background: rgba(0, 0, 0, 0.02);
  }

  &--active {
    color: var(--theme-primary);
    border-bottom-color: var(--theme-primary);
    background: transparent;
  }

  &--disabled {
    opacity: 0.5;
    cursor: not-allowed;

    &:hover {
      background: transparent;
      color: #666;
    }
  }

  &__icon {
    width: 14px;
    height: 14px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  &__label {
    font-size: 13px;
  }
}
</style>
