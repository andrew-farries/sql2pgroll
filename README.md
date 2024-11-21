# sql2pgroll

Experiments around converting SQL to `pgroll` migrations.

### Usage

Use the `sql2pgroll` command to convert SQL to `pgroll` migrations.

```bash
$ go run ./cmd/sql2pgroll "CREATE TABLE foo (id SERIAL PRIMARY KEY, name TEXT UNIQUE);"
```

Output:

```json
{
  "name": "inferred_from_sql",
  "operations": [
    {
      "create_table": {
        "columns": [
          {
            "name": "id",
            "pk": true,
            "type": "serial"
          },
          {
            "name": "name",
            "nullable": true,
            "type": "text",
            "unique": true
          }
        ],
        "name": "foo"
      }
    }
  ]
}
```

Take a look at the tests in the `pkg/sql2pgroll` package to see which SQL statements are supported.
