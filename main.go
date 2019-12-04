package main

import  (
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"context"
	"reflect"
	"fmt"
	"os"
	"time"

	Data "github.com/mike-webster/go-etl/data"
	Models "github.com/mike-webster/go-etl/models"
)

func main() {
	for i := 0; i < 13; i++ {
		fmt.Println("...sleeping for db setup... {", i, "} ")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("App starting")
	sdb := os.Getenv("SOURCE_DB_URL")
	fmt.Println("SOURCE_DB_URL:", sdb)
	ddb := os.Getenv("DESTINATION_DB_URL")
    fmt.Println("DESTINATION_DB_URL:", ddb)
	ctx := context.Background()

	conn := Data.Connection {
		SourceDBDriverName: "mysql",
		SourceDBConnectionString: sdb,
		DestinationDBDriverName: "mysql",
		DestinationDBConnectionString: ddb,
	}
	errs := conn.Initialize()
	if errs != nil {
		for _, e := range *errs {
			fmt.Println(e)
		}
		panic("fatal error initializing")
	}

	err := conn.SourceDB.Ping()
	if err != nil {
		fmt.Println("Fatal error pinging source db")
		panic(err)
	}

	err = conn.DestinationDB.Ping()
	if err != nil {
		fmt.Println("Fatal error pinging destination db")
		panic(err)
	}

	queries := []Models.Queryable {
		&Models.Example1{},
		&Models.Example2{},
	}

	ch := make(chan Models.Queryable)
	go func(){
		for _, q := range queries {
			ch <- q
		}
		close(ch)
	}()

	ch2 := selecting(ctx, &conn, ch)
	ch3 := persist(ctx, &conn, ch2)
	report := retrieveErrors(ctx, ch3)

	//report := retrieveErrors(ctx, persist(ctx, conn, select(ctx, ch)))
	for _, r := range report {
		fmt.Println("-- Error: \n\t", r)
	}
}

func selecting(ctx context.Context, conn *Data.Connection, queries <-chan Models.Queryable) <-chan Models.Queryable {
	output := make(chan Models.Queryable)
	
	go func(){
		for q := range queries {
			res, err := conn.SourceSelect(q.SourceQuery())
			if err != nil {
				q.Error(err)
				continue
			}
	
			for res.Next() {
				switch reflect.TypeOf(q) {
				case reflect.TypeOf(&Models.Example1{}):
					var e Models.Example1
					err := res.StructScan(&e)
					if err != nil {
						q.Error(err)
						continue
					}
					
					fmt.Println("---- Adding: \n---- ", e)
					output <- &e
				case reflect.TypeOf(&Models.Example2{}):
					var e Models.Example2
					err := res.StructScan(&e)
					if err != nil {
						q.Error(err)
						continue
					}
	
					fmt.Println("---- Adding: \n---- ", e)
					output <- &e
				default:
					panic(fmt.Sprint("query attempting to be selected that is not handled: ", reflect.TypeOf(q)))
				}
			}
		}
		close(output)
	}()
	return output
}

func persist(ctx context.Context, conn *Data.Connection, queries <-chan Models.Queryable) <-chan Models.Queryable {
	output := make(chan Models.Queryable)
	go func(){
		for q := range queries {
			for _, sql := range q.DestinationSQL(ctx) {
				fmt.Println("---- Query\n\t", sql)
				_, err := conn.DestinationInsert(sql)
				if err != nil {
					q.Error(err)
				}
			}

			output <- q
		}
		close(output)
	}()
	return output
}

func retrieveErrors(ctx context.Context, queries <-chan Models.Queryable) []string {
	ret := []string{}
	for q := range queries {
		ret = append(ret, q.Errors()...)
	}
	return ret
}