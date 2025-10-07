<script lang="ts" setup>
import { onMounted, onUnmounted } from 'vue'
import { Icon } from '@iconify/vue'

interface ModalProps {
  visible: boolean
  title?: string
  closable?: boolean
  maskClosable?: boolean
  showCloseButton?: boolean
}

const props = withDefaults(defineProps<ModalProps>(), {
  title: '',
  closable: true,
  maskClosable: true,
  showCloseButton: true
})

const emit = defineEmits<{
  close: []
  'update:visible': [value: boolean]
}>()

const handleClose = () => {
  if (props.closable) {
    emit('close')
    emit('update:visible', false)
  }
}

const handleMaskClick = () => {
  if (props.maskClosable) {
    handleClose()
  }
}

const handleKeydown = (event: KeyboardEvent) => {
  if (event.key === 'Escape' && props.visible && props.closable) {
    handleClose()
  }
}

onMounted(() => {
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
})
</script>

<template>
  <Teleport to="body">
    <Transition name="modal">
      <div
        v-if="props.visible"
        class="modal-overlay"
        @click="handleMaskClick"
      >
        <div
          class="modal-container"
          @click.stop
        >
          <div v-if="title || showCloseButton" class="modal-header">
            <h2 v-if="title" class="modal-title">{{ title }}</h2>
            <button
              v-if="showCloseButton"
              class="close-btn"
              @click="handleClose"
              aria-label="关闭"
            >
              <Icon icon="lucide:x" :width="24" :height="24" />
            </button>
          </div>

          <div class="modal-body">
            <slot></slot>
          </div>

          <div v-if="$slots.footer" class="modal-footer">
            <slot name="footer"></slot>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style lang="less" scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.7);
  backdrop-filter: blur(12px);
  z-index: 1000;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 1rem;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);

  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: radial-gradient(circle at center,
      color-mix(in srgb, var(--theme-primary) 10%, transparent) 0%,
      transparent 70%);
    pointer-events: none;
  }
}

.modal-container {
  background: white;
  border-radius: 1rem;
  box-shadow:
    0 20px 40px rgba(0, 0, 0, 0.15),
    0 8px 16px rgba(0, 0, 0, 0.1),
    0 0 0 1px rgba(255, 255, 255, 0.1);
  width: 100%;
  max-width: 600px;
  max-height: none;
  overflow: hidden;
  position: relative;
  display: flex;
  flex-direction: column;
  border: 1px solid rgba(255, 255, 255, 0.2);

  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 1px;
    background: linear-gradient(90deg,
      transparent 0%,
      color-mix(in srgb, var(--theme-primary) 30%, transparent) 50%,
      transparent 100%);
  }

}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.875rem 1rem 0.625rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  background: var(--theme-primary);
  color: white;
  flex-shrink: 0;
  position: relative;

  &::after {
    content: '';
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    height: 1px;
    background: linear-gradient(90deg,
      transparent 0%,
      rgba(255, 255, 255, 0.3) 50%,
      transparent 100%);
  }
}

.modal-title {
  font-size: 1.375rem;
  font-weight: 700;
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  letter-spacing: -0.025em;
}

.close-btn {
  background: rgba(255, 255, 255, 0.15);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 0.5rem;
  width: 2.5rem;
  height: 2.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: white;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);

  &:hover {
    background: rgba(255, 255, 255, 0.25);
    transform: scale(1.1) rotate(90deg);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
    border-color: rgba(255, 255, 255, 0.4);
  }

  &:active {
    transform: scale(0.95) rotate(90deg);
    transition: all 0.1s ease;
  }
}

.modal-body {
  padding: 0.875rem;
  flex: 1;
  overflow-y: auto;
  background: #fafbfc;
}

.modal-footer {
  padding: 0.625rem 0.875rem 0.875rem;
  border-top: 1px solid color-mix(in srgb, var(--theme-primary) 10%, transparent);
  display: flex;
  justify-content: flex-end;
  gap: 0.625rem;
  background: #f8fafc;
  flex-shrink: 0;
  position: relative;

  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    height: 1px;
    background: linear-gradient(90deg,
      transparent 0%,
      color-mix(in srgb, var(--theme-primary) 20%, transparent) 50%,
      transparent 100%);
  }

  > * {
  }
}

.modal-enter-active,
.modal-leave-active {
  transition: all 0.4s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
  backdrop-filter: blur(0px);
}

.modal-enter-to,
.modal-leave-from {
  backdrop-filter: blur(12px);
}

</style>
