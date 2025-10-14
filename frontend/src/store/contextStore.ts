import { reactive } from 'vue'

export interface ContextMenuItem {
  icon: string
  label: string
  action: () => void
  separator?: boolean
}

export interface ContextMenuState {
  visible: boolean
  menuItems: ContextMenuItem[]
  position: { x: number; y: number }
}

export interface ContextStore {
  contextMenu: ContextMenuState
  showContextMenu: (e: MouseEvent, menuItems: ContextMenuItem[]) => void
  hideContextMenu: () => void
}

export const contextStore = reactive<ContextStore>({
  contextMenu: {
    visible: false,
    menuItems: [],
    position: { x: 0, y: 0 }
  },
  
  showContextMenu(e: MouseEvent, menuItems: ContextMenuItem[]) {
    e.preventDefault()
    this.contextMenu.menuItems = menuItems
    this.contextMenu.position = { x: e.clientX, y: e.clientY }
    this.contextMenu.visible = true
  },
  
  hideContextMenu() {
    this.contextMenu.visible = false
    setTimeout(() => {
      this.contextMenu.menuItems = []
    }, 200)
  }
})
