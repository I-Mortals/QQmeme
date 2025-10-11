import type { Directive, DirectiveBinding } from 'vue'

// 点击外部区域指令
export const clickOutside: Directive = {
  mounted(el: HTMLElement, binding: DirectiveBinding) {
    const handler = (event: Event) => {
      if (!el.contains(event.target as Node)) {
        if (typeof binding.value === 'function') {
          binding.value(event)
        }
      }
    }
    
    document.addEventListener('click', handler, true)
    el._clickOutsideHandler = handler
  },
  
  unmounted(el: HTMLElement) {
    if (el._clickOutsideHandler) {
      document.removeEventListener('click', el._clickOutsideHandler, true)
      delete el._clickOutsideHandler
    }
  }
}

declare global {
  interface HTMLElement {
    _clickOutsideHandler?: (event: Event) => void
  }
}
