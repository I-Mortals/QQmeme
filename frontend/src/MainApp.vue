<script lang="ts" setup>
import { ref } from 'vue'
import { Tab, TabPanel } from '@/components/tab'
import MemeBox from '@/pages/memes/MemeBox.vue'
import SettingPane from '@/pages/settings/SettingPane.vue'
import UsageGuide from '@/pages/UsageGuide.vue'
import type { TabItem } from '@/components/tab/types'

const mainTabsConfig = [
  {
    key: 'memes',
    label: '表情包',
    icon: '😀',
    component: MemeBox
  },
  {
    key: 'settings',
    label: '设置',
    icon: '⚙️',
    component: SettingPane
  },
  {
    key: 'usage',
    label: '使用说明',
    icon: '📖',
    component: UsageGuide
  }
]

const mainTabs: TabItem[] = mainTabsConfig.map(config => ({
  key: config.key,
  label: config.label,
  icon: config.icon
}))

const activeTab = ref('memes')
</script>

<template>
  <div class="main-app">
    <Tab
      v-model="activeTab"
      :tabs="mainTabs"
    >
      <TabPanel
        v-for="value in mainTabsConfig"
        :key="value.key"
        :tab-key="value.key"
        :active="activeTab === value.key"
      >
        <component :is="value.component" />
      </TabPanel>
    </Tab>
  </div>
</template>

<style lang="less" scoped>
.main-app {
  width: 100%;
  height: 100vh;
  overflow: hidden;
  padding: 0;
  margin: 0;
}
</style>
