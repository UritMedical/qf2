package define

import (
	"fmt"
	"github.com/UritMedical/qf2/util/qdate"
	"strconv"
	"strings"
	"time"
)

var (
	dateFormat     string // 日期掩码
	dateTimeFormat string // 日期时间掩码
)

type Date uint32

//
// FromTime
//  @Description: 通过原生的time赋值
//  @param time
//
//goland:noinspection GoMixedReceiverTypes
func (d *Date) FromTime(time time.Time) {
	t := time.Local()
	s := fmt.Sprintf("%04d%02d%02d", t.Year(), t.Month(), t.Day())
	v, _ := strconv.ParseUint(s, 10, 32)
	*d = Date(v)
}

//
// ToString
//  @Description: 根据全局format格式化输出
//  @return string
//
//goland:noinspection GoMixedReceiverTypes
func (d Date) ToString() string {
	return qdate.ToString(d.ToTime(), dateFormat)
}

//
// ToTime
//  @Description: 转为原生时间对象
//  @return time.Time
//
//goland:noinspection GoMixedReceiverTypes
func (d Date) ToTime() time.Time {
	if d == 0 {
		return time.Time{}
	}
	str := fmt.Sprintf("%d", d)
	if len(str) != 8 {
		str = str + strings.Repeat("0", 8-len(str))
	}
	year, _ := strconv.Atoi(str[0:4])
	month, _ := strconv.Atoi(str[4:6])
	day, _ := strconv.Atoi(str[6:8])
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

//
// MarshalJSON
//  @Description: 复写json转换
//  @return []byte
//  @return error
//
//goland:noinspection GoMixedReceiverTypes
func (d Date) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("\"%s\"", d.ToString())
	return []byte(str), nil
}

//
// UnmarshalJSON
//  @Description: 复写json转换
//  @param data
//  @return error
//
//goland:noinspection GoMixedReceiverTypes
func (d *Date) UnmarshalJSON(data []byte) error {
	v, err := qdate.ToNumber(string(data), dateFormat)
	if err == nil {
		*d = Date(v)
	}
	return err
}

type DateTime uint64

//
// NowDateTime
//  @Description: 当前系统时间
//  @return DateTime
//
func NowDateTime() DateTime {
	var now DateTime
	now.FromTime(time.Now().Local())
	return now
}

//
// FromTime
//  @Description: 通过原生的time赋值
//  @param time
//
//goland:noinspection GoMixedReceiverTypes
func (d *DateTime) FromTime(time time.Time) {
	t := time.Local()
	s := fmt.Sprintf("%04d%02d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	v, _ := strconv.ParseUint(s, 10, 64)
	*d = DateTime(v)
}

//
// Date
//  @Description: 转为日期
//  @return Date
//
//goland:noinspection GoMixedReceiverTypes
func (d DateTime) Date() Date {
	var date Date
	date.FromTime(d.ToTime())
	return date
}

//
// ToString
//  @Description: 根据全局format格式化输出
//  @return string
//
//goland:noinspection GoMixedReceiverTypes
func (d DateTime) ToString() string {
	return qdate.ToString(d.ToTime(), dateTimeFormat)
}

//
// ToTime
//  @Description: 转为原生时间对象
//  @return time.Time
//
//goland:noinspection GoMixedReceiverTypes
func (d DateTime) ToTime() time.Time {
	if d == 0 {
		return time.Time{}
	}
	str := fmt.Sprintf("%d", d)
	if len(str) != 14 {
		str = str + strings.Repeat("0", 14-len(str))
	}
	year, _ := strconv.Atoi(str[0:4])
	month, _ := strconv.Atoi(str[4:6])
	day, _ := strconv.Atoi(str[6:8])
	hour, _ := strconv.Atoi(str[8:10])
	minute, _ := strconv.Atoi(str[10:12])
	second, _ := strconv.Atoi(str[12:14])
	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
}

//
// MarshalJSON
//  @Description: 复写json转换
//  @return []byte
//  @return error
//
//goland:noinspection GoMixedReceiverTypes
func (d DateTime) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("\"%s\"", d.ToString())
	return []byte(str), nil
}

//
// UnmarshalJSON
//  @Description: 复写json转换
//  @param data
//  @return error
//
//goland:noinspection GoMixedReceiverTypes
func (d *DateTime) UnmarshalJSON(data []byte) error {
	v, err := qdate.ToNumber(string(data), dateTimeFormat)
	if err == nil {
		*d = DateTime(v)
	}
	return err
}
