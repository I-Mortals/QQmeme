<script lang="ts" setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import loadingGif from '@/assets/loading.gif'

export interface LazyLoadImgProps {
  placeholderSrc?: string
  src: string
  alt?: string
  title?: string
  className?: string
}

const props = withDefaults(defineProps<LazyLoadImgProps>(), {
  placeholderSrc: loadingGif,
  alt: '',
  title: '',
  className: ''
})

const imgRef = ref<HTMLImageElement>()
const imgLoaded = ref(false)
const inView = ref(false)
const isVisible = ref(false)

let observer: IntersectionObserver | null = null

const showRealImg = computed(() => inView.value && imgLoaded.value)

onMounted(() => {
  if (!imgRef.value) return

  observer = new IntersectionObserver(
    (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          inView.value = true
          setTimeout(() => {
            isVisible.value = true
          }, 100)
          observer?.unobserve(entry.target)
        }
      })
    },
    {
      threshold: 0.1,
      // -80px tabbar 的高度
      rootMargin: '0px 0px -80px 0px'
    }
  )

  observer.observe(imgRef.value)
})

onUnmounted(() => {
  if (observer) {
    observer.disconnect()
  }
})

const handleLoad = () => {
  imgLoaded.value = true
}

const handleError = () => {
  imgLoaded.value = true
}
</script>

<template>
  <img
    ref="imgRef"
    :class="[
      'lazy-image',
       className,
       [isVisible ? 'visible' : 'hidden']
    ]"
    :src="showRealImg ? src : placeholderSrc"
    :alt="alt"
    :title="title"
    @load="handleLoad"
    @error="handleError"
  />
</template>

<style lang="less" scoped>
.lazy-image {
  opacity: 0;
  transform: translateY(20px) scale(0.95);
  transition: all 0.6s cubic-bezier(0.4, 0, 0.2, 1);

  &.visible {
    opacity: 1;
    transform: translateY(0) scale(1);
  }

  &.hidden {
    height: 150px;
  }
}
</style>
