import { defaultThemeConfig, presetThemes, ThemeColorOptions, variablesNameMap } from '../src/styles/themes'
import { getInterpolateByContrast, parseRgbStr } from '../src/utils/color'
import * as fs from 'node:fs'

const convertColorWithVariable = (theme: ThemeColorOptions) => {
  if (!theme || typeof theme !== 'object') return theme
  
  const convertResult = {}
  
  const base1 = theme.base1 || defaultThemeConfig.base1
  const primary = theme.primary || defaultThemeConfig.primary
  const info = theme.info || defaultThemeConfig.info
  const success = theme.success || defaultThemeConfig.success
  const warning = theme.warning || defaultThemeConfig.warning
  const error = theme.error || defaultThemeConfig.error
  
  const defaultColorMap: ThemeColorOptions = {
    base1: parseRgbStr(base1),
    base2: parseRgbStr(getInterpolateByContrast(base1, 0.07)),
    base3: parseRgbStr(getInterpolateByContrast(base1, 0.14)),
    baseContent: parseRgbStr(getInterpolateByContrast(base1, 0.8)),
    primary: parseRgbStr(defaultThemeConfig.primary),
    primaryContent: parseRgbStr(getInterpolateByContrast(primary, 0.8)),
    info: parseRgbStr(defaultThemeConfig.info),
    infoContent: parseRgbStr(getInterpolateByContrast(info, 0.8)),
    success: parseRgbStr(defaultThemeConfig.success),
    successContent: parseRgbStr(getInterpolateByContrast(success, 0.8)),
    warning: parseRgbStr(defaultThemeConfig.warning),
    warningContent: parseRgbStr(getInterpolateByContrast(warning, 0.8)),
    error: parseRgbStr(defaultThemeConfig.error),
    errorContent: parseRgbStr(getInterpolateByContrast(error, 0.8))
  }
  
  for (const [k, v] of Object.entries(defaultColorMap)) {
    convertResult[variablesNameMap[k]] = k in theme ? parseRgbStr(theme[k]) : v
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

fs.writeFile(`src/styles/themes-${new Date().getTime()}.less`, themes.join('\n'), (err) => {
  if (err) {
    console.error(err)
  } else {
    console.log('Generate theme style success')
  }
})
