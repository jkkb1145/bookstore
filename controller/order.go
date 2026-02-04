package controller

import (
	"demo02/service"
)

type CreateOrderRequest struct {
	UserID int          `json:"user_id"`
	Items  []OrderItems `json:"items"`
}

type OrderItems struct {
	BookID   int `json:"book_id"`
	Quantity int `json:"quantity"`
	Price    int `json:"price"`
}

type OrderController struct {
	orderservice *service.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{
		orderservice: service.NewOrderService(),
	}
}

//func (o *OrderController) CreateOrder(c *gin.Context) {
//	var req CreateOrderRequest
//	if err := c.ShouldBindJSON(&req); err != nil {
//		c.JSON(500, gin.H{
//			"code": -1,
//			"msg":  "绑定参数失败",
//		})
//		return
//	}
//	userID, exists := c.Get("userID")
//	if !exists {
//		c.JSON(400, gin.H{
//			"code": -1,
//			"msg":  "未登录",
//		})
//		return
//	}
//	req.UserID = userID.(int)
//	o.orderservice.CreateOrder(&req)
//}
