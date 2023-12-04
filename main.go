package main

import (
	"database/sql"
	"dbsample/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// データベースへの接続情報を宣言する
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	// データベースに接続
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	const sqlStr = `select title, contents, username, nice from articles;`
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	// rows に存在するレコードそれぞれに対して、繰り返し処理を実行する
	for rows.Next() {
		// 変数 article の各フィールドに、取得レコードのデータを入れる
		// (SQL クエリの select 句から、タイトル・本文・ユーザー名・いいね数が返ってくることはわかっている)
		var article models.Article
		err := rows.Scan(&article.Title, &article.Contents, &article.UserName,
			&article.NiceNum)

		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}
	fmt.Printf("%+v\n", articleArray)
}
