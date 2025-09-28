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
import { store } from '../store'
import { Icon } from '@iconify/vue'

export type ToastType = 'success' | 'error' | 'warning' | 'info'

export interface ToastItem {
  id: string
  message: string
  type: ToastType
}
</script>

<style scoped>
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
  background: rgba(0, 0, 0, 0.8);
  color: white;
  padding: 12px 20px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  gap: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(10px);
  min-width: 200px;
  max-width: 100%;
}

.toast-icon {
  background: #4ade80;
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

.toast-icon-success {
  background: #4ade80;
}

.toast-icon-error {
  background: #ef4444;
}

.toast-icon-warning {
  background: #f59e0b;
}

.toast-icon-info {
  background: #3b82f6;
}

.toast-message {
  font-size: 14px;
  font-weight: 500;
  word-wrap: break-word;
  overflow-wrap: break-word;
}

.toast-enter-active {
  transition: all 0.3s ease-out;
}

.toast-leave-active {
  transition: all 0.3s ease-in;
}

.toast-enter-from {
  opacity: 0;
  transform: translateX(100%);
}

.toast-leave-to {
  opacity: 0;
  transform: translateX(100%);
}

.toast-move {
  transition: transform 0.3s ease;
}
</style>
