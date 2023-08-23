// Date: 2023/5/2
// Author:
// Description：

package hztime

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Println("NowMonday", NowMonday0Clock())
	fmt.Println("——————————————————————")
	tm := Today0Clock()
	fmt.Println(tm)
	fmt.Println("week", tm.Weekday())
	fmt.Println("Day0", TsDay0Clock(tm.Unix()))
	fmt.Println("Day0", time.Unix(int64(TsCycle0Clock(0, tm.Unix())), 0))
	fmt.Println("Monday0", TsMonday0Clock(tm.Unix()))
	fmt.Println("Monday0", time.Unix(int64(TsCycle0Clock(1, tm.Unix())), 0))
	fmt.Println("Month1st0", TsMonth1st0Clock(tm.Unix()))
	fmt.Println("Month1st0", time.Unix(int64(TsCycle0Clock(2, tm.Unix())), 0))
	fmt.Println("——————————————————————")
	tm1 := time.Unix(1672588800, 0)
	fmt.Println(tm1)
	fmt.Println("week", tm1.Weekday())
	fmt.Println("Day0", TsDay0Clock(tm1.Unix()))
	fmt.Println("UponDay0", time.Unix(int64(TsUponCycle0Clock(0, tm1.Unix())), 0))
	fmt.Println("Monday0", TsMonday0Clock(tm1.Unix()))
	fmt.Println("UponMonday0", time.Unix(int64(TsUponCycle0Clock(1, tm1.Unix())), 0))
	fmt.Println("Month1st0", TsMonth1st0Clock(tm1.Unix()))
	fmt.Println("UponMonth1st0", time.Unix(int64(TsUponCycle0Clock(2, tm1.Unix())), 0))
	// https://api.debank.com/history/list?user_addr=0x553ead49101d764b8a7234ae45c358be53f5aa31&chain=&start_time=1672228416&page_count=20
}
