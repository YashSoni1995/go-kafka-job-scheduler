package jobs

import (
	"go-kafka-job-scheduler/config"
)

type Jobs struct {
	TableName  	struct{} 	`sql:"jobs" json:"-"`
	Id 		   	string 		`sql:"id"` 
	Task 		string 		`sql:"task"`
	Topic 		string 		`sql:"topic"`	
	Schedule 	string 		`sql:"schedule"`
}

func ReadJobs(c config.Config) (jobs []Jobs, err error){
	err = c.PG.Model(&Jobs{}).Select(&jobs)
	return
}
