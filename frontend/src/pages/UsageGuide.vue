<template>
  <div class="usage-guide-page">
    <div class="guide-container">
      <div class="markdown-content" v-html="renderedMarkdown"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { marked } from 'marked'

const renderedMarkdown = ref('')

onMounted(async () => {
  try {
    marked.setOptions({
      breaks: true,
      gfm: true,
    })

    const response = await fetch('/usage-guide.md')

    const markdownText = await response.text()
    renderedMarkdown.value = await marked(markdownText)
  } catch (error) {
    console.error('Failed to load markdown:', error)
    renderedMarkdown.value = '<p>无法加载使用说明，请检查文件是否存在。</p>'
  }
})
</script>

<style lang="less" scoped>
.usage-guide-page {
  height: 100%;
  background: #f8fafc;
  overflow-y: auto;
}

.guide-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 1rem;
}

.markdown-content {
  background: white;
  border-radius: 0.5rem;
  padding: 1.5rem;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
  border: 1px solid #e2e8f0;
  line-height: 1.5;
  color: #4a5568;

  :deep(h1) {
    font-size: 2rem;
    font-weight: 700;
    color: #1a202c;
    margin: 0 0 0.75rem 0;
    text-align: center;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
  }

  :deep(h2) {
    font-size: 1.25rem;
    font-weight: 600;
    color: #1a202c;
    margin: 1.5rem 0 0.75rem 0;
    padding-bottom: 0.375rem;
    border-bottom: 2px solid #e2e8f0;
  }

  :deep(h3) {
    font-size: 1.125rem;
    font-weight: 600;
    color: #2d3748;
    margin: 1.25rem 0 0.5rem 0;
  }

  :deep(h4) {
    font-size: 1rem;
    font-weight: 600;
    color: #2d3748;
    margin: 0.75rem 0 0.375rem 0;
  }

  :deep(p) {
    margin: 0 0 0.75rem 0;
    color: #4a5568;
  }

  :deep(ul), :deep(ol) {
    margin: 0.75rem 0;
    padding-left: 1.5rem;
  }

  :deep(li) {
    margin-bottom: 0.375rem;
    color: #4a5568;
  }

  :deep(blockquote) {
    margin: 0.75rem 0;
    padding: 0.75rem 1rem;
    border-left: 4px solid #667eea;
    background: #f7fafc;
    border-radius: 0 0.375rem 0.375rem 0;

    p {
      margin: 0;
      font-style: italic;
      color: #2d3748;
    }
  }

  :deep(code) {
    background: #f1f5f9;
    padding: 0.2rem 0.4rem;
    border-radius: 0.25rem;
    font-family: 'Courier New', monospace;
    font-size: 0.875rem;
    color: #e53e3e;
  }

  :deep(pre) {
    background: #1a202c;
    color: #e2e8f0;
    padding: 0.75rem;
    border-radius: 0.375rem;
    overflow-x: auto;
    margin: 0.75rem 0;

    code {
      background: transparent;
      padding: 0;
      color: inherit;
      font-size: 0.875rem;
    }
  }

  :deep(strong) {
    font-weight: 600;
    color: #2d3748;
  }

  :deep(em) {
    font-style: italic;
    color: #4a5568;
  }

  :deep(a) {
    color: #667eea;
    text-decoration: none;

    &:hover {
      text-decoration: underline;
    }
  }

  :deep(hr) {
    border: none;
    height: 1px;
    background: #e2e8f0;
    margin: 1.5rem 0;
  }

  :deep(table) {
    width: 100%;
    border-collapse: collapse;
    margin: 0.75rem 0;

    th, td {
      padding: 0.5rem;
      text-align: left;
      border-bottom: 1px solid #e2e8f0;
    }

    th {
      background: #f7fafc;
      font-weight: 600;
      color: #2d3748;
    }
  }

  :deep(.emoji) {
    font-size: 1.2em;
    margin-right: 0.5rem;
  }
}
</style>
