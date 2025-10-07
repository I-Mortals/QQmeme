<script lang="ts" setup>
import { ref, onMounted, computed } from 'vue'
import FileManagement from './components/FileManagement.vue'
import NetworkSettings from './components/NetworkSettings.vue'
import AppearanceSettings from './components/AppearanceSettings.vue'

const settingsConfig = [
  {
    id: 'file-management',
    title: '文件管理',
    icon: '📁',
    component: FileManagement
  },
  {
    id: 'network-settings',
    title: '网络设置',
    icon: '🔗',
    component: NetworkSettings
  },
  {
    id: 'appearance-settings',
    title: '外观设置',
    icon: '🎨',
    component: AppearanceSettings
  }
]

const activeSection = ref('file-management')

const currentComponent = computed(() => {
  const config = settingsConfig.find(item => item.id === activeSection.value)
  return config?.component
})

const switchSection = (sectionId: string) => {
  activeSection.value = sectionId
}

onMounted(() => {
  activeSection.value = 'file-management'
})
</script>

<template>
  <div class="settings-page">
    <div class="settings-container">
      <!-- 左侧导航 -->
      <div class="settings-sidebar">
        <div
          v-for="item in settingsConfig"
          :key="item.id"
          class="nav-item"
          :class="{ active: activeSection === item.id }"
          @click="switchSection(item.id)">
          <span class="nav-icon">{{ item.icon }}</span>
          <span class="nav-text">{{ item.title }}</span>
        </div>
      </div>

      <!-- 右侧内容区 -->
      <div class="settings-content">
        <transition name="fade-slide" mode="out-in">
          <component
            :key="activeSection"
            :is="currentComponent" />
        </transition>
      </div>
    </div>
  </div>
</template>


<style lang="less" scoped>
.settings-page {
  height: 100%;
  background: #f5f5f5;
  display: flex;
  flex-direction: column;
}

.settings-container {
  display: flex;
  flex: 1;
  background: #ffffff;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.08);
  width: 100%;
  position: relative;
  min-height: 0;
}

.settings-sidebar {
  width: 180px;
  min-width: 180px;
  background: #ffffff;
  border-right: 1px solid #e5e7eb;
  padding: 0.5rem 0;
  flex-shrink: 0;
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  overflow-y: auto;
  z-index: 10;
}

.nav-item {
  display: flex;
  align-items: center;
  padding: 0.5rem 1rem;
  cursor: pointer;
  transition: all 0.2s ease;
  border-left: 3px solid transparent;
  color: #374151;

  &:hover {
    background: #f3f4f6;
    color: #1f2937;
  }

  &.active {
    background: #eff6ff;
    border-left-color: #3b82f6;
    color: #1d4ed8;
    font-weight: 600;
  }
}

.nav-icon {
  font-size: 1rem;
  margin-right: 0.5rem;
  opacity: 0.8;
}

.nav-item.active .nav-icon {
  opacity: 1;
}

.nav-text {
  font-size: 0.8rem;
  font-weight: 500;
}

.settings-content {
  flex: 1;
  padding: 1rem;
  overflow-y: auto;
  margin-left: 180px;
  min-height: 0;
  position: relative;
}

.fade-slide-enter-active,
.fade-slide-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.fade-slide-enter-from {
  opacity: 0;
  transform: translateX(-10px);
}

.fade-slide-leave-to {
  opacity: 0;
  transform: translateX(10px);
}

.fade-slide-enter-to,
.fade-slide-leave-from {
  opacity: 1;
  transform: translateX(0);
}



</style>
