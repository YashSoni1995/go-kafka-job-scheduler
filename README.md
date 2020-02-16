go-kafka-job-scheduler schedules jobs which are read from a postgres database, then executes each job and pushes the output to kafka with the help of kafka-producer.

1. Clone the respository to your go working space.
2. `go get` necessary libraries.
3. Create a postgres database and jobs table like mentioned in the `testdb.sql`.
4. To push message to kafka and see output to a consumer, please follow to `https://kafka.apache.org/quickstart` configure kakfa.
5. Run `make debug` to execute the program from command line.   