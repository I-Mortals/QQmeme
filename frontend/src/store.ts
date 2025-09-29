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
  themeColor: '110,71,148,1', // 默认紫色，因为紫色更有韵味，格式为 r,g,b,a
  themeBackgroundColor: '255,255,255,1',
  setThemeColor(color: string) {
    this.themeColor = color
    // 更新CSS变量，支持带透明度的颜色
    if (color.includes(',')) {
      const [r, g, b, a = '1'] = color.split(',')
      document.documentElement.style.setProperty('--theme-primary', `rgba(${r},${g},${b},${a})`)
    } else {
      document.documentElement.style.setProperty('--theme-primary', color)
    }
  },
  setThemeBackgroundColor(color: string) {
    this.themeBackgroundColor = color
    // 更新CSS变量，支持带透明度的颜色
    if (color.includes(',')) {
      const [r, g, b, a = '1'] = color.split(',')
      document.documentElement.style.setProperty('--theme-background', `rgba(${r},${g},${b},${a})`)
    } else {
      document.documentElement.style.setProperty('--theme-background', color)
    }
  },
  // Toast弹窗
  toasts: [] as ToastItem[],
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


watch(() => store.rootPath, (newValue) => {
  if (window) {    
    window.localStorage.setItem('rootPath', newValue)
  }
})

watch(() => store.allMemesPath, (newValue) => {
  if (window) {
    window.localStorage.setItem('allMemesPath', JSON.stringify(newValue))
  }
},{deep: true})

watch(() => store.themeColor, (newValue) => {
  if (window) {
    window.localStorage.setItem('themeColor', newValue)
  }
})

watch(() => store.themeBackgroundColor, (newValue) => {
  if (window) {
    window.localStorage.setItem('themeBackgroundColor', newValue)
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
    
    // 恢复 themeBackgroundColor
    const cachedThemeBackgroundColor = window.localStorage.getItem('themeBackgroundColor')
    if (cachedThemeBackgroundColor) {
      store.setThemeBackgroundColor(cachedThemeBackgroundColor)
    } else {
      // 恢复默认主题色
      store.setThemeBackgroundColor(store.themeBackgroundColor)
    }
  }
}
