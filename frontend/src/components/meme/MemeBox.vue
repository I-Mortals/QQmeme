<script lang="ts" setup>
import { store } from '../../store'
import MemePane from './MemePane.vue'

</script>

<template>
  <main class="meme-box dark-bg">
    <template v-for="memes in store.allMemesPath" :key="memes.code">
      <MemePane :memeInfo="memes"/>
    </template>
  </main>
</template>

<style scoped>
.meme-box {
  width: 100%;
  height: 100%;
  overflow-y: auto;
  overflow-x: hidden;
  position: relative;
  scroll-behavior: smooth;
}

.meme-box::before {
  content: '';
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background:
    radial-gradient(circle at 20% 80%, rgba(120, 119, 198, 0.3) 0%, transparent 50%),
    radial-gradient(circle at 80% 20%, rgba(255, 119, 198, 0.3) 0%, transparent 50%),
    radial-gradient(circle at 40% 40%, rgba(120, 219, 255, 0.2) 0%, transparent 50%);
  pointer-events: none;
  z-index: 0;
}

.meme-box::-webkit-scrollbar {
  width: 12px;
}

/* 滚动指示器 */
.meme-box::after {
  content: '';
  position: fixed;
  top: 50%;
  right: 2px;
  width: 4px;
  height: 40px;
  background: rgba(255, 255, 255, 0.4);
  border-radius: 2px;
  transform: translateY(-50%);
  opacity: 0;
  transition: opacity 0.3s ease;
  z-index: 10;
  pointer-events: none;
}

.meme-box:hover::after {
  opacity: 1;
}

@media (max-width: 768px) {
  .meme-box::-webkit-scrollbar {
    width: 8px;
  }

  .meme-box::-webkit-scrollbar-track {
    border-radius: 4px;
    margin: 4px 0;
  }

  .meme-box::-webkit-scrollbar-thumb {
    border-radius: 4px;
    border-width: 1px;
  }
}

@media (max-width: 480px) {
  .meme-box::-webkit-scrollbar {
    width: 6px;
  }

  .meme-box::after {
    right: 1px;
    width: 3px;
    height: 30px;
  }
}
</style>
