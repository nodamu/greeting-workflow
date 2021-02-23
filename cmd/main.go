package main

import (
	"context"
	"fmt"
	app "github.com/myworkflow"
	"github.com/myworkflow/workflows"
	"go.temporal.io/sdk/client"
	"log"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal Client ",err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		TaskQueue: app.GreetingTaskQueue,
		ID: "greeting-workflow-id",
	}
	var name = "Nick"
	we , err := c.ExecuteWorkflow(context.Background(), options,workflows.MyWorkflow,name)
	if err != nil {
		log.Fatalln("Unable to complete workflow ", err)
	}
	var greeting string
	err = we.Get(context.Background(), greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())
	fmt.Printf("\n%s\n\n", greeting)
}
