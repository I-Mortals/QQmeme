<script setup lang="ts">
import { computed, useSlots } from 'vue'
import { Icon } from '@iconify/vue'

interface ButtonProps {
  variant?: 'primary' | 'secondary' | 'danger' | 'success' | 'warning'
  disabled?: boolean
  loading?: boolean
  block?: boolean
  icon?: string
  iconPosition?: 'left' | 'right'
}

const props = withDefaults(defineProps<ButtonProps>(), {
  variant: 'primary',
  disabled: false,
  loading: false,
  block: false,
  iconPosition: 'left'
})

const emit = defineEmits<{
  click: [event: MouseEvent]
}>()

const slots = useSlots()

const buttonClasses = computed(() => {
  return [
    'btn',
    `btn-${props.variant}`,
    {
      'btn-block': props.block,
      'btn-loading': props.loading,
      'btn-disabled': props.disabled || props.loading,
      'btn-icon-only': !slots.default && props.icon,
      'btn-with-icon': slots.default && props.icon
    }
  ]
})

const handleClick = (event: MouseEvent) => {
  if (!props.disabled && !props.loading) {
    emit('click', event)
  }
}
</script>

<template>
  <button
    :class="buttonClasses"
    :disabled="disabled || loading"
    @click="handleClick"
  >
    <span v-if="icon && iconPosition === 'left'" class="btn-icon btn-icon-left">
      <Icon :icon="icon" :width="18" :height="18"/>
    </span>

    <span v-if="loading" class="btn-loading-spinner"></span>

    <span v-if="slots.default" class="btn-content">
      <slot></slot>
    </span>

    <span v-if="icon && iconPosition === 'right'" class="btn-icon btn-icon-right">
      {{ icon }}
    </span>
  </button>
</template>

<style lang="less" scoped>.button-variant(@type, @bg, @text, @shadow: 0.1, @hover-shadow: 0.2) {
  .btn-@{type} {
    background: @bg;
    color: @text;
    box-shadow: 0 4px 6px fade(@bg, percentage(@shadow));

    &:hover:not(.btn-disabled) {
      background: darken(@bg, 10%);
      transform: translateY(-2px);
      box-shadow: 0 8px 15px fade(@bg, percentage(@hover-shadow));
    }

    &:active:not(.btn-disabled) {
      transform: translateY(0);
    }

    &:focus-visible {
      outline: 2px solid @bg;
      outline-offset: 2px;
    }
  }
}
.button-variant(primary, #005eec, white, 0.1, 0.2);
.button-variant(secondary, #f1f5f9, #475569, 0, 0);
.button-variant(danger, #ef4444, white, 0.3, 0.4);
.button-variant(success, #10b981, white, 0.3, 0.4);
.button-variant(warning, #f59e0b, white, 0.3, 0.4);

.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.375rem;
  border: none;
  border-radius: 0.375rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  text-decoration: none;
  outline: none;
  position: relative;
  overflow: hidden;
  white-space: nowrap;
  user-select: none;
  padding: 0.5rem 1rem;
  font-size: 0.875rem;
  min-height: 2rem;


  &-disabled {
    opacity: 0.6;
    cursor: not-allowed;
    transform: none !important;
    box-shadow: none !important;
  }

  &-loading {
    cursor: wait;
  }

  &-block {
    width: 100%;
  }

  &-icon {
    display: flex;
    align-items: center;
    justify-content: center;
  }

  &-icon-only {
    padding: 0.5rem;
    aspect-ratio: 1;

    &.btn-small {
      padding: 0.375rem;
    }

    &.btn-large {
      padding: 0.75rem;
    }
  }

  &-with-icon .btn-content {
    display: flex;
    align-items: center;
  }

  &-loading-spinner {
    width: 1rem;
    height: 1rem;
    border-top: 2px solid currentColor;
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

</style>
