# go-etl
A generic ETL template, written in golang, utilizing pipelines. 

## What is this?
This was meant to be a proof of concept for building an ETL using golang pipelines

## Why does this exist?
I was curious to see how much I could improve an existing ETL by using pipelines.  The answer was "a lot".

## Does it work?
Yup.

## How do I run it?
1. Clone the repo locally `git clone https://github.com/mike-webster/go-etl.git`
2. Move into the directory `cd go-etl`
3. Kick off the process `docker-compose up --build`

## How do I know it worked?
After the logs quiet down, the process is done.  Then you can check the destination database.
`docker exec -it example_app sh`

Once you're in the container, you can log into mysql.
`mysql -u example -p pass`

Run the queries to check the data.
`select * from example.example1`
`select * from example.example2`

## How does it work?
The docker-compose file will spin up two containers, one for the source database and another for the destination. We are also mounting the appropriate directory to provide a seed sql script that will get copied into the docker-entrypoint-initdb.d directory (this will cause the script to be run on database creation).  The source script will also populate some information into the tables.

The docker-compose also stands up the etl app.  It provides connection strings to both DBs and the app runs which selects the data, transforms it, and persists it in the destination DB.

There are two models defined, that both satisfy the Queryable interface. There are methods implemented that will define how we select, transform, and persist the data from each table.  Since the models will all satisfy the Queryable interface, the infrastructure is set up to take Queryable interfaces instead of specific models.