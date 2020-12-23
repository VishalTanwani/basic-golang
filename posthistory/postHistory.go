package posthistory

/*
	name: post_history
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

type PostHistory struct {
	ID                int
	PostHistoryTypeID int
	PostID            int
	RevisionGUID      string
	CreationDate      string
	UserID            int
	Text              string
	UserDisplayName   string
	Comment           string
}

var postHistory []PostHistory

func init() {
	data, err := ioutil.ReadFile("../Json Data/PostHistory.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &postHistory)
	if err != nil {
		panic(err)
	}

	fmt.Println("data extracted from file")
}

func (value PostHistory) add() {
	fmt.Println("# Inserting values")
	var lastInsertID int
	err := db.DB.QueryRow("INSERT INTO post_history(id,post_history_type_id,post_id,revision_guid,creation_date,text,user_id,user_display_name,comment)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id;", value.ID, value.PostHistoryTypeID, value.PostID, value.RevisionGUID, value.CreationDate, value.Text, value.UserID, value.UserDisplayName, value.Comment).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}
	fmt.Println("last inserted id =", lastInsertID)
}

func Add() {
	for _, value := range postHistory {
		value.add()
		// fmt.Println(value.LastEditorUserID)
	}
}

func DeleteAll() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DELETE FROM post_history")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted All")
}

func DropTable() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DROP TABLE post_history")
	if err != nil {
		panic(err)
	}
	fmt.Println("Table Deleted")
}

func ShowIDS() {
	fmt.Println("# Querying")
	rows, err := db.DB.Query("SELECT id FROM post_history")
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
	fmt.Println("# Querying post_history")
	rows, err := db.DB.Query("SELECT id FROM post_history LIMIT 5")
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
