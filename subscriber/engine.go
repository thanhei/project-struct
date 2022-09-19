package subscriber

import (
	"context"
	"fmt"
	"go-training/common"
	"go-training/component/app_context"
	"go-training/component/asyncjob"
	"go-training/pubsub"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx app_context.AppContext
}

func NewEngine(appCtx app_context.AppContext) *consumerEngine {
	return &consumerEngine{
		appCtx: appCtx,
	}
}

func (e *consumerEngine) Start() error {
	e.startSubTopic(
		common.TopicUserLikeRestaurant,
		true,
		RunIncreaseLikeCountAfterUserLikeRestaurant(e.appCtx),
	)

	return nil
}

func (e *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, jobs ...consumerJob) error {
	c, _ := e.appCtx.GetPubsub().Subscribe(context.Background(), topic)

	for _, item := range jobs {
		fmt.Println("Setup consumer for: ", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			fmt.Println("Running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for {
			msg := c

			jobHdlArr := make([]asyncjob.Job, len(jobs))

			for i := range jobs {
				jobHdlArr[i] = asyncjob.NewJob(getJobHandler(&jobs[i], <-msg))
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				fmt.Println("Error when run job group: ", err)
				fmt.Println(err)
			}
		}
	}()

	return nil
}
