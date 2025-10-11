<script lang="ts" setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { store } from '@/store'
import Button from '@/components/Button.vue'
import Input from '@/components/Input.vue'
import Select from '@/components/Select.vue'
import { DownloadTgStickerSet } from '../../../wailsjs/go/memeFile/MemeFile'
import { EventsOn } from '../../../wailsjs/runtime'

// 类型定义
interface StickerSetInfo {
  name: string
  title: string
  description: string
  stickerCount: number
  isAnimated: boolean
  isVideo: boolean
}

interface DownloadProgress {
  current: number
  total: number
  status: string
}

interface ProgressUpdateData {
  stickerSetName: string
  progress: DownloadProgress
}

// API 相关
const API_BASE = computed(() => `https://api.telegram.org/bot${store.botToken}`)

// 搜索相关状态
const searchQuery = ref<string>('')
const isSearching = ref<boolean>(false)
const searchResults = ref<StickerSetInfo[]>([])

// 下载相关状态
const downloadingSets = ref<Set<string>>(new Set())
const downloadedSets = ref<Set<string>>(new Set())
const downloadProgress = ref<Record<string, DownloadProgress>>({})
const selectedFolders = ref<Record<string, string>>({})

// 进度更新相关
let progressUnsubscribe: (() => void) | null = null

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
  if (!searchQuery.value.trim()) {
    store.showToast('请输入获取关键词或链接', 'error')
    return
  }

  if (!store.botToken) {
    store.showToast('请先在设置中配置 Telegram Bot Token', 'error')
    return
  }

  isSearching.value = true
  try {
    const stickerSetName = getStickerSetName(searchQuery.value)

    if (!stickerSetName) {
      store.showToast('无法解析 sticker 集合名称', 'error')
      return
    }

    searchQuery.value = stickerSetName

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
      const stickerSetInfo: StickerSetInfo = {
        name: stickerSetName,
        title: data.result.title,
        description: data.result.description,
        stickerCount: data.result.stickers.length,
        isAnimated: data.result.is_animated,
        isVideo: data.result.is_video
      }

      const existingIndex = searchResults.value.findIndex(item => item.name === stickerSetName)
      if (existingIndex >= 0) {
        searchResults.value[existingIndex] = stickerSetInfo
      } else {
        searchResults.value.unshift(stickerSetInfo)
      }

      store.showToast(`获取到表情包集合: ${stickerSetInfo.title} (${stickerSetInfo.stickerCount} 个)`)
    } else {
      store.showToast(`未获取到名为 "${stickerSetName}" 的表情包集合`, 'error')
    }
  } catch (error) {
    console.error('获取失败:', error)
    store.showToast('获取失败，请稍后重试', 'error')
  } finally {
    isSearching.value = false
  }
}

const downloadStickerSet = async (stickerSet: StickerSetInfo) => {
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

    let finalSavePath = store.rootPath
    const selectedFolderCode = selectedFolders.value[stickerSet.name]
    if (selectedFolderCode) {
      const selectedFolderInfo = store.allMemesPath.find(folder => folder.code === selectedFolderCode)
      if (selectedFolderInfo) {
        // 直接使用选中文件夹的完整路径
        finalSavePath = selectedFolderInfo.parentPath
      }
    } else {
      // 没有选择文件夹时，直接用 sticker集合的名称
      finalSavePath = `${store.rootPath}/${stickerSet.name}`
    }

    await DownloadTgStickerSet(stickerSet.name, finalSavePath, store.botToken, store.proxyURL, store.proxyEnabled)

    // 更新进度为完成状态
    downloadProgress.value[stickerSet.name] = {
      current: stickerSet.stickerCount,
      total: stickerSet.stickerCount,
      status: `下载完成: ${stickerSet.title}`
    }

    downloadedSets.value.add(stickerSet.name)
    store.showToast(`下载完成: ${stickerSet.title}`, 'success')

    await store.refreshMemes()
    store.forceRefreshCurrentTab()

    setTimeout(() => {
      downloadingSets.value.delete(stickerSet.name)
      delete downloadProgress.value[stickerSet.name]
    }, 2000)
  } catch (error) {
    console.error('下载失败:', error)
    const errorMessage = error instanceof Error ? error.message : '未知错误'
    store.showToast(`下载失败: ${errorMessage}`, 'error')

    // 下载失败时立即清理
    downloadingSets.value.delete(stickerSet.name)
    delete downloadProgress.value[stickerSet.name]
  }
}

const hasResults = computed(() => searchResults.value.length > 0)
const isGetDisabled = computed(() => isSearching.value || !searchQuery.value.trim())

const folderOptions = computed(() => {
  const options = [
    { value: '', label: '默认（使用sticker集合名称）' }
  ]

  store.allMemesPath.forEach(folder => {
    options.push({
      value: folder.code,
      label: folder.name
    })
  })

  return options
})

const isDownloading = (stickerSetName: string) => downloadingSets.value.has(stickerSetName)
const isDownloaded = (stickerSetName: string) => downloadedSets.value.has(stickerSetName)

const getDownloadProgress = (stickerSetName: string): DownloadProgress => {
  return downloadProgress.value[stickerSetName] || { current: 0, total: 0, status: '' }
}

// 计算进度百分比
const getProgressPercentage = (stickerSetName: string): number => {
  const progress = getDownloadProgress(stickerSetName)
  return progress.total > 0 ? Math.round((progress.current / progress.total) * 100) : 0
}

const removeStickerSet = (stickerSetName: string) => {
  const index = searchResults.value.findIndex(item => item.name === stickerSetName)
  if (index >= 0) {
    searchResults.value.splice(index, 1)
    downloadingSets.value.delete(stickerSetName)
    downloadedSets.value.delete(stickerSetName)
    delete downloadProgress.value[stickerSetName]
    delete selectedFolders.value[stickerSetName]
  }
}

onMounted(() => {
  progressUnsubscribe = EventsOn('telegram-download-progress', (data: ProgressUpdateData) => {
    const { stickerSetName, progress } = data

    if (stickerSetName && downloadingSets.value.has(stickerSetName)) {
      // 直接更新进度
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
        <Input
          v-model="searchQuery"
          type="text"
          placeholder="输入链接或名称"
          :disabled="isSearching"
          @keyup.enter="getStickerSet"
        >
          <template #append>
            <Button
              @click="getStickerSet"
              :disabled="isGetDisabled"
              :loading="isSearching"
            >
              {{ isSearching ? '获取中...' : '获取' }}
            </Button>
          </template>
        </Input>
      </div>
    </div>


    <!-- 获取结果列表 -->
    <div v-if="hasResults" class="get-results">
      <div
        v-for="stickerSet in searchResults"
        :key="stickerSet.name"
        class="sticker-set-item"
      >
        <div class="set-info">
          <div class="set-content">
            <div class="set-metadata">
              <div class="set-header">
                <h4>{{ stickerSet.title }}</h4>
                <div class="title-tags">
                  <span class="sticker-count">{{ stickerSet.stickerCount }} 个表情包</span>
                  <span v-if="stickerSet.isAnimated" class="format-tag animated">动画</span>
                  <span v-if="stickerSet.isVideo" class="format-tag video">视频</span>
                  <span v-if="!stickerSet.isAnimated && !stickerSet.isVideo" class="format-tag static">静态</span>
                </div>
              </div>
              <p v-if="stickerSet.description" class="set-description">
                {{ stickerSet.description }}
              </p>
              <div class="set-meta">
                <span class="set-name">集合名称: {{ stickerSet.name }}</span>
              </div>
            </div>

            <div class="set-actions">
              <div class="action-buttons">
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

              <div class="item-folder-selector">
                <label class="item-folder-label">下载到文件夹：</label>
                <Select
                  v-model="selectedFolders[stickerSet.name]"
                  :options="folderOptions"
                  :disabled="isDownloading(stickerSet.name)"
                  placeholder="选择文件夹"
                />
              </div>
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
                :class="{ 'progress-complete': getDownloadProgress(stickerSet.name).current === getDownloadProgress(stickerSet.name).total }"
                :style="{ width: `${getProgressPercentage(stickerSet.name)}%` }"
              ></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="!hasResults && !isSearching" class="empty-state">
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
@import '@/styles/variables.less';
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
    color: @rgb-p;
    margin: 0 0 0.5rem 0;
  }
}

.get-controls {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}


.item-folder-selector {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem;
  background: rgba(var(--bc), 0.02);
  border-radius: 0.5rem;
  border: 1px solid @rgb-b3;
  width: 100%;
  min-width: 200px;
}

.item-folder-label {
  font-size: 0.8rem;
  font-weight: 500;
  color: @rgb-bc;
  white-space: nowrap;
  flex-shrink: 0;
}

.get-results {
  display: flex;
  flex-direction: column;
  gap: 0.875rem;
  margin-bottom: 1.5rem;
}

.sticker-set-item {
  background: @rgb-b1;
  border-radius: 0.75rem;
  padding: 1.25rem;
  border: 1px solid @rgb-b3;
  transition: all 0.3s ease;

  &:hover {
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    border-color: @rgb-p;
  }
}

.set-content {
  display: flex;
  gap: 1.5rem;
  align-items: flex-start;
}

.set-metadata {
  flex: 1;
  min-width: 0;
}

.set-header {
  margin-bottom: 0.5rem;

  h4 {
    margin: 0 0 0.5rem 0;
    color: @rgb-bc;
    font-size: 1.25rem;
  }
}

.set-actions {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  align-items: flex-end;
  min-width: 200px;
}

.action-buttons {
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

.title-tags {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  flex-wrap: wrap;
}

.sticker-count {
  background: @rgb-p;
  color: @rgb-pc;
  padding: 0.25rem 0.75rem;
  border-radius: 1rem;
  font-size: 0.875rem;
  font-weight: 500;
  align-self: flex-start;
}

.set-description {
  margin: 0 0 0.75rem 0;
  color: @rgb-bc;
  line-height: 1.5;
  opacity: 0.8;
}

.set-meta {
  font-size: 0.875rem;
  margin-top: 0.5rem;
}

.set-name {
  color: @rgb-bc;
  opacity: 0.8;
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
  color: @rgb-bc;
}

.example-links {
  margin: 0 auto;
  margin-top: 1rem;
  padding: 1rem;
  background: rgba(var(--bc), 0.06);
  border-radius: 0.5rem;
  border: 1px solid @rgb-b3;
  max-width: 400px;
  display: flex;
  flex-direction: column;
  opacity: 0.7;

  p {
    margin: 0.25rem 0;
    font-size: 0.875rem;
    color: @rgb-bc;
  }

  code {
    background: @rgb-b3;
    padding: 0.125rem 0.25rem;
    border-radius: 0.25rem;
    font-size: 0.8rem;
    color: @rgb-bc;
  }
}

.download-progress {
  margin-top: 1rem;
  padding: 0.75rem;
  background: rgba(var(--bc), 0.06);
  border-radius: 0.5rem;
  border: 1px solid @rgb-b3;
}

.progress-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
}

.progress-text {
  color: @rgb-bc;
  font-weight: 500;
}

.progress-count {
  color: @rgb-p;
  font-weight: 600;
  font-family: 'Courier New', monospace;
}

.progress-bar {
  width: 100%;
  height: 6px;
  background: @rgb-b3;
  border-radius: 3px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background-color: @rgb-p;
  border-radius: 3px;
  transition: width 0.5s cubic-bezier(0.4, 0, 0.2, 1);

  &.progress-complete {
    background: @rgb-s;
  }
}
</style>
