package assignment

import (
	"fmt"
	"mini/db"
)

func First() {
	fmt.Println("Get all questions with all the details")
	rows, err := db.DB.Query("SELECT id, view_count, answer_count, comment_count, favourite_count, closed_date, title FROM posts where post_type_id = 1")
	if err != nil {
		panic(err)
	}

	// fmt.Println(rows)
	fmt.Println()
	fmt.Println()
	fmt.Println("id  ||  view_count  ||  answer_count  ||  comment_count  ||  favourite_count  ||  closed_date  ||  title")
	fmt.Println()
	fmt.Println()

	for rows.Next() {
		var id int
		var view_count int
		var answer_count int
		var comment_count int
		var favourite_count int
		var closed_date string
		var title string
		err = rows.Scan(&id, &view_count, &answer_count, &comment_count, &favourite_count, &closed_date, &title)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v  ||  %v  ||  %v  ||  %v  ||  %v  ||  %v  ||  %v\n", id, view_count, answer_count, comment_count, favourite_count, closed_date, title)
	}
}

func Second(postId int) {
	fmt.Println("Get all comments for a post")
	rows, err := db.DB.Query("SELECT * FROM comments where post_id = $1", postId)
	if err != nil {
		panic(err)
	}

	// fmt.Println(rows)
	fmt.Println()
	fmt.Println()
	fmt.Println("id  ||  post_id  ||  creation_date  ||  text  ||  user_id  ||  score  ||  user_display_name")
	fmt.Println()
	fmt.Println()

	for rows.Next() {
		var id int
		var post_id int
		var creation_date string
		var text string
		var user_id int
		var score int
		var user_display_name string
		err = rows.Scan(&id, &post_id, &creation_date, &text, &user_id, &score, &user_display_name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v  ||  %v  ||  %v  ||  %v  ||  %v  ||  %v  ||  %v\n", id, post_id, creation_date, text, user_id, score, user_display_name)
	}
}

func ThirdSortByScore() {
	fmt.Println("Get answers for these questions")
	rows, err := db.DB.Query("SELECT id, body FROM posts where post_type_id = 1")
	if err != nil {
		panic(err)
	}

	// fmt.Println(rows)
	fmt.Println()
	fmt.Println()
	fmt.Println("id  ||  Questions")
	fmt.Println()
	fmt.Println()

	for rows.Next() {
		var id int
		var body string
		err = rows.Scan(&id, &body)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("\n%v", id)
		rows1, err := db.DB.Query("SELECT id, body FROM posts where post_type_id = 2 AND parent_id = $1 ORDER BY score", id)
		fmt.Printf("\t%v  ||  %v\n", id, body)
		fmt.Println("\t")
		fmt.Println("\t")
		fmt.Println("\tid  ||  Answers")
		fmt.Println("\t")
		fmt.Println("\t")
		for rows1.Next() {
			var id int
			var body string
			err = rows1.Scan(&id, &body)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\t%v  ||  %v\n", id, body)
		}
	}
}

func ThirdSortByCreationDate() {
	fmt.Println("Get answers for these questions")
	rows, err := db.DB.Query("SELECT id, body FROM posts where post_type_id = 1")
	if err != nil {
		panic(err)
	}

	// fmt.Println(rows)
	fmt.Println()
	fmt.Println()
	fmt.Println("id  ||  Questions")
	fmt.Println()
	fmt.Println()

	for rows.Next() {
		var id int
		var body string
		err = rows.Scan(&id, &body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\t%v  ||  %v\n", id, body)
		rows1, err := db.DB.Query("SELECT id, body FROM posts where post_type_id = 2 AND parent_id = $1 ORDER BY creation_date", id)
		fmt.Println("\t")
		fmt.Println("\t")
		fmt.Println("\tid  ||  Answers")
		fmt.Println("\t")
		fmt.Println("\t")
		for rows1.Next() {
			var id int
			var body string
			err = rows1.Scan(&id, &body)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\t%v  ||  %v\n", id, body)
		}

	}
}

func ThirdSortByLastEditDate() {
	fmt.Println("Get answers for these questions")
	rows, err := db.DB.Query("SELECT id, body FROM posts where post_type_id = 1")
	if err != nil {
		panic(err)
	}

	// fmt.Println(rows)
	fmt.Println()
	fmt.Println()
	fmt.Println("id  ||  Questions")
	fmt.Println()
	fmt.Println()

	for rows.Next() {
		var id int
		var body string
		err = rows.Scan(&id, &body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\t%v  ||  %v\n", id, body)
		rows1, err := db.DB.Query("SELECT id, body FROM posts where post_type_id = 2 AND parent_id = $1 ORDER BY last_edit_date", id)
		fmt.Println("\t")
		fmt.Println("\t")
		fmt.Println("\tid  ||  Answers")
		fmt.Println("\t")
		fmt.Println("\t")
		for rows1.Next() {
			var id int
			var body string
			err = rows1.Scan(&id, &body)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\t%v  ||  %v\n", id, body)
		}

	}
}

func ThirdSortByScoreId(postId int) {
	fmt.Println("Get answers for these questions")
	rows, err := db.DB.Query("SELECT id, body FROM posts where post_type_id = 1 AND id = $1", postId)
	if err != nil {
		panic(err)
	}

	// fmt.Println(rows)
	fmt.Println()
	fmt.Println()
	fmt.Println("id  ||  Questions")
	fmt.Println()
	fmt.Println()

	for rows.Next() {
		var id int
		var body string
		err = rows.Scan(&id, &body)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("\n%v", id)
		rows1, err := db.DB.Query("SELECT id, body, score FROM posts where post_type_id = 2 AND parent_id = $1 ORDER BY score", id)
		fmt.Printf("\t%v  ||  %v\n", id, body)
		fmt.Println("\t")
		fmt.Println("\t")
		fmt.Println("\tid  ||  Answers")
		fmt.Println("\t")
		fmt.Println("\t")
		for rows1.Next() {
			var id int
			var body string
			var score int
			err = rows1.Scan(&id, &body, &score)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\t%v  ||  %v  ||  %v\n", id, body, score)
		}
	}
}

func ThirdSortByCreationDateId(postId int) {
	fmt.Println("Get answers for these questions")
	rows, err := db.DB.Query("SELECT id, body FROM posts where post_type_id = 1 AND id = $1", postId)
	if err != nil {
		panic(err)
	}

	// fmt.Println(rows)
	fmt.Println()
	fmt.Println()
	fmt.Println("id  ||  Questions")
	fmt.Println()
	fmt.Println()

	for rows.Next() {
		var id int
		var body string
		err = rows.Scan(&id, &body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\t%v  ||  %v\n", id, body)
		rows1, err := db.DB.Query("SELECT id, body, creation_date FROM posts where post_type_id = 2 AND parent_id = $1 ORDER BY creation_date", id)
		fmt.Println("\t")
		fmt.Println("\t")
		fmt.Println("\tid  ||  Answers")
		fmt.Println("\t")
		fmt.Println("\t")
		for rows1.Next() {
			var id int
			var body string
			var creation_date int
			err = rows1.Scan(&id, &body, &creation_date)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\t%v  ||  %v  ||  %v\n", id, body, creation_date)
		}

	}
}

func ThirdSortByLastEditDateId(postId int) {
	fmt.Println("Get answers for these questions")
	rows, err := db.DB.Query("SELECT id, body FROM posts where post_type_id = 1 AND id = $1", postId)
	if err != nil {
		panic(err)
	}

	// fmt.Println(rows)
	fmt.Println()
	fmt.Println()
	fmt.Println("id  ||  Questions")
	fmt.Println()
	fmt.Println()

	for rows.Next() {
		var id int
		var body string
		err = rows.Scan(&id, &body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\t%v  ||  %v\n", id, body)
		rows1, err := db.DB.Query("SELECT id, body, last_edit_date FROM posts where post_type_id = 2 AND parent_id = $1 ORDER BY last_edit_date", id)
		fmt.Println("\t")
		fmt.Println("\t")
		fmt.Println("\tid  ||  Answers")
		fmt.Println("\t")
		fmt.Println("\t")
		for rows1.Next() {
			var id int
			var body string
			var last_edit_date int
			err = rows1.Scan(&id, &body, &last_edit_date)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\t%v  ||  %v  ||  %v\n", id, body, last_edit_date)
		}

	}
}

func ForthById(uid int) {
	fmt.Println("Details of a user")
	var id int
	var display_name string
	var reputation string
	var location string
	var totalQ int
	var totalQA int
	// var rank int
	var allB []string
	res, err := db.DB.Query("SELECT id,display_name,reputation,location FROM users WHERE id = $1", uid)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	for res.Next() {
		err = res.Scan(&id, &display_name, &reputation, &location)
		if err != nil {
			panic(err)
		}
	}
	row, err := db.DB.Query("SELECT name FROM badges WHERE user_id = $1", uid)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	for row.Next() {
		var name string
		err = row.Scan(&name)
		if err != nil {
			panic(err)
		}
		allB = append(allB, name)
	}
	row1, err := db.DB.Query("SELECT COUNT(id) FROM posts WHERE owner_user_id = $1 AND post_type_id = 1", uid)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	var count int
	for row1.Next() {
		err = row1.Scan(&count)
		if err != nil {
			panic(err)
		}
	}
	totalQ = count
	row2, err := db.DB.Query("SELECT COUNT(id) FROM posts WHERE owner_user_id = $1 AND post_type_id = 2", id)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	// var count int
	for row2.Next() {
		err = row2.Scan(&count)
		if err != nil {
			panic(err)
		}
	}
	totalQA = count
	fmt.Println(id, display_name, reputation, location, totalQ, totalQA, allB)

}

func FortnByName(uname string) {
	fmt.Println("Details of a user")
	var id int
	var display_name string
	var reputation string
	var location string
	var totalQ int
	var totalQA int
	// var rank int
	var allB []string
	res, err := db.DB.Query("SELECT id,display_name,reputation,location FROM users WHERE display_name = $1", uname)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	for res.Next() {
		err = res.Scan(&id, &display_name, &reputation, &location)
		if err != nil {
			panic(err)
		}
	}
	row, err := db.DB.Query("SELECT name FROM badges WHERE user_id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	for row.Next() {
		var name string
		err = row.Scan(&name)
		if err != nil {
			panic(err)
		}
		allB = append(allB, name)
	}
	row1, err := db.DB.Query("SELECT COUNT(id) FROM posts WHERE owner_user_id = $1 AND post_type_id = 1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	var count int
	for row1.Next() {
		err = row1.Scan(&count)
		if err != nil {
			panic(err)
		}
	}
	totalQ = count
	row2, err := db.DB.Query("SELECT COUNT(id) FROM posts WHERE owner_user_id = $1 AND post_type_id = 2", id)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	// var count int
	for row2.Next() {
		err = row2.Scan(&count)
		if err != nil {
			panic(err)
		}
	}
	totalQA = count
	fmt.Println(id, display_name, reputation, location, totalQ, totalQA, allB)

}
func Fifth(name ...string) {
	query := "SELECT Body FROM posts WHERE "
	for i, v := range name {
		if i > 0 {
			query = query + " AND "
		}
		query = query + "tags LIKE " + "'%" + v + "%'"
	}
	fmt.Println(query)
	row, err := db.DB.Query(query)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	// var count int
	for row.Next() {
		var body string
		err = row.Scan(&body)
		if err != nil {
			panic(err)
		}
		fmt.Println(body)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}

}
