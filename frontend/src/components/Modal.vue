<script lang="ts" setup>
import { onMounted, onUnmounted } from 'vue'
import { Icon } from '@iconify/vue'
import Button from './Button.vue'

interface ModalProps {
  visible: boolean
  title?: string
  closable?: boolean
  maskClosable?: boolean
  showCloseButton?: boolean
  cancelBtn?: {
    text?: string
    disabled?: boolean
    loading?: boolean
  }
  confirmBtn?: {
    text?: string
    disabled?: boolean
    loading?: boolean
  }
  showCancel?: boolean
  showConfirm?: boolean
}

const props = withDefaults(defineProps<ModalProps>(), {
  title: '',
  closable: true,
  maskClosable: true,
  showCloseButton: true,
  cancelBtn: () => ({
    text: '取消',
    disabled: false,
    loading: false
  }),
  confirmBtn: () => ({
    text: '确定',
    disabled: false,
    loading: false
  }),
  showCancel: true,
  showConfirm: true
})

const emit = defineEmits<{
  close: []
  'update:visible': [value: boolean]
  cancel: []
  confirm: []
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

const handleCancel = () => {
  emit('cancel')
}

const handleConfirm = () => {
  emit('confirm')
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
            <div v-if="title" class="modal-title">{{ title }}</div>
            <button
              v-if="showCloseButton"
              class="close-btn"
              @click="handleClose"
              aria-label="关闭"
            >
              <Icon icon="lucide:x" :width="24" :height="24"/>
            </button>
          </div>

          <div class="modal-body">
            <slot></slot>
          </div>

          <div v-if="$slots.footer || showCancel || showConfirm" class="modal-footer">
            <slot name="footer">
              <div v-if="showCancel || showConfirm" class="modal-actions">
                <Button
                  v-if="showCancel"
                  variant="secondary"
                  :disabled="cancelBtn?.disabled || cancelBtn?.loading"
                  :loading="cancelBtn?.loading"
                  @click="handleCancel"
                >
                  {{ cancelBtn?.text }}
                </Button>
                <Button
                  v-if="showConfirm"
                  variant="primary"
                  :loading="confirmBtn?.loading"
                  :disabled="confirmBtn?.disabled || confirmBtn?.loading"
                  @click="handleConfirm"
                >
                  {{ confirmBtn?.text }}
                </Button>
              </div>
            </slot>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

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
}

.modal-container {
  background: @rgb-b1;
  border-radius: 1rem;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15),
  0 8px 16px rgba(0, 0, 0, 0.1),
  0 0 0 1px rgba(255, 255, 255, 0.1);
  width: 100%;
  max-width: 600px;
  max-height: none;
  overflow: hidden;
  position: relative;
  display: flex;
  flex-direction: column;
  border: 1px solid @rgb-b3;
}

.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.75rem 1rem 0.5rem;
  border-bottom: 1px solid @rgb-b3;
  background: @rgb-p;
  color: @rgb-pc;
  flex-shrink: 0;
  position: relative;
}

.modal-title {
  font-size: 1.25rem;
  font-weight: bold;
}

.close-btn {
  background: rgba(@pc, 0.15);
  border: 1px solid rgba(@pc, 0.2);
  border-radius: 0.5rem;
  width: 2.5rem;
  height: 2.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  color: @rgb-pc;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);

  &:hover {
    background: rgba(@pc, 0.25);
    transform: scale(1.1) rotate(90deg);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.12), 0 2px 4px rgba(0, 0, 0, 0.08);
    border-color: rgba(@pc, 0.4);
  }

  &:active {
    transform: scale(0.95) rotate(90deg);
  }
}

.modal-body {
  padding: 0.75rem;
  flex: 1;
  overflow-y: auto;
  background: @rgb-b1;
  color: @rgb-bc;
}

.modal-footer {
  padding: 0.5rem 0.875rem 0.75rem;
  border-top: 1px solid @rgb-b3;
  display: flex;
  justify-content: flex-end;
  gap: 0.625rem;
  background: @rgb-b1;
  flex-shrink: 0;
  position: relative;
}

.modal-actions {
  display: flex;
  gap: 0.75rem;
  justify-content: flex-end;
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
