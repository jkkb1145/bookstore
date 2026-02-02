package model

import "database/sql"

//TODO: 数据库搜索追求效率,若收藏表定义如下,搜索收藏夹书籍信息时,需要建立多表连接查询,请尝试直接将书籍信息包含在收藏表中

type Favourite struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}

type FavouriteInfo struct {
	BookID    int           // 书籍ID
	BookName  string        // 书籍名称
	Author    string        // 书籍作者
	Margin    int64         //书籍余量
	Sales     int64         //书籍销量
	Price     int           // 书籍价格
	CollectID sql.NullInt64 // 收藏ID（未收藏为NULL，使用sql.NullXXX处理NULL值）
	UserID    sql.NullInt64 // 收藏用户ID（未收藏为NULL）
}
