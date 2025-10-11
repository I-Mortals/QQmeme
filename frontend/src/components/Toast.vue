<template>
  <div class="toast-container">
    <TransitionGroup name="toast">
      <div
        v-for="toast in store.toasts"
        :key="toast.id"
        class="toast"
        :class="`toast-${toast.type}`"
      >
        <div class="toast-icon" :class="`toast-icon-${toast.type}`">
          <Icon icon="lucide:check" :width="12" :height="12" v-if="toast.type === 'success'" />
          <Icon icon="lucide:x" :width="12" :height="12" v-else-if="toast.type === 'error'" />
          <Icon icon="lucide:triangle-alert" :width="12" :height="12" v-else-if="toast.type === 'warning'" />
          <Icon icon="lucide:info" :width="12" :height="12" v-else-if="toast.type === 'info'" />
        </div>
        <div class="toast-message">{{ toast.message }}</div>
      </div>
    </TransitionGroup>
  </div>
</template>

<script setup lang="ts">
import { store } from '@/store'
import { Icon } from '@iconify/vue'

export type ToastType = 'success' | 'error' | 'warning' | 'info'

export interface ToastItem {
  id: string
  message: string
  type: ToastType
}
</script>

<style lang="less" scoped>
@import '@/styles/variables.less';

.toast-container {
  position: fixed;
  top: 20px;
  right: 20px;
  z-index: 9999;
  pointer-events: none;
  display: flex;
  flex-direction: column;
  gap: 10px;
  max-width: 300px;
}

.toast {
  background: @rgb-b1;
  color: @rgb-bc;
  padding: 12px 20px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  gap: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(10px);
  min-width: 200px;
  max-width: 100%;
  border: 1px solid @rgb-b3;

  &-success {
    .toast-icon {
      background: @rgb-s;
    }
  }

  &-error {
    .toast-icon {
      background: @rgb-e;
    }
  }

  &-warning {
    .toast-icon {
      background: @rgb-w;
    }
  }

  &-info {
    .toast-icon {
      background: @rgb-i;
    }
  }
}

.toast-icon {
  background: @rgb-s;
  color: white;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: bold;
  flex-shrink: 0;
}

.toast-message {
  font-size: 14px;
  font-weight: 500;
  word-wrap: break-word;
  overflow-wrap: break-word;
}

.toast-move,
.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s ease;
}

.toast-enter-from,
.toast-leave-to {
  opacity: 0;
  transform: translateX(100%);
}

.toast-leave-active {
  position: absolute;
}
</style>
