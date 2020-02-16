CREATE DATABASE testdb;
CREATE USER testuser PASSWORD 'testuserpassword';
GRANT ALL PRIVILEGES ON DATABASE testdb TO testuser;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE jobs (
	id uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
  	task text not null,
  	schedule text not null,
  	topic text not null,
	uts timestamp default current_timestamp
);

INSERT INTO jobs (task, topic, schedule) VALUES ('Say Hello to Kafka!', 'test', '*/10 * * * * *');

 
