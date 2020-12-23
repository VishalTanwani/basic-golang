package dump

import (
	"encoding/json"
	"fmt"
	"mini/db"
	"os"
)

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
type test struct {
	name string
	id   int
}

var rows []Post

func Run() {
	row, err := db.DB.Query("SELECT id, post_type_id, score, view_count, tags, answer_count, comment_count, favourite_count, creation_date, body, closed_date, accepted_answer_id, parent_id, owner_user_id, owner_display_name, last_edit_date, last_activity_date, title, community_owned_date FROM posts WHERE post_type_id = 1 AND answer_count = 0")
	if err != nil {
		panic(err)
	}
	for row.Next() {
		var r Post
		row.Scan(&r.ID, &r.PostTypeID, &r.Score, &r.ViewCount, &r.Tags, &r.AnswerCount, &r.CommentCount, &r.FavoriteCount, &r.CreationDate, &r.Body, &r.ClosedDate, &r.AcceptedAnswerID, &r.ParentID, &r.OwnerUserID, &r.OwnerDisplayName, &r.LastEditDate, &r.LastActivityDate, &r.Title, &r.CommunityOwnedDate)
		// fmt.Println(r)
		rows = append(rows, r)
	}
	fmt.Println(rows)
	//1
	fileName := "unanswered_ques.json"
	file, err := os.Create(fileName)
	encode := json.NewEncoder(file)
	encode.SetIndent("", "  ")
	encode.Encode(rows)
	//2
	// rankingsJson, _ := json.MarshalIndent(rows, "", "  ")
	// err = ioutil.WriteFile("unanswered_ques.json", rankingsJson, 0644)
	// fmt.Printf("%+v", rows)
	//3
	// buf := new(bytes.Buffer)
	// encode := json.NewEncoder(buf)
	// encode.Encode(rr)
	// file, err := os.Create("unanswered_ques.json")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(rr)
	// defer file.Close()
	// io.Copy(file, buf)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// io.Copy(os.Stdout, buf)

}
