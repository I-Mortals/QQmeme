import { reactive, watch } from 'vue'

export const store = reactive({
    // meme选项卡
    tabCurrent: '',
    handleTabClick(index: string) {
        this.tabCurrent = index
    },
    // meme显示的列数
    showColumn:6,
    // 当前加载的meme
    currentMemeData: [] as string[],
    memeDataList: [] as {key: string, data: string[]}[],
    clearCache(){
        this.memeDataList = []
    },
    // meme目录
    rootPath:'',
    rootMemeDirs: [] as string[],
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

watch(() => store.rootMemeDirs, (newValue, oldValue) => {
    if (window) {
        window.localStorage.setItem('rootMemeDirs', newValue.toString())
        // 当key为rootMemeDirs时，触发回调
        storeOn.forEach((cb, key) => {
            if (key === 'rootMemeDirs') {
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
        
        // 恢复 rootMemeDirs
        const cachedRootMemeDirs = window.localStorage.getItem('rootMemeDirs')
        if (cachedRootMemeDirs && cachedRootMemeDirs.length > 0) {
            store.rootMemeDirs = cachedRootMemeDirs.split(',')
        }
    }
}
