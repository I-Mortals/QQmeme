<script lang="ts" setup>
import { computed, ref } from 'vue'
import ThemeSwitch from '@/components/common/ThemeSwitch.vue';
import SettingItem from './setting/SettingItem.vue'
import SettingGroup from './setting/SettingGroup.vue'
import SettingsSection from './setting/SettingsSection.vue'
import ColorPicker from '@/components/ColorPicker.vue'
import { themeStore } from '@/store/themeStore'
import { type ThemeColorOptions } from '@/styles/themes'
import { parseRgbStr } from '@/utils/color'
import Button from '@/components/Button.vue'

const themeVariables = [
  { key: 'base1', label: '背景色', desc: '应用的主要背景颜色' },
  { key: 'baseContent', label: '文本色', desc: '主要文本内容的颜色' },
  { key: 'primary', label: '主色调', desc: '应用的主要强调色' },
  { key: 'info', label: '信息色', desc: '信息提示的颜色' },
  { key: 'success', label: '成功色', desc: '成功状态的颜色' },
  { key: 'warning', label: '警告色', desc: '警告状态的颜色' },
  { key: 'error', label: '错误色', desc: '错误状态的颜色' }
] as const

const customTheme = ref<Partial<ThemeColorOptions>>({})

const currentThemeConfig = computed(() => {
  return themeStore.getThemeConfig()
})

const keyToCssVar: Record<string, string> = {
  base1: 'b1',
  baseContent: 'bc',
  primary: 'p',
  info: 'i',
  success: 's',
  warning: 'w',
  error: 'e'
}

const getThemeVariableValue = (key: string) => {
  if (customTheme.value[key as keyof ThemeColorOptions]) {
    return customTheme.value[key as keyof ThemeColorOptions]!
  }
  const config = currentThemeConfig.value
  return config[key as keyof ThemeColorOptions] || '#000000'
}

const updateThemeVariable = (key: string, value: string) => {
  customTheme.value = {
    ...customTheme.value,
    [key]: value
  }

  const cssVar = keyToCssVar[key]
  if (cssVar) {
    // 主题使用 rgb 如 --p: 0,102,255
    const rgb = parseRgbStr(value)
    document.documentElement.style.setProperty(`--${cssVar}`, rgb)
  }
}

const resetThemeVariable = (key: string) => {
  const originalValue = currentThemeConfig.value[key as keyof ThemeColorOptions]
  if (originalValue) {
    delete customTheme.value[key as keyof ThemeColorOptions]
    updateThemeVariable(key, originalValue)
  }
}

const resetAllCustomTheme = () => {
  customTheme.value = {}
  Object.values(keyToCssVar).forEach((cssVar) => {
    document.documentElement.style.removeProperty(`--${cssVar}`)
  })
  themeStore.setTheme(themeStore.currentTheme)
}
</script>

<template>
  <SettingsSection title="外观设置" section-id="appearance-settings">
    <SettingGroup>
      <SettingItem>
        <template #text>主题方案</template>
        <template #desc>选择应用的主题配色方案</template>
        <template #actions>
          <div class="theme-control">
            <ThemeSwitch />
          </div>
        </template>
      </SettingItem>
    </SettingGroup>


    <SettingGroup>
      <template #title>主题变量</template>
      <template #desc>自定义当前主题的关键颜色变量</template>

      <SettingItem v-for="item in themeVariables" :key="item.key">
        <template #text>{{ item.label }}</template>
        <template #desc>{{ item.desc }}</template>
        <template #actions>
          <div class="color-trigger">
            <ColorPicker :value="getThemeVariableValue(item.key)"
              @update:value="(val: string) => updateThemeVariable(item.key, val)">
              <template #trigger="{ currentColor }">
                <div class="color-preview-wrapper" :title="currentColor">
                  <div class="color-preview" :style="{ backgroundColor: currentColor }" />
                </div>
              </template>
            </ColorPicker>
            <Button variant="primary" @click="resetThemeVariable(item.key)">重置</Button>
          </div>
        </template>
      </SettingItem>

      <SettingItem>
        <template #actions>
          <Button variant="primary" @click="resetAllCustomTheme">重置全部</Button>
        </template>
      </SettingItem>
    </SettingGroup>
  </SettingsSection>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

.theme-control {
  display: flex;
  align-items: center;
  gap: .5rem;
}

.color-trigger {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.color-preview-wrapper {
  position: relative;
  width: 1.5rem;
  height: 1.5rem;
  border-radius: 0.25rem;
  border: 1px solid rgba(0, 0, 0, 0.2);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  background-image: url("data:image/svg+xml,%3csvg width='100' height='100' xmlns='http://www.w3.org/2000/svg'%3e%3cdefs%3e%3cpattern id='a' patternUnits='userSpaceOnUse' width='10' height='10'%3e%3crect width='10' height='10' fill='%23fff'/%3e%3crect width='5' height='5' fill='%23ccc'/%3e%3crect x='5' y='5' width='5' height='5' fill='%23ccc'/%3e%3c/pattern%3e%3c/defs%3e%3crect width='100' height='100' fill='url(%23a)'/%3e%3c/svg%3e");
  overflow: hidden;
  cursor: pointer;
  transition: all 0.2s ease;

  &:hover {
    transform: scale(1.05);
  }
}

.color-preview {
  width: 100%;
  height: 100%;
  border-radius: 0.25rem;
  transition: transform 0.1s ease;
}

.color-hex {
  font-family: 'SF Mono', 'Monaco', 'Inconsolata', 'Roboto Mono', monospace;
  font-size: 0.875rem;
  font-weight: 500;
  color: @rgb-bc;
  background: @rgb-b1;
  padding: 0.375rem 0.625rem;
  border-radius: 0.375rem;
  border: 1px solid @rgb-b3;
  min-width: 4.5rem;
  text-align: center;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}
</style>
