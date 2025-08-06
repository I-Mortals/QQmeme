<script lang="ts" setup>
import { defineProps, ref, withDefaults } from 'vue';
import { GetDirs, SelectRootDir } from '../../wailsjs/go/memeFile/memeFile'
import {store} from '../store'

interface SettingProps {
    visible?: boolean;
    closeCall?: () => void;
}

const props = withDefaults(defineProps<SettingProps>(), {
    title: '设置',
    visible: false,
    closeCall: () => { },
})
// 存入缓存
const rootPath = ref('')
const memeDirPathList = ref<string[]>([])

const selectRoot = async () => {
    const path = await SelectRootDir()
    if(path){
        rootPath.value = path
        store.rootPath = path
    }
    const memeDIrs = getMEMEDirList()
}
const getMEMEDirList = async () => {
    // 获取MEME主目录下的所有文件夹
    const paths = await GetDirs(rootPath.value)
    if (paths) {
        memeDirPathList.value = paths
        store.rootMemeDirs = paths
    }
}

const closeSetting = () => {
    props.closeCall()
}

const clearCache = () => {
    store.clearCache()
}
</script>

<template>
    <main class="Setting">
        <div class="Centent">
            <h1>设置</h1>
            <div>
                <button @click="selectRoot">选择MEME主目录</button>
                <p>MEME主目录:{{ store.rootPath }}</p>
            </div>
            <div>
                <input type="number" v-model="store.showColumn" />
                <p>显示列数:{{ store.showColumn }}</p>
            </div>
            <div>
                <p>清除图片内存缓存</p>
                <button @click="clearCache">清除</button>
            </div>
            <button @click="closeSetting">关闭</button>
        </div>
    </main>
</template>


<style scoped>
.Setting {
    width: 100%;
    height: 100%;
    background-color: #f5f5f590;
    position: absolute;
    z-index: 100;
    top: 0;
    left: 0;
}

.Centent {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    color: #333;
}
</style>