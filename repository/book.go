package repository

import (
	"database/sql"
	"demo02/global"
	"demo02/model"
	"fmt"
	"log"
)

type BookDAO struct {
	BookDB *sql.DB
}

func NewBookDAO() *BookDAO {
	return &BookDAO{
		BookDB: global.GetDb(),
	}
}

func (b *BookDAO) GetPopBooks() (error, *[]model.Book) {
	var books []model.Book
	//从书籍表中查找销量排名前十的书籍,也就是书籍销量按降序排序后前十本书的基本信息
	querySQL := `
		SELECT 书籍ID,书籍名,书籍作者,书籍余量,书籍销量,书籍价格 
		FROM 书籍 
		ORDER BY 书籍销量 DESC 
		LIMIT 5;
	`

	// 5. 执行查询并绑定结果到books变量
	rows, err := global.Db.Query(querySQL)
	if err != nil {
		log.Fatalf("查询失败: %v", err)
		return err, nil
	}
	defer rows.Close()

	// 遍历查询结果并赋值给books
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.BookID, &book.BookName, &book.Author, &book.Margin, &book.Sales, &book.Price)
		if err != nil {
			log.Fatalf("结果扫描失败: %v", err)
			return err, nil
		}
		books = append(books, book)
	}

	// 检查遍历过程中是否有错误
	if err = rows.Err(); err != nil {
		log.Fatalf("遍历结果失败: %v", err)
		return err, nil
	}
	return nil, &books
}

// 根据keywords查询书籍,分页逻辑的实现
func (b *BookDAO) SearchBooks(keywords string, page, pageSize int) (*[]model.Book, int, error) {
	var books []model.Book

	//实现分页
	// 1. 查询符合条件的总条数（用于计算总页数）
	var total int
	countSQL := `SELECT COUNT(*) FROM 书籍 WHERE 书籍名 LIKE ? OR 书籍作者 LIKE ?`
	// 参数化查询，防止SQL注入，%是模糊匹配通配符
	err := global.Db.QueryRow(countSQL, "%"+keywords+"%", "%"+keywords+"%").Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("获取相关书籍总数失败：%v", err)
	}

	// 2. 计算分页偏移量（OFFSET从0开始，所以page-1）
	offset := (page - 1) * pageSize
	// 3. 分页查询书籍列表
	querySQL := `
		SELECT 书籍ID,书籍名,书籍作者,书籍余量,书籍销量,书籍价格
		FROM 书籍 
		WHERE 书籍名 LIKE ? OR 书籍作者 LIKE ? 
		LIMIT ? OFFSET ?
	`
	rows, err := global.Db.Query(querySQL, "%"+keywords+"%", "%"+keywords+"%", pageSize, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("查询书籍列表失败：%v", err)
	}
	defer rows.Close() // 关闭结果集

	// 4. 解析查询结果到Book切片
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.BookID, &book.BookName, &book.Author, &book.Margin, &book.Sales, &book.Price)
		if err != nil {
			return nil, 0, fmt.Errorf("解析书籍数据失败：%v", err)
		}
		books = append(books, book)
	}
	// 检查遍历结果集时的错误
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("遍历结果集失败：%v", err)
	}

	return &books, total, nil

}
