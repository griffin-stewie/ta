package command

import (
	"github.com/codegangsta/cli"
	backlog "github.com/griffin-stewie/go-backlog"
	chatwork "github.com/griffin-stewie/go-chatwork"

	"fmt"
	"log"
	"os"
)

// CmdList is
func CmdList(c *cli.Context) {
	os.Exit(CmdListConcreate(c))
}

// CmdListConcreate is main part
func CmdListConcreate(c *cli.Context) int {

	// Backlog
	backlogToken := c.GlobalString("backlog-token")
	if len(backlogToken) == 0 {
		log.Printf("[ERROR] %s", "You need Backlog API-Key")
		return 1
	}

	backlogEndpoint := c.GlobalString("backlog-endpoint")
	if len(backlogEndpoint) == 0 {
		log.Printf("[ERROR] %s", "You need Backlog API endpoint")
		return 1
	}

	backlogClient := backlog.NewClient(backlogEndpoint, backlogToken)
	me, err := backlogClient.Myself()

	opt := &backlog.IssuesOption{
		AssigneeIds: []int{*me.ID},
		Statuses:    []backlog.IssueStatus{1, 2, 3},
	}

	issues, err := backlogClient.IssuesWithOption(opt)

	if err != nil {
		log.Printf("[ERROR] Backlog request fails '%s'", err)
		return 1
	}

	log.Printf("[DEBUG] issues: %#+v", issues)

	for _, item := range issues {
		url := backlogEndpoint + "/view" + "/" + *item.IssueKey

		due := ""
		if item.DueDate != nil {
			time := *item.DueDate
			due = time.String()
		}

		fmt.Printf("%v, %v, %v\n", *item.Summary, url, due)
	}

	// ChatWork
	chatworkToken := c.GlobalString("chatwork-token")
	if len(chatworkToken) == 0 {
		log.Printf("[ERROR] %s", "You need Chatwork API-Key")
		return 1
	}

	chatworkClient := chatwork.NewClient(chatworkToken)
	myTasks := chatworkClient.MyTasks(nil)
	for _, item := range myTasks {
		timeString := ""
		if item.LimitTime > 0 {
			timeString = item.LimitDate().String()
		}
		fmt.Printf("%v, %v\n", item.Body, timeString)
	}

	return 0
}
