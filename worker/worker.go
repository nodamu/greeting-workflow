package main

import (
	app "github.com/myworkflow"
	"github.com/myworkflow/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Could not start Temporal Client",err)
	}
	defer c.Close()

	w := worker.New(c,app.GreetingTaskQueue,worker.Options{})
	w.RegisterActivity(app.ComposeGreeting)
	w.RegisterWorkflow(workflows.MyWorkflow)
	// Start listening to task queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker ", err)
	}
}
