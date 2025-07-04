//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var WebTokens = newWebTokensTable("public", "web__tokens", "")

type webTokensTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	UserID    postgres.ColumnInteger
	Token     postgres.ColumnString
	CreatedAt postgres.ColumnTimestamp
	ExpiresAt postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
	DefaultColumns postgres.ColumnList
}

type WebTokensTable struct {
	webTokensTable

	EXCLUDED webTokensTable
}

// AS creates new WebTokensTable with assigned alias
func (a WebTokensTable) AS(alias string) *WebTokensTable {
	return newWebTokensTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new WebTokensTable with assigned schema name
func (a WebTokensTable) FromSchema(schemaName string) *WebTokensTable {
	return newWebTokensTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new WebTokensTable with assigned table prefix
func (a WebTokensTable) WithPrefix(prefix string) *WebTokensTable {
	return newWebTokensTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new WebTokensTable with assigned table suffix
func (a WebTokensTable) WithSuffix(suffix string) *WebTokensTable {
	return newWebTokensTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newWebTokensTable(schemaName, tableName, alias string) *WebTokensTable {
	return &WebTokensTable{
		webTokensTable: newWebTokensTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newWebTokensTableImpl("", "excluded", ""),
	}
}

func newWebTokensTableImpl(schemaName, tableName, alias string) webTokensTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		UserIDColumn    = postgres.IntegerColumn("user_id")
		TokenColumn     = postgres.StringColumn("token")
		CreatedAtColumn = postgres.TimestampColumn("created_at")
		ExpiresAtColumn = postgres.TimestampColumn("expires_at")
		allColumns      = postgres.ColumnList{IDColumn, UserIDColumn, TokenColumn, CreatedAtColumn, ExpiresAtColumn}
		mutableColumns  = postgres.ColumnList{UserIDColumn, TokenColumn, CreatedAtColumn, ExpiresAtColumn}
		defaultColumns  = postgres.ColumnList{IDColumn, CreatedAtColumn}
	)

	return webTokensTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		UserID:    UserIDColumn,
		Token:     TokenColumn,
		CreatedAt: CreatedAtColumn,
		ExpiresAt: ExpiresAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
		DefaultColumns: defaultColumns,
	}
}
