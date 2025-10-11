<template>
  <div class="usage-guide-page">
    <div class="guide-container">
      <Markdown :content="markdownContent" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import Markdown from '@/components/Markdown.vue'

const markdownContent = ref('')

onMounted(async () => {
  try {
    const response = await fetch('/usage-guide.md')
    if (response.ok) {
      const text = await response.text()
      markdownContent.value = text
    } else {
      markdownContent.value = '无法加载使用说明，请检查文件是否存在。'
    }
  } catch (error) {
    console.error('加载使用说明失败:', error)
    markdownContent.value = '无法加载使用说明，请检查文件是否存在。'
  }
})
</script>

<style lang="less" scoped>
@import '@/styles/variables.less';

.usage-guide-page {
  height: 100%;
  background: @rgb-b1;
  overflow-y: auto;
}

.guide-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 1rem;
}
</style>
