package postlink

/*
	name: post_link
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

type PostLink struct {
	ID            int
	CreationDate  string
	PostID        int
	RelatedPostID int
	LinkTypeID    int
}

var postLink []PostLink

func init() {
	data, err := ioutil.ReadFile("../Json Data/PostLinks.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &postLink)
	if err != nil {
		panic(err)
	}

	fmt.Println("data extracted from file")
}

func (value PostLink) add() {
	fmt.Println("# Inserting values")
	var lastInsertID int
	err := db.DB.QueryRow("INSERT INTO post_link(id,post_id,creation_date,related_post_id,link_type_id)VALUES($1,$2,$3,$4,$5) returning id;", value.ID, value.PostID, value.CreationDate, value.RelatedPostID, value.LinkTypeID).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}
	fmt.Println("last inserted id =", lastInsertID)
}

func Add() {
	for _, value := range postLink {
		value.add()
		// fmt.Println(value.LastEditorUserID)
	}
}

func DeleteAll() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DELETE FROM post_link")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted All")
}

func DropTable() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DROP TABLE post_link")
	if err != nil {
		panic(err)
	}
	fmt.Println("Table Deleted")
}

func ShowIDS() {
	fmt.Println("# Querying")
	rows, err := db.DB.Query("SELECT id FROM post_link")
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
	fmt.Println("# Querying post_link")
	rows, err := db.DB.Query("SELECT id FROM post_link LIMIT 5")
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
