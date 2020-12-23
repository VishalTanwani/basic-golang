package badge

/*
	name: badges
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

type Badge struct {
	ID       int
	UserID   int
	Date     string
	Name     string
	Class    int
	TagBased string
}

var badge []Badge

func init() {
	data, err := ioutil.ReadFile("../Json Data/Badges.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &badge)
	if err != nil {
		panic(err)
	}

	fmt.Println("data extracted from file")
}

func (value Badge) add() {
	fmt.Println("# Inserting values")
	var lastInsertID int
	err := db.DB.QueryRow("INSERT INTO badges(id, user_id, date, name, class, tagbased) VALUES($1,$2,$3,$4,$5,$6) returning id;", value.ID, value.UserID, value.Date, value.Name, value.Class, value.TagBased).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}
	fmt.Println("last inserted id =", lastInsertID)
}

func Add() {
	for _, value := range badge {
		value.add()
		// fmt.Println(value.LastEditorUserID)
	}
}

func DeleteAll() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DELETE FROM badges")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted All")
}

func DropTable() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DROP TABLE badges")
	if err != nil {
		panic(err)
	}
	fmt.Println("Table Deleted")
}

func ShowIDS() {
	fmt.Println("# Querying")
	rows, err := db.DB.Query("SELECT id FROM badges")
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
	fmt.Println("# Querying badges")
	rows, err := db.DB.Query("SELECT id FROM badges LIMIT 5")
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
