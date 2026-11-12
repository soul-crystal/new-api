package model

import (
	"github.com/QuantumNous/new-api/common"
	"gorm.io/gorm"
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

// saveOption writes an Option row with upsert semantics on every database.
// GORM's Save falls back to INSERT ... ON CONFLICT when its UPDATE matches
// no rows, which openGauss does not support, so there the write runs as an
// explicit update-then-insert keyed on the option key.
func saveOption(tx *gorm.DB, option *Option) error {
	if common.UsingOpenGauss() {
		result := tx.Model(&Option{}).
			Where(commonKeyCol+" = ?", option.Key).
			Updates(map[string]any{"value": option.Value})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected > 0 {
			return nil
		}
		return tx.Create(option).Error
	}
	return tx.Save(option).Error
}
