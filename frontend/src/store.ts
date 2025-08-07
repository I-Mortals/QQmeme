import { reactive, watch } from 'vue'

type MemeInfoType = {
  Name: string
  Code: string
  ParentPath: string
  Icon: string
  Memes: string[]
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
  }
})

// TODO 监听store变化
const storeOn = new Map<string, (value: any) => void[]>()

watch(() => store.rootPath, (newValue, oldValue) => {
  if (window) {
    window.localStorage.setItem('rootPath', newValue)
    storeOn.forEach((cb, key) => {
      if (key === 'rootPath') {
        cb(newValue)
      }
    })
  }
})

watch(() => store.allMemesPath, (newValue, oldValue) => {
  if (window) {
    window.localStorage.setItem('allMemesPath', JSON.stringify(newValue))
    storeOn.forEach((cb, key) => {
      if (key === 'allMemesPath') {
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
        store.tabCurrent = store.allMemesPath[0].Code
      }
    }
  }
}
