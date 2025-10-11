<script lang="ts" setup>
import { ref, watch, onMounted } from 'vue'
import type { TabPanelProps } from './types'

const props = withDefaults(defineProps<TabPanelProps>(), {
  active: false,
  lazy: true,
  loaded: false
})

const emit = defineEmits<{
  'load': []
}>()

const shouldShow = ref(false)
const hasLoaded = ref(props.loaded)

watch(() => props.active, (newActive) => {
  if (newActive && props.lazy && !hasLoaded.value) {
    shouldShow.value = true
    hasLoaded.value = true
    emit('load')
  } else if (!props.lazy) {
    shouldShow.value = newActive
  } else if (hasLoaded.value) {
    shouldShow.value = newActive
  }
}, { immediate: true })

onMounted(() => {
  if (!props.lazy) {
    shouldShow.value = props.active
  }
})
</script>

<template>
  <div
    v-show="shouldShow"
    :class="[
      'tab-panel',
      {
        'tab-panel-active': props.active,
        'tab-panel-loaded': hasLoaded
      },
      props.class
    ]"
  >
    <slot />
  </div>
</template>

<style lang="less" scoped>
.tab-panel {
  width: 100%;
  height: 100%;
  overflow: auto;
  position: relative;
  animation: fadeIn 0.3s ease-in-out;

  &.tab-panel-active {
    display: block;
  }

  &.tab-panel-loaded {
    opacity: 1;
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
