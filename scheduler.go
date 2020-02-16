package main

import (
	"fmt"
	"log"
	"go-kafka-job-scheduler/config"
	"gopkg.in/robfig/cron.v2"
	"net/http"
	"go-kafka-job-scheduler/models/jobs"
	"go-kafka-job-scheduler/kafka"
)

var jobIds = []cron.EntryID{0}


func ExecuteJob (config config.Config, job jobs.Jobs) (string, error){
	return fmt.Sprintf(job.Task), nil
}

func ExecuteJobAndPushToKafka(config config.Config, job jobs.Jobs) {
	message, err := ExecuteJob(config, job)
	topic := job.Topic
	log.Println("Push message to kafka\n", "topic - ", topic, "| message - ", message, "with error", err)
	//push message received after executing the job to kafka
	kafka.PushMessageToKafka(config, topic, message)
}

func RemoveJobsFromCron(c *cron.Cron, jobIds []cron.EntryID) {
	log.Println("removing previously added jobs and rescheduling again...")
	for _,id := range jobIds {
		c.Remove(id)
	}
}

func AddJobsToCron (c *cron.Cron, config config.Config, jobs []jobs.Jobs) {
	//first removing previously added jobs to avoid duplication and ensure updated jobs run everytime
	RemoveJobsFromCron(c, jobIds)
	jobIds = []cron.EntryID{0}
	//executing each job acc to schedule and pushing data to kafka
	for _,job := range jobs {
		jobId,_ := c.AddFunc(job.Schedule, func() {
			ExecuteJobAndPushToKafka(config, job)
		})
		jobIds = append(jobIds, jobId)
	}
}

func ScheduleJobs(cron *cron.Cron, config config.Config) {
	//read jobs from jobs table
	jobs,_  := jobs.ReadJobs(config)
	if len(jobs) > 0 {
		AddJobsToCron(cron, config, jobs)
	}
	return 
}

func main() {
	
	env := "local"
	port := "5000"
	pg := config.InitPG(env)
	defer pg.Close()

	config := config.Config{
		PG:     pg,
		Port:   port,
		Env:    env,
		Kafka:  config.InitKafkaProducer(env),
	}

	cron := cron.New()
	//cron job to fetch and schedule jobs from db every 5 mins
	cron.AddFunc("0 */5 * * * *", func() { 
			ScheduleJobs(cron, config) 
		})
	
	cron.Start()
	
	defer cron.Stop()
	
	log.Fatal(http.ListenAndServe(":"+"5000", nil))

}