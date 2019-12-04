package main

import  (
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"context"
	"reflect"
	"fmt"
    "os"
	"www.github.com/mike-webster/go-etl/data"
	"www.github.com/mike-webster/go-etl/models"
)

func main() {
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
	err := conn.Initialize()
	if err != nil {
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
	
					output <- q
				case reflect.TypeOf(&Models.Example2{}):
					var e Models.Example2
					err := res.StructScan(&e)
					if err != nil {
						q.Error(err)
						continue
					}
	
					output <- q
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