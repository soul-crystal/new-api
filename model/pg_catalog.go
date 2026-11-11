package model

import "github.com/QuantumNous/new-api/common"

// openGauss is based on PostgreSQL 9.2 and lacks to_regclass() and
// cardinality(), both introduced in PostgreSQL 9.4. The helpers below render
// equivalent SQL for the pg_catalog queries used by the migrations, so those
// queries keep a single parameter layout on every database.

// pgTableOidMatch renders "<column> = <table oid>", where the table name is
// the next ? parameter and is resolved in the current schema.
func pgTableOidMatch(column string) string {
	if common.UsingOpenGauss() {
		return column + ` = (SELECT rel_class.oid
          FROM pg_catalog.pg_class AS rel_class
          JOIN pg_catalog.pg_namespace AS rel_namespace
            ON rel_namespace.oid = rel_class.relnamespace
          WHERE rel_class.relname = ? AND rel_namespace.nspname = current_schema()
          LIMIT 1)`
	}
	return column + " = to_regclass(?)"
}

// pgSingleColumnConkey renders a check that a pg_constraint.conkey array holds
// exactly one column.
func pgSingleColumnConkey(column string) string {
	if common.UsingOpenGauss() {
		return "array_length(" + column + ", 1) = 1"
	}
	return "cardinality(" + column + ") = 1"
}
