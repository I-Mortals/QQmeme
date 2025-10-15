<script lang="ts" setup>
import { ref, watch } from 'vue'
import { ColorPicker as VueColorPicker } from 'vue-color-kit'
import 'vue-color-kit/dist/vue-color-kit.css'
import Popup from './popup/Popup.vue'

interface ColorPickerProps {
  value: string
}

interface Emits {
  (e: 'update:value', value: string): void
}

const props = defineProps<ColorPickerProps>()
const emit = defineEmits<Emits>()

const currentColor = ref(props.value)
const showPicker = ref(false)
const triggerRef = ref<HTMLElement>()

const changeColor = (color: any) => {
  currentColor.value = color.hex
  
  emit('update:value', currentColor.value)
}

const togglePicker = () => {
  showPicker.value = !showPicker.value
}

const closePicker = () => {
  showPicker.value = false
}

// 监听外部值变化
watch(() => props.value, (newValue) => {
  currentColor.value = newValue
}, { immediate: true })
</script>

<template>
  <div class="color-picker">
    <div ref="triggerRef" @click="togglePicker">
      <slot name="trigger" :currentColor="currentColor">
        <div class="color-preview-container">
          <div
            class="color-preview"
            :style="{ backgroundColor: currentColor }" />
        </div>
      </slot>
    </div>
    
    <Popup
      :visible="showPicker"
      :trigger="triggerRef"
      placement="bottom-start"
      :click-outside="true"
      @close="closePicker"
      @update:visible="(visible) => showPicker = visible">
      <div class="color-picker-container">
        <VueColorPicker
          style="box-sizing: content-box;"
          :color="currentColor"
          :sucker-hide="true"
          @changeColor="changeColor"
        />
      </div>
    </Popup>
  </div>
</template>
<style lang="less" scoped>
@import '@/styles/variables.less';

.color-picker {
  position: relative;
  display: inline-block;
}

.color-preview-container {
  position: relative;
  width: 1.5rem;
  height: 1.5rem;
  border-radius: 0.25rem;
  border: 1px solid rgba(0, 0, 0, 0.2);
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.08), 0 1px 2px rgba(0, 0, 0, 0.04);
  background-image: url("data:image/svg+xml,%3csvg width='100' height='100' xmlns='http://www.w3.org/2000/svg'%3e%3cdefs%3e%3cpattern id='a' patternUnits='userSpaceOnUse' width='10' height='10'%3e%3crect width='10' height='10' fill='%23fff'/%3e%3crect width='5' height='5' fill='%23ccc'/%3e%3crect x='5' y='5' width='5' height='5' fill='%23ccc'/%3e%3c/pattern%3e%3c/defs%3e%3crect width='100' height='100' fill='url(%23a)'/%3e%3c/svg%3e");
  overflow: hidden;
  cursor: pointer;
  transition: all 0.2s ease;

  &:hover {
    transform: scale(1.05);
  }
}

.color-preview {
  width: 100%;
  height: 100%;
  border-radius: 0.25rem;
  transition: transform 0.1s ease;
}

.color-picker-container {
  display: flex;
  flex-direction: column;
  overflow: visible;
}

.popup-body {
  padding: 0;
}
</style>
