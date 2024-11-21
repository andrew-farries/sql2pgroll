package sql2pgroll_test

import (
	"testing"

	"github.com/andrew-farries/pgroll2sql/pkg/sql2pgroll"
	"github.com/andrew-farries/pgroll2sql/pkg/sql2pgroll/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xataio/pgroll/pkg/migrations"
)

func TestConvertCreateTableStatements(t *testing.T) {
	t.Parallel()

	tests := []struct {
		sql        string
		expectedOp migrations.Operation
	}{
		{
			sql:        "CREATE TABLE foo(a int)",
			expectedOp: testdata.CreateTableOp1,
		},
		{
			sql:        "CREATE TABLE foo(a int NOT NULL)",
			expectedOp: testdata.CreateTableOp2,
		},
		{
			sql:        "CREATE TABLE foo(a varchar(255))",
			expectedOp: testdata.CreateTableOp3,
		},
		{
			sql:        "CREATE TABLE foo(a numeric(10, 2))",
			expectedOp: testdata.CreateTableOp4,
		},
		{
			sql:        "CREATE TABLE foo(a int UNIQUE)",
			expectedOp: testdata.CreateTableOp5,
		},
		{
			sql:        "CREATE TABLE foo(a int PRIMARY KEY)",
			expectedOp: testdata.CreateTableOp6,
		},
		{
			sql:        "CREATE TABLE foo(a text[])",
			expectedOp: testdata.CreateTableOp7,
		},
		{
			sql:        "CREATE TABLE foo(a text[5])",
			expectedOp: testdata.CreateTableOp8,
		},
		{
			sql:        "CREATE TABLE foo(a text[5][3])",
			expectedOp: testdata.CreateTableOp9,
		},
	}

	for _, tc := range tests {
		t.Run(tc.sql, func(t *testing.T) {
			ops, err := sql2pgroll.Convert(tc.sql)
			require.NoError(t, err)

			require.Len(t, ops, 1)

			createTableOp, ok := ops[0].(*migrations.OpCreateTable)
			require.True(t, ok)

			assert.Equal(t, tc.expectedOp, createTableOp)
		})
	}
}
