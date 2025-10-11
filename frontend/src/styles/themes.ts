type BaseOptions = {
  base: string
  fill: string
}

type StatusOptions = {
  primary?: string
  info?: string
  success?: string
  warning?: string
  error?: string
}

export type ThemeColorOptions = {
  base1: string
  baseContent?: string
} & {
  [K in keyof BaseOptions as `${K}1` | `${K}2` | `${K}3`]?: string
} & StatusOptions & {
  [K in keyof StatusOptions as `${K}Content`]?: string
}

export const defaultThemeConfig: ThemeColorOptions = {
  base1: '#FFFFFF',
  primary: '#0066FF',
  info: '#00B5FF',
  success: '#00A96E',
  warning: '#FFBE00',
  error: '#FF5861'
}

export const presetThemes: Record<string, ThemeColorOptions> = {
  light: {
    base1: '#FFFFFF',
    baseContent: '#1F2937',
    primary: '#0066FF',
    info: '#00B5FF',
    success: '#00A96E',
    warning: '#FFBE00',
    error: '#FF5861'
  },
  dark: {
    base1: '#1d232a',
    baseContent: '#A6ADBB',
    primary: '#4a9d9c'
  },
  business: {
    base1: '#202020',
    baseContent: '#CDCDCD',
    primary: '#FF9900',
    info: '#0091D5',
    success: '#6BB187',
    warning: '#DBAE59',
    error: '#AC3E31'
  }
}

export const getThemeColors = (
  theme: string,
  ...color: (keyof ThemeColorOptions)[]
) => {
  return color.map((item) => getThemeColor(theme, item))
}

export const getThemeColor = (
  theme: string,
  color: keyof ThemeColorOptions
) => {
  const themesColor = presetThemes[theme]
  
  return themesColor[color] || defaultThemeConfig[color]
}
