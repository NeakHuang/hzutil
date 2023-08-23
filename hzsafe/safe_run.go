// Date: 2023/2/11
// Author:
// Description：

package hzsafe

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"runtime/debug"
)

type RunFunc func()

var runFuncList []RunFunc // 协程管理list
func RegRunFunc(f RunFunc) {
	runFuncList = append(runFuncList, f)
}

// SafeRunFuncList 防崩溃调用
func SafeRunFuncList() {
	for _, run := range runFuncList {
		go SafeRun(run)()
	}
}

func SafeRun(runFunc RunFunc) RunFunc {
	return func() {
		defer CheckFatal("协程崩溃：%+v", runFunc)
		runFunc()
	}
}

// 检查并记录异常日志
func CheckFatal(content string, args ...interface{}) {
	if err, ok := recover().(error); ok {
		logx.ErrorStackf("%v(%v)\n%v",
			fmt.Sprintf(content, args...), err, string(debug.Stack()))
	}
}
