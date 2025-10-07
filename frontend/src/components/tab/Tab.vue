<script lang="ts" setup>
import { ref, computed, provide, watch } from 'vue'
import TabHeader from './TabHeader.vue'
import type { TabProps, TabEmits, TabItem } from './types'

const props = withDefaults(defineProps<TabProps>(), {
  headerPosition: 'top',
  lazy: true,
  showHeader: true
})

const emit = defineEmits<TabEmits>()

const activeKey = ref(props.modelValue)
const loadedTabs = ref(new Set<string>())
watch(() => props.modelValue, (newValue) => {
  activeKey.value = newValue
})

watch(activeKey, (newKey, oldKey) => {
  emit('update:modelValue', newKey)
  if (oldKey) {
    emit('tab-change', newKey, oldKey)
  }
})

const handleTabClick = (tab: TabItem, event: MouseEvent) => {
  if (tab.disabled) return

  activeKey.value = tab.key
  emit('tab-click', tab, event)
}

const handleTabLoad = (tabKey: string) => {
  loadedTabs.value.add(tabKey)
}

const activeTab = computed(() => {
  return props.tabs.find(tab => tab.key === activeKey.value)
})

provide('tabContext', {
  activeKey: computed(() => activeKey.value),
  tabs: computed(() => props.tabs),
  lazy: computed(() => props.lazy),
  loadedTabs: computed(() => loadedTabs.value),
  onTabLoad: handleTabLoad
})

const hasCustomHeaderSlot = computed(() => {
  return !!slots.header
})

const hasFixedLeftSlot = computed(() => {
  return !!slots.fixedLeft
})

const hasFixedRightSlot = computed(() => {
  return !!slots.fixedRight
})

const hasScrollableSlot = computed(() => {
  return !!slots.scrollable
})

const slots = defineSlots<{
  header?: () => any
  fixedLeft?: () => any
  fixedRight?: () => any
  scrollable?: () => any
  default?: () => any
}>()
</script>

<template>
  <div
    :class="[
      'tab-container',
      `tab-container--${headerPosition}`,
      props.class
    ]"
    :style="style"
  >
    <!-- Header在顶部 -->
    <TabHeader
      v-if="headerPosition === 'top' && showHeader"
      :tabs="tabs"
      :active-key="activeKey"
      :position="headerPosition"
      :show="showHeader"
      :custom-slot="hasCustomHeaderSlot"
      :fixed-left-slot="hasFixedLeftSlot"
      :fixed-right-slot="hasFixedRightSlot"
      :scrollable-slot="hasScrollableSlot"
      @tab-click="handleTabClick"
    >
      <!-- 整个header自定义插槽 -->
      <template #header v-if="hasCustomHeaderSlot">
        <slot name="header" />
      </template>

      <!-- 左侧固定部分插槽 -->
      <template #fixedLeft v-if="hasFixedLeftSlot">
        <slot name="fixedLeft" />
      </template>

      <!-- 右侧固定部分插槽 -->
      <template #fixedRight v-if="hasFixedRightSlot">
        <slot name="fixedRight" />
      </template>

      <!-- 可滚动部分插槽 -->
      <template #scrollable v-if="hasScrollableSlot">
        <slot name="scrollable" />
      </template>
    </TabHeader>

    <!-- Tab内容区域 -->
    <div class="tab-content">
      <slot />
    </div>

    <!-- Header在底部 -->
    <TabHeader
      v-if="headerPosition === 'bottom' && showHeader"
      :tabs="tabs"
      :active-key="activeKey"
      :position="headerPosition"
      :show="showHeader"
      :custom-slot="hasCustomHeaderSlot"
      :fixed-left-slot="hasFixedLeftSlot"
      :fixed-right-slot="hasFixedRightSlot"
      :scrollable-slot="hasScrollableSlot"
      @tab-click="handleTabClick"
    >
      <!-- 整个header自定义插槽 -->
      <template #header v-if="hasCustomHeaderSlot">
        <slot name="header" />
      </template>

      <!-- 左侧固定部分插槽 -->
      <template #fixedLeft v-if="hasFixedLeftSlot">
        <slot name="fixedLeft" />
      </template>

      <!-- 右侧固定部分插槽 -->
      <template #fixedRight v-if="hasFixedRightSlot">
        <slot name="fixedRight" />
      </template>

      <!-- 可滚动部分插槽 -->
      <template #scrollable v-if="hasScrollableSlot">
        <slot name="scrollable" />
      </template>
    </TabHeader>
  </div>
</template>

<style lang="less" scoped>
.tab-container {
  display: flex;
  flex-direction: column;
  width: 100%;
  height: 100%;
  background: var(--theme-background);
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.08);

  &--bottom {
    flex-direction: column-reverse;
  }
}

.tab-content {
  flex: 1;
  overflow: hidden;
  position: relative;
  background: var(--theme-background);
}

</style>
