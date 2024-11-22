package sql2pgroll_test

import (
	"testing"

	"github.com/andrew-farries/pgroll2sql/pkg/sql2pgroll"
	"github.com/andrew-farries/pgroll2sql/pkg/sql2pgroll/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xataio/pgroll/pkg/migrations"
)

func TestConvertAlterTableStatements(t *testing.T) {
	t.Parallel()

	tests := []struct {
		sql        string
		expectedOp migrations.Operation
	}{
		{
			sql:        "ALTER TABLE foo ALTER COLUMN a SET NOT NULL",
			expectedOp: testdata.AlterTableOp1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.sql, func(t *testing.T) {
			ops, err := sql2pgroll.Convert(tc.sql)
			require.NoError(t, err)

			require.Len(t, ops, 1)

			alterColumnOps, ok := ops[0].(*migrations.OpAlterColumn)
			require.True(t, ok)

			assert.Equal(t, tc.expectedOp, alterColumnOps)
		})
	}
}
