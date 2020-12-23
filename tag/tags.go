package tag

/*
	name: tags
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

type Tag struct {
	ID            int
	TagName       string
	Count         int
	ExcerptPostID int
	WikiPostID    int
}

var tag []Tag

func init() {
	data, err := ioutil.ReadFile("../Json Data/Tags.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &tag)
	if err != nil {
		panic(err)
	}

	fmt.Println("data extracted from file")
}

func (value Tag) add() {
	fmt.Println("# Inserting values")
	var lastInsertID int
	err := db.DB.QueryRow("INSERT INTO tags(id,tag_name,count,excerpt_post_id,wiki_post_id)VALUES($1,$2,$3,$4,$5) returning id;", value.ID, value.TagName, value.Count, value.ExcerptPostID, value.WikiPostID).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}
	fmt.Println("last inserted id =", lastInsertID)
}

func Add() {
	for _, value := range tag {
		value.add()
		// fmt.Println(value.LastEditorUserID)
	}
}

func DeleteAll() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DELETE FROM tags")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted All")
}

func DropTable() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DROP TABLE tags")
	if err != nil {
		panic(err)
	}
	fmt.Println("Table Deleted")
}

func ShowIDS() {
	fmt.Println("# Querying")
	rows, err := db.DB.Query("SELECT id FROM tags")
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
	fmt.Println("# Querying tags")
	rows, err := db.DB.Query("SELECT id FROM tags LIMIT 5")
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
