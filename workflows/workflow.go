package workflows

import (
	app "github.com/myworkflow"
	"go.temporal.io/sdk/workflow"
	"time"
)

func MyWorkflow(ctx workflow.Context,name string )(string,error){
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second *5,
	}

	ctx = workflow.WithActivityOptions(ctx,options)
	var result string
	err := workflow.ExecuteActivity(ctx,app.ComposeGreeting,name).Get(ctx,&result)
	return result, err
}