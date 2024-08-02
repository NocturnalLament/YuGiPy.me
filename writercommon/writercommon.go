package writercommon

import "database/sql"

type GetSqlStmtInterface interface {
	GetSqlStmt() (*sql.Stmt, error)
}
