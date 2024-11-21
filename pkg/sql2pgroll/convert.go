package sql2pgroll

import (
	"fmt"

	pgq "github.com/pganalyze/pg_query_go/v5"
	"github.com/xataio/pgroll/pkg/migrations"
)

var ErrMultipleStatements = fmt.Errorf("expected only one statement")

func Convert(sql string) (migrations.Operations, error) {
	op, err := convert(sql)
	if err != nil {
		return nil, fmt.Errorf("unable to convert SQL statement: %w", err)
	}

	if op == nil {
		return makeRawSQLOperation(sql), nil
	}

	return op, nil
}

func convert(sql string) (migrations.Operations, error) {
	tree, err := pgq.Parse(sql)
	if err != nil {
		return nil, err
	}

	stmts := tree.GetStmts()
	if len(stmts) > 1 {
		return nil, fmt.Errorf("%w: got %d statements", ErrMultipleStatements, len(stmts))
	}
	node := stmts[0].GetStmt().GetNode()

	switch node := (node).(type) {
	case *pgq.Node_CreateStmt:
		return convertCreateStmt(node.CreateStmt)
	case *pgq.Node_AlterTableStmt:
		return convertAlterTableStmt(node.AlterTableStmt)
	default:
		return makeRawSQLOperation(sql), nil
	}
}

func makeRawSQLOperation(sql string) migrations.Operations {
	return migrations.Operations{
		&migrations.OpRawSQL{Up: sql},
	}
}

func ptr[T any](x T) *T {
	return &x
}
