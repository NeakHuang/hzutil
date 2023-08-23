// Date: 2023/4/10
// Author:
// Description：

package hzstr

import (
	"fmt"
	"testing"
)

func TestSliceToString(t *testing.T) {
	whereList := []string{
		"a = 1",
		"b = 2",
		"c = 3",
	}
	sql := StrSliceToString(whereList, ",")
	fmt.Println(sql)
}
