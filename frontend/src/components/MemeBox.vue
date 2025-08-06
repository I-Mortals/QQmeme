<script lang="ts" setup>
import { ref } from 'vue';
import { store } from '../store';
import { joinPath,joinShowImgPath } from '../utils';


const outColumns = () => {
  return Array(store.showColumn).fill('1fr').toString().replaceAll(',', ' ')
}

// 定义图片列表数据
const images = ref<string[]>([])

</script>

<template>
  <main class="MemeBox">
    <div class="meme-grid" :style="{gridTemplateColumns:outColumns()}">
      <div v-for="(image,id) in store.currentMemeData" :key="id" class="meme-item">
        <img loading="lazy" :src="joinShowImgPath(store.rootPath+`\\`+store.tabCurrent,image)" :alt="image">
      </div>
    </div>
  </main>
</template>

<style scoped>
.MemeBox {
  width: 100%;
  height: 100%;
  overflow-y: scroll;
}
.meme-grid {
  display: grid;
  /* grid-template-columns: ; */
  grid-column: 1fr 1fr;
  gap: 16px;
  padding: 16px;
  background-color: antiquewhite;
}

.meme-item {
  aspect-ratio: 1;
  overflow: hidden;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s ease;
  background-color: #f3f4f7;
}

.meme-item:hover {
  transform: scale(1.05);
}

.meme-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
</style>
