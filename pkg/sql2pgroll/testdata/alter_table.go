package testdata

import (
	"github.com/andrew-farries/pgroll2sql/pkg/sql2pgroll"
	"github.com/xataio/pgroll/pkg/migrations"
)

var AlterTableOp1 = &migrations.OpAlterColumn{
	Table:    "foo",
	Column:   "a",
	Nullable: ptr(false),
	Up:       sql2pgroll.PlaceHolderSQL,
	Down:     sql2pgroll.PlaceHolderSQL,
}
