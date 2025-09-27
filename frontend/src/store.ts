import { reactive, watch } from 'vue'
import type { ToastItem, ToastType } from './components/Toast.vue'

type MemeInfoType = {
  name: string
  code: string
  parentPath: string
  icon: string
  memes: string[]
}

export const store = reactive({
  // meme选项卡
  tabCurrent: '',
  handleTabClick(index: string) {
    this.tabCurrent = index
  },
  // meme目录
  rootPath: '',
  allMemesPath: [] as MemeInfoType[],
  clearCache() {
    this.allMemesPath = []
  },
  // 主题色
  themeColor: '110,71,148', // 默认蓝色
  themeBackgroundColor: '255,255,255',
  setThemeColor(color: string) {
    this.themeColor = color
    // 更新CSS变量
    document.documentElement.style.setProperty('--theme-primary', color)
  },
  setThemeBackgroundColor(color: string) {
    this.themeBackgroundColor = color
    document.documentElement.style.setProperty('--theme-background', color)
  },
  // Toast弹窗
  toasts: [] as ToastItem[],
  showToast(message: string, type: ToastType = 'success') {
    const id = Date.now().toString() + Math.random().toString(36).substr(2, 9)
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

// TODO 监听store变化
const storeOn = new Map<string, (value: any) => void[]>()

watch(() => store.rootPath, (newValue) => {
  if (window) {
    window.localStorage.setItem('rootPath', newValue)
    storeOn.forEach((cb, key) => {
      if (key === 'rootPath') {
        cb(newValue)
      }
    })
  }
})

watch(() => store.allMemesPath, (newValue) => {
  if (window) {
    window.localStorage.setItem('allMemesPath', JSON.stringify(newValue))
    storeOn.forEach((cb, key) => {
      if (key === 'allMemesPath') {
        cb(newValue)
      }
    })
  }
})

watch(() => store.themeColor, (newValue) => {
  if (window) {
    window.localStorage.setItem('themeColor', newValue)
    storeOn.forEach((cb, key) => {
      if (key === 'themeColor') {
        cb(newValue)
      }
    })
  }
})

// init
export function initializeStoreFromCache() {
  if (window) {
    // 恢复 rootPath
    const cachedRootPath = window.localStorage.getItem('rootPath')
    if (cachedRootPath) {
      store.rootPath = cachedRootPath
    }
    
    // 恢复 allMemesPath
    const cachedAllMemesPath = window.localStorage.getItem('allMemesPath')
    if (cachedAllMemesPath && cachedAllMemesPath.length > 0) {
      store.allMemesPath = JSON.parse(cachedAllMemesPath)
      // 初始化tabCurrent
      if (store.allMemesPath.length > 0) {
        store.tabCurrent = store.allMemesPath[0].code
      }
    }
    
    // 恢复 themeColor
    const cachedThemeColor = window.localStorage.getItem('themeColor')
    if (cachedThemeColor) {
      store.setThemeColor(cachedThemeColor)
    } else {
      // 初始化默认主题色
      store.setThemeColor(store.themeColor)
    }
  }
}
