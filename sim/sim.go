package sim

import (
	"net/http"
)

// HandlerFunc sim使用的处理函数
type HandlerFunc func(ctx *Context)

// H 提供便利的简写形式
type H map[string]interface{}

type Engine struct {
	// 把engine抽象成了一个顶级的RouterGroup
	// 路由注册的api直接放入group中
	// 由于使用了匿名嵌套, engine也能使用对应的api
	*RouterGroup

	// 实现路由功能,存储路由信息
	router *router

	// 存储所有group
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 实现http服务接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w, req)
	e.router.handle(ctx)
}

// Run 启动server
func (e *Engine) Run() error {
	return http.ListenAndServe(":8080", e)
}
