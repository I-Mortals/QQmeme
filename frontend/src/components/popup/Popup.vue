<script lang="ts" setup>
import { ref, computed, watch, nextTick, onMounted, onUnmounted } from 'vue'
import { useFloating, autoUpdate, offset, flip, shift, arrow, hide as hideMiddleware } from '@floating-ui/vue'
import type { PopupProps, PopupEmits, PopupExpose } from './types'

const props = withDefaults(defineProps<PopupProps>(), {
  placement: 'bottom',
  strategy: 'absolute',
  offset: 8,
  showArrow: false,
  clickOutside: true,
  duration: 200
})

const emit = defineEmits<PopupEmits>()

const popupRef = ref<HTMLElement>()
const arrowRef = ref<HTMLElement>()

const triggerElement = computed(() => {
  if (props.trigger) {
    return props.trigger
  }
  return popupRef.value?.parentElement
})

const middleware = computed(() => {
  const middlewares: any[] = [
    offset(props.offset),
    flip(),
    shift({ padding: 8 })
  ]

  if (props.showArrow) {
    middlewares.push(arrow({ element: arrowRef }))
  }

  middlewares.push(hideMiddleware())

  if (props.middleware) {
    middlewares.push(...props.middleware)
  }

  return middlewares
})

const { 
  floatingStyles, 
  placement, 
  middlewareData,
  isPositioned,
  update
} = useFloating(triggerElement, popupRef, {
  placement: props.placement,
  strategy: props.strategy,
  middleware: middleware,
  whileElementsMounted: autoUpdate,
  open: () => props.visible
})

const arrowStyles = computed(() => {
  if (!props.showArrow || !middlewareData.value.arrow) {
    return {}
  }

  const { x, y } = middlewareData.value.arrow
  const side = placement.value.split('-')[0]

  const staticSide = {
    top: 'bottom',
    right: 'left',
    bottom: 'top',
    left: 'right'
  }[side] as string

  return {
    position: 'absolute' as const,
    left: x != null ? `${x}px` : '',
    top: y != null ? `${y}px` : '',
    [staticSide]: '-4px',
    transform: 'rotate(45deg)',
    width: '8px',
    height: '8px',
    backgroundColor: 'inherit',
    border: 'inherit'
  }
})

const customStyle = computed(() => {
  const style: Record<string, any> = {}
  
  if (props.width) {
    style.width = typeof props.width === 'number' ? `${props.width}px` : props.width
  }
  
  if (props.height) {
    style.height = typeof props.height === 'number' ? `${props.height}px` : props.height
  }
  
  if (props.maxWidth) {
    style.maxWidth = typeof props.maxWidth === 'number' ? `${props.maxWidth}px` : props.maxWidth
  }
  
  if (props.maxHeight) {
    style.maxHeight = typeof props.maxHeight === 'number' ? `${props.maxHeight}px` : props.maxHeight
  }
  
  if (props.zIndex) {
    style.zIndex = props.zIndex
  }
  
  return style
})

const customClass = computed(() => {
  const classes: string[] = []
  
  if (props.className) {
    classes.push(props.className)
  }
  
  if (props.size) {
    classes.push(`popup-${props.size}`)
  }
  
  return classes
})

const handleClose = () => {
  emit('close')
  emit('update:visible', false)
}


const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Escape' && props.visible) {
    handleClose()
  }
}


onMounted(() => {
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
})


const show = () => {
  emit('update:visible', true)
}

const hide = () => {
  emit('update:visible', false)
}

const toggle = () => {
  emit('update:visible', !props.visible)
}

const updatePosition = () => {
  update()
}

defineExpose<PopupExpose>({
  show,
  hide,
  toggle,
  updatePosition
})
</script>

<template>
  <Teleport to="body">
    <Transition
      name="fade"
      appear
    >
      <div
        v-show="isPositioned"
        ref="popupRef"
        v-click-outside="props.clickOutside ? handleClose : undefined"
        class="popup-content"
        :class="[
          customClass,
          `popup-${placement}`,
          { 'popup-content-with-arrow': showArrow }
        ]"
        :style="{ ...floatingStyles, ...customStyle }"
        tabindex="-1"
      >
        
        <div
          v-if="showArrow"
          ref="arrowRef"
          class="popup-arrow"
          :style="arrowStyles"
        />
        
        <div class="popup-body">
          <slot />
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style lang="less">
@import '@/styles/variables.less';

.popup-content {
  background: @rgb-b1;
  border-radius: 0.25rem;
  box-shadow:
    0 8px 20px rgba(0, 0, 0, 0.12),
    0 2px 6px rgba(0, 0, 0, 0.08);
  border: 1px solid @rgb-b3;
  color: @rgb-bc;
  pointer-events: auto;
  max-width: 90vw;
  max-height: 90vh;
  overflow: hidden;
  position: relative;
  z-index: 1000;

  &-with-arrow {
    &::before {
      content: '';
      position: absolute;
      width: 8px;
      height: 8px;
      background: @rgb-b1;
      border: 1px solid @rgb-b3;
      transform: rotate(45deg);
      z-index: -1;
    }
  }
}

.popup-arrow {
  width: 8px;
  height: 8px;
  background: @rgb-b1;
  border: 1px solid @rgb-b3;
  transform: rotate(45deg);
  z-index: -1;
}

.popup-body {
  padding: 0.25rem;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease-out;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>