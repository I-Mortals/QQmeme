package main

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp 创建新的 App 应用程序结构
func NewApp() *App {
	return &App{}
}

// startup 启动时调用启动。保存上下文
// 因此，我们可以调用运行时方法(单独使用onStartup时有效)
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
// 销毁所有资源
func (a *App) shutdown(ctx context.Context) {
	// a.ctx = ctx
	a.ctx = nil
}
// 从顶层传入app的ctx
func (a *App) SetContext(ctx context.Context) {
	a.ctx = ctx
}
