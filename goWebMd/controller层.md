## controller层基本功能（路由层）

定义供router层调用的函数，调用service层方法

## UserControl

```go
type UserControl struct {
    UserService *service.UserService
}
```

controller层需要调用service层方法，定义一个包含service层结构体的结构体

## (b *UserControl)

令该方法属于UserControl