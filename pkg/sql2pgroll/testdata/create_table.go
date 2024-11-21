package testdata

import "github.com/xataio/pgroll/pkg/migrations"

var CreateTableOp1 = &migrations.OpCreateTable{
	Name: "foo",
	Columns: []migrations.Column{
		{
			Name:     "a",
			Type:     "int4",
			Nullable: ptr(true),
		},
	},
}

var CreateTableOp2 = &migrations.OpCreateTable{
	Name: "foo",
	Columns: []migrations.Column{
		{
			Name: "a",
			Type: "int4",
		},
	},
}

var CreateTableOp3 = &migrations.OpCreateTable{
	Name: "foo",
	Columns: []migrations.Column{
		{
			Name:     "a",
			Type:     "varchar(255)",
			Nullable: ptr(true),
		},
	},
}

var CreateTableOp4 = &migrations.OpCreateTable{
	Name: "foo",
	Columns: []migrations.Column{
		{
			Name:     "a",
			Type:     "numeric(10,2)",
			Nullable: ptr(true),
		},
	},
}

var CreateTableOp5 = &migrations.OpCreateTable{
	Name: "foo",
	Columns: []migrations.Column{
		{
			Name:     "a",
			Type:     "int4",
			Nullable: ptr(true),
			Unique:   ptr(true),
		},
	},
}

var CreateTableOp6 = &migrations.OpCreateTable{
	Name: "foo",
	Columns: []migrations.Column{
		{
			Name: "a",
			Type: "int4",
			Pk:   ptr(true),
		},
	},
}

var CreateTableOp7 = &migrations.OpCreateTable{
	Name: "foo",
	Columns: []migrations.Column{
		{
			Name:     "a",
			Type:     "text[]",
			Nullable: ptr(true),
		},
	},
}

var CreateTableOp8 = &migrations.OpCreateTable{
	Name: "foo",
	Columns: []migrations.Column{
		{
			Name:     "a",
			Type:     "text[5]",
			Nullable: ptr(true),
		},
	},
}

var CreateTableOp9 = &migrations.OpCreateTable{
	Name: "foo",
	Columns: []migrations.Column{
		{
			Name:     "a",
			Type:     "text[5][3]",
			Nullable: ptr(true),
		},
	},
}

func ptr[T any](x T) *T {
	return &x
}
