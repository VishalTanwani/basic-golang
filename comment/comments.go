package comment

/*
	name: comments
*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	db "mini/db"
)

type basic interface {
	add()
}

type Comment struct {
	ID              int
	PostID          int
	Score           int
	Text            string
	UserID          int
	CreationDate    string
	UserDisplayName string
}

var comment []Comment
var commentJSON []Comment

func init() {
	data, err := ioutil.ReadFile("../Json Data/Comments.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &comment)
	if err != nil {
		panic(err)
	}

	fmt.Println("data extracted from file")
}

func (value Comment) add() {
	fmt.Println("# Inserting values")
	var lastInsertID int
	err := db.DB.QueryRow("INSERT INTO comments(id, post_id, creation_date, text, user_id, score, user_display_name) VALUES($1,$2,$3,$4,$5,$6,$7) returning id;", value.ID, value.PostID, value.CreationDate, value.Text, value.UserID, value.Score, value.UserDisplayName).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}
	fmt.Println("last inserted id =", lastInsertID)
}

func Add() {
	for _, value := range comment {
		value.add()
		// fmt.Println(value.LastEditorUserID)
	}
}

func DeleteAll() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DELETE FROM comments")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted All")
}

func DropTable() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DROP TABLE comments")
	if err != nil {
		panic(err)
	}
	fmt.Println("Table Deleted")
}

func ShowIDS() {
	fmt.Println("# Querying")
	rows, err := db.DB.Query("SELECT id FROM comments")
	if err != nil {
		panic(err)
	}

	fmt.Println(rows)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println(id)
	}
}

func Show5IDS() {
	fmt.Println("# Querying comments")
	rows, err := db.DB.Query("SELECT id FROM comments LIMIT 5")
	if err != nil {
		panic(err)
	}

	fmt.Println(rows)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Println(id)
	}
}

// func SecondJSON(){
// 	fmt.Println("Get all comments for a post")
// 	rows, err := db.DB.Query("SELECT * FROM comments where post_id = $1", postId)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// fmt.Println(rows)
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println("id  ||  post_id  ||  creation_date  ||  text  ||  user_id  ||  score  ||  user_display_name")
// 	fmt.Println()
// 	fmt.Println()

// 	for rows.Next() {
// 		var id int
// 		var post_id int
// 		var creation_date string
// 		var text string
// 		var user_id int
// 		var score int
// 		var user_display_name string
// 		err = rows.Scan(&id, &post_id, &creation_date, &text, &user_id, &score, &user_display_name)
// 		if err != nil {
// 			panic(err)
// 		}
// 		c := Comment{ID:id,PostID:post_id,Score:score,Text:text,UserID:user_id,CreationDate:`creation_date`,UserDisplayName:`$user_display_name`}
// 	}
// }
