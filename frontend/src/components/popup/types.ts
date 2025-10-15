import type { Placement, Strategy, Middleware } from '@floating-ui/vue'

export interface PopupProps {
  /** 是否显示popup */
  visible: boolean
  /** 触发元素 */
  trigger?: HTMLElement | null
  /** 弹出位置 */
  placement?: Placement
  /** 定位策略 */
  strategy?: Strategy
  /** 偏移量 */
  offset?: number
  /** 是否显示箭头 */
  showArrow?: boolean
  /** 是否可点击外部关闭 */
  clickOutside?: boolean
  /** 自定义中间件 */
  middleware?: Middleware[]
  /** 动画持续时间 */
  duration?: number
  /** 自定义类名 */
  className?: string
  /** 自定义样式 */
  customStyle?: Record<string, any>
  /** 尺寸 */
  size?: 'small' | 'medium' | 'large'
}

export interface PopupEmits {
  /** 关闭事件 */
  close: []
  /** 显示状态更新 */
  'update:visible': [value: boolean]
  /** 显示事件 */
  show: []
  /** 隐藏事件 */
  hide: []
}

export interface PopupExpose {
  /** 显示popup */
  show: () => void
  /** 隐藏popup */
  hide: () => void
  /** 切换显示状态 */
  toggle: () => void
  /** 更新位置 */
  updatePosition: () => void
}
