<script lang="ts" setup>
import { ref } from 'vue';
import { store  } from '../store';
import MeSetting from './Setting.vue';
import { GetImages, Hello } from '../../wailsjs/go/memeFile/memeFile'
import { joinPath } from '../utils';

const isShowSetting = ref(false)

const handleTabClick = (tab: string) => {
    store.handleTabClick(tab)
    selectMemeData(tab)
}

// setting
const openSetting = async () => {
    isShowSetting.value = true
    const h = await Hello()
    console.log(h)
}
const closeSetting = () => {
    isShowSetting.value = false
}


// 获取当前选中的meme数据集
const selectMemeData = async (index: string) => {
    const memePaths = joinPath(store.rootPath, index)
    const existingItem = store.memeDataList.find((v) => index == v.key)
    console.log('缓存meme数据集', store.memeDataList);
    
    if (existingItem) {
        // 如果存在，使用内存数据更新
        store.currentMemeData = existingItem.data
    } else {
        const memes = await GetImages(memePaths)
        if (memes) {
            store.currentMemeData = memes
            store.memeDataList.push({
                key: index,
                data: memes
            })
        }
        console.log(memes)
    }
}

</script>


<template>
    <main class="BottomBar">
        <div class="content">
            <button style="color: black;" @click="openSetting()">设置</button>
            <div style="display: flex; gap: 5px;height: 100%;">
                <div v-for="(image,key) in store.rootMemeDirs" :key="key" @click="handleTabClick(image)"
                    :class="'meme-item' + ' ' + (image == store.tabCurrent ? 'meme-item-action' : '')">
                    <!-- <img :src="image.src" :alt="image.alt" /> -->
                    <p>{{ image }}</p>
                </div>
            </div>
        </div>
        <MeSetting v-if="isShowSetting" :closeCall="closeSetting" />

    </main>


</template>

<style scoped>
.BottomBar {
    /* 固定底部 */
    /* position: fixed; */
    bottom: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: #f3f4f7;
    z-index: 1000;
}

.content {
    height: 100%;
    display: flex;
    flex-wrap: nowrap;
    align-items: center;
    gap: 5px;
    margin-left: 10px;
}

.meme-item {
    width: 60px;
    height: 100%;

    flex: 0 0 auto;
    overflow-y: hidden;
    overflow-x: hidden;
    background-color: cadetblue;
}

.meme-item-action {
    background-color: aqua;
}

.meme-item img {
    width: 100%;
    height: 100%;
    /* 禁止拖拽 */
    user-drag: none;
    -webkit-user-drag: none;
    user-select: none;
    -moz-user-select: none;
    -webkit-user-select: none;
    -ms-user-select: none;
    /* 隐藏alt提示 */
    font-size: 0;
    color: transparent;

}
</style>
