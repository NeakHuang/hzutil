// Date: 2023/6/30
// Author:
// Description：

package hzmodel

import (
	"database/sql"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var jsonstd = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	FatalJsonMarshalFailed   = "Json data marshal to string failed"
	FatalSqlStringScanFailed = "Sql.NullString Scan to value failed"
)

func JsonToSqlString(jsonData any, sqlString *sql.NullString) error {
	if str, err := jsonstd.MarshalToString(jsonData); err != nil {
		sqlString.String = FatalJsonMarshalFailed
		sqlString.Valid = true
		return errors.Errorf("%v %v", FatalJsonMarshalFailed, err)
	} else {
		if err = sqlString.Scan(str); err != nil {
			sqlString.String = FatalSqlStringScanFailed
			sqlString.Valid = true
			return errors.Errorf("%v %v", FatalSqlStringScanFailed, err)
		}
	}
	return nil
}
