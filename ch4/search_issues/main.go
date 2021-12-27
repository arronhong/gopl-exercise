package main

import (
	"fmt"
	"gopl-exercise/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var lt1mon []*github.Issue
	var lt1year []*github.Issue
	var gte1year []*github.Issue
	now := time.Now()
	for _, iss := range result.Items {
		switch {
		case iss.CreatedAt.After(now.AddDate(0, -1, 0)):
			lt1mon = append(lt1mon, iss)
		case iss.CreatedAt.After(now.AddDate(-1, 0, 0)):
			lt1year = append(lt1year, iss)
		default:
			gte1year = append(gte1year, iss)
		}
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("less 1 month:")
	for _, item := range lt1mon {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Println("\nless 1 year:")
	for _, item := range lt1year {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	fmt.Println("\ngreater then 1 year:")
	for _, item := range gte1year {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
