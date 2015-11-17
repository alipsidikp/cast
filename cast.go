package cast

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
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

/*
DateFormat legend:
2		d = date
02		dd = date 2 digit
1		M = month
01		MM = month 2 digit
Jan		MMM = month in name, 3 chars
January	MMMM = month in name, full
06		YY = Year 2 digit
2006	YYYY = Year 4 digit
3	h = hour
03	hh = hour 2 digit
	H = hour in 24 cycle
15	HH = hour in 24 cycle 2 digit
4	m = minute
04	mm = minute 2 digits
5	s = Second
05	ss = second 2 digit
PM	A = AMPM
MST	T = Timezone
	L = Location
*/

func Date2String(t time.Time, dateFormat string) string {

	var dateMap = map[string]string{
		"dd":   "02",
		"d":    "2",
		"MMMM": "January",
		"MMM":  "Jan",
		"MM":   "01",
		"M":    "1",
		"YYYY": "2006",
		"YY":   "06",
		"hh":   "03",
		"h":    "3",
		"HH":   "15",
		"mm":   "04",
		"m":    "4",
		"ss":   "05",
		"s":    "5",
		"A":    "PM",
		"T":    "MST",
	}
	// "H":    "",
	// "L":  "",

	var dateOrder = map[int]string{1: "dd", 2: "d", 3: "MMMM", 4: "MMM", 5: "MM", 6: "M", 7: "YYYY", 8: "YY",
		9: "hh", 10: "h", 11: "HH", 12: "mm", 13: "m", 14: "ss", 15: "s", 16: "A", 17: "T",
	}

	var keys []int
	for k := range dateOrder {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	dateFormat = strings.Replace(dateFormat, "y", "Y", -1)
	for _, k := range keys {
		dateFormat = strings.Replace(dateFormat, dateOrder[k], dateMap[dateOrder[k]], -1)
	}

	return t.Format(dateFormat)
}

func ToInt(o interface{}, rounding string) int {
	var ret int
	k := Kind(o)
	v := Value(o)

	if k == reflect.String {
		i := strings.Index(v.String(), ".")
		if i >= 0 {
			f, _ := strconv.ParseFloat(v.String(), 64)
			ret = ToInt(f, rounding)
		} else {
			if i, e := strconv.Atoi(v.String()); e == nil {
				return i
			} else {
				return 0
			}
		}
	} else if k == reflect.Int || k == reflect.Int8 ||
		k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 {
		return int(v.Int())
	} else if k == reflect.Float32 || k == reflect.Float64 {
		f := ToF64(v.Float(), 0, rounding)
		return int(int64(f))
	}

	return ret
}

func ToF32(o interface{}, decimalPoint int, rounding string) float32 {
	var f float64

	k := Kind(o)
	v := Value(o)

	if k == reflect.String {
		f = ToF64(v.String(), 0, rounding)
	} else if k == reflect.Int || k == reflect.Int8 ||
		k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 {
		f = ToF64(v.Int(), 0, rounding)
	} else if k == reflect.Float32 || k == reflect.Float64 {
		f = ToF64(v.Float(), 0, rounding)
	}

	return float32(f)
}

func ToF64(o interface{}, decimalPoint int, rounding string) float64 {
	var f float64
	var e error

	k := Kind(o)
	v := Value(o)

	if k == reflect.String {
		f, e = strconv.ParseFloat(v.String(), 64)
		if e != nil {
			return 0
		}
	} else if k == reflect.Int || k == reflect.Int8 ||
		k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64 {
		f = float64(v.Int())
	} else if k == reflect.Float32 || k == reflect.Float64 {
		f = float64(v.Float())
	}

	switch rounding {
	case RoundingAuto:
		return RoundingAuto64(f, decimalPoint)
	case RoundingDown:
		return RoundingDown64(f, decimalPoint)
	case RoundingUp:
		return RoundingUp64(f, decimalPoint)
	}

	return f
}

func RoundingAuto64(f float64, decimalPoint int) (retValue float64) {

	tempPow := math.Pow(10, float64(decimalPoint))
	f = f * tempPow

	if f < 0 {
		f = math.Ceil(f - 0.5)
	} else {
		f = math.Floor(f + 0.5)
	}

	retValue = f / tempPow
	return
}

func RoundingDown64(f float64, decimalPoint int) (retValue float64) {
	tempPow := math.Pow(10, float64(decimalPoint))
	f = f * tempPow
	f = math.Floor(f)
	retValue = f / tempPow
	return
}

func RoundingUp64(f float64, decimalPoint int) (retValue float64) {
	tempPow := math.Pow(10, float64(decimalPoint))
	f = f * tempPow
	f = math.Ceil(f)
	retValue = f / tempPow
	return
}

func ToDate(o interface{}) time.Time {
	return time.Now()
}

func ToDuration(o interface{}) time.Duration {
	return (1 * time.Second)
}
