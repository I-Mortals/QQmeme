import { converter, formatHex, formatHex8, Hsl, hsl, parse, Rgb } from 'culori'

export const DEFAULT_COLOR: Hsl = {
  mode: 'hsl',
  h: 0,
  s: 0,
  l: 1,
  alpha: 1
}

// 1,1,1 是 culori 里的白色
export const WHITE_RGB = { r: 1, g: 1, b: 1 }
// 0,0,0 是 culori 里的黑色
export const BLACK_RGB = { r: 0, g: 0, b: 0 }

export const hexToHsl = (hex: string): Hsl => {
  return hsl(hex) || DEFAULT_COLOR
}

export const parseHsl = (hslColor: Hsl): { h: number; s: number; l: number; alpha: number } => {
  const { h, s, l, alpha } = hslColor
  return {
    h: Number(h?.toFixed(2)) || 0,
    s: Number((s * 100).toFixed(2)),
    l: Number((l * 100).toFixed(2)),
    alpha: alpha ?? 1
  }
}

export const hslToHex = (h: number = 0, s: number, l: number): string => {
  return formatHex(`hsl(${h},${s}%,${l}%)`) || formatHex(DEFAULT_COLOR)
}

export const hslToHex8 = (h: number = 0, s: number, l: number, alpha: number = 1): string => {
  return formatHex8(`hsl(${h},${s}%,${l}%,${alpha})`) || formatHex8(DEFAULT_COLOR)
}

export const hslToRgb = (h: number = 0, s: number, l: number): Rgb => {
  return (converter('rgb')(`hsl(${h},${s}%,${l}%)`) as Rgb) || (converter('rgb')('white') as Rgb)
}

export const parseRgb = (rgbColor: Rgb) => {
  const { r, g, b, alpha } = rgbColor
  return {
    r: Math.round(r * 255) || 255,
    g: Math.round(g * 255) || 255,
    b: Math.round(b * 255) || 255,
    alpha: alpha ?? 1
  }
}

export const colorToHsl = (colorValue: string) => {
  return parseHsl(converter('hsl')(colorValue) || DEFAULT_COLOR)
}

export const colorToRgb = (colorValue: string) => {
  return parseRgb(parse(colorValue) as Rgb)
}

export const clampValue = (value: number, min: number, max: number): number => Math.max(min, Math.min(max, value))
