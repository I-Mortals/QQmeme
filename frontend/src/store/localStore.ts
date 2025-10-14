import { reactive, watch } from 'vue'
import { memeStore } from './memeStore'
import { themeStore } from './themeStore'
import { ALL_MEMES_PATH_KEY, ROOT_PATH_KEY, STAR_MEMES_KEY, BOT_TOKEN_KEY, PROXY_ENABLED_KEY, PROXY_URL_KEY } from '@/utils/common'
import { applicationStore } from './applicationStore'

export interface LocalStore {}

export const localStore = reactive<LocalStore>({})

// 监听数据变化并保存到localStorage
watch(() => memeStore.rootPath, (newValue) => {
  if (window) {
    window.localStorage.setItem(ROOT_PATH_KEY, newValue)
  }
})

watch(() => memeStore.allMemesPath, (newValue) => {
  if (window) {
    window.localStorage.setItem(ALL_MEMES_PATH_KEY, JSON.stringify(newValue))
  }
}, { deep: true })

watch(() => themeStore.currentTheme, (newValue) => {
  if (window) {
    window.localStorage.setItem('meme-theme', newValue)
  }
})

watch(() => memeStore.starMemes, (newValue) => {
  if (window) {
    window.localStorage.setItem(STAR_MEMES_KEY, JSON.stringify(newValue))
  }
}, { deep: true })

// 应用配置持久化（迁回 localStore）
watch(() => applicationStore.botToken, (newValue) => {
  if (window) {
    window.localStorage.setItem(BOT_TOKEN_KEY, newValue)
  }
})

watch(() => applicationStore.proxyEnabled, (newValue) => {
  if (window) {
    window.localStorage.setItem(PROXY_ENABLED_KEY, newValue.toString())
  }
})

watch(() => applicationStore.proxyURL, (newValue) => {
  if (window) {
    window.localStorage.setItem(PROXY_URL_KEY, newValue)
  }
})

// 从缓存初始化数据
export function initializeStoreFromCache() {
  if (window) {
    const cachedRootPath = window.localStorage.getItem(ROOT_PATH_KEY)
    if (cachedRootPath) {
      memeStore.rootPath = cachedRootPath
    }
    
    const cachedAllMemesPath = window.localStorage.getItem(ALL_MEMES_PATH_KEY)
    if (cachedAllMemesPath && cachedAllMemesPath.length > 0) {
      memeStore.allMemesPath = JSON.parse(cachedAllMemesPath)
      memeStore.tabCurrent = STAR_MEMES_KEY
    }
    
    const cachedTheme = window.localStorage.getItem('meme-theme')
    if (cachedTheme && themeStore.availableThemes.includes(cachedTheme)) {
      themeStore.setTheme(cachedTheme)
    } else {
      themeStore.setTheme(themeStore.currentTheme)
    }
    
    const cachedStarMemes = window.localStorage.getItem(STAR_MEMES_KEY)
    if (cachedStarMemes && cachedStarMemes.length > 0) {
      memeStore.starMemes = JSON.parse(cachedStarMemes)
    }
    
    // 恢复应用网络配置
    const cachedBotToken = window.localStorage.getItem(BOT_TOKEN_KEY)
    if (cachedBotToken) {
      applicationStore.botToken = cachedBotToken
    }

    const cachedProxyEnabled = window.localStorage.getItem(PROXY_ENABLED_KEY)
    if (cachedProxyEnabled) {
      applicationStore.proxyEnabled = cachedProxyEnabled === 'true'
    }

    const cachedProxyURL = window.localStorage.getItem(PROXY_URL_KEY)
    if (cachedProxyURL) {
      applicationStore.proxyURL = cachedProxyURL
    }
  }
}
