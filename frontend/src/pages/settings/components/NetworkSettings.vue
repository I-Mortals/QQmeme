<script setup lang="ts">
import { ref } from 'vue'
import Button from '@/components/Button.vue'
import { store } from '@/store'

const botTokenInput = ref(store.botToken)
const proxyEnabled = ref(store.proxyEnabled)
const proxyURLInput = ref(store.proxyURL)

const saveBotToken = () => {
  const token = botTokenInput.value.trim()

  if (!token) {
    store.showToast('Bot Token 不能为空', 'error')
    return
  }

  if (!/^\d+:[A-Za-z0-9_-]+$/.test(token)) {
    store.showToast('Bot Token 格式不正确，应为 "数字:字符串" 格式', 'error')
    return
  }

  store.setBotToken(token)
  store.showToast('Telegram Bot Token 保存成功！', 'success')
}

const saveProxySettings = () => {
  if (proxyEnabled.value) {
    const url = proxyURLInput.value.trim()

    if (!url) {
      store.showToast('启用代理时，代理地址不能为空', 'error')
      return
    }

    try {
      new URL(url)
    } catch {
      store.showToast('代理地址格式不正确，请输入有效的 URL', 'error')
      return
    }
  }

  store.setProxySettings(proxyEnabled.value, proxyURLInput.value.trim())
  store.showToast('代理设置保存成功！', 'success')
}

const toggleProxy = () => {
  proxyEnabled.value = !proxyEnabled.value
  store.setProxySettings(proxyEnabled.value, store.proxyURL)
}
</script>

<template>
  <div class="content-section">
    <h3 class="section-title">网络设置</h3>
    <div class="setting-group">
      <!-- Telegram 配置 -->
      <div class="setting-item">
        <div class="setting-label">
          <span class="label-text">Telegram Bot Token</span>
          <span class="label-desc">用于获取Telegram贴纸包的机器人Token</span>
        </div>
        <div class="setting-control">
          <div class="input-group">
            <input
              v-model="botTokenInput"
              type="text"
              placeholder="请输入Bot Token"
              class="config-input"
            />
            <Button variant="primary" @click="saveBotToken">
              保存
            </Button>
          </div>
        </div>
      </div>

      <!-- 代理设置 -->
      <div class="setting-item">
        <div class="setting-label">
          <span class="label-text">代理设置</span>
          <span class="label-desc">配置网络代理以访问Telegram服务</span>
        </div>
        <div class="setting-control">
          <div class="proxy-controls">
            <div class="toggle-switch" @click="toggleProxy">
              <div class="toggle-slider" :class="{ active: proxyEnabled }"></div>
            </div>
            <span class="toggle-label">{{ proxyEnabled ? '已启用' : '已禁用' }}</span>
          </div>
        </div>
      </div>

      <div v-if="proxyEnabled" class="setting-item">
        <div class="setting-label">
          <span class="label-text">代理地址</span>
          <span class="label-desc">代理服务器的完整地址</span>
        </div>
        <div class="setting-control">
          <div class="input-group">
            <input
              v-model="proxyURLInput"
              type="text"
              placeholder="http://127.0.0.1:7890"
              class="config-input"
            />
            <Button variant="primary" @click="saveProxySettings">
              保存
            </Button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>
.content-section {
  margin-bottom: 1.5rem;

  &:last-child {
    margin-bottom: 0;
  }
}

.section-title {
  font-size: 1.25rem;
  font-weight: 600;
  color: #1f2937;
  margin: 0 0 1rem 0;
  padding-bottom: 0.5rem;
  border-bottom: 1px solid #e5e7eb;
}

.setting-group {
  background: #ffffff;
  border-radius: 0.5rem;
  border: 1px solid #e5e7eb;
  overflow: hidden;
}

.setting-item {
  display: flex;
  align-items: center;
  padding: 0.875rem 1.25rem;
  border-bottom: 1px solid #f3f4f6;

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
  color: #374151;
  margin-bottom: 0.25rem;
}

.label-desc {
  display: block;
  font-size: 0.75rem;
  color: #6b7280;
  line-height: 1.4;
}

.setting-control {
  flex-shrink: 0;
  margin-left: 1rem;
}

.input-group {
  display: flex;
  align-items: stretch;
  gap: 0.75rem;
  min-width: 300px;
}

.config-input {
  flex: 1;
  min-width: 0;
}

.proxy-controls {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.toggle-switch {
  position: relative;
  width: 44px;
  height: 24px;
  background: #d1d5db;
  border-radius: 12px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.toggle-slider {
  position: absolute;
  top: 1.5px;
  left: 1.5px;
  width: 20px;
  height: 20px;
  background: #ffffff;
  border-radius: 50%;
  transition: transform 0.3s ease;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);

  &.active {
    transform: translateX(20px);
  }
}

.toggle-switch:has(.toggle-slider.active) {
  background: #005eec;
}

.toggle-label {
  font-size: 0.875rem;
  color: #374151;
  font-weight: 500;
}
</style>
