import { reactive } from 'vue'
import { ToastItem, ToastType } from '../components/Toast.vue'

export interface ToastStore {
  toasts: ToastItem[]
  showToast: (message: string, type?: ToastType) => void
}

export const toastStore = reactive<ToastStore>({
  toasts: [],
  
  showToast(message: string, type: ToastType = 'success') {
    const id = Math.random().toString(36).substring(2, 15)
    this.toasts.push({
      id,
      message,
      type
    })
    
    // 5秒后直接从数组中删除
    setTimeout(() => {
      const index = this.toasts.findIndex(toast => toast.id === id)
      if (index > -1) {
        this.toasts.splice(index, 1)
      }
    }, 5000)
  }
})
