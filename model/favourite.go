package model

//TODO: 数据库搜索追求效率,若收藏表定义如下,搜索收藏夹书籍信息时,需要建立多表连接查询,请尝试直接将书籍信息包含在收藏表中

type Favourite struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}
