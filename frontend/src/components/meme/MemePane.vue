<script setup lang="ts">
import { memeFile } from '../../../wailsjs/go/models'
import MemeInfo = memeFile.MemeInfo
import { joinShowImgPath, joinPath } from '../../utils'
import { store } from '../../store'
import { ref, watch } from 'vue'
import { WriteFileToClipboard } from '../../../wailsjs/go/memeFile/memeFile'

interface MemePaneProps {
  memeInfo: MemeInfo
}
const {
  memeInfo
} = defineProps<MemePaneProps>()

const shouldRender = ref(false)
const isActive = ref(
  store.tabCurrent === memeInfo.code
)

watch(() => store.tabCurrent, (newVal) => {
  isActive.value = newVal === memeInfo.code
  shouldRender.value = shouldRender.value || isActive.value
}, {
  deep: true,
  immediate: true
})

const handleClick = async (image: string) => {
  const realPath = joinPath(store.rootPath+`\\`+memeInfo.code,image)
  if (realPath) {
    try {
      await WriteFileToClipboard(realPath)
      store.showToast('复制成功！')
    } catch (error) {      
      store.showToast(`复制失败：${error}`, 'error')
    }
  }
}

</script>

<template>
  <div
    v-if="shouldRender"
    v-show="isActive"
    class="meme-grid"
  >
    <div
      v-for="image in memeInfo.memes"
      :key="image"
      class="meme-item"
    >
      <img
       loading="lazy"
        :src="joinShowImgPath(store.rootPath+`\\`+memeInfo.code,image)"
        :alt="image"
        @click="handleClick(image)">
    </div>
  </div>
</template>

<style scoped>
.meme-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  justify-content: start;
  align-content: start;
  gap: 1.25rem;
  background: var(--theme-background);
  margin-bottom: 5rem;
  width: 100%;
  padding: 1.5rem;
  min-height: calc(100vh - 5rem);
  position: relative;
}

.meme-grid::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, 
    transparent 0%, 
    color-mix(in srgb, var(--theme-primary) 20%, transparent) 50%, 
    transparent 100%);
}

.meme-item {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  display: flex;
  align-items: center;
  border-radius: 1rem;
  position: relative;
}

.meme-item img {
  width: 100%;
  height: auto;
  object-fit: cover;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  position: relative;
  z-index: 2;
  cursor: pointer;
  border-radius: 0.75rem;
  background: #ffffff;
  box-shadow:
      0 4px 6px rgba(0, 0, 0, 0.05),
      0 1px 3px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(232, 232, 232, 0.8);
}

.meme-item img:hover {
  transform: translateY(-4px) scale(1.1);
  border-color: color-mix(in srgb, var(--theme-primary) 20%, transparent);
}

.meme-item:active img {
  transform: scale(0.98);
}

@media (max-width: 1200px) {
  .meme-grid {
    grid-template-columns: repeat(auto-fill, minmax(130px, 1fr));
    gap: 1rem;
    padding: 1.25rem;
  }
}

@media (max-width: 768px) {
  .meme-grid {
    grid-template-columns: repeat(auto-fill, minmax(110px, 1fr));
    gap: 0.875rem;
    padding: 1rem;
    margin-bottom: 5rem;
    min-height: calc(100vh - 5rem);
  }
  
  .meme-item {
    border-radius: 0.875rem;
  }
  
  .meme-item img {
    border-radius: 0.625rem;
  }
}

@media (max-width: 480px) {
  .meme-grid {
    grid-template-columns: repeat(auto-fill, minmax(90px, 1fr));
    gap: 0.625rem;
    padding: 0.875rem;
    margin-bottom: 4.5rem;
    min-height: calc(100vh - 4.5rem);
  }
  
  .meme-item {
    border-radius: 0.625rem;
  }
  
  .meme-item img {
    border-radius: 0.5rem;
  }
  
  .meme-item:hover {
    transform: translateY(-6px) scale(1.02);
  }
}

@media (max-width: 360px) {
  .meme-grid {
    grid-template-columns: repeat(auto-fill, minmax(80px, 1fr));
    gap: 0.5rem;
    padding: 0.625rem;
  }
}
</style>
