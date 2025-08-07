<script setup lang="ts">
import { memeFile } from '../../../wailsjs/go/models'
import MemeInfo = memeFile.MemeInfo
import { joinShowImgPath } from '../../utils'
import { store } from '../../store'
import { ref, watch } from 'vue'

interface MemePaneProps {
  memeInfo: MemeInfo
}

const {
  memeInfo
} = withDefaults(defineProps<MemePaneProps>(), {
  memeInfo: undefined
})

const shouldRender = ref(false)
const isActive = ref(
  store.tabCurrent === memeInfo.Code
)

watch(() => store.tabCurrent, (newVal) => {
  isActive.value = newVal === memeInfo.Code
  shouldRender.value = shouldRender.value || isActive.value
}, {
  deep: true,
  immediate: true
})

</script>

<template>
  <div
    v-if="shouldRender"
    v-show="isActive"
    class="meme-grid"
  >
    <div
      v-for="image in memeInfo.Memes"
      :key="image"
      class="meme-item"
    >
      <img loading="lazy" :src="joinShowImgPath(store.rootPath+`\\`+memeInfo.Code,image)" :alt="image">
    </div>
  </div>
</template>

<style scoped>
.meme-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(120px, 1fr));
  gap: .5rem;
  background-color: antiquewhite;
  margin-bottom: 5rem;
  width: 100%;
  flex-shrink: 0;
  padding: .5rem;
}

.meme-item {
  transition: transform 0.2s ease;
  display: flex;
  align-items: center;
}

.meme-item:hover {
  transform: scale(1.05);
}

.meme-item img {
  width: 100%;
  object-fit: cover;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  background-color: #f3f4f7;
  border-radius: .25rem;
}
</style>
