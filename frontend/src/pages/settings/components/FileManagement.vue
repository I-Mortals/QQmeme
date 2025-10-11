<script lang="ts" setup>
import { GenerateAllMemePath, SelectRootDir } from '../../../../wailsjs/go/memeFile/MemeFile'
import { store } from '@/store'
import Button from '@/components/Button.vue'
import Input from '@/components/Input.vue'

const selectRoot = async () => {
  const path = await SelectRootDir()
  if (path) {
    store.rootPath = path
    store.allMemesPath = await GenerateAllMemePath(store.rootPath)
  }
}

const clearCache = () => {
  store.clearCache()
  store.rootPath = ''
  store.showToast('缓存清除成功！', 'success')
}

const refreshMemes = () => {
  store.refreshMemes()
  store.showToast('刷新列表成功！', 'success')
}
</script>

<template>
  <div id="file-management" class="content-section">
    <h3 class="section-title">文件管理</h3>
    <div class="setting-group">
      <div class="setting-item">
        <div class="setting-label">
          <span class="label-text">MEME主目录</span>
          <span class="label-desc">选择表情包存储的主目录</span>
        </div>
        <div class="setting-control">
          <Input
            :model-value="store.rootPath || '请选择主目录'"
            readonly
            placeholder="请选择主目录"
            class="path-input"
          >
            <template #append>
              <Button
                variant="primary"
                icon="lucide:folder-up"
                @click="selectRoot">
                选择目录
              </Button>
            </template>
          </Input>
        </div>
      </div>

      <div class="setting-item">
        <div class="setting-label">
          <span class="label-text">缓存管理</span>
          <span class="label-desc">管理应用缓存和刷新数据</span>
        </div>
        <div class="setting-control">
          <div class="button-group">
            <Button
              variant="danger"
              icon="lucide:trash-2"
              @click="clearCache">
              清除缓存
            </Button>
            <Button
              variant="primary"
              icon="lucide:refresh-ccw"
              @click="refreshMemes">
              刷新列表
            </Button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

.content-section {
  margin-bottom: 1.5rem;

  &:last-child {
    margin-bottom: 0;
  }
}

.section-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: @rgb-bc;
  margin: 0 0 1rem 0;
  padding-bottom: 0.5rem;
  border-bottom: 1px solid @rgb-b3;
}

.setting-group {
  background: @rgb-b1;
  border-radius: 0.5rem;
  border: 1px solid @rgb-b3;
  overflow: hidden;
}

.setting-item {
  display: flex;
  align-items: center;
  padding: 0.875rem 1.25rem;
  border-bottom: 1px solid @rgb-b3;

  &:last-child {
    border-bottom: none;
  }
}

.setting-label {
  flex: 1;
  min-width: 0;
}

.label-text {
  display: block;
  font-size: 0.875rem;
  font-weight: 500;
  color: @rgb-bc;
  margin-bottom: 0.25rem;
}

.label-desc {
  display: block;
  font-size: 0.75rem;
  color: @rgb-bc;
  opacity: 0.7;
  line-height: 1.4;
}

.setting-control {
  flex-shrink: 0;
  margin-left: 1rem;
}

.path-input {
  min-width: 300px;
}

.button-group {
  display: flex;
  gap: 0.5rem;
}
</style>
