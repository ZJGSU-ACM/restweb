restweb
=======

light web framwork for go

##Features

1. 支持路由配置文件, 支持控制器方法参数传入
2. 模板自动渲染
3. session 管理
4. 表单验证
5. 过滤和拦截
	1. 过滤器函数返回值为true则为拦截，支持控制器方法调用前和后拦截、过滤
	2. 过滤器按注册的顺序安排优先级，注册早的优先级高
	3. 对于一个url，如果一个拦截器被执行，将立即停止执行其后的过滤器和控制器方法

##TODO List:

4. 配置文件