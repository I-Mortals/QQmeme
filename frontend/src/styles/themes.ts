type BaseOptions = {
  base: string
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

export type ThemeColorOptionsKey = keyof ThemeColorOptions

export const variablesNameMap = {
  base1: '--b1',
  base2: '--b2',
  base3: '--b3',
  baseContent: '--bc',
  primary: '--p',
  primaryContent: '--pc',
  info: '--i',
  infoContent: '--ic',
  success: '--s',
  successContent: '--sc',
  warning: '--w',
  warningContent: '--wc',
  error: '--e',
  errorContent: '--ec'
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
  // 原有主题
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
  },

  // DaisyUI 经典主题
  cupcake: {
    base1: '#FAF7F5',
    baseContent: '#291334',
    primary: '#65C3C8',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  forest: {
    base1: '#1A202C',
    baseContent: '#A7ACB1',
    primary: '#1EB854',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  synthwave: {
    base1: '#0D0A0F',
    baseContent: '#F9F7FF',
    primary: '#E779C1',
    info: '#58C7F3',
    success: '#25D0AB',
    warning: '#F7D51D',
    error: '#F25C54'
  },
  cyberpunk: {
    base1: '#0B0B0B',
    baseContent: '#00FF00',
    primary: '#FF0090',
    info: '#00FFFF',
    success: '#00FF00',
    warning: '#FFFF00',
    error: '#FF0000'
  },
  valentine: {
    base1: '#FFFFFF',
    baseContent: '#1F2937',
    primary: '#E96D7B',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  halloween: {
    base1: '#212121',
    baseContent: '#F7F3E9',
    primary: '#F28C18',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  garden: {
    base1: '#FFFFFF',
    baseContent: '#1F2937',
    primary: '#66CC8A',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  aqua: {
    base1: '#F0F9FF',
    baseContent: '#1F2937',
    primary: '#00D4FF',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  lofi: {
    base1: '#FFFFFF',
    baseContent: '#0D0D0D',
    primary: '#0D0D0D',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  pastel: {
    base1: '#FFFFFF',
    baseContent: '#1F2937',
    primary: '#D1C1D7',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  fantasy: {
    base1: '#FFFFFF',
    baseContent: '#1F2937',
    primary: '#6D28D9',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  wireframe: {
    base1: '#FAFAFA',
    baseContent: '#000000',
    primary: '#B8B8B8',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  black: {
    base1: '#000000',
    baseContent: '#FFFFFF',
    primary: '#FFFFFF',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  luxury: {
    base1: '#0A0A0B',
    baseContent: '#FAFAFA',
    primary: '#FFFFFF',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  dracula: {
    base1: '#282A36',
    baseContent: '#F8F8F2',
    primary: '#FF79C6',
    info: '#8BE9FD',
    success: '#50FA7B',
    warning: '#F1FA8C',
    error: '#FF5555'
  },
  cmyk: {
    base1: '#FFFFFF',
    baseContent: '#1F2937',
    primary: '#45AEEE',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  autumn: {
    base1: '#2D1B0E',
    baseContent: '#F7F3E9',
    primary: '#A78BFA',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  acid: {
    base1: '#0B0F14',
    baseContent: '#F0F0F0',
    primary: '#FF00FF',
    info: '#00FFFF',
    success: '#00FF00',
    warning: '#FFFF00',
    error: '#FF0000'
  },
  lemonade: {
    base1: '#FFFFFF',
    baseContent: '#1F2937',
    primary: '#E0E7FF',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  night: {
    base1: '#0F1729',
    baseContent: '#A6ADBB',
    primary: '#3B82F6',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  coffee: {
    base1: '#221551',
    baseContent: '#F9F7FF',
    primary: '#DB924B',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },
  winter: {
    base1: '#FFFFFF',
    baseContent: '#1F2937',
    primary: '#047AFF',
    info: '#3ABFF8',
    success: '#36D399',
    warning: '#FBBD23',
    error: '#F87272'
  },

  // 自定义主题
  ocean: {
    base1: '#F0F9FF',
    baseContent: '#0C4A6E',
    primary: '#0891B2',
    info: '#0284C7',
    success: '#059669',
    warning: '#D97706',
    error: '#DC2626'
  },
  sunset: {
    base1: '#FFF7ED',
    baseContent: '#431407',
    primary: '#EA580C',
    info: '#0284C7',
    success: '#059669',
    warning: '#D97706',
    error: '#DC2626'
  },
  mint: {
    base1: '#F0FDF4',
    baseContent: '#14532D',
    primary: '#10B981',
    info: '#0284C7',
    success: '#059669',
    warning: '#D97706',
    error: '#DC2626'
  },
  rose: {
    base1: '#FFF1F2',
    baseContent: '#881337',
    primary: '#E11D48',
    info: '#0284C7',
    success: '#059669',
    warning: '#D97706',
    error: '#DC2626'
  },
  purple: {
    base1: '#FAF5FF',
    baseContent: '#581C87',
    primary: '#8B5CF6',
    info: '#0284C7',
    success: '#059669',
    warning: '#D97706',
    error: '#DC2626'
  },
  emerald: {
    base1: '#ECFDF5',
    baseContent: '#064E3B',
    primary: '#059669',
    info: '#0284C7',
    success: '#059669',
    warning: '#D97706',
    error: '#DC2626'
  },
  amber: {
    base1: '#FFFBEB',
    baseContent: '#92400E',
    primary: '#F59E0B',
    info: '#0284C7',
    success: '#059669',
    warning: '#D97706',
    error: '#DC2626'
  },
  slate: {
    base1: '#F8FAFC',
    baseContent: '#0F172A',
    primary: '#475569',
    info: '#0284C7',
    success: '#059669',
    warning: '#D97706',
    error: '#DC2626'
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
