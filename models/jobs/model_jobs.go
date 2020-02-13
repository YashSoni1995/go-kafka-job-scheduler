package jobs

import (
	"goyash/golang-kafka-job-scheduler/config"
)

type Jobs struct {
	TableName  	struct{} 	`sql:"jobs" json:"-"`
	Id 		   	string 		`sql:"id"` 
	Query 		string 		`sql:"query"`
	Schedule 	string 		`sql:"schedule"`
}

func ReadJobs(c config.Config) (jobs []Jobs, err error){
	err = c.PG.Model(&Jobs{}).
		Limit(1).
		Select(&jobs)
	return
}
