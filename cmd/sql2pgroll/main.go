package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/andrew-farries/pgroll2sql/pkg/sql2pgroll"
	"github.com/xataio/pgroll/pkg/migrations"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <sql string>\n", os.Args[0])
		os.Exit(1)
	}

	ops, err := sql2pgroll.Convert(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	mig := migrations.Migration{
		Name:       "inferred_from_sql",
		Operations: ops,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(mig); err != nil {
		log.Fatal(err)
	}
}
