# controller层基本功能（路由层）

定义供router层调用的函数，调用service层方法

## 发生错误时的处理

发生错误，执行err != nil 下的函数, 返回JSON格式的错误信息, 并return结束函数,函数正常结束不return

## c.JSON显示错误信息

200为请求处理成功, 40X表示错误源自用户, 50X表示错误源自服务器

gin.H是存储相关信息的map类型数据, code为0表示无问题-1表示有问题, msg表示处理此次请求是否成功, error表示错误信息没有则为nil

## UserControl

```go
type UserControl struct {
    UserService *service.UserService
}
```

controller层需要调用service层方法，定义一个包含service层结构体的结构体

- ### func (b *UserControl)


令该方法属于UserControl

