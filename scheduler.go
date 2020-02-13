package main

import (
	"fmt"
	"log"
	"goyash/golang-kafka-job-scheduler/config"
	//"github.com/jasonlvhit/gocron"
	"os"
	"gopkg.in/robfig/cron.v2"
	"time"
	//"sync"
	"net/http"
	"goyash/golang-kafka-job-scheduler/models/jobs"
)

const ONE_SECOND = 1*time.Second + 10*time.Millisecond

func task() {
	fmt.Println("I am running task.")
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}





var jobIds = []cron.EntryID{0}

func RemovePreviousJobs(c *cron.Cron, jobIds []cron.EntryID) {
	log.Println("removing previously added jobs...")
	for _,id := range jobIds {
		c.Remove(id)
	}
}

func AddJobsToCron (c *cron.Cron, config config.Config, jobs []jobs.Jobs) {
	RemovePreviousJobs(c, jobIds)
	jobIds = []cron.EntryID{0}
	for _,job := range jobs {
		jobId,_ := c.AddFunc(job.Schedule, RunJobTask)
		jobIds = append(jobIds, jobId)
	}
}

func RunJobTask () {
	log.Println("Hello")
}

func ScheduleJobs(cron *cron.Cron, config config.Config) {
	
	jobs,_  := jobs.ReadJobs(config)
	
	if len(jobs) > 0 {
		AddJobsToCron(cron, config, jobs)
	}
	
	return 
}

func main() {
	
	env := "local"
	pg := config.InitPG(env)
	defer pg.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	config := config.Config{
		PG:     pg,
		Port:   port,
		Env:    env,
		//Kafka:  bfConfig.InitKafkaProducer(env),
	}

	cron := cron.New()
	cron.AddFunc("*/30 * * * * *", func() { 
			ScheduleJobs(cron, config) 
		})
	
	cron.Start()
	
	defer cron.Stop()
	
	log.Fatal(http.ListenAndServe(":"+"5000", nil))

}