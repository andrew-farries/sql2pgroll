package main

import (
	"fmt"
	"log"
	"os"

	pg_query "github.com/pganalyze/pg_query_go/v5"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <sql string>\n", os.Args[0])
		os.Exit(1)
	}

	tree, err := pg_query.ParseToJSON(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tree)
}
