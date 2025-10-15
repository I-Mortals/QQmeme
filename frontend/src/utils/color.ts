import { converter, interpolate, Rgb, wcagContrast } from 'culori'

export const parseRgb = (rgbColor: string | Rgb): Rgb => {
  let rgbStr: Rgb
  if (typeof rgbColor === 'string') {
    rgbStr = converter()(rgbColor) as Rgb
  } else {
    rgbStr = rgbColor
  }
  const { r, g, b, alpha } = rgbStr
  return {
    r: Math.round(r * 255) || 0,
    g: Math.round(g * 255) || 0,
    b: Math.round(b * 255) || 0,
    alpha: alpha ?? 1,
    mode: 'rgb'
  }
}

export const parseRgbStrWithCSS = (rgbColor: string | Rgb) => {
  const { r, g, b, alpha } = parseRgb(rgbColor)
  return `rgba(${r},${g},${b},${alpha})`
}

export const parseRgbStr = (rgbColor: string | Rgb) => {
  const { r, g, b} = parseRgb(rgbColor)
  return `${r},${g},${b}`
}


export const isMoreDark = (color: string) => {
  /*
  * 对比度指的是颜色之间的差异程度。高对比度表示两种颜色之间的差异很大，低对比度则表示颜色之间的差异较小。
  * 如果待判断的颜色与亮色（白色）之间的对比度更高，则该颜色趋近于暗色。
  * 如果待判断的颜色与暗色（黑色）之间的对比度更高，则该颜色趋近于亮色。
  * color 与 black 的对比度越低，越趋近于暗色，否则更趋近于亮色
  */
  return wcagContrast(color, 'black') < wcagContrast(color, 'white')
}

export const getInterpolateByContrast = (
  color: string,
  percent: number
): Rgb => {
  return interpolate([color, isMoreDark(color) ? 'white' : 'black'], 'rgb')(percent)
}
