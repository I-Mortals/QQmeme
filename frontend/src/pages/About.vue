<script lang="ts" setup>
import { Icon } from '@iconify/vue'
import { BrowserOpenURL } from '@wailsjs/runtime/runtime'
import logo from '@/assets/images/logo-universal.png'

const appInfo = {
  name: 'QQmeme',
  subtitle: '表情包管理桌面应用',
  description:
    '一个基于 Wails + Vue 3 构建的表情包管理工具，支持表情包整理与收藏、标签页排序、Telegram 贴纸下载等功能。',
  url: 'https://github.com/I-Mortals/QQmeme'
}

const authors = [
  {
    name: 'I-Mortals',
    url: 'https://github.com/I-Mortals',
    role: '作者'
  },
  {
    name: 'archonk',
    url: 'https://github.com/archonk',
    role: '作者'
  }
]

const openLink = (url: string) => {
  BrowserOpenURL(url)
}
</script>

<template>
  <div class="about-page">
    <div class="about-container">
      <div class="app-header">
        <img :src="logo" class="app-logo" alt="QQmeme logo" />
        <div class="app-meta">
          <h1 class="app-name">{{ appInfo.name }}</h1>
          <p class="app-subtitle">{{ appInfo.subtitle }}</p>
        </div>
      </div>

      <p class="app-description">{{ appInfo.description }}</p>

      <div class="info-card">
        <div class="info-row">
          <div class="info-label">
            <Icon icon="lucide:link" :width="16" :height="16" />
            <span>软件地址</span>
          </div>
          <button class="link-button" @click="openLink(appInfo.url)">
            {{ appInfo.url }}
            <Icon icon="lucide:external-link" :width="14" :height="14" />
          </button>
        </div>
      </div>

      <h2 class="section-heading">作者</h2>
      <div class="author-list">
        <button
          v-for="author in authors"
          :key="author.name"
          class="author-card"
          @click="openLink(author.url)"
        >
          <Icon icon="lucide:github" :width="28" :height="28" class="author-icon" />
          <div class="author-info">
            <span class="author-name">{{ author.name }}</span>
            <span class="author-role">{{ author.role }}</span>
          </div>
          <Icon icon="lucide:external-link" :width="14" :height="14" class="author-arrow" />
        </button>
      </div>

      <p class="footer">感谢使用 QQmeme ❤️</p>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

.about-page {
  height: 100%;
  background: @rgb-b1;
  overflow-y: auto;
}

.about-container {
  max-width: 720px;
  margin: 0 auto;
  padding: 2rem 1.5rem;
}

.app-header {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.app-logo {
  width: 72px;
  height: 72px;
  object-fit: contain;
  border-radius: 1rem;
  background: rgba(var(--bc), 0.06);
  padding: 0.5rem;
}

.app-meta {
  display: flex;
  flex-direction: column;
}

.app-name {
  margin: 0;
  font-size: 1.75rem;
  font-weight: 700;
  color: @rgb-bc;
}

.app-subtitle {
  margin: 0.15rem 0 0;
  font-size: 0.9rem;
  color: @rgb-p;
}

.app-description {
  margin: 1.25rem 0;
  font-size: 0.9rem;
  line-height: 1.7;
  color: @rgb-bc;
  opacity: 0.8;
}

.info-card {
  background: rgba(var(--bc), 0.04);
  border: 1px solid @rgb-b3;
  border-radius: 0.75rem;
  padding: 0.5rem 0.75rem;
}

.info-row {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  flex-wrap: wrap;
}

.info-label {
  display: flex;
  align-items: center;
  gap: 0.4rem;
  font-size: 0.875rem;
  font-weight: 500;
  color: @rgb-bc;
}

.link-button {
  display: inline-flex;
  align-items: center;
  gap: 0.35rem;
  border: none;
  background: transparent;
  color: @rgb-p;
  font-size: 0.85rem;
  cursor: pointer;
  padding: 0.35rem 0.5rem;
  border-radius: 0.5rem;
  transition: background-color 0.2s ease;

  &:hover {
    background: rgba(@p, 0.1);
  }
}

.section-heading {
  margin: 1.75rem 0 0.75rem;
  font-size: 1.1rem;
  font-weight: 600;
  color: @rgb-bc;
}

.author-list {
  display: flex;
  flex-direction: column;
  gap: 0.65rem;
}

.author-card {
  display: flex;
  align-items: center;
  gap: 0.85rem;
  width: 100%;
  padding: 0.85rem 1rem;
  border: 1px solid @rgb-b3;
  border-radius: 0.75rem;
  background: rgba(var(--bc), 0.04);
  cursor: pointer;
  text-align: left;
  transition: all 0.2s ease;

  &:hover {
    border-color: @rgb-p;
    background: rgba(@p, 0.08);
    transform: translateY(-1px);

    .author-arrow {
      opacity: 1;
      transform: translateX(2px);
    }
  }
}

.author-icon {
  color: @rgb-bc;
  flex-shrink: 0;
}

.author-info {
  display: flex;
  flex-direction: column;
  flex: 1;
  min-width: 0;
}

.author-name {
  font-size: 0.95rem;
  font-weight: 600;
  color: @rgb-bc;
}

.author-role {
  font-size: 0.75rem;
  color: @rgb-bc;
  opacity: 0.6;
}

.author-arrow {
  color: @rgb-p;
  opacity: 0;
  transition: all 0.2s ease;
}

.footer {
  margin-top: 2rem;
  text-align: center;
  font-size: 0.8rem;
  color: @rgb-bc;
  opacity: 0.5;
}
</style>
