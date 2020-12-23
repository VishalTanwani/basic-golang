package cli

import (
	"fmt"
	Ass "mini/assignment"
	"strings"
)

func Run() {
	fmt.Println("this is go lang mini assignment")
	fmt.Println("Select the option")
	fmt.Println("1\tGet all questions with all the details")
	fmt.Println("2\tGet all comments for a post")
	fmt.Println("3\tGet answers for these questions")
	fmt.Println("4\tDetails of a user")
	fmt.Println("5\tlist questions by tags")
	var op int
	var postID int
	fmt.Scanln(&op)
	switch op {
	case 1:
		Ass.First()
	case 2:
		fmt.Println("Please enter post id to comments on that post")
		fmt.Scanln(&postID)
		Ass.Second(postID)
	case 3:
		var op int
		fmt.Println("select the option \n1\tsort by Score\n2\tsort by date of creation\n3\tsort by last activity date\n4\tsort by Score\n5\tsort by date of creation\n6\tsort by last activity date")
		fmt.Scanln(&op)
		switch op {
		case 1:
			Ass.ThirdSortByScore()
		case 2:
			Ass.ThirdSortByCreationDate()
		case 3:
			Ass.ThirdSortByLastEditDate()
		case 4:
			Ass.ThirdSortByScoreId(postID)
		case 5:
			Ass.ThirdSortByCreationDateId(postID)
		case 6:
			Ass.ThirdSortByLastEditDateId(postID)
		default:
			fmt.Println("select proper option")
		}
	case 4:
		var op int
		fmt.Println("\n1\tsearch by id\n2\tsearch by name")
		fmt.Scanln(&op)
		switch op {
		case 1:
			var uid int
			fmt.Println("Please enter user id for search")
			fmt.Scanln(&uid)
			Ass.ForthById(uid)
		case 2:
			var uname string
			fmt.Println("Please enter user name for search")
			fmt.Scanln(&uname)
			Ass.FortnByName(uname)
		default:
			fmt.Println("select proper option")
		}
	case 5:
		var names string
		fmt.Println("enter tag name seprated by ,(coma)")
		fmt.Scanln(&names)
		values := strings.Split(names, ",")
		Ass.Fifth(values...)
	default:
		fmt.Println("there are more easy so be happy")
	}
}
