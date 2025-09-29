# Context
首先我们要明白什么是 Context，其本质就是请求的上下文。里边包含我们取消当前请求的一些信息等。
我们框架中有 `gin.Context`，这个是我们框架给我们封装的一个 **Context**。主要是来自于用户端的请求，我们可以直接使用
里边的方法`ctx.Request.Context()`来获取，里边会有我们客户端取消的函数，我们不用担心请求会一直挂着导致我们的 CPU 崩溃。

## 场景
1️⃣ 场景一：定时任务（Cron Job）  

2️⃣ 场景二：消费队列（Kafka/RabbitMQ）  

3️⃣ 场景三：Go Routine  

4️⃣ 场景四：CLI 工具 / 命令行脚本  

5️⃣ 场景五：内部 RPC 调用（无外部上下文）  

6️⃣ 场景六：健康检查（Health Check）

## 总结
| 条件	                                            | 是否需要加 NewContext() |
|------------------------------------------------|--------------------|
| 没有 http.Request                                | 	✅ 需要              |
| 没有 context.Context 参数                          | 	✅ 需要              |
| 是后台任务、定时任务、消息消费                                | 	✅ 需要              |
| 是 CLI 脚本、工具函数                                  | 	✅ 需要              |
| 有 c.Request.Context() 或 ctx context.Context 参数 | 	❌ 不需要，直接用         |

## 自定义
````go
type Context struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewContext() *Context {
	// 增加一个 2 秒钟的取消
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	return &Context{
		ctx:        ctx,
		cancelFunc: cancelFunc,
	}
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

func (c *Context) Done() <-chan struct{} { return c.ctx.Done() }

func (c *Context) Err() error { return c.ctx.Err() }

func (c *Context) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

func (c *Context) Cancel() {
	c.cancelFunc()
}
````