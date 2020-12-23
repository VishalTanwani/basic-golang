package Post

/*
	name: posts
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

type Post struct {
	ID                    int
	PostTypeID            int
	Score                 int
	ViewCount             int
	Tags                  string
	AnswerCount           int
	CommentCount          int
	FavoriteCount         int
	CreationDate          string
	Body                  string
	ClosedDate            string
	AcceptedAnswerID      int
	ParentID              int
	OwnerUserID           int
	OwnerDisplayName      string
	LastEditorUserID      int
	LastEditorDisplayName string
	LastEditDate          string
	LastActivityDate      string
	Title                 string
	CommunityOwnedDate    string
}

var post []Post

func init() {
	data, err := ioutil.ReadFile("../Json Data/Posts.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &post)
	if err != nil {
		panic(err)
	}

	fmt.Println("data extracted from file")
}

func (value Post) add() {
	fmt.Println("# Inserting values")
	var lastInsertID int
	err := db.DB.QueryRow("INSERT INTO posts(id,post_type_id,score,view_count,tags,answer_count,comment_count,favourite_count,creation_date,body,closed_date,accepted_answer_id,parent_id,owner_user_id,owner_display_name,last_editor_user_id,last_editor_display_name,last_edit_date,last_activity_date,title,community_owned_date) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21) returning id;", value.ID, value.PostTypeID, value.Score, value.ViewCount, value.Tags, value.AnswerCount, value.CommentCount, value.FavoriteCount, value.CreationDate, value.Body, value.ClosedDate, value.AcceptedAnswerID, value.ParentID, value.OwnerUserID, value.OwnerDisplayName, value.LastEditorUserID, value.LastEditorDisplayName, value.LastEditDate, value.LastActivityDate, value.Title, value.CommunityOwnedDate).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}
	fmt.Println("last inserted id =", lastInsertID)
}

func Add() {
	for _, value := range post {
		value.add()
		// fmt.Println(value.LastEditorUserID)
	}
}

func DeleteAll() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DELETE FROM posts")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted All")
}

func DropTable() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DROP TABLE posts")
	if err != nil {
		panic(err)
	}
	fmt.Println("Table Deleted")
}

func ShowIDS() {
	fmt.Println("# Querying")
	rows, err := db.DB.Query("SELECT id FROM posts")
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
	fmt.Println("# Querying posts")
	rows, err := db.DB.Query("SELECT id FROM posts LIMIT 5")
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
