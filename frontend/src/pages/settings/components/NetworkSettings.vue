<script setup lang="ts">
import { ref } from 'vue'
import Button from '@/components/Button.vue'
import Input from '@/components/Input.vue'
import { applicationStore, toastStore } from '@/store'
import SettingItem from './setting/SettingItem.vue'
import SettingsSection from './setting/SettingsSection.vue'
import SettingGroup from './setting/SettingGroup.vue'

const botTokenInput = ref(applicationStore.botToken)
const proxyEnabled = ref(applicationStore.proxyEnabled)
const proxyURLInput = ref(applicationStore.proxyURL)

const saveBotToken = () => {
  const token = botTokenInput.value.trim()

  if (!token) {
    toastStore.showToast('Bot Token 不能为空', 'error')
    return
  }

  if (!/^\d+:[A-Za-z0-9_-]+$/.test(token)) {
    toastStore.showToast('Bot Token 格式不正确，应为 "数字:字符串" 格式', 'error')
    return
  }

  applicationStore.setBotToken(token)
  toastStore.showToast('Telegram Bot Token 保存成功！', 'success')
}

const saveProxySettings = () => {
  if (proxyEnabled.value) {
    const url = proxyURLInput.value.trim()

    if (!url) {
      toastStore.showToast('启用代理时，代理地址不能为空', 'error')
      return
    }

    try {
      new URL(url)
    } catch {
      toastStore.showToast('代理地址格式不正确，请输入有效的 URL', 'error')
      return
    }
  }

  applicationStore.setProxySettings(proxyEnabled.value, proxyURLInput.value.trim())
  toastStore.showToast('代理设置保存成功！', 'success')
}

const toggleProxy = () => {
  proxyEnabled.value = !proxyEnabled.value
  applicationStore.setProxySettings(proxyEnabled.value, applicationStore.proxyURL)
}
</script>

<template>
  <SettingsSection title="网络设置">
    <SettingGroup>
      <SettingItem>
        <template #text>Telegram Bot Token</template>
        <template #desc>用于获取Telegram贴纸包的机器人Token</template>
        <template #actions>
          <Input v-model="botTokenInput" type="text" placeholder="请输入Bot Token" class="config-input">
          <template #append>
            <Button variant="primary" @click="saveBotToken">
              保存
            </Button>
          </template>
          </Input>
        </template>
      </SettingItem>
      <SettingItem>
        <template #text>代理设置</template>
        <template #desc>配置网络代理以访问Telegram服务</template>
        <template #actions>
          <div class="proxy-controls">
            <div class="toggle-switch" @click="toggleProxy">
              <div class="toggle-slider" :class="{ active: proxyEnabled }"></div>
            </div>
            <span class="toggle-label">{{ proxyEnabled ? '已启用' : '已禁用' }}</span>
          </div>
        </template>
      </SettingItem>
      <SettingItem v-if="proxyEnabled">
        <template #text>代理地址</template>
        <template #desc>代理服务器的完整地址</template>
        <template #actions>
          <Input v-model="proxyURLInput" type="text" placeholder="http://127.0.0.1:7890" class="config-input">
          <template #append>
            <Button variant="primary" @click="saveProxySettings">
              保存
            </Button>
          </template>
          </Input>
        </template>
      </SettingItem>
    </SettingGroup>
  </SettingsSection>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

.config-input {
  min-width: 300px;
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
  background: @rgb-b3;
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
  background: @rgb-b1;
  border-radius: 50%;
  transition: transform 0.3s ease;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.08), 0 1px 2px rgba(0, 0, 0, 0.04);

  &.active {
    transform: translateX(20px);
  }
}

.toggle-switch:has(.toggle-slider.active) {
  background: @rgb-p;
}

.toggle-label {
  font-size: 0.875rem;
  color: @rgb-bc;
  font-weight: 500;
}
</style>
