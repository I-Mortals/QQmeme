<script lang="ts" setup>
import { GenerateAllMemePath, SelectRootDir } from '@wailsjs/go/memeFile/MemeFile'
import { memeStore, toastStore } from '@/store'
import Button from '@/components/Button.vue'
import Input from '@/components/Input.vue'
import SettingItem from './setting/SettingItem.vue'
import SettingsSection from './setting/SettingsSection.vue'
import SettingGroup from './setting/SettingGroup.vue'

const selectRoot = async () => {
  const path = await SelectRootDir()
  if (path) {
    memeStore.rootPath = path
    memeStore.allMemesPath = await GenerateAllMemePath(memeStore.rootPath)
  }
}

const clearCache = () => {
  memeStore.clearCache()
  memeStore.rootPath = ''
  toastStore.showToast('缓存清除成功！', 'success')
}

const refreshMemes = () => {
  memeStore.refreshMemes()
  toastStore.showToast('刷新列表成功！', 'success')
}
</script>

<template>
  <SettingsSection title="文件管理" section-id="file-management">
    <SettingGroup>
      <SettingItem>
        <template #text>MEME主目录</template>
        <template #desc>选择表情包存储的主目录</template>
        <template #actions>
          <Input :model-value="memeStore.rootPath || '请选择主目录'" readonly placeholder="请选择主目录" class="path-input">
          <template #append>
            <Button variant="primary" icon="lucide:folder-up" @click="selectRoot">
              选择目录
            </Button>
          </template>
          </Input>
        </template>
      </SettingItem>

      <SettingItem>
        <template #text>缓存管理</template>
        <template #desc>管理应用缓存和刷新数据</template>
        <template #actions>
          <div class="button-group">
            <Button variant="danger" icon="lucide:trash-2" @click="clearCache">
              清除缓存
            </Button>
            <Button variant="primary" icon="lucide:refresh-ccw" @click="refreshMemes">
              刷新列表
            </Button>
          </div>
        </template>
      </SettingItem>
    </SettingGroup>
  </SettingsSection>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

.path-input {
  min-width: 300px;
}

.button-group {
  display: flex;
  gap: 0.5rem;
}
</style>
