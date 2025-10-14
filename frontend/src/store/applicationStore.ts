import { reactive } from 'vue'

export interface ApplicationStore {
  // 应用配置
  botToken: string
  proxyEnabled: boolean
  proxyURL: string

  // 配置设置方法
  setBotToken: (token: string) => void
  setProxySettings: (enabled: boolean, url: string) => void
}

export const applicationStore = reactive<ApplicationStore>({
  // 应用配置
  botToken: '',
  proxyEnabled: false,
  proxyURL: 'http://127.0.0.1:7890',

  // 配置设置方法
  setBotToken(token: string) {
    this.botToken = token
  },

  setProxySettings(enabled: boolean, url: string) {
    this.proxyEnabled = enabled
    this.proxyURL = url
  }
})

