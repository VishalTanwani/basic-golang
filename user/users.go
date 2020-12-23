package user

/*
	name: users
*/
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	db "mini/db"
	"net/http"
)

var user []User

type basic interface {
	add()
}

type User struct {
	ID              int
	Reputation      int
	CreationDate    string
	DisplayName     string
	LastAccessDate  string
	WebsiteURL      string
	Location        string
	AboutMe         string
	Views           int
	UpVotes         int
	DownVotes       int
	AccountID       int
	ProfileImageURL string
}

func init() {
	data, err := ioutil.ReadFile("../Json Data/Users.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &user)
	if err != nil {
		panic(err)
	}
	fmt.Println("data extracted from file")

}

func (value User) add() {
	fmt.Println("# Inserting values")
	var lastInsertID int
	err := db.DB.QueryRow("INSERT INTO users VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) returning id;", value.ID, value.Reputation, value.CreationDate, value.DisplayName, value.LastAccessDate, value.WebsiteURL, value.Location, value.AboutMe, value.Views, value.UpVotes, value.DownVotes, value.AccountID, value.ProfileImageURL).Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}
	fmt.Println("last inserted id =", lastInsertID)
}

func Add() {
	for _, value := range user {
		value.add()
	}
}

func DeleteAll() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DELETE FROM users")
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted All")
}

func DropTable() {
	fmt.Println("# Querying")
	_, err := db.DB.Query("DROP TABLE users")
	if err != nil {
		panic(err)
	}
	fmt.Println("Table Deleted")
}

func ShowIDS() {
	fmt.Println("# Querying")
	rows, err := db.DB.Query("SELECT id FROM users")
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
	fmt.Println("# Querying")
	rows, err := db.DB.Query("SELECT id FROM users LIMIT 5")
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

func AllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
