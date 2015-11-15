package cast

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

const (
	RoundingAuto = "RoundAuto"
	RoundingUp   = "RoundUp"
	RoundingDown = "RoundDown"
)

func Value(o interface{}) reflect.Value {
	return reflect.ValueOf(o)
}

func Kind(o interface{}) reflect.Kind {
	return Value(o).Kind()
}

func ToString(o interface{}) string {
	v := Value(o)
	k := v.Kind()
	if k == reflect.Interface && v.IsNil() {
		return ""
	} else if k == reflect.String {
		return o.(string)
	} else if k == reflect.Int || k == reflect.Int8 ||
		k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 {
		return fmt.Sprintf("%d", o)
	} else if k == reflect.Float32 || k == reflect.Float64 {
		return fmt.Sprintf("%f", o)
	} else {
		return ""
	}
}

func Date2String(t time.Time, dateFormat string) string {
	return ""
}

func ToInt(o interface{}, rounding string) int {
	var ret int
	k := Kind(o)
	v := Value(o)

	//fmt.Printf("Data %v Kind %v \n", o, k)
	if k == reflect.String {
		if i, e := strconv.Atoi(v.String()); e == nil {
			return i
		} else {
			return 0
		}
	} else if k == reflect.Int || k == reflect.Int8 ||
		k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 {
		return int(v.Int())
	} else if k == reflect.Float32 || k == reflect.Float64 {
		return int(int64(v.Float()))
	}

	return ret
}

func ToF32(o interface{}, decimalPoint int, rounding string) float32 {
	return 0
}

func ToF64(o interface{}, decimalPoint int, rounding string) float64 {
	return 0
}

/*
DateFormat legend:
	d = date
	dd = date 2 digit
	M = month
	MM = month 2 digit
	MMM = month in name, 3 chars
	MMMM = month in name, full
	YY = Year 2 digit
	YYYY = Year 4 digit
	h = hour
	hh = hour 2 digit
	H = hour in 24 cycle
	HH = hour in 24 cycle 2 digit
	m = minute
	mm = minute 2 digits
	s = Second
	ss = second 2 digit
	A = AMPM
	T = Timezone
	L = Location
*/
func ToDate(o interface{}) time.Time {
	return time.Now()
}

func ToDuration(o interface{}) time.Duration {
	return (1 * time.Second)
}
