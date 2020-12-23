package Post

/*
	name: votes
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

type Vote struct {
	ID           int
	PostID       int
	VoteTypeID   int
	CreationDate string
	UserID       int
	BountyAmount int
}

var vote []Vote

func init() {
	data, err := ioutil.ReadFile("../Json Data/Votes.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &vote)
	if err != nil {
		panic(err)
	}

	fmt.Println("data extracted from file")
}

func (value Vote) add() {
	fmt.Println("# Inserting values")
	var lastInsertID int
	err := db.DB.QueryRow("INSERT INTO votes(id,post_id,creation_date,vote_type_id,user_id,bounty_amount) VALUES($1,$2,$3,$4,$5,$6) returning id;", value.ID, value.PostID, value.CreationDate, value.VoteTypeID, value.UserID, value.BountyAmount).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}
	fmt.Println("last inserted id =", lastInsertID)
}

func Add() {
	for _, value := range vote {
		value.add()
		// fmt.Println(value.LastEditorUserID)
	}
}

func DeleteAll() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DELETE FROM votes")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted All")
}

func DropTable() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DROP TABLE votes")
	if err != nil {
		panic(err)
	}
	fmt.Println("Table Deleted")
}

func ShowIDS() {
	fmt.Println("# Querying")
	rows, err := db.DB.Query("SELECT id FROM votes")
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
	fmt.Println("# Querying votes")
	rows, err := db.DB.Query("SELECT id FROM votes LIMIT 5")
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
