// 拼接路径
export const joinPath = (path: string, dir: string) => {
  if (path === '') {
    return dir
  }
  
  // 使用正斜杠作为路径分隔符，这在所有平台上都有效
  return path + '/' + dir
}

export const joinShowImgPath = (path: string, dir: string) => {
  // 检查参数是否有效
  if (!path || path === '') {
    return dir
  }
  
  if (path.match(/^[a-zA-Z]:\\/)) {
    // Windows格式: 将 c:\ 格式修改为 \c\ 格式
    const drive = path.charAt(0)
    const restPath = path.substring(2) // 跳过盘符和冒号
    
    return '\\' + drive + restPath + '\\' + dir
  }
  
  // Unix/Mac格式: 直接使用正斜杠拼接
  return path + '/' + dir
}
