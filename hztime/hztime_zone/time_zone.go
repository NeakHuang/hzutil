// Date: 2023/6/5
// Author:
// Description：

package hztime_zone

import (
	"os"
	"time"
)

// SetTimezone
// Deprecated: 只要time被调用过后，SetEnv修改TZ的方式就不再生效了，因此使用场景很少
// 推荐使用 SetTimezoneLocal 和 SetTimezoneOffset 代替该方法
func SetTimezone(zone Timezone) error {
	location, err := time.LoadLocation(zone.String())
	if err != nil {
		return err
	}
	err = os.Setenv("TZ", location.String())
	return err
}

// SetTimezoneLocal
// Asia/Shanghai
// UTC
// and so on
func SetTimezoneLocal(zone Timezone) error {
	loc, err := time.LoadLocation(zone.String())
	if err != nil {
		return err
	}
	time.Local = loc
	return nil
}

// SetTimezoneOffset
// Asia/Shanghai
// UTC
// and so on
func SetTimezoneOffset(zone ...Timezone) {
	// 设置程序的默认时区为 "Asia/Shanghai" 东8区
	offsetHour := int(Timezone8_AsiaShanghai)
	if len(zone) > 0 {
		offsetHour = int(zone[0])
	}
	time.Local = time.FixedZone("CST", offsetHour*60*60)
}
