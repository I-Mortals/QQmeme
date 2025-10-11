<script setup lang="ts">
import { Icon } from '@iconify/vue'
import { computed, useSlots } from 'vue'

export interface InputProps {
  modelValue?: string | number
  type?: 'text' | 'password' | 'email' | 'number' | 'tel' | 'url' | 'search'
  placeholder?: string
  disabled?: boolean
  readonly?: boolean
  required?: boolean
  prefix?: string
  suffix?: string
  prepend?: string
  append?: string
  clearable?: boolean
}

const props = withDefaults(defineProps<InputProps>(), {
  type: 'text',
  clearable: true
})

const emit = defineEmits<{
  'update:modelValue': [value: string | number]
  'focus': [event: FocusEvent]
  'blur': [event: FocusEvent]
  'input': [event: Event]
  'change': [event: Event]
  'keydown': [event: KeyboardEvent]
  'keyup': [event: KeyboardEvent]
  'clear': []
}>()

const slots = useSlots()

const inputValue = computed({
  get: () => props.modelValue ?? '',
  set: (value: string | number) => {
    emit('update:modelValue', value)
  }
})

const inputClasses = computed(() => [
  'input-field',
  {
    'input-field-disabled': props.disabled,
    'input-field-readonly': props.readonly,
    'input-field-with-prefix': props.prefix || slots.prefix,
    'input-field-with-suffix': props.suffix || slots.suffix
  }
])


const hasPrepend = computed(() => props.prepend || slots.prepend)
const hasAppend = computed(() => props.append || slots.append)

const showClearButton = computed(() => {
  return props.clearable && !props.disabled && !props.readonly && inputValue.value
})


const handleInput = (event: Event) => {
  const target = event.target as HTMLInputElement
  inputValue.value = target.value
  emit('input', event)
}

const handleChange = (event: Event) => {
  emit('change', event)
}

const handleFocus = (event: FocusEvent) => {
  emit('focus', event)
}

const handleBlur = (event: FocusEvent) => {
  emit('blur', event)
}

const handleKeydown = (event: KeyboardEvent) => {
  emit('keydown', event)
}

const handleKeyup = (event: KeyboardEvent) => {
  emit('keyup', event)
}

const handleClear = () => {
  inputValue.value = ''
  emit('clear')
}
</script>

<template>
  <div class="input-wrapper" :class="{
    'input-with-suffix': props.suffix || slots.suffix,
    'input-with-clear': showClearButton
  }">
    <!-- 外部前缀 -->
    <div v-if="hasPrepend" class="input-prepend">
      <span v-if="prepend" class="input-prepend-text">{{ prepend }}</span>
      <slot v-else name="prepend"/>
    </div>

    <div class="input-inner">
      <!-- 内部前缀 -->
      <div v-if="prefix || slots.prefix" class="input-prefix">
        <span v-if="prefix" class="input-prefix-text">{{ prefix }}</span>
        <slot v-else name="prefix"/>
      </div>

      <div class="input-content">
        <!-- 输入框 -->
        <input
          v-model="inputValue"
          :type="type"
          :placeholder="placeholder"
          :disabled="disabled"
          :readonly="readonly"
          :required="required"
          :class="inputClasses"
          @input="handleInput"
          @change="handleChange"
          @focus="handleFocus"
          @blur="handleBlur"
          @keydown="handleKeydown"
          @keyup="handleKeyup"
        />
      </div>

      <!-- 清除按钮 -->
      <button
        v-if="showClearButton"
        type="button"
        class="input-clear"
        @click="handleClear"
      >
        <Icon icon="lucide:x"/>
      </button>


      <!-- 内部后缀 -->
      <div v-if="suffix || slots.suffix" class="input-suffix">
        <span v-if="suffix" class="input-suffix-text">{{ suffix }}</span>
        <slot v-else name="suffix"/>
      </div>
    </div>

    <!-- 外部后缀 -->
    <div v-if="hasAppend" class="input-append">
      <span v-if="append" class="input-append-text">{{ append }}</span>
      <slot v-else name="append"/>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import '@/styles/variables.less';

.input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
  width: 100%;
}

.input-inner {
  position: relative;
  flex: 1;
}

.input-content {
  position: relative;
  flex: 1;
}

.input-field {
  width: 100%;
  border: 1px solid @rgb-b3;
  border-radius: 0.375rem;
  font-size: 0.875rem;
  font-family: inherit;
  background: @rgb-b1;
  color: @rgb-bc;
  transition: all 0.2s ease;
  outline: none;
    padding: 0.5rem 0.75rem;
    font-size: 0.875rem;

  &::placeholder {
    color: @rgb-bc;
    opacity: 0.5;
  }

  &:focus {
    border-color: @rgb-p;
    box-shadow: 0 0 0 2px rgba(@p, 0.1);
  }

  &:disabled {
    background: rgba(var(--bc), 0.06);
    color: @rgb-bc;
    opacity: 0.5;
    cursor: not-allowed;
  }

  &:read-only {
    background: rgba(var(--bc), 0.04);
    cursor: default;
  }


  // 前缀后缀样式
  &.input-field-with-prefix {
    padding-left: 2.5rem;
  }

  &.input-field-with-suffix {
    padding-right: 2.5rem;
  }
}

.input-prefix,
.input-suffix {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  display: flex;
  align-items: center;
  color: @rgb-bc;
  font-size: 0.875rem;
  pointer-events: none;
  z-index: 1;
}

.input-prefix {
  left: 0.75rem;
}

.input-suffix {
  right: 0.75rem;
}

.input-prefix-text,
.input-suffix-text {
  font-weight: 500;
}

.input-clear {
  position: absolute;
  right: 0.25rem;
  top: 50%;
  transform: translateY(-50%);
  display: flex;
  align-items: center;
  justify-content: center;
  padding: .25rem;
  border: none;
  background: none;
  color: @rgb-bc;
  opacity: 0.6;
  cursor: pointer;
  border-radius: 0.25rem;
  transition: all 0.2s ease;
  z-index: 2;

  &:hover {
    opacity: 1;
    background: rgba(var(--bc), 0.1);
  }

  &:active {
    transform: translateY(-50%) scale(0.95);
  }
}


.input-prepend,
.input-append {
  display: flex;
  align-items: center;
}

.input-prepend {
  margin-right: 0.5rem;
}

.input-append {
  margin-left: 0.5rem;
}

.input-prepend-text,
.input-append-text {
  font-size: 0.875rem;
  color: @rgb-bc;
  font-weight: 500;
  white-space: nowrap;
}

.input-with-clear .input-field {
  padding-right: 1.5rem;
}

.input-with-suffix.input-with-clear .input-suffix {
  right: 2.5rem;
}

</style>
