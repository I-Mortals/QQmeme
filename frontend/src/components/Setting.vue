<script lang="ts" setup>
import { defineProps, withDefaults } from 'vue'
import { GenerateAllMemePath, SelectRootDir } from '../../wailsjs/go/memeFile/memeFile'
import { store } from '../store'
import Modal from './Modal.vue'
import SetColor from './SetColor.vue'

interface SettingProps {
  visible?: boolean;
  closeCall?: () => void;
}

const props = withDefaults(defineProps<SettingProps>(), {
  visible: false,
  closeCall: () => { }
})

const selectRoot = async () => {
  const path = await SelectRootDir()
  if (path) {
    store.rootPath = path
    store.allMemesPath = await GenerateAllMemePath(store.rootPath)
  }
}

const closeSetting = () => {
  props.closeCall()
}

const clearCache = () => {
  store.clearCache()
  store.rootPath = ''
}
</script>

<template>
  <Modal 
    :visible="props.visible" 
    title="设置"
    size="medium"
    @close="closeSetting"
  >
    <div class="setting-content">
      <div class="setting-section">
        <h3 class="section-title">目录设置</h3>
        <div class="setting-item">
          <label class="setting-label">MEME主目录</label>
          <div class="path-display">
            <span class="path-text">{{ store.rootPath || "请选择主目录" }}</span>
            <button class="select-btn" @click="selectRoot">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                <polyline points="7,10 12,15 17,10"></polyline>
                <line x1="12" y1="15" x2="12" y2="3"></line>
              </svg>
              选择目录
            </button>
          </div>
        </div>
      </div>
      
      <div class="setting-section">
        <h3 class="section-title">缓存管理</h3>
        <div class="setting-item">
          <button class="action-btn danger" @click="clearCache">
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <polyline points="3,6 5,6 21,6"></polyline>
              <path d="M19,6v14a2,2 0 0,1 -2,2H7a2,2 0 0,1 -2,-2V6m3,0V4a2,2 0 0,1 2,-2h4a2,2 0 0,1 2,2v2"></path>
            </svg>
            清除图片内存缓存
          </button>
        </div>
      </div>

      <div class="setting-section">
        <h3 class="section-title">主题设置</h3>
        <div class="setting-item">
          <SetColor></SetColor>
        </div>
      </div>



    </div>
    
    <template #footer>
      <button class="btn secondary" @click="closeSetting">关闭</button>
    </template>
  </Modal>
</template>


<style scoped>
.setting-content {
  display: flex;
  flex-direction: column;
  gap: 0;
}

.setting-section {
  margin-bottom: 2rem;
  position: relative;
}

.setting-section:last-child {
  margin-bottom: 0;
}

.setting-section::before {
  content: '';
  position: absolute;
  top: -1rem;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, 
    transparent 0%, 
    rgba(102, 126, 234, 0.2) 50%, 
    transparent 100%);
}

.section-title {
  font-size: 1.125rem;
  font-weight: 700;
  color: #1f2937;
  margin: 0 0 1.25rem 0;
  padding-bottom: 0.625rem;
  border-bottom: 3px solid transparent;
  text-align: center;
  position: relative;
}

.section-title::after {
  content: '';
  position: absolute;
  bottom: -3px;
  left: 0;
  width: 100%;
  height: 1px;
  background: linear-gradient(90deg, 
    transparent 0%, 
    rgba(102, 126, 234, 0.3) 50%, 
    transparent 100%);
}

.setting-item {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.setting-label {
  font-weight: 600;
  color: #374151;
  font-size: 0.9375rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.setting-label::before {
  content: '';
  width: 4px;
  height: 4px;
  background: rgb(var(--theme-primary));
  border-radius: 50%;
}

.path-display {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  background: #f8fafc;
  border: 2px solid #e2e8f0;
  border-radius: 0.75rem;
  padding: 0.75rem;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.path-display:hover {
  border-color: #cbd5e1;
  background: #f1f5f9;
  transform: translateY(-2px);
  box-shadow: 
    0 8px 25px rgba(0, 0, 0, 0.08),
    0 4px 12px rgba(var(--theme-primary), 0.1);
}

.path-text {
  flex: 1;
  color: #64748b;
  font-size: 0.875rem;
  word-break: break-all;
  line-height: 1.5;
  font-family: 'SF Mono', 'Monaco', 'Inconsolata', 'Roboto Mono', monospace;
  background: rgba(255, 255, 255, 0.7);
  padding: 0.5rem 0.75rem;
  border-radius: 0.5rem;
  border: 1px solid rgba(var(--theme-primary), 0.1);
}

.select-btn {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  background: rgb(var(--theme-primary));
  color: white;
  border: none;
  border-radius: 0.625rem;
  padding: 0.625rem 1rem;
  font-size: 0.875rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  white-space: nowrap;
  box-shadow: 0 4px 15px rgba(var(--theme-primary), 0.3);
}

.select-btn:hover {
  transform: translateY(-3px) scale(1.02);
  box-shadow: 
    0 8px 25px rgba(var(--theme-primary), 0.4),
    0 4px 12px rgba(var(--theme-primary), 0.3);
}

.select-btn:active {
  transform: translateY(-1px) scale(0.98);
  transition: all 0.1s ease;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 0.625rem;
  background: #ef4444;
  color: white;
  border: none;
  border-radius: 0.625rem;
  padding: 0.75rem 1.25rem;
  font-size: 0.875rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  width: 100%;
  justify-content: center;
  box-shadow: 0 4px 15px rgba(239, 68, 68, 0.3);
}

.action-btn:hover {
  background: #dc2626;
  transform: translateY(-3px) scale(1.02);
  box-shadow: 
    0 8px 25px rgba(239, 68, 68, 0.4),
    0 4px 12px rgba(220, 38, 38, 0.3);
}

.action-btn:active {
  transform: translateY(-1px) scale(0.98);
  transition: all 0.1s ease;
}

.btn {
  padding: 0.75rem 1.5rem;
  border-radius: 0.625rem;
  font-size: 0.875rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  border: none;
}

.btn.secondary {
  background: #6b7280;
  color: white;
  box-shadow: 0 4px 15px rgba(107, 114, 128, 0.3);
}

.btn.secondary:hover {
  background: #4b5563;
  transform: translateY(-3px) scale(1.02);
  box-shadow: 
    0 8px 25px rgba(107, 114, 128, 0.4),
    0 4px 12px rgba(75, 85, 99, 0.3);
}

.btn.secondary:active {
  transform: translateY(-1px) scale(0.98);
  transition: all 0.1s ease;
}

/* 响应式设计 */
@media (max-width: 640px) {
  .setting-section {
    margin-bottom: 2rem;
  }
  
  .section-title {
    font-size: 1.125rem;
    margin-bottom: 1.25rem;
  }
  
  .path-display {
    flex-direction: column;
    align-items: stretch;
    gap: 0.75rem;
  }
  
  .select-btn {
    justify-content: center;
    padding: 0.875rem 1rem;
  }
  
  .action-btn {
    padding: 0.875rem 1rem;
  }
  
  .btn {
    padding: 0.75rem 1.5rem;
  }
}

@media (max-width: 480px) {
  .setting-content {
    gap: 0.5rem;
  }
  
  .setting-section {
    margin-bottom: 1.5rem;
  }
  
  .section-title {
    font-size: 1rem;
    margin-bottom: 1rem;
  }
  
  .path-display {
    padding: 0.75rem;
  }
  
  .path-text {
    font-size: 0.8125rem;
    padding: 0.375rem 0.5rem;
  }
}
</style>
