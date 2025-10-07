<script lang="ts" setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { store } from '@/store'
import Button from '@/components/Button.vue'
import { DownloadTgStickerSet } from '../../../wailsjs/go/memeFile/MemeFile'
import { EventsOn } from '../../../wailsjs/runtime'

const API_BASE = computed(() => `https://api.telegram.org/bot${store.botToken}`)

const getQuery = ref('')
const isLoading = ref(false)
const getResults = ref<any[]>([])
const downloadingSets = ref<Set<string>>(new Set())
const downloadedSets = ref<Set<string>>(new Set())
const downloadProgress = ref<Record<string, { current: number; total: number; status: string }>>({})

const getStickerSetName = (input: string): string => {
  const patterns = [
    // https://t.me/addstickers/name
    /https?:\/\/t\.me\/addstickers\/([^\/\s]+)/i,
    // t.me/addstickers/name
    /t\.me\/addstickers\/([^\/\s]+)/i,
    // @name 格式
    /@([^\/\s]+)/i,
    // 直接输入名称
    /^([a-zA-Z0-9_]+)$/
  ]

  for (const pattern of patterns) {
    const match = input.match(pattern)
    if (match) {
      return match[1]
    }
  }

  return input.trim()
}

const getStickerSet = async () => {
  if (!getQuery.value.trim()) {
    store.showToast('请输入获取关键词或链接', 'error')
    return
  }

  if (!store.botToken) {
    store.showToast('请先在设置中配置 Telegram Bot Token', 'error')
    return
  }

  isLoading.value = true
  try {
    const stickerSetName = getStickerSetName(getQuery.value)

    if (!stickerSetName) {
      store.showToast('无法解析 sticker 集合名称', 'error')
      return
    }

    getQuery.value = stickerSetName

    const response = await fetch(`${API_BASE.value}/getStickerSet?name=${stickerSetName}`, {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
      },
      signal: AbortSignal.timeout(30000)
    })

    if (!response.ok) {
      store.showToast(`获取失败: ${response.status} ${response.statusText}`, 'error')
      return
    }

    const data = await response.json()

    if (data.ok) {
      const stickerSetInfo = {
        name: stickerSetName,
        title: data.result.title,
        description: data.result.description,
        stickerCount: data.result.stickers.length,
        isAnimated: data.result.is_animated,
        isVideo: data.result.is_video
      }

      const existingIndex = getResults.value.findIndex(item => item.name === stickerSetName)
      if (existingIndex >= 0) {
        getResults.value[existingIndex] = stickerSetInfo
      } else {
        getResults.value.unshift(stickerSetInfo)
      }

      store.showToast(`获取到表情包集合: ${stickerSetInfo.title} (${stickerSetInfo.stickerCount} 个)`)
    } else {
      store.showToast(`未获取到名为 "${stickerSetName}" 的表情包集合`, 'error')
    }
  } catch (error) {
    console.error('获取失败:', error)
    store.showToast('获取失败，请稍后重试', 'error')
  } finally {
    isLoading.value = false
  }
}

const downloadStickerSet = async (stickerSet: any) => {
  if (downloadingSets.value.has(stickerSet.name) || downloadedSets.value.has(stickerSet.name)) {
    return
  }

  downloadingSets.value.add(stickerSet.name)

  try {
    downloadProgress.value[stickerSet.name] = {
      current: 0,
      total: stickerSet.stickerCount,
      status: '准备下载...'
    }

    const savePath = store.rootPath

    await DownloadTgStickerSet(stickerSet.name, savePath, store.botToken, store.proxyURL, store.proxyEnabled)

    downloadedSets.value.add(stickerSet.name)
    store.showToast(`下载完成: ${stickerSet.title}`, 'success')

    await store.refreshMemes()
    store.forceRefreshCurrentTab()
  } catch (error) {
    console.error('下载失败:', error)
    const errorMessage = error instanceof Error ? error.message : '未知错误'
    store.showToast(`下载失败: ${errorMessage}`, 'error')
  } finally {
    downloadingSets.value.delete(stickerSet.name)
    delete downloadProgress.value[stickerSet.name]
  }
}

const hasResults = computed(() => getResults.value.length > 0)
const isGetDisabled = computed(() => isLoading.value || !getQuery.value.trim())

const isDownloading = (stickerSetName: string) => downloadingSets.value.has(stickerSetName)
const isDownloaded = (stickerSetName: string) => downloadedSets.value.has(stickerSetName)

const getDownloadProgress = (stickerSetName: string) => {
  return downloadProgress.value[stickerSetName] || { current: 0, total: 0, status: '' }
}

const removeStickerSet = (stickerSetName: string) => {
  const index = getResults.value.findIndex(item => item.name === stickerSetName)
  if (index >= 0) {
    getResults.value.splice(index, 1)
    downloadingSets.value.delete(stickerSetName)
    downloadedSets.value.delete(stickerSetName)
    delete downloadProgress.value[stickerSetName]
  }
}

let progressUnsubscribe: (() => void) | null = null

onMounted(() => {
  progressUnsubscribe = EventsOn('telegram-download-progress', (data: any) => {
    const { stickerSetName, progress } = data

    if (stickerSetName && downloadingSets.value.has(stickerSetName)) {
      downloadProgress.value[stickerSetName] = {
        current: progress.current,
        total: progress.total,
        status: progress.status
      }
    }
  })
})

onUnmounted(() => {
  if (progressUnsubscribe) {
    progressUnsubscribe()
  }
})

</script>

<template>
  <div class="telegram-sticker-pane">
    <!-- 获取区域 -->
    <div class="get-section">
      <div class="get-header">
        <h3>📱 Telegram 表情包获取</h3>
      </div>

      <div class="get-controls">
        <div class="get-input-group">
          <input
            v-model="getQuery"
            type="text"
            placeholder="输入链接或名称"
            class="get-input"
            @keyup.enter="getStickerSet"
            :disabled="isLoading"
          />
          <Button
            @click="getStickerSet"
            :disabled="isGetDisabled"
            :loading="isLoading"
            class="get-button"
          >
            {{ isLoading ? '获取中...' : '获取' }}
          </Button>
        </div>

      </div>
    </div>

    <!-- 获取结果列表 -->
    <div v-if="hasResults" class="get-results">
      <div
        v-for="stickerSet in getResults"
        :key="stickerSet.name"
        class="sticker-set-item"
      >
        <div class="set-info">
          <div class="set-header">
            <div class="set-title-section">
              <h4>{{ stickerSet.title }}</h4>
              <span class="sticker-count">{{ stickerSet.stickerCount }} 个表情包</span>
            </div>
            <div class="set-actions">
              <Button
                @click="downloadStickerSet(stickerSet)"
                :disabled="isDownloading(stickerSet.name) || isDownloaded(stickerSet.name)"
                :loading="isDownloading(stickerSet.name)"
                :class="{ 'downloaded': isDownloaded(stickerSet.name) }"
              >
                {{
                  isDownloading(stickerSet.name) ? '下载中...' :
                  isDownloaded(stickerSet.name) ? '已下载' : '下载'
                }}
              </Button>
              <Button
                @click="removeStickerSet(stickerSet.name)"
                variant="danger"
                class="delete-button"
              >
                删除
              </Button>
            </div>
          </div>
          <p v-if="stickerSet.description" class="set-description">
            {{ stickerSet.description }}
          </p>
          <div class="set-meta">
            <span class="set-name">集合名称: {{ stickerSet.name }}</span>
            <div class="format-tags">
              <span v-if="stickerSet.isAnimated" class="format-tag animated">动画</span>
              <span v-if="stickerSet.isVideo" class="format-tag video">视频</span>
              <span v-if="!stickerSet.isAnimated && !stickerSet.isVideo" class="format-tag static">静态</span>
            </div>
          </div>

          <!-- 下载进度显示 -->
          <div v-if="isDownloading(stickerSet.name)" class="download-progress">
            <div class="progress-info">
              <span class="progress-text">{{ getDownloadProgress(stickerSet.name).status }}</span>
              <span class="progress-count">
                {{ getDownloadProgress(stickerSet.name).current }}/{{ getDownloadProgress(stickerSet.name).total }}
              </span>
            </div>
            <div class="progress-bar">
              <div
                class="progress-fill"
                :style="{ width: `${(getDownloadProgress(stickerSet.name).current / getDownloadProgress(stickerSet.name).total) * 100}%` }"
              ></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 空状态 -->
    <div v-if="!hasResults && !isLoading" class="empty-state">
      <div class="example-links">
        <p><strong>支持的格式：</strong></p>
        <p>• <code>https://t.me/addstickers/capoo_sp_animated</code></p>
        <p>• <code>t.me/addstickers/capoo_sp_animated</code></p>
        <p>• <code>capoo_sp_animated</code></p>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
.telegram-sticker-pane {
  padding: 1rem;
  max-width: 1200px;
  margin: 0 auto;
}

.get-section {
  margin-bottom: 1.5rem;
}

.get-header {
  text-align: center;
  margin-bottom: 1.5rem;

  h3 {
    font-size: 1.5rem;
    font-weight: 600;
    color: var(--theme-primary);
    margin: 0 0 0.5rem 0;
  }
}

.get-controls {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.get-input-group {
  display: flex;
  gap: 0.75rem;
  align-items: center;
}

.get-button {
  flex-shrink: 0;
}

.get-results {
  display: flex;
  flex-direction: column;
  gap: 0.875rem;
  margin-bottom: 1.5rem;
}

.sticker-set-item {
  background: #f8f9fa;
  border-radius: 0.75rem;
  padding: 1.25rem;
  border: 1px solid #e9ecef;
  transition: all 0.3s ease;

  &:hover {
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    border-color: var(--theme-primary);
  }
}

.set-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 0.5rem;
  gap: 1rem;

  h4 {
    margin: 0;
    color: #333;
    font-size: 1.25rem;
  }
}

.set-actions {
  display: flex;
  gap: 0.5rem;
  align-items: center;
}

.set-title-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.sticker-count {
  background: var(--theme-primary);
  color: white;
  padding: 0.25rem 0.75rem;
  border-radius: 1rem;
  font-size: 0.875rem;
  font-weight: 500;
  align-self: flex-start;
}

.set-description {
  margin: 0 0 0.75rem 0;
  color: #666;
  line-height: 1.5;
}

.set-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 0.875rem;
}

.set-name {
  color: #666;
  font-family: 'Courier New', monospace;
}

.format-tags {
  display: flex;
  gap: 0.5rem;
}

.format-tag {
  padding: 0.125rem 0.5rem;
  border-radius: 0.25rem;
  font-size: 0.75rem;
  font-weight: 600;

  &.animated {
    background: rgba(255, 107, 53, 0.1);
    color: #ff6b35;
  }

  &.video {
    background: rgba(111, 66, 193, 0.1);
    color: #6f42c1;
  }

  &.static {
    background: rgba(40, 167, 69, 0.1);
    color: #28a745;
  }
}

.empty-state {
  text-align: center;
  color: #666;
}

.example-links {
  margin-top: 1rem;
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 0.5rem;
  border: 1px solid #e9ecef;
  text-align: left;
  max-width: 400px;
  margin-left: auto;
  margin-right: auto;

  p {
    margin: 0.25rem 0;
    font-size: 0.875rem;
  }

  code {
    background: #e9ecef;
    padding: 0.125rem 0.25rem;
    border-radius: 0.25rem;
    font-family: 'Courier New', monospace;
    font-size: 0.8rem;
    color: #495057;
  }
}

.download-progress {
  margin-top: 1rem;
  padding: 0.75rem;
  background: #f8f9fa;
  border-radius: 0.5rem;
  border: 1px solid #e9ecef;
}

.progress-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
}

.progress-text {
  color: #666;
  font-weight: 500;
}

.progress-count {
  color: var(--theme-primary);
  font-weight: 600;
  font-family: 'Courier New', monospace;
}

.progress-bar {
  width: 100%;
  height: 6px;
  background: #e9ecef;
  border-radius: 3px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, var(--theme-primary), #4fc3f7);
  border-radius: 3px;
  transition: width 0.3s ease;
}

</style>
