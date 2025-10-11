export interface TabItem {
  /** tab的唯一标识 */
  key: string
  /** tab的标题 */
  label: string
  /** tab是否禁用 */
  disabled?: boolean
  /** tab的图标 */
  icon?: string
  /** 自定义数据 */
  data?: any
  /** tab对应的组件 */
  component?: any
  /** tab对应的组件props */
  props?: any
}

export interface TabProps {
  /** tab数据列表 */
  tabs: TabItem[]
  /** 当前激活的tab key */
  modelValue: string
  /** header位置 */
  headerPosition?: 'top' | 'bottom'
  /** 是否启用延迟加载 */
  lazy?: boolean
  /** 是否显示header */
  showHeader?: boolean
  /** 自定义class */
  class?: string
  /** 自定义style */
  style?: string | Record<string, any>
}

export interface TabHeaderProps {
  /** tab数据列表 */
  tabs: TabItem[]
  /** 当前激活的tab key */
  activeKey: string
  /** header位置 */
  position?: 'top' | 'bottom'
  /** 是否显示header */
  show?: boolean
  /** 左侧固定部分的插槽内容 */
  fixedLeftSlot?: boolean
  /** 右侧固定部分的插槽内容 */
  fixedRightSlot?: boolean
  /** 可滚动部分的插槽内容 */
  scrollableSlot?: boolean
  /** 整个header自定义插槽 */
  customSlot?: boolean
  /** 自定义class */
  class?: string
}

export interface TabPanelProps {
  /** tab的唯一标识 */
  tabKey: string
  /** 是否激活 */
  active: boolean
  /** 是否启用延迟加载 */
  lazy?: boolean
  /** 是否已加载过 */
  loaded?: boolean
  /** 自定义class */
  class?: string
}

export interface TabEmits {
  /** tab切换事件 */
  'update:modelValue': [key: string]
  /** tab点击事件 */
  'tab-click': [tab: TabItem, event: MouseEvent]
  /** tab改变事件 */
  'tab-change': [newKey: string, oldKey: string]
}
