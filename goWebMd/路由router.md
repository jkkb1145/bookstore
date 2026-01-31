## 路由层功能：

定义供客户端访问的路由，访问路由时，调用controller层函数

v1:=r. Group(“/API/v1”)
---

路由分组v1,所属该分组的所有路由在访问时,都要在自己的路由前加上v1的路由

## user := v1.Group("/user")

user属于v1,属于user的分组也同时属于v1,访问时需要加上user与v1两者的分组

##  auth.Use(midleWare.AdminAuthMiddleware())

auth分组中使用**中间件**实现鉴权

- ### **`Use`方法的核心作用**:

`Use`方法的核心功能是：**将指定的中间件绑定到当前路由分组（`auth`）上**，让该分组下的**所有路由**在处理请求时，**统一执行该中间件的逻辑**

`Use`方法会让**这 4 个路由成为`AdminAuthMiddleware`中间件的 “作用目标”** —— 当客户端发起对这 4 个接口的请求时，Gin 框架会**先执行`AdminAuthMiddleware`的鉴权逻辑**，再执行后续的控制器处理函数（`GetUserProfile`/`UpdateUserProfile`等）。 