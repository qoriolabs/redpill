package redpill

import (
	"math/rand"
	"time"
)

func GetDashboardEventFeed() <-chan *DashboardEvent {
	events := make(chan *DashboardEvent)

	go func() {
		for {
			event := examples[rand.Intn(len(examples))]
			event.User = users[rand.Intn(len(users))]
			event.Timestamp = time.Now().Unix()

			send := event
			events <- &send
			time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
		}
	}()

	return events
}

var (
	examples = []DashboardEvent{
		DashboardEvent{
			Type:        "git_commit",
			Title:       "Git commit",
			Description: "Commit by user to repo",
			Url:         "https://github.com/project/tree/branch/12345edcba",
		},
		DashboardEvent{
			Type:        "env_change",
			Title:       "Environment variable change",
			Description: "Environment variable updated by user",
			Url:         "http://server.com/domain/config/myapp/env/var1",
		},
		DashboardEvent{
			Type:        "deployment",
			Title:       "Deployment to Staging",
			Description: "New code pushed to Staging",
			Url:         "http://server.com/domain/deploy/12345",
		},
		DashboardEvent{
			Type:        "deployment",
			Title:       "Deployment to Production",
			Description: "New code pushed to production",
			Url:         "http://server.com/domain/deploy/1112345",
		},
		DashboardEvent{
			Type:        "setlive",
			Title:       "Load balancer updated in Staging",
			Description: "Staging load balancer updated",
			Url:         "http://server.com/domain/setlive/1112345",
		},
		DashboardEvent{
			Type:        "setlive",
			Title:       "Load balancer updated in Production",
			Description: "Production load balancer updated",
			Url:         "http://server.com/domain/setlive/2222345",
		},
		DashboardEvent{
			Type:        "run_instances",
			Title:       "Instances run",
			Description: "New instances are launched",
			Url:         "http://server.com/domain/resources/compute/start/1124",
		},
		DashboardEvent{
			Type:        "terminate_instances",
			Title:       "Instances terminated",
			Description: "EC2 instances are terminated in subnet-12cbd212",
			Url:         "http://server.com/domain/resources/compute/terminate/1124",
		},
	}

	users = []string{
		"john", "jason", "sergey", "philip", "david", "jane", "bob",
	}
)
