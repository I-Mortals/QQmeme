// 拼接路径
export const joinPath = (path: string, dir: string) => {
  if (path === '') {
    return dir
  }
  
  return path + '\\' + dir
}

export const joinShowImgPath = (path: string, dir: string) => {
  // 检查参数是否有效
  if (!path || path === '') {
    return dir
  }
  
  // 将 c:\ 格式修改为 \c\ 格式
  if (path.match(/^[a-zA-Z]:\\/)) {
    // 提取盘符并重新格式化路径
    const drive = path.charAt(0)
    const restPath = path.substring(2) // 跳过盘符和冒号
    
    return '\\' + drive + restPath + '\\' + dir
  }
  
  // 如果路径格式不匹配，返回原始拼接
  return path + '\\' + dir
}
