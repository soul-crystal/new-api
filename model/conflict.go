package model

import (
	"github.com/QuantumNous/new-api/common"
	"gorm.io/gorm/clause"
)

// ConflictClauses returns the ON CONFLICT clauses to attach to a Create call.
//
// openGauss speaks the PostgreSQL protocol but does not support the
// PostgreSQL ON CONFLICT syntax, so there the clause is dropped and the
// insert runs as a plain INSERT; callers must tolerate duplicate-key errors
// in that case. SQLite, MySQL and PostgreSQL keep the given clause.
//
// Use this for both the shared DB/LOG_DB handles and gorm.DB values passed
// in from other packages (such as service/authz).
func ConflictClauses(onConflict clause.OnConflict) []clause.Expression {
	if common.UsingOpenGauss() {
		return nil
	}
	return []clause.Expression{onConflict}
}
