package model

type Book struct {
	BookID   int    `json:"book_id"`                    //书籍ID
	BookName string `form:"book_name" json:"book_name"` //书籍名
	Author   string `form:"author" json:"author"`       //书籍作者
	Margin   int64  `form:"margin" json:"margin"`       //书籍余量
	Sales    int64  `form:"sales" json:"sales"`         //书籍销量
	Price    int    `json:"price"`                      //书籍价格
}
