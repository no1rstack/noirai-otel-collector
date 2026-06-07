package schemamigrator

import "time"

type MigrationSchemaMigrationRecord struct {
	MigrationID uint64    `ch:"migration_id"`
	Status      string    `ch:"status"`
	Error       string    `ch:"error"`
	CreatedAt   time.Time `ch:"created_at"`
	UpdatedAt   time.Time `ch:"updated_at"`
}

var V2MigrationTablesLogs = []SchemaMigrationRecord{
	{
		MigrationID: 1,
		UpItems: []Operation{
			CreateTableOperation{
				Database: "noirai_logs",
				Table:    "schema_migrations_v2",
				Columns: []Column{
					{Name: "migration_id", Type: ColumnTypeUInt64},
					{Name: "status", Type: ColumnTypeString},
					{Name: "error", Type: ColumnTypeString},
					{Name: "created_at", Type: DateTime64ColumnType{Precision: 9}},
					{Name: "updated_at", Type: DateTime64ColumnType{Precision: 9}},
				},
				Engine: ReplacingMergeTree{
					MergeTree{
						OrderBy:    "migration_id",
						PrimaryKey: "migration_id",
					},
				},
			},
		},
	},
	{
		MigrationID: 2,
		UpItems: []Operation{
			CreateTableOperation{
				Database: "noirai_logs",
				Table:    "distributed_schema_migrations_v2",
				Columns: []Column{
					{Name: "migration_id", Type: ColumnTypeUInt64},
					{Name: "status", Type: ColumnTypeString},
					{Name: "error", Type: ColumnTypeString},
					{Name: "created_at", Type: DateTime64ColumnType{Precision: 9}},
					{Name: "updated_at", Type: DateTime64ColumnType{Precision: 9}},
				},
				Engine: Distributed{
					Database:    "noirai_logs",
					Table:       "schema_migrations_v2",
					ShardingKey: "rand()",
				},
			},
		},
	},
}

var V2MigrationTablesTraces = []SchemaMigrationRecord{
	{
		MigrationID: 3,
		UpItems: []Operation{
			CreateTableOperation{
				Database: "noirai_traces",
				Table:    "schema_migrations_v2",
				Columns: []Column{
					{Name: "migration_id", Type: ColumnTypeUInt64},
					{Name: "status", Type: ColumnTypeString},
					{Name: "error", Type: ColumnTypeString},
					{Name: "created_at", Type: DateTime64ColumnType{Precision: 9}},
					{Name: "updated_at", Type: DateTime64ColumnType{Precision: 9}},
				},
				Engine: ReplacingMergeTree{
					MergeTree{
						OrderBy:    "migration_id",
						PrimaryKey: "migration_id",
					},
				},
			},
		},
	},
	{
		MigrationID: 4,
		UpItems: []Operation{
			CreateTableOperation{
				Database: "noirai_traces",
				Table:    "distributed_schema_migrations_v2",
				Columns: []Column{
					{Name: "migration_id", Type: ColumnTypeUInt64},
					{Name: "status", Type: ColumnTypeString},
					{Name: "error", Type: ColumnTypeString},
					{Name: "created_at", Type: DateTime64ColumnType{Precision: 9}},
					{Name: "updated_at", Type: DateTime64ColumnType{Precision: 9}},
				},
				Engine: Distributed{
					Database:    "noirai_traces",
					Table:       "schema_migrations_v2",
					ShardingKey: "rand()",
				},
			},
		},
	},
}

var V2MigrationTablesMetrics = []SchemaMigrationRecord{
	{
		MigrationID: 5,
		UpItems: []Operation{
			CreateTableOperation{
				Database: "noirai_metrics",
				Table:    "schema_migrations_v2",
				Columns: []Column{
					{Name: "migration_id", Type: ColumnTypeUInt64},
					{Name: "status", Type: ColumnTypeString},
					{Name: "error", Type: ColumnTypeString},
					{Name: "created_at", Type: DateTime64ColumnType{Precision: 9}},
					{Name: "updated_at", Type: DateTime64ColumnType{Precision: 9}},
				},
				Engine: ReplacingMergeTree{
					MergeTree{
						OrderBy:    "migration_id",
						PrimaryKey: "migration_id",
					},
				},
			},
		},
	},
	{
		MigrationID: 6,
		UpItems: []Operation{
			CreateTableOperation{
				Database: "noirai_metrics",
				Table:    "distributed_schema_migrations_v2",
				Columns: []Column{
					{Name: "migration_id", Type: ColumnTypeUInt64},
					{Name: "status", Type: ColumnTypeString},
					{Name: "error", Type: ColumnTypeString},
					{Name: "created_at", Type: DateTime64ColumnType{Precision: 9}},
					{Name: "updated_at", Type: DateTime64ColumnType{Precision: 9}},
				},
				Engine: Distributed{
					Database:    "noirai_metrics",
					Table:       "schema_migrations_v2",
					ShardingKey: "rand()",
				},
			},
		},
	},
}

var V2MigrationTablesMetadata = []SchemaMigrationRecord{
	{
		MigrationID: 7,
		UpItems: []Operation{
			CreateTableOperation{
				Database: "noirai_metadata",
				Table:    "schema_migrations_v2",
				Columns: []Column{
					{Name: "migration_id", Type: ColumnTypeUInt64},
					{Name: "status", Type: ColumnTypeString},
					{Name: "error", Type: ColumnTypeString},
					{Name: "created_at", Type: DateTime64ColumnType{Precision: 9}},
					{Name: "updated_at", Type: DateTime64ColumnType{Precision: 9}},
				},
				Engine: ReplacingMergeTree{
					MergeTree{
						OrderBy:    "migration_id",
						PrimaryKey: "migration_id",
					},
				},
			},
		},
	},
	{
		MigrationID: 8,
		UpItems: []Operation{
			CreateTableOperation{
				Database: "noirai_metadata",
				Table:    "distributed_schema_migrations_v2",
				Columns: []Column{
					{Name: "migration_id", Type: ColumnTypeUInt64},
					{Name: "status", Type: ColumnTypeString},
					{Name: "error", Type: ColumnTypeString},
					{Name: "created_at", Type: DateTime64ColumnType{Precision: 9}},
					{Name: "updated_at", Type: DateTime64ColumnType{Precision: 9}},
				},
				Engine: Distributed{
					Database:    "noirai_metadata",
					Table:       "schema_migrations_v2",
					ShardingKey: "rand()",
				},
			},
		},
	},
}

var V2MigrationTablesAnalytics = []SchemaMigrationRecord{
	{
		MigrationID: 9,
		UpItems: []Operation{
			CreateTableOperation{
				Database: "noirai_analytics",
				Table:    "schema_migrations_v2",
				Columns: []Column{
					{Name: "migration_id", Type: ColumnTypeUInt64},
					{Name: "status", Type: ColumnTypeString},
					{Name: "error", Type: ColumnTypeString},
					{Name: "created_at", Type: DateTime64ColumnType{Precision: 9}},
					{Name: "updated_at", Type: DateTime64ColumnType{Precision: 9}},
				},
				Engine: ReplacingMergeTree{
					MergeTree{
						OrderBy:    "migration_id",
						PrimaryKey: "migration_id",
					},
				},
			},
			CreateTableOperation{
				Database: "noirai_analytics",
				Table:    "distributed_schema_migrations_v2",
				Columns: []Column{
					{Name: "migration_id", Type: ColumnTypeUInt64},
					{Name: "status", Type: ColumnTypeString},
					{Name: "error", Type: ColumnTypeString},
					{Name: "created_at", Type: DateTime64ColumnType{Precision: 9}},
					{Name: "updated_at", Type: DateTime64ColumnType{Precision: 9}},
				},
				Engine: Distributed{
					Database:    "noirai_analytics",
					Table:       "schema_migrations_v2",
					ShardingKey: "rand()",
				},
			},
		},
	},
}

var V2MigrationTablesMeter = []SchemaMigrationRecord{
	{
		MigrationID: 10,
		UpItems: []Operation{
			CreateTableOperation{
				Database: "noirai_meter",
				Table:    "schema_migrations_v2",
				Columns: []Column{
					{Name: "migration_id", Type: ColumnTypeUInt64},
					{Name: "status", Type: ColumnTypeString},
					{Name: "error", Type: ColumnTypeString},
					{Name: "created_at", Type: DateTime64ColumnType{Precision: 9}},
					{Name: "updated_at", Type: DateTime64ColumnType{Precision: 9}},
				},
				Engine: ReplacingMergeTree{
					MergeTree{
						OrderBy:    "migration_id",
						PrimaryKey: "migration_id",
					},
				},
			},
			CreateTableOperation{
				Database: "noirai_meter",
				Table:    "distributed_schema_migrations_v2",
				Columns: []Column{
					{Name: "migration_id", Type: ColumnTypeUInt64},
					{Name: "status", Type: ColumnTypeString},
					{Name: "error", Type: ColumnTypeString},
					{Name: "created_at", Type: DateTime64ColumnType{Precision: 9}},
					{Name: "updated_at", Type: DateTime64ColumnType{Precision: 9}},
				},
				Engine: Distributed{
					Database:    "noirai_meter",
					Table:       "schema_migrations_v2",
					ShardingKey: "rand()",
				},
			},
		},
	},
}
