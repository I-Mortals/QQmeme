<script lang="ts" setup>
import { computed, nextTick, onMounted, ref, watch } from 'vue'
import { autoUpdate, flip, offset, useFloating } from '@floating-ui/vue'
import { Hsl, Rgb, convertHslToRgb } from 'culori'
import {
  WHITE_RGB,
  BLACK_RGB,
  hexToHsl,
  parseHsl,
  hslToHex,
  hslToHex8,
  hslToRgb,
  parseRgb,
  colorToHsl,
  colorToRgb,
  clampValue
} from '@/utils/color'

interface ColorPickerProps {
  value: string
}

interface Emits {
  (e: 'update:value', value: string): void
}

type HslState = Omit<Hsl, 'mode'>

type RgbState = Omit<Rgb, 'mode'>

type SliderPositions = Omit<Hsl, 'mode'>

type DragType = 'h' | 'sl' | 'alpha' | null

const props = defineProps<ColorPickerProps>()
const emit = defineEmits<Emits>()

const hsl = ref<HslState>({
  h: 0,
  s: 100,
  l: 50,
  alpha: 1
})

const rgb = ref<RgbState>({
  r: 0,
  g: 0,
  b: 0
})

const showPicker = ref(false)
const buttonRef = ref<HTMLElement>()
const popupRef = ref<HTMLElement>()

const isDragging = ref(false)
const dragType = ref<DragType>(null)

const sliderPositions = ref<SliderPositions>({
  h: 0,
  s: 0,
  l: 0,
  alpha: 0
})
const isFirstRender = ref(true)

const hSliderRef = ref<HTMLElement>()
const slAreaRef = ref<HTMLElement>()
const alphaSliderRef = ref<HTMLElement>()

const { floatingStyles, placement, isPositioned } = useFloating(buttonRef, popupRef, {
  placement: 'bottom-start',
  middleware: [offset(8), flip()],
  whileElementsMounted: autoUpdate,
  open: showPicker
})


const currentColor = computed(() => {
  const { h, s, l, alpha } = hsl.value
  const alphaValue = alpha ?? 1
  // 透明度小于1显示hex8格式，否则显示hex格式
  return alphaValue < 1 ? hslToHex8(h, s, l, alphaValue) : hslToHex(h, s, l)
})

// 色相 h
const hBackground = computed(() => 'linear-gradient(to right, hsl(0, 100%, 50%), hsl(60, 100%, 50%), hsl(120, 100%, 50%), hsl(180, 100%, 50%), hsl(240, 100%, 50%), hsl(300, 100%, 50%), hsl(360, 100%, 50%))')

// 饱和度 s、亮度 l
const slBackground = computed(() => {
  const normalizedHue = hsl.value.h || 0 % 360
  return `
    linear-gradient(to right, white 0%, hsl(${normalizedHue}, 100%, 50%) 100%),
    linear-gradient(to bottom, transparent 0%, black 100%)
  `
})

// 透明度滑块背景
const alphaBackground = computed(() => {
  const { r, g, b } = rgb.value
  return `
    linear-gradient(to right,
      rgba(${r}, ${g}, ${b}, 0) 0%,
      rgba(${r}, ${g}, ${b}, 1) 100%
    ),
    url("data:image/svg+xml,%3csvg width='100' height='100' xmlns='http://www.w3.org/2000/svg'%3e%3cdefs%3e%3cpattern id='a' patternUnits='userSpaceOnUse' width='10' height='10'%3e%3crect width='10' height='10' fill='%23fff'/%3e%3crect width='5' height='5' fill='%23ccc'/%3e%3crect x='5' y='5' width='5' height='5' fill='%23ccc'/%3e%3c/pattern%3e%3c/defs%3e%3crect width='100' height='100' fill='url(%23a)'/%3e%3c/svg%3e")
  `
})

const updateRgb = (): void => {
  const { h, s, l } = hsl.value
  rgb.value = parseRgb(hslToRgb(h, s, l))
}

watch(
  hsl,
  () => {
    updateRgb()
    emit('update:value', currentColor.value)
  },
  { deep: true }
)

watch(isPositioned, (newValue) => {
  if (newValue && isFirstRender.value) {
    nextTick(() => {
      updateSliderPositions()
      isFirstRender.value = false
    })
  }
})

const handleRgbInput = (channel: keyof RgbState, value: number): void => {
  rgb.value[channel] = clampValue(value, 0, 255)

  const { r, g, b } = rgb.value
  const hslColor = hexToHsl(`rgb(${r}, ${g}, ${b})`)
  hsl.value = parseHsl(hslColor)
  updateSliderPositions()
}

const handleHslInput = (channel: keyof HslState, value: number): void => {
  const max = channel === 'h' ? 360 : 100
  hsl.value[channel] = clampValue(value, 0, max)
  updateSliderPositions()
}

const handleAlphaInput = (value: number): void => {
  hsl.value.alpha = clampValue(value / 100, 0, 1)
  updateSliderPositions()
}

/**
 * 更新滑块位置
 */
const updateSliderPositions = (): void => {
  const hueEl = hSliderRef.value
  const slEl = slAreaRef.value
  const alphaEl = alphaSliderRef.value

  if (!hueEl || !slEl || !alphaEl) return

  const hueRect = hueEl.getBoundingClientRect()
  const slRect = slEl.getBoundingClientRect()
  const alphaRect = alphaEl.getBoundingClientRect()

  if (hueRect.width > 0 && slRect.width > 0 && slRect.height > 0 && alphaRect.width > 0) {
    sliderPositions.value = {
      h: ((hsl.value.h || 0) / 360) * hueRect.width,
      s: (hsl.value.s / 100) * slRect.width,
      l: (1 - hsl.value.l / 100) * slRect.height,
      alpha: (hsl.value.alpha ?? 1) * alphaRect.width
    }
  }
}


const updateHueFromEvent = (event: MouseEvent): void => {
  const hueEl = hSliderRef.value
  if (!hueEl) return

  const rect = hueEl.getBoundingClientRect()
  const x = clampValue(event.clientX - rect.left, 0, rect.width)

  sliderPositions.value.h = x

  const percentage = x / rect.width
  hsl.value.h = Math.round(percentage * 360)
}

// 拖拽
const createDragHandlers = (type: DragType, updateFn: (event: MouseEvent) => void) => {
  const handleMouseMove = (e: MouseEvent): void => {
    if (isDragging.value && dragType.value === type) {
      e.preventDefault()
      updateFn(e)
    }
  }

  const handleMouseUp = (e: MouseEvent): void => {
    e.preventDefault()
    isDragging.value = false
    dragType.value = null
    document.removeEventListener('mousemove', handleMouseMove)
    document.removeEventListener('mouseup', handleMouseUp)
  }

  return { handleMouseMove, handleMouseUp }
}

const handleHueMouseDown = (event: MouseEvent): void => {
  event.preventDefault()
  isDragging.value = true
  dragType.value = 'h'
  updateHueFromEvent(event)

  const { handleMouseMove, handleMouseUp } = createDragHandlers('h', updateHueFromEvent)

  document.addEventListener('mousemove', handleMouseMove)
  document.addEventListener('mouseup', handleMouseUp)
}

// 计算渐变区域指定位置的颜色
const calculateColorAtPosition = (x: number, y: number, hue: number): { r: number; g: number; b: number; alpha?: number } => {
  const hueColor = convertHslToRgb({
    h: hue,
    s: 1,
    l: 0.5
  })

  // 水平插值：白色 -> 色相颜色
  const baseColor = {
    r: WHITE_RGB.r + (hueColor.r - WHITE_RGB.r) * x,
    g: WHITE_RGB.g + (hueColor.g - WHITE_RGB.g) * x,
    b: WHITE_RGB.b + (hueColor.b - WHITE_RGB.b) * x
  }

  // 垂直插值：基础颜色 -> 黑色
  const finalColor = {
    r: baseColor.r + (BLACK_RGB.r - baseColor.r) * y,
    g: baseColor.g + (BLACK_RGB.g - baseColor.g) * y,
    b: baseColor.b + (BLACK_RGB.b - baseColor.b) * y
  }

  return parseRgb(finalColor as Rgb)
}

const updateSLFromEvent = (event: MouseEvent): void => {
  const slEl = slAreaRef.value
  if (!slEl) return

  const rect = slEl.getBoundingClientRect()
  const x = clampValue(event.clientX - rect.left, 0, rect.width)
  const y = clampValue(event.clientY - rect.top, 0, rect.height)

  sliderPositions.value.s = x
  sliderPositions.value.l = y

  const relativeX = x / rect.width
  const relativeY = y / rect.height

  const rgbColor = calculateColorAtPosition(relativeX, relativeY, hsl.value.h || 0)
  const rgbStr = `rgb(${rgbColor.r}, ${rgbColor.g}, ${rgbColor.b})`

  const hslColor = parseHsl(hexToHsl(rgbStr))

  hsl.value.l = hslColor.l
  hsl.value.s = hslColor.s
  rgb.value = rgbColor
}

const handleSLAreaMouseDown = (event: MouseEvent): void => {
  event.preventDefault()
  isDragging.value = true
  dragType.value = 'sl'
  updateSLFromEvent(event)

  const { handleMouseMove, handleMouseUp } = createDragHandlers('sl', updateSLFromEvent)

  document.addEventListener('mousemove', handleMouseMove)
  document.addEventListener('mouseup', handleMouseUp)
}

const updateAlphaFromEvent = (event: MouseEvent): void => {
  const alphaEl = alphaSliderRef.value
  if (!alphaEl) return

  const rect = alphaEl.getBoundingClientRect()
  const x = clampValue(event.clientX - rect.left, 0, rect.width)

  sliderPositions.value.alpha = x

  hsl.value.alpha = x / rect.width
}

const handleAlphaMouseDown = (event: MouseEvent): void => {
  event.preventDefault()
  isDragging.value = true
  dragType.value = 'alpha'
  updateAlphaFromEvent(event)

  const { handleMouseMove, handleMouseUp } = createDragHandlers('alpha', updateAlphaFromEvent)

  document.addEventListener('mousemove', handleMouseMove)
  document.addEventListener('mouseup', handleMouseUp)
}

onMounted(() => {
  const colorValue = props.value

  rgb.value = colorToRgb(colorValue)
  hsl.value = colorToHsl(colorValue)

  updateSliderPositions()
})
</script>

<template>
  <div class="color-picker">
    <div
      class="color-preview-button"
      ref="buttonRef"
      @click="showPicker = !showPicker">
      <slot name="trigger" :currentColor="currentColor">
        <div class="color-preview-container">
          <div
            class="color-preview"
            :style="{ backgroundColor: currentColor }" />
        </div>
      </slot>
    </div>

    <Teleport to="body">
      <Transition
        name="fade"
        appear>
        <div
          v-show="isPositioned"
          ref="popupRef"
          v-click-outside="() => (showPicker = false)"
          class="color-picker-popup"
          :class="`popup-${placement}`"
          :style="floatingStyles">
          <div class="color-picker-content">
            <!-- s,l -->
            <div
              class="saturation-lightness-area"
              ref="slAreaRef"
              :style="{ background: slBackground }"
              @mousedown="handleSLAreaMouseDown">
              <div
                class="color-selector"
                :style="{
                  left: `${sliderPositions.s}px`,
                  top: `${sliderPositions.l}px`,
                  backgroundColor: currentColor
                }" />
            </div>

            <!-- h -->
            <div class="hue-slider-container">
              <div
                class="hue-slider"
                ref="hSliderRef"
                :style="{ background: hBackground }"
                @mousedown="handleHueMouseDown">
                <div
                  class="hue-selector"
                  :style="{ left: `${sliderPositions.h}px` }"></div>
              </div>
            </div>

            <!-- a -->
            <div class="alpha-slider-container">
              <div
                class="alpha-slider"
                ref="alphaSliderRef"
                :style="{ background: alphaBackground }"
                @mousedown="handleAlphaMouseDown">
                <div
                  class="alpha-selector"
                  :style="{ left: `${sliderPositions.alpha}px` }"></div>
              </div>
            </div>

            <!-- rgba -->
            <div class="rgb-inputs">
              <div class="input-group">
                <label>R</label>
                <input
                  :value="rgb.r"
                  @input="(e) => handleRgbInput('r', Number((e.target as HTMLInputElement).value))"
                  min="0"
                  max="255"
                  class="rgb-input" />
              </div>
              <div class="input-group">
                <label>G</label>
                <input
                  :value="rgb.g"
                  @input="(e) => handleRgbInput('g', Number((e.target as HTMLInputElement).value))"
                  min="0"
                  max="255"
                  class="rgb-input" />
              </div>
              <div class="input-group">
                <label>B</label>
                <input
                  :value="rgb.b"
                  @input="(e) => handleRgbInput('b', Number((e.target as HTMLInputElement).value))"
                  min="0"
                  max="255"
                  class="rgb-input" />
              </div>
              <div class="input-group">
                <label>A</label>
                <input
                  :value="Math.round((hsl.alpha ?? 1) * 100)"
                  @input="(e) => handleAlphaInput(Number((e.target as HTMLInputElement).value))"
                  min="0"
                  max="100"
                  class="alpha-input-field" />
              </div>
            </div>

            <!-- hsla -->
            <div class="hsl-inputs">
              <div class="input-group">
                <label>H</label>
                <input
                  :value="Math.round(hsl.h || 0)"
                  @input="(e) => handleHslInput('h', Number((e.target as HTMLInputElement).value))"
                  min="0"
                  max="360"
                  class="hsl-input" />
              </div>
              <div class="input-group">
                <label>S</label>
                <input
                  :value="Math.round(hsl.s)"
                  @input="(e) => handleHslInput('s', Number((e.target as HTMLInputElement).value))"
                  min="0"
                  max="100"
                  class="hsl-input" />
              </div>
              <div class="input-group">
                <label>L</label>
                <input
                  :value="Math.round(hsl.l)"
                  @input="(e) => handleHslInput('l', Number((e.target as HTMLInputElement).value))"
                  min="0"
                  max="100"
                  class="hsl-input" />
              </div>
              <div class="input-group">
                <label>A</label>
                <input
                  :value="Math.round((hsl.alpha ?? 1) * 100)"
                  @input="(e) => handleAlphaInput(Number((e.target as HTMLInputElement).value))"
                  min="0"
                  max="100"
                  class="alpha-input-field" />
              </div>
            </div>

            <div class="current-color-preview">
              <div class="preview-label">当前颜色</div>
              <div class="preview-color-container">
                <div
                  class="preview-color"
                  :style="{ backgroundColor: currentColor }"></div>
              </div>
              <div class="preview-values">
                <div class="preview-hex">{{ currentColor.toUpperCase() }}</div>
              </div>
            </div>
          </div>
        </div>
      </Transition>
    </Teleport>
  </div>
</template>
<style lang="less" scoped>
.color-picker {
  position: relative;
  display: inline-block;
}

.color-preview-button {
  display: inline-block;
  cursor: pointer;
  transition: all 0.2s ease;

  &:hover .color-preview {
    transform: scale(1.05);
  }
}

.color-preview-container {
  position: relative;
  width: 1.5rem;
  height: 1.5rem;
  border-radius: 0.25rem;
  border: 1px solid rgba(0, 0, 0, 0.2);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  background-image: url("data:image/svg+xml,%3csvg width='100' height='100' xmlns='http://www.w3.org/2000/svg'%3e%3cdefs%3e%3cpattern id='a' patternUnits='userSpaceOnUse' width='10' height='10'%3e%3crect width='10' height='10' fill='%23fff'/%3e%3crect width='5' height='5' fill='%23ccc'/%3e%3crect x='5' y='5' width='5' height='5' fill='%23ccc'/%3e%3c/pattern%3e%3c/defs%3e%3crect width='100' height='100' fill='url(%23a)'/%3e%3c/svg%3e");
  overflow: hidden;
}

.color-preview {
  width: 100%;
  height: 100%;
  border-radius: 0.25rem;
  transition: transform 0.1s ease;
}

.color-value {
  font-size: 0.75rem;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-weight: 500;
  user-select: none;
}

.color-picker-popup {
  z-index: 9999;
  background: white;
  border-radius: 0.5rem;
  box-shadow: 0 12px 24px rgba(0, 0, 0, 0.15), 0 4px 8px rgba(0, 0, 0, 0.1);
  overflow: hidden;
  min-width: 240px;
  max-width: 280px;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease-out;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.color-picker-content {
  padding: 0.75rem;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.saturation-lightness-area {
  position: relative;
  width: 100%;
  height: 120px;
  border-radius: 0.375rem;
  cursor: crosshair;
  background-blend-mode: multiply;

  &:hover .color-selector {
    transform: translate(-50%, -50%) scale(1.1);
  }
}

.color-selector {
  position: absolute;
  width: 12px;
  height: 12px;
  border: 2px solid white;
  border-radius: 50%;
  transform: translate(-50%, -50%);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  pointer-events: none;
  transition: transform 0.1s ease;
}

.hue-slider-container {
  width: 100%;
}

.hue-slider {
  position: relative;
  width: 100%;
  height: 16px;
  border-radius: 8px;
  cursor: pointer;
  border: 1px solid rgba(0, 0, 0, 0.1);
  transition: all 0.2s ease;

  &:hover {
    border-color: rgba(0, 0, 0, 0.2);

    .hue-selector {
      transform: translate(-50%, -50%) scale(1.1);
    }
  }
}

.hue-selector {
  position: absolute;
  top: 50%;
  width: 12px;
  height: 12px;
  background: white;
  border: 2px solid #374151;
  border-radius: 50%;
  transform: translate(-50%, -50%);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  pointer-events: none;
  transition: transform 0.1s ease;
}

.alpha-slider-container {
  width: 100%;
}

.alpha-slider {
  position: relative;
  width: 100%;
  height: 16px;
  border-radius: 8px;
  cursor: pointer;
  border: 1px solid rgba(0, 0, 0, 0.1);
  transition: all 0.2s ease;
  background-size: 8px 8px, 100% 100%;
  background-position: 0 0, 0 0;

  &:hover {
    border-color: rgba(0, 0, 0, 0.2);

    .alpha-selector {
      transform: translate(-50%, -50%) scale(1.1);
    }
  }
}

.alpha-selector {
  position: absolute;
  top: 50%;
  width: 12px;
  height: 12px;
  background: white;
  border: 2px solid #374151;
  border-radius: 50%;
  transform: translate(-50%, -50%);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.3);
  pointer-events: none;
  transition: transform 0.1s ease;
}

.rgb-inputs,
.hsl-inputs {
  display: flex;
  gap: 0.375rem;
}

.input-group {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 0.25rem;

  label {
    font-size: 0.75rem;
    font-weight: 500;
    color: #6b7280;
    text-align: center;
  }
}

.rgb-input,
.hsl-input,
.alpha-input-field {
  width: 100%;
  padding: 0.375rem;
  border: 1px solid #d1d5db;
  border-radius: 0.25rem;
  font-size: 0.75rem;
  text-align: center;
  transition: all 0.2s ease;
  background: white;

  &:focus {
    outline: none;
    border-color: #3b82f6;
    box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
  }

  &:hover:not(:focus) {
    border-color: #9ca3af;
  }
}

.current-color-preview {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem;
  background: #f9fafb;
  border-radius: 0.375rem;
  border: 1px solid #e5e7eb;
}

.preview-label {
  font-size: 0.875rem;
  font-weight: 500;
  color: #374151;
}

.preview-color-container {
  position: relative;
  width: 1.5rem;
  height: 1.5rem;
  border-radius: 0.25rem;
  border: 1px solid rgba(0, 0, 0, 0.1);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  background-image: url("data:image/svg+xml,%3csvg width='100' height='100' xmlns='http://www.w3.org/2000/svg'%3e%3cdefs%3e%3cpattern id='a' patternUnits='userSpaceOnUse' width='10' height='10'%3e%3crect width='10' height='10' fill='%23fff'/%3e%3crect width='5' height='5' fill='%23ccc'/%3e%3crect x='5' y='5' width='5' height='5' fill='%23ccc'/%3e%3c/pattern%3e%3c/defs%3e%3crect width='100' height='100' fill='url(%23a)'/%3e%3c/svg%3e");
  overflow: hidden;
}

.preview-color {
  width: 100%;
  height: 100%;
  border-radius: 0.25rem;
}

.preview-values {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
  margin-left: auto;
  text-align: right;
}

.preview-hex {
  font-size: 0.75rem;
  font-family: 'Monaco', 'Menlo', 'Ubuntu Mono', monospace;
  font-weight: 500;
  color: #6b7280;
}

</style>
