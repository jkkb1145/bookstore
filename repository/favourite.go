package repository

import (
	"database/sql"
	"demo02/global"
	"demo02/model"
	"fmt"
)

type FavouriteDAO struct {
	FavouriteDB *sql.DB
}

func NewFavouriteDAO() *FavouriteDAO {
	return &FavouriteDAO{
		FavouriteDB: global.GetDb(),
	}
}

func (f *FavouriteDAO) AddFavourite(userID, bookID int) error {
	// 1. 预处理SQL语句：使用?作为占位符（PostgreSQL用$1、$2）
	// 插入字段明确指定为用户ID书籍ID，匹配建表结构
	stmt, err := global.Db.Prepare("INSERT INTO 收藏 (用户ID,书籍ID) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("预处理SQL失败：%w", err)
	}
	defer stmt.Close() // 延迟关闭预处理语句，释放资源

	// 2. 执行预处理语句，传入参数（按占位符顺序）
	_, err = stmt.Exec(userID, bookID)
	if err != nil {
		return fmt.Errorf("执行插入失败：%w", err)
	}
	return nil
}

func (f *FavouriteDAO) RemoveFavourite(userID, bookID int) (bool, error) {
	// 1. 预处理SQL语句：使用?占位符，分离SQL模板与参数，从根本防止SQL注入
	// 条件精准匹配user_id和book_id，确保仅删除指定的一条关联记录
	stmt, err := global.Db.Prepare("DELETE FROM 收藏 WHERE 用户ID = ? AND 书籍ID = ? LIMIT 1")
	if err != nil {
		return false, fmt.Errorf("预处理删除SQL失败：%w", err)
	}
	defer stmt.Close() // 延迟关闭预处理语句，释放数据库资源

	// 2. 执行预处理语句，按占位符顺序传入参数
	result, err := stmt.Exec(userID, bookID)
	if err != nil {
		return false, fmt.Errorf("执行删除操作失败：%w", err)
	}

	// 3. 获取受影响的行数，校验是否实际删除了数据
	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("获取受影响行数失败：%w", err)
	}

	// 受影响行数=1 → 删除成功；=0 → 无匹配的用户ID+书籍ID记录
	return affectedRows == 1, nil
}

func (f *FavouriteDAO) GetUserFavourite(userID int) (*[]model.FavouriteInfo, error) {
	// 核心左连接SQL语句（参数化查询，防止SQL注入）
	sqlStr := `
	SELECT 
	  书籍.书籍ID ,
	  书籍.书籍名,
	  书籍.书籍作者,
	  书籍.书籍余量,
	  书籍.书籍销量,
	  书籍.书籍价格,
	  收藏.ID,
	  收藏.用户ID
	FROM 
	  书籍
	LEFT JOIN 
	  收藏
	ON 
	  书籍.书籍ID=收藏.书籍ID
	WHERE 
	  收藏.用户ID = ?
	`

	// 执行查询，传入用户ID作为参数
	rows, err := global.Db.Query(sqlStr, userID)
	if err != nil {
		return nil, fmt.Errorf("查询失败：%w", err)
	}
	// 必须延迟关闭rows，避免数据库连接泄漏
	defer rows.Close()

	// 遍历结果集，封装到切片中
	var result []model.FavouriteInfo
	for rows.Next() {
		var bc model.FavouriteInfo
		// 扫描行数据到结构体，字段顺序必须与SQL查询结果一致
		err := rows.Scan(
			&bc.BookID,
			&bc.BookName,
			&bc.Author,
			&bc.Price,
			&bc.CollectID,
			&bc.UserID,
			&bc.CollectID,
			&bc.UserID,
		)
		if err != nil {
			return nil, fmt.Errorf("解析结果失败：%w", err)
		}
		result = append(result, bc)
	}

	// 检查遍历结果集时是否发生错误
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历结果失败：%w", err)
	}

	return &result, nil
}
