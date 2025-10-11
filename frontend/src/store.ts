import { reactive, watch } from 'vue'
import { GenerateAllMemePath } from '../wailsjs/go/memeFile/MemeFile'
import { ToastItem, ToastType } from './components/Toast.vue'
import { presetThemes, type ThemeColorOptions } from './styles/themes'
import {
  ALL_MEMES_PATH_KEY,
  BOT_TOKEN_KEY,
  PROXY_ENABLED_KEY,
  PROXY_URL_KEY,
  ROOT_PATH_KEY,
  STAR_MEMES_KEY
} from './utils/common'

interface ContextMenuItem {
  icon: string
  label: string
  action: () => void
  separator?: boolean
}

interface ContextMenuState {
  visible: boolean
  menuItems: ContextMenuItem[]
  position: { x: number; y: number }
}

export type MemeTabItem = {
  name: string
  code: string
  parentPath: string
  icon: string
  memes: string[]
  orderChanged?: boolean
}

export type StarMemeItem = {
  id: string
  fileName: string
  fromFolder: string
  addedAt: number
}

export const store = reactive({
  // meme选项卡
  tabCurrent: '',
  handleTabClick(index: string) {
    this.tabCurrent = index
  },
  // 强制刷新
  forceRefreshKey: 0,
  forceRefreshCurrentTab() {
    this.forceRefreshKey++
  },
  // meme目录
  rootPath: '',
  allMemesPath: [] as MemeTabItem[],
  tabOrderChanged: false,
  clearCache() {
    this.allMemesPath = []
    this.starMemes = []
  },
  setMemeOrderChanged(memeCode: string, changed: boolean = true) {
    const meme = this.allMemesPath.find(m => m.code === memeCode)
    if (meme) {
      meme.orderChanged = changed
    }
  },
  setTabOrderChanged(changed: boolean = true) {
    this.tabOrderChanged = changed
  },
  // 收藏夹
  starMemes: [] as StarMemeItem[],
  addToStarMemes(fileName: string, fromFolder: string) {
    const id = Math.random().toString(36).substring(2, 15)

    const starMeme = this.starMemes.find(item => item.fileName === fileName)
    if (starMeme) {
      this.showToast('已收藏！')
      return
    }

    const starMemeItem: StarMemeItem = {
      id,
      fileName,
      fromFolder,
      addedAt: Date.now()
    }
    
    this.starMemes.unshift(starMemeItem)
    this.showToast('已收藏！')
  },
  removeFromStarMemes(id: string) {
    const index = this.starMemes.findIndex(item => item.id === id)
    if (index > -1) {
      this.starMemes.splice(index, 1)
      this.showToast('已取消收藏！')
    }
  },
  // 主题系统
  currentTheme: 'dark', // 当前主题
  availableThemes: Object.keys(presetThemes),
  // 应用配置
  botToken: '',
  proxyEnabled: false,
  proxyURL: 'http://127.0.0.1:7890',
  setTheme(theme: string) {
    if (!this.availableThemes.includes(theme)) {
      console.warn(`Theme "${theme}" is not available`)
      return
    }
    
    this.currentTheme = theme
    document.documentElement.dataset.theme = theme
  },
  getThemeConfig(theme?: string): ThemeColorOptions {
    const targetTheme = theme || this.currentTheme
    return presetThemes[targetTheme] || presetThemes.dark
  },
  // 配置设置方法
  setBotToken(token: string) {
    this.botToken = token
  },
  setProxySettings(enabled: boolean, url: string) {
    this.proxyEnabled = enabled
    this.proxyURL = url
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
  },
  // 全局右键菜单
  contextMenu: {
    visible: false,
    menuItems: [],
    position: { x: 0, y: 0 }
  } as ContextMenuState,
  showContextMenu(e: MouseEvent, menuItems: ContextMenuItem[]) {
    e.preventDefault()
    this.contextMenu.menuItems = menuItems
    this.contextMenu.position = { x: e.clientX, y: e.clientY }
    this.contextMenu.visible = true
  },
  hideContextMenu() {
    this.contextMenu.visible = false
    setTimeout(() => {
      this.contextMenu.menuItems = []
    }, 200)
  },
  
  async refreshMemes() {
    if (this.rootPath) {
      try {
        this.allMemesPath = await GenerateAllMemePath(this.rootPath)
      } catch (error) {
        console.error('刷新失败:', error)
      }
    } else {
      this.showToast('请先选择主目录', 'error')
    }
  }
  
})

watch(() => store.rootPath, (newValue) => {
  if (window) {
    window.localStorage.setItem(ROOT_PATH_KEY, newValue)
  }
})

watch(() => store.allMemesPath, (newValue) => {
  if (window) {
    window.localStorage.setItem(ALL_MEMES_PATH_KEY, JSON.stringify(newValue))
  }
}, { deep: true })

watch(() => store.currentTheme, (newValue) => {
  if (window) {
    window.localStorage.setItem('meme-theme', newValue)
  }
})

watch(() => store.starMemes, (newValue) => {
  if (window) {
    window.localStorage.setItem(STAR_MEMES_KEY, JSON.stringify(newValue))
  }
}, { deep: true })

watch(() => store.botToken, (newValue) => {
  if (window) {
    window.localStorage.setItem(BOT_TOKEN_KEY, newValue)
  }
})

watch(() => store.proxyEnabled, (newValue) => {
  if (window) {
    window.localStorage.setItem(PROXY_ENABLED_KEY, newValue.toString())
  }
})

watch(() => store.proxyURL, (newValue) => {
  if (window) {
    window.localStorage.setItem(PROXY_URL_KEY, newValue)
  }
})

export function initializeStoreFromCache() {
  if (window) {
    const cachedRootPath = window.localStorage.getItem(ROOT_PATH_KEY)
    if (cachedRootPath) {
      store.rootPath = cachedRootPath
    }
    
    const cachedAllMemesPath = window.localStorage.getItem(ALL_MEMES_PATH_KEY)
    if (cachedAllMemesPath && cachedAllMemesPath.length > 0) {
      store.allMemesPath = JSON.parse(cachedAllMemesPath)
      store.tabCurrent = STAR_MEMES_KEY
    }
    
    const cachedTheme = window.localStorage.getItem('meme-theme')
    if (cachedTheme && store.availableThemes.includes(cachedTheme)) {
      store.setTheme(cachedTheme)
    } else {
      store.setTheme(store.currentTheme)
    }
    
    const cachedStarMemes = window.localStorage.getItem(STAR_MEMES_KEY)
    if (cachedStarMemes && cachedStarMemes.length > 0) {
      store.starMemes = JSON.parse(cachedStarMemes)
    }
    
    // 加载配置
    const cachedBotToken = window.localStorage.getItem(BOT_TOKEN_KEY)
    if (cachedBotToken) {
      store.botToken = cachedBotToken
    }
    
    const cachedProxyEnabled = window.localStorage.getItem(PROXY_ENABLED_KEY)
    if (cachedProxyEnabled) {
      store.proxyEnabled = cachedProxyEnabled === 'true'
    }
    
    const cachedProxyURL = window.localStorage.getItem(PROXY_URL_KEY)
    if (cachedProxyURL) {
      store.proxyURL = cachedProxyURL
    }
  }
}
