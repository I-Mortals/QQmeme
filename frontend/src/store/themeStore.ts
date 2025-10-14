import { reactive } from 'vue'
import { presetThemes, type ThemeColorOptions } from '@/styles/themes'

export interface ThemeStore {
  currentTheme: string
  availableThemes: string[]
  setTheme: (theme: string) => void
  getThemeConfig: (theme?: string) => ThemeColorOptions
}

export const themeStore = reactive<ThemeStore>({
  currentTheme: 'dark',
  availableThemes: Object.keys(presetThemes),
  
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
  }
})
