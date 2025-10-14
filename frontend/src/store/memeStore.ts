import { reactive } from 'vue'
import { GenerateAllMemePath } from '@wailsjs/go/memeFile/MemeFile'

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

export interface MemeStore {
  // meme选项卡
  tabCurrent: string
  handleTabClick: (index: string) => void
  
  // 强制刷新
  forceRefreshKey: number
  forceRefreshCurrentTab: () => void
  
  // meme目录
  rootPath: string
  allMemesPath: MemeTabItem[]
  tabOrderChanged: boolean
  
  // 收藏夹
  starMemes: StarMemeItem[]
  
  // 缓存管理
  clearCache: () => void
  refreshMemes: () => Promise<void>
  
  // 表情包顺序管理
  setMemeOrderChanged: (memeCode: string, changed?: boolean) => void
  setTabOrderChanged: (changed?: boolean) => void
  
  // 收藏夹管理
  addToStarMemes: (fileName: string, fromFolder: string) => void
  removeFromStarMemes: (id: string) => void
}

export const memeStore = reactive<MemeStore>({
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
  allMemesPath: [],
  tabOrderChanged: false,
  
  // 收藏夹
  starMemes: [],
  
  // 缓存管理
  clearCache() {
    this.allMemesPath = []
    this.starMemes = []
  },
  
  async refreshMemes() {
    if (this.rootPath) {
      try {
        this.allMemesPath = await GenerateAllMemePath(this.rootPath)
      } catch (error) {
        console.error('刷新失败:', error)
      }
    }
  },
  
  // 表情包顺序管理
  setMemeOrderChanged(memeCode: string, changed: boolean = true) {
    const meme = this.allMemesPath.find(m => m.code === memeCode)
    if (meme) {
      meme.orderChanged = changed
    }
  },
  
  setTabOrderChanged(changed: boolean = true) {
    this.tabOrderChanged = changed
  },
  
  // 收藏夹管理
  addToStarMemes(fileName: string, fromFolder: string) {
    const id = Math.random().toString(36).substring(2, 15)

    const starMeme = this.starMemes.find(item => item.fileName === fileName)
    if (starMeme) {
      return // 不显示toast，由调用方处理
    }

    const starMemeItem: StarMemeItem = {
      id,
      fileName,
      fromFolder,
      addedAt: Date.now()
    }
    
    this.starMemes.unshift(starMemeItem)
  },
  
  removeFromStarMemes(id: string) {
    const index = this.starMemes.findIndex(item => item.id === id)
    if (index > -1) {
      this.starMemes.splice(index, 1)
    }
  }
})
