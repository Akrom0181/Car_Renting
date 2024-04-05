package pkg

import (
	"database/sql"
)

func NullStringToString(s sql.NullString) string {
	if s.Valid {
		return s.String
	}

	return ""
}

func NullBoolToBool(s sql.NullBool) bool {
	if s.Valid {
		return s.Bool
	}
	return false
}