<script lang="ts" setup>
import { defineProps, defineEmits, onMounted, onUnmounted } from 'vue'
import { IconClose } from './icons'

interface ModalProps {
  visible: boolean
  title?: string
  size?: 'small' | 'medium' | 'large' | 'full'
  closable?: boolean
  maskClosable?: boolean
  showCloseButton?: boolean
}

const props = withDefaults(defineProps<ModalProps>(), {
  title: '',
  size: 'medium',
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
          :class="[`modal-${size}`]"
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
              <IconClose />
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

<style scoped>
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
  padding: 1.5rem;
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

.modal-overlay::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: radial-gradient(circle at center,
    rgba(var(--theme-primary), 0.1) 0%,
    transparent 70%);
  pointer-events: none;
}

.modal-container {
  background: white;
  border-radius: 1.5rem;
  box-shadow:
    0 32px 64px rgba(0, 0, 0, 0.2),
    0 16px 32px rgba(0, 0, 0, 0.1),
    0 0 0 1px rgba(255, 255, 255, 0.1);
  width: 100%;
  max-height: 90vh;
  overflow: hidden;
  position: relative;
  display: flex;
  flex-direction: column;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.modal-container::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg,
    transparent 0%,
    rgba(var(--theme-primary), 0.3) 50%,
    transparent 100%);
}

.modal-small {
  max-width: 420px;
}

.modal-medium {
  max-width: 520px;
}

.modal-large {
  max-width: 720px;
}

.modal-full {
  max-width: 95vw;
  max-height: 95vh;
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 1.5rem 1.5rem 1.25rem;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  background: rgb(var(--theme-primary));
  color: white;
  flex-shrink: 0;
  position: relative;
}

.modal-header::after {
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

.modal-title {
  font-size: 1.5rem;
  font-weight: 700;
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  letter-spacing: -0.025em;
}

.close-btn {
  background: rgba(255, 255, 255, 0.15);
  border: 1px solid rgba(255, 255, 255, 0.2);
  border-radius: 0.625rem;
  width: 2.75rem;
  height: 2.75rem;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: white;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.25);
  transform: scale(1.1) rotate(90deg);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
  border-color: rgba(255, 255, 255, 0.4);
}

.close-btn:active {
  transform: scale(0.95) rotate(90deg);
  transition: all 0.1s ease;
}

.modal-body {
  padding: 1.5rem;
  flex: 1;
  overflow-y: auto;
  background: #fafbfc;
}


.modal-footer {
  padding: 1.25rem 1.5rem 1.5rem;
  border-top: 1px solid rgba(var(--theme-primary), 0.1);
  display: flex;
  justify-content: flex-end;
  gap: 0.75rem;
  background: #f8fafc;
  flex-shrink: 0;
  position: relative;
}

.modal-footer::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg,
    transparent 0%,
    rgba(var(--theme-primary), 0.2) 50%,
    transparent 100%);
}

/* Animations */
.modal-enter-active {
  transition: all 0.4s cubic-bezier(0.16, 1, 0.3, 1);
}

.modal-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 1, 1);
}

.modal-enter-from {
  opacity: 0;
  backdrop-filter: blur(0px);
}

.modal-leave-to {
  opacity: 0;
  backdrop-filter: blur(0px);
}

.modal-enter-from .modal-container {
  transform: translateY(60px) scale(0.8) rotateX(10deg);
  opacity: 0;
  filter: blur(8px);
}

.modal-leave-to .modal-container {
  transform: translateY(40px) scale(0.85) rotateX(5deg);
  opacity: 0;
  filter: blur(4px);
}

.modal-enter-to .modal-container,
.modal-leave-from .modal-container {
  transform: translateY(0) scale(1) rotateX(0deg);
  opacity: 1;
  filter: blur(0px);
}

.modal-enter-to {
  backdrop-filter: blur(12px);
}

.modal-leave-from {
  backdrop-filter: blur(12px);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .modal-overlay {
    padding: 1rem;
  }

  .modal-container {
    border-radius: 1.25rem;
  }

  .modal-header {
    padding: 1.5rem 1.5rem 1.25rem;
  }

  .modal-title {
    font-size: 1.5rem;
  }

  .close-btn {
    width: 2.75rem;
    height: 2.75rem;
  }

  .modal-body {
    padding: 1.5rem;
  }

  .modal-footer {
    padding: 1.25rem 1.5rem 1.5rem;
  }
}

@media (max-width: 640px) {
  .modal-overlay {
    padding: 0.75rem;
  }

  .modal-container {
    margin: 0;
    max-width: none;
    border-radius: 1rem;
  }

  .modal-header {
    padding: 1.25rem 1.25rem 1rem;
  }

  .modal-title {
    font-size: 1.375rem;
  }

  .close-btn {
    width: 2.5rem;
    height: 2.5rem;
  }

  .modal-body {
    padding: 1.25rem;
  }

  .modal-footer {
    padding: 1rem 1.25rem 1.25rem;
    flex-direction: column;
  }

  .modal-footer > * {
    width: 100%;
  }
}

@media (max-width: 480px) {
  .modal-overlay {
    padding: 0.5rem;
  }

  .modal-container {
    border-radius: 0.75rem;
  }

  .modal-header {
    padding: 1rem;
  }

  .modal-title {
    font-size: 1.25rem;
  }

  .modal-body {
    padding: 1rem;
  }

  .modal-footer {
    padding: 1rem;
  }
}
</style>
