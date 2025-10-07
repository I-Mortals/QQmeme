<script setup lang="ts">
import { nextTick, ref, watch } from 'vue'
import { store } from '@/store'

interface MenuItem {
  icon: string
  label: string
  action: () => void
  separator?: boolean
}

const contextMenu = ref<HTMLElement>()

// 计算菜单位置
const calculateMenuPosition = (x: number, y: number) => {
  if (!contextMenu.value) {
    return { left: `${x}px`, top: `${y}px` }
  }

  const menuRect = contextMenu.value.getBoundingClientRect()
  const menuWidth = menuRect.width
  const menuHeight = menuRect.height

  if (x + menuWidth > window.innerWidth) {
    x = x - menuWidth
  }

  if (y + menuHeight > window.innerHeight) {
    y = y - menuHeight
  }

  return { left: `${x}px`, top: `${y}px` }
}

watch([() => store.contextMenu.visible, () => store.contextMenu.position], ([visible, position]) => {
  if (visible && position && contextMenu.value) {
    nextTick(()=>{
      if (!contextMenu.value) return
      const { left, top } = calculateMenuPosition(position.x, position.y)
      contextMenu.value.style.left = left
      contextMenu.value.style.top = top
    })
  }
}, { deep: true })

const hideMenu = () => {
  store.hideContextMenu()
}

const handleMenuItemClick = (action: () => void) => {
  action?.()
  hideMenu()
}

</script>

<template>
  <Teleport to="body">
    <Transition
      name="fade"
      appear>
      <div
        ref="contextMenu"
        class="context-menu"
        v-show="store.contextMenu.visible"
        v-click-outside="hideMenu">
        <template
          v-for="(item, index) in store.contextMenu.menuItems"
          :key="index">
          <!-- 分隔符 -->
          <div v-if="item.separator" class="menu-separator"></div>
          <!-- 菜单项 -->
          <div
            class="menu-item"
            @click="handleMenuItemClick(item.action)">
            <span class="menu-icon">{{ item.icon }}</span>
            <span>{{ item.label }}</span>
          </div>
        </template>
      </div>
    </Transition>
  </Teleport>
</template>

<style lang="less" scoped>
.context-menu {
  position: fixed;
  background: rgba(255, 255, 255);
  border-radius: 8px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.2);
  z-index: 1001;
  min-width: 120px;
  padding: 4px;
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.menu-item {
  padding: 6px;
  border-radius: 6px;
  color: #374151;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 16px;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  font-size: 14px;
  font-weight: 500;

  &:hover {
    background: rgba(59, 130, 246, 0.08);
    color: #1d4ed8;
  }
}

.menu-icon {
  width: 16px;
  text-align: center;
}

.menu-separator {
  height: 1px;
  background: rgba(0, 0, 0, 0.1);
  margin: 4px 8px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease-out;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
