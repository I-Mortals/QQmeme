<script setup lang="ts">
import { store } from '@/store'
import { ref } from 'vue'

const lists = ref([
  {
    title: 'Emojiall',
    url: 'https://www.emojiall.com/',
    icon: 'favicon.ico'
  },
  {
    title: '斗图吧',
    url: 'https://www.doutub.com/',
    icon: 'favicon.ico'
  },
  {
    title: '发表情',
    url: 'https://fabiaoqing.com/',
    icon: 'Public/img/logonew.png'
  },
  {
    title: 'petpet',
    url: 'https://benisland.neocities.org/petpet/',
    icon: 'favicon.ico'
  },
  {
    title: 'Emojiall',
    url: 'https://www.emojiall.com/',
    icon: 'favicon.ico'
  }
])

const copy = async (url: string) => {
  try {
    await navigator.clipboard.writeText(url)
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
@import '@/styles/variables.less';

.container {
  background: linear-gradient(135deg, @rgb-p 0%, @rgb-i 100%);
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
  background: rgba(@b1, 0.1);
  backdrop-filter: blur(10px);
  border-radius: 15px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
  overflow: hidden;
  border: 1px solid rgba(@b1, 0.2);
  padding: 30px;
}

h2 {
  color: @rgb-b1;
  font-size: 1.5rem;
  font-weight: 600;
  margin: 0;
  text-align: center;
}

hr {
  margin: 20px 0;
  color: @rgb-b1;
  border: none;
  height: 1px;
  background: linear-gradient(to right, rgba(@b1, 0), rgba(@b1, 0.5), rgba(@b1, 0));
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
  background: rgba(@b1, 0.15);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  transition: all 0.3s ease;
  border: 1px solid rgba(@b1, 0.2);
  cursor: pointer;
  min-width: 200px;
}

.card-list:hover {
  transform: translateY(-3px);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.2);
  background: rgba(@b1, 0.2);
}

.card-list-item {
  border-bottom: 1px solid rgba(@b1, 0.1);
  align-items: center;
  display: flex;
  gap: 10px;
  flex-direction: column;
  color: @rgb-b1;
  text-align: center;

  img {
    width: 32px;
    height: 32px;
    border-radius: 4px;
    object-fit: contain;
  }

  div {
    font-size: 0.875rem;
    line-height: 1.4;

    &:first-of-type {
      font-weight: 600;
      color: @rgb-b1;
    }

    &:last-of-type {
      color: rgba(@b1, 0.8);
      font-size: 0.75rem;
    }
  }
}
</style>
