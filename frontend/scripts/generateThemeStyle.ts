import { defaultThemeConfig, presetThemes, ThemeColorOptions } from '../src/styles/themes'
import { colorToRgb, getInterpolateByContrast, parseRgbStr } from '../src/utils/color'
import * as fs from 'node:fs'

const variablesName = {
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

const formatRGB = (rgb: { r: number, g: number, b: number }) => {
  return `${rgb.r},${rgb.g},${rgb.b}`
}

const convertColorWithVariable = (theme: ThemeColorOptions) => {
  if (!theme || typeof theme !== 'object') return theme
  
  const convertResult = {}
  
  const base1 = theme.base1 || defaultThemeConfig.base1
  const primary = theme.primary || defaultThemeConfig.primary
  const info = theme.info || defaultThemeConfig.info
  const success = theme.success || defaultThemeConfig.success
  const warning = theme.warning || defaultThemeConfig.warning
  const error = theme.info || defaultThemeConfig.error
  
  const defaultColorMap: ThemeColorOptions = {
    base1: parseRgbStr(base1),
    base2: formatRGB(getInterpolateByContrast(base1, 0.07)),
    base3: formatRGB(getInterpolateByContrast(base1, 0.14)),
    baseContent: formatRGB(getInterpolateByContrast(base1, 0.8)),
    primary: parseRgbStr(defaultThemeConfig.primary),
    primaryContent: formatRGB(getInterpolateByContrast(primary, 0.8)),
    info: parseRgbStr(defaultThemeConfig.info),
    infoContent: formatRGB(getInterpolateByContrast(info, 0.8)),
    success: parseRgbStr(defaultThemeConfig.success),
    successContent: formatRGB(getInterpolateByContrast(success, 0.8)),
    warning: parseRgbStr(defaultThemeConfig.warning),
    warningContent: formatRGB(getInterpolateByContrast(warning, 0.8)),
    error: parseRgbStr(defaultThemeConfig.error),
    errorContent: formatRGB(getInterpolateByContrast(error, 0.8))
  }
  
  for (const [k, v] of Object.entries(defaultColorMap)) {
    convertResult[variablesName[k]] = k in theme ? formatRGB(colorToRgb(theme[k])) : v
  }
  
  return convertResult
}

const toCSS = (k: string, v: object) => {
  return `:root[data-theme=${k}] {${Object.entries(v).map(([k, v]) => `${k}: ${v}`).join(';')}}`
}

const themes = Object.entries(presetThemes).map(([k, v]) => {
  const colorWithVariable = convertColorWithVariable(v)
  return toCSS(k, colorWithVariable)
})

fs.writeFile('src/styles/themes.less', themes.join('\n'), (err) => {
  if (err) {
    console.error(err)
  } else {
    console.log('Generate theme style success')
  }
})
