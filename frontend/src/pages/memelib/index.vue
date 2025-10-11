<script setup lang="ts">
import { store } from '@/store';
import { ref } from 'vue'

const lists = ref([
    {
        title: "Emojiall",
        url: 'https://www.emojiall.com/',
        icon: 'favicon.ico'
    },
    {
        title: "Cool",
        url: 'https://bqb.cool/',
        icon: 'favicon.ico'
    },
    {
        title: "斗图吧",
        url: 'https://www.doutub.com/',
        icon: 'favicon.ico'
    },
]);

const copy = async (url: string) => {
    try {
        await navigator.clipboard.writeText(url);
        store.showToast('复制网址成功！')
    } catch (error) {
        store.showToast('复制网址失败！')
    }

}

</script>

<template>
    <div class="container">
        <div class="container-body">
            <h2 class="">表情库导航</h2>
            <hr>
            <div class="card-body">

                <div class="card-list" @click="() => copy(item.url)" v-for="(item, index) in lists" :key="index">
                    <div class="card-list-item">
                        <img :src="item.url + item.icon" alt="">
                        <div>名称：{{ item.title }}</div>
                        <div>网址：{{ item.url }}</div>
                    </div>
                </div>

            </div>
        </div>
    </div>
</template>

<style lang="less" scoped>
.container {
    background: linear-gradient(135deg, #6a11cb 0%, #2575fc 100%);
    height: 100%;
    display: flex;
    justify-content: center;
    padding: 20px;
}

.container-body {
    display: flex;
    flex-direction: column;
    justify-content: start;
    align-items: center;
    width: 100%;
    max-width: 800px;
    background: rgba(255, 255, 255, 0.1);
    backdrop-filter: blur(10px);
    border-radius: 15px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
    overflow: hidden;
    border: 1px solid rgba(255, 255, 255, 0.2);
    padding: 30px;
}

hr {
    margin: 20px 0;
    color: #fff;
    border: none;
    height: 1px;
    background: linear-gradient(to right, rgba(255, 255, 255, 0), rgba(255, 255, 255, 0.5), rgba(255, 255, 255, 0));
    width: 100%;
}

.components-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
}

.card-body {
    display: flex;
    flex-direction: row;
    justify-content: center;
    flex-wrap: wrap;
    gap: 20px;
}

.card-list {
    padding: 15px;
    display: flex;
    background: rgba(255, 255, 255, 0.15);
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
    transition: all 0.3s ease;
    border: 1px solid rgba(255, 255, 255, 0.2);
}

.card-list:hover {
    transform: translateY(-3px);
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.2);
}

.card-list-item {
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
    align-items: center;
    display: flex;
    gap: 10px;
    flex-direction: column;
}
</style>