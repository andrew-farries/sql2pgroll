package sql2pgroll

import (
	pgq "github.com/pganalyze/pg_query_go/v5"
	"github.com/xataio/pgroll/pkg/migrations"
)

const PlaceHolderSQL = "TODO: Implement SQL data migration"

func convertAlterTableStmt(stmt *pgq.AlterTableStmt) (migrations.Operations, error) {
	if stmt.Objtype != pgq.ObjectType_OBJECT_TABLE {
		return nil, nil
	}

	var ops migrations.Operations
	for _, cmd := range stmt.Cmds {
		alterTableCmd := cmd.GetAlterTableCmd()
		if alterTableCmd == nil {
			continue
		}

		switch alterTableCmd.Subtype {
		case pgq.AlterTableType_AT_SetNotNull:
			ops = append(ops, convertAlterTableSetNotNull(stmt, alterTableCmd))
		}
	}

	return ops, nil
}

func convertAlterTableSetNotNull(stmt *pgq.AlterTableStmt, cmd *pgq.AlterTableCmd) migrations.Operation {
	return &migrations.OpAlterColumn{
		Table:    stmt.GetRelation().GetRelname(),
		Column:   cmd.GetName(),
		Nullable: ptr(false),
		Up:       PlaceHolderSQL,
		Down:     PlaceHolderSQL,
	}
}
