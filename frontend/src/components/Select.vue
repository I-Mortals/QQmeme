<script lang="ts" setup>
import { ref, computed, onMounted, onUnmounted, nextTick, watch } from 'vue'
import { Icon } from '@iconify/vue'

export interface SelectOption {
  value: string | number
  label: string
  disabled?: boolean
}

interface Props {
  modelValue?: string | number
  options: SelectOption[]
  placeholder?: string
  disabled?: boolean
  clearable?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  modelValue: '',
  placeholder: '请选择',
  disabled: false,
  clearable: false
})

const emit = defineEmits<{
  'update:modelValue': [value: string | number]
  'change': [value: string | number, option: SelectOption]
}>()

const isOpen = ref(false)
const selectRef = ref<HTMLElement>()
const dropdownRef = ref<HTMLElement>()
const dropdownPosition = ref({ top: 0, left: 0, width: 0 })

const selectedOption = computed(() => {
  return props.options.find(option => option.value === props.modelValue)
})

const displayText = computed(() => {
  if (selectedOption.value) {
    return selectedOption.value.label
  }
  return props.placeholder
})

const selectClasses = computed(() => ({
  'select': true,
  'select-open': isOpen.value,
  'select-disabled': props.disabled
}))

const updateDropdownPosition = async () => {
  if (!selectRef.value) return
  
  await nextTick()
  const rect = selectRef.value.getBoundingClientRect()
  dropdownPosition.value = {
    top: rect.bottom + window.scrollY,
    left: rect.left + window.scrollX,
    width: rect.width
  }
}

const handleClick = async () => {
  if (props.disabled) return
  if (!isOpen.value) {
    await updateDropdownPosition()
    isOpen.value = true
  } else {
    isOpen.value = false
  }
}

const handleOptionClick = (option: SelectOption) => {
  if (option.disabled) return
  
  emit('update:modelValue', option.value)
  emit('change', option.value, option)
  isOpen.value = false
}

const handleClear = (e: Event) => {
  e.stopPropagation()
  emit('update:modelValue', '')
  emit('change', '', { value: '', label: '' })
}


const handleClickOutside = () => {
  isOpen.value = false
}

const handleResize = () => {
  if (isOpen.value) {
    updateDropdownPosition()
  }
}

const handleScroll = () => {
  if (isOpen.value) {
    updateDropdownPosition()
  }
}


onMounted(() => {
  window.addEventListener('resize', handleResize)
  window.addEventListener('scroll', handleScroll, true)
})

onUnmounted(() => {
  window.removeEventListener('resize', handleResize)
  window.removeEventListener('scroll', handleScroll, true)
})
</script>

<template>
  <div 
    ref="selectRef"
    :class="selectClasses"
    @click="handleClick"
    v-click-outside="handleClickOutside"
    tabindex="0"
  >
    <div class="select-trigger">
      <span class="select-value" :class="{ 'select-placeholder': !selectedOption }">
        {{ displayText }}
      </span>
      <div class="select-actions">
        <button
          v-if="clearable && selectedOption && !disabled"
          @click="handleClear"
          class="select-clear"
          type="button"
        >
          <Icon icon="mdi:close" />
        </button>
        <Icon 
          icon="mdi:chevron-down" 
          class="select-arrow" 
          :class="{ 'select-arrow-open': isOpen }"
        />
      </div>
    </div>
    
    <Teleport to="body">
      <Transition name="dropdown">
        <div 
          v-if="isOpen"
          ref="dropdownRef"
          class="select-dropdown"
          :style="{
            position: 'fixed',
            top: dropdownPosition.top + 'px',
            left: dropdownPosition.left + 'px',
            width: dropdownPosition.width + 'px',
            zIndex: 1000
          }"
        >
          <div class="select-options">
            <div
              v-for="option in options"
              :key="option.value"
              :class="{
                'select-option': true,
                'select-option-selected': option.value === modelValue,
                'select-option-disabled': option.disabled
              }"
              @click="handleOptionClick(option)"
            >
              {{ option.label }}
            </div>
          </div>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

.select {
  position: relative;
  display: inline-block;
  width: 100%;
  min-width: 120px;
  background: @rgb-b1;
  border: 1px solid @rgb-b3;
  border-radius: 0.375rem;
  cursor: pointer;
  transition: all 0.2s ease;
  outline: none;

  &:hover:not(.select-disabled) {
    border-color: @rgb-bc;
  }

  &:focus-within {
    border-color: @rgb-p;
    box-shadow: 0 0 0 2px rgba(@p, 0.1);
  }

  &.select-disabled {
    background: rgba(var(--bc), 0.06);
    color: @rgb-bc;
    opacity: 0.6;
    cursor: not-allowed;
  }

  &.select-open {
    border-color: @rgb-p;
    box-shadow: 0 0 0 2px rgba(@p, 0.1);
  }

}

.select-trigger {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
  color: @rgb-bc;
  user-select: none;
}

.select-value {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;

  &.select-placeholder {
    color: @rgb-bc;
    opacity: 0.6;
  }
}

.select-actions {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-left: 0.5rem;
}

.select-clear {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 1rem;
  height: 1rem;
  border: none;
  background: @rgb-b3;
  color: @rgb-bc;
  border-radius: 50%;
  cursor: pointer;
  font-size: 0.75rem;
  line-height: 1;
  transition: all 0.2s ease;

  &:hover {
    background: @rgb-bc;
    color: @rgb-b1;
  }
}

.select-arrow {
  width: 1rem;
  height: 1rem;
  color: @rgb-bc;
  transition: transform 0.2s ease;

  &.select-arrow-open {
    transform: rotate(180deg);
  }
}

.select-dropdown {
  margin-top: 0.25rem;
  background: @rgb-b1;
  border: 1px solid @rgb-b3;
  border-radius: 0.375rem;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  overflow: hidden;
}

.select-options {
  max-height: 200px;
  overflow-y: auto;
}

.select-option {
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
  color: @rgb-bc;
  cursor: pointer;
  transition: background-color 0.2s ease;
  user-select: none;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;

  &:hover:not(.select-option-disabled) {
    background: rgba(var(--bc), 0.06);
  }

  &.select-option-selected {
    background: rgba(@p, 0.1);
    color: @rgb-p;
    font-weight: 500;
  }

  &.select-option-disabled {
    color: @rgb-bc;
    opacity: 0.5;
    cursor: not-allowed;
  }
}

// 下拉动画
.dropdown-enter-active,
.dropdown-leave-active {
  transition: all 0.2s ease;
  transform-origin: top;
}

.dropdown-enter-from,
.dropdown-leave-to {
  opacity: 0;
  transform: scaleY(0.8);
}
</style>
