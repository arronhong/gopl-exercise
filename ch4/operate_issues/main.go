package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"gopl-exercise/ch4/github"
)

func CreateIssue(owner, repo string) (*github.Issue, error) {
	tmpfile, err := ioutil.TempFile("./", "*")
	if err != nil {
		return nil, fmt.Errorf("create tmp file fail: %s", err)
	}
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	cmd := exec.Command("vim", tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("open vim fail: %s", err)
	}

	tmpfile, err = os.Open(tmpfile.Name())
	if err != nil {
		return nil, fmt.Errorf("reopen tmp file fail: %s", err)
	}
	defer tmpfile.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST",
		fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo),
		tmpfile,
	)
	if err != nil {
		return nil, fmt.Errorf("new request fail: %s", err)
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "bearer "+os.Getenv("PAT"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request github fail: %s", err)
	}
	defer resp.Body.Close()

	issue := new(github.Issue)
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(issue); err != nil {
		return nil, fmt.Errorf("decode body fail: %s", err)
	}
	return issue, nil
}

func UpdateIssue(owner, repo string, number int) (*github.Issue, error) {
	tmpfile, err := ioutil.TempFile("./", "*")
	if err != nil {
		return nil, fmt.Errorf("create tmp file fail: %s", err)
	}
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	cmd := exec.Command("vim", tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("open vim fail: %s", err)
	}

	tmpfile, err = os.Open(tmpfile.Name())
	if err != nil {
		return nil, fmt.Errorf("reopen tmp file fail: %s", err)
	}
	defer tmpfile.Close()

	client := &http.Client{}
	req, err := http.NewRequest("PATCH",
		fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%d",
			owner, repo, number),
		tmpfile,
	)
	if err != nil {
		return nil, fmt.Errorf("new request fail: %s", err)
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "bearer "+os.Getenv("PAT"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request github fail: %s", err)
	}
	defer resp.Body.Close()

	issue := new(github.Issue)
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(issue); err != nil {
		return nil, fmt.Errorf("decode body fail: %s", err)
	}
	return issue, nil
}

func ListIssues(owner, repo string) ([]github.Issue, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo),
		nil)
	if err != nil {
		return nil, fmt.Errorf("new request fail: %s", err)
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "bearer "+os.Getenv("PAT"))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request github fail: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body fail: %s", err)
	}

	issues := []github.Issue{}
	if err = json.Unmarshal(body, &issues); err != nil {
		return nil, fmt.Errorf("unmarshal github resp fail: %s", err)
	}
	return issues, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("invalid arguments number")
	}
	switch os.Args[1] {
	case "create":
		if len(os.Args) != 4 {
			log.Fatal(
				"invalid arguments number, argument like: " +
					"create <owner> <repo>")
		}
		issue, err := CreateIssue(os.Args[2], os.Args[3])
		if err != nil {
			log.Fatalf("create issue fail: %s", err)
		}
		fmt.Printf("issue created, see %s", issue.HTMLURL)
	case "list":
		if len(os.Args) != 4 {
			log.Fatal(
				"invalid arguments number, argument like: " +
					"list <owner> <repo>")
		}
		issues, err := ListIssues(os.Args[2], os.Args[3])
		if err != nil {
			log.Fatalf("list issues fail: %s", err)
		}
		for _, issue := range issues {
			fmt.Printf("title: %s\tstate: %s\tcreated by: %s\n%s\n\n%s\n\n",
				issue.Title,
				issue.State,
				issue.User.Login,
				issue.HTMLURL,
				issue.Body)
		}
	case "update":
		if len(os.Args) != 5 {
			log.Fatal(
				"invalid arguments number, argument like: " +
					"update <owner> <repo> <number>")
		}
		number, err := strconv.Atoi(os.Args[4])
		if err != nil {
			log.Fatalf("invalid number argument")
		}
		issue, err := UpdateIssue(os.Args[2], os.Args[3], number)
		if err != nil {
			log.Fatalf("update issue fail: %s", err)
		}
		fmt.Printf("issue updated, see %s", issue.HTMLURL)
	default:
		log.Fatal("unknown operation")
	}
}
