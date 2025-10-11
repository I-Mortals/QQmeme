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
    'tab-header-item-active': tab.key === props.activeKey,
    'tab-header-item-disabled': tab.disabled
  }
}
</script>

<template>
  <div
    v-if="show"
    :class="[
      'tab-header',
      `tab-header-${position}`,
      props.class
    ]"
  >
    <!-- header自定义插槽 -->
    <slot name="header" v-if="customSlot"/>

    <template v-else-if="showDefaultHeader">
      <!-- 左侧固定部分 -->
      <div v-if="fixedLeftSlot" class="tab-header-fixed tab-header-fixed-left">
        <slot name="fixedLeft"/>
      </div>

      <!-- 分隔符 -->
      <div v-if="fixedLeftSlot && scrollableSlot" class="tab-header-separator"></div>

      <!-- 可滚动部分 -->
      <div v-if="scrollableSlot" class="tab-header-scrollable">
        <slot name="scrollable"/>
      </div>

      <!-- 默认tab列表 -->
      <div v-else class="tab-header-default">
        <div
          v-for="tab in tabs"
          :key="tab.key"
          :class="getTabClass(tab)"
          @click="handleTabClick(tab, $event)"
        >
          <span v-if="tab.icon" class="tab-header-item-icon">{{ tab.icon }}</span>
          <span class="tab-header-item-label">{{ tab.label }}</span>
        </div>
      </div>

      <!-- 分隔符 -->
      <div v-if="fixedRightSlot" class="tab-header-separator"></div>

      <!-- 右侧固定部分 -->
      <div v-if="fixedRightSlot" class="tab-header-fixed tab-header-fixed-right">
        <slot name="fixedRight"/>
      </div>
    </template>
  </div>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

.tab-header {
  display: flex;
  gap: .25rem;
  align-items: center;
  background: @rgb-b1;
  border-bottom: 1px solid @rgb-b3;
  padding: 0 0.25rem;
  position: relative;
  z-index: 10;
  backdrop-filter: blur(20px);
  order: 0;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);

  &.tab-header-bottom {
    order: -1;
    border-top: 1px solid @rgb-b3;
    border-bottom: none;

    .tab-header-item {
      border-bottom: none;
      border-top: 2px solid transparent;

      &.tab-header-item-active {
        border-top-color: @rgb-p;
      }
    }
  }

  .tab-header-fixed {
    display: flex;
    align-items: center;
    flex-shrink: 0;
  }

  .tab-header-separator {
    width: 4px;
    height: 1.5rem;
    background: rgba(var(--bc), 0.3);
    border-radius: 4px;
    flex-shrink: 0;
  }

  .tab-header-scrollable {
    flex: 1;
    overflow-x: auto;
    display: flex;
    gap: 0.5rem;
    align-items: center;
  }

  .tab-header-default {
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
  color: @rgb-bc;
  font-weight: 400;
  position: relative;
  border-bottom: 2px solid transparent;
  background: transparent;

  &:hover {
    color: @rgb-p;
    background: rgba(var(--bc), 0.06);
  }

  &.tab-header-item-active {
    color: @rgb-p;
    border-bottom-color: @rgb-p;
    background: transparent;
  }

  &.tab-header-item-disabled {
    opacity: 0.5;
    cursor: not-allowed;

    &:hover {
      background: transparent;
      color: @rgb-bc;
    }
  }

  .tab-header-item-icon {
    width: 14px;
    height: 14px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .tab-header-item-label {
    font-size: 13px;
  }
}
</style>
