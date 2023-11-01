package subscriber

import (
	"context"
	"fmt"
	"go-training/internal/common"
	"go-training/internal/component/asyncjob"
	"go-training/internal/component/pubsub"
	restaurantrepo "go-training/internal/modules/restaurant/repository"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type ConsumerEngine struct {
	pubsub         pubsub.Pubsub
	restaurantRepo restaurantrepo.RestaurantRepository
}

func NewEngine(pubsub pubsub.Pubsub, restaurantRepo restaurantrepo.RestaurantRepository) *ConsumerEngine {
	return &ConsumerEngine{
		pubsub:         pubsub,
		restaurantRepo: restaurantRepo,
	}
}

func (e *ConsumerEngine) Start() error {
	e.startSubTopic(
		common.TopicUserLikeRestaurant,
		true,
		RunIncreaseLikeCountAfterUserLikeRestaurant(e.restaurantRepo),
		NotificationAfterUserLikeRestaurant(),
	)

	e.startSubTopic(
		common.TopicUserDislikeRestaurant,
		true,
		DecreaseLikeCountAfterUserUnLikeRestaurant(e.restaurantRepo),
	)

	return nil
}

func (e *ConsumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, jobs ...consumerJob) error {
	c, _ := e.pubsub.Subscribe(context.Background(), topic)

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
			msg := <-c

			jobHdlArr := make([]asyncjob.Job, len(jobs))

			for i := range jobs {
				jobHdlArr[i] = asyncjob.NewJob(getJobHandler(&jobs[i], msg))
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
