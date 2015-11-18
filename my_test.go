package cast

import (
	"fmt"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	var f float64 = 200.30

	s := ToString(f)
	if s == "" {
		t.Errorf("Unable to cast %f to string \n", f)
		return
	} else {
		fmt.Printf("%f is casted to string = %s \n", f, s)
	}
}

func TestNo(t *testing.T) {
	s := "209.20"

	i := ToInt(s, RoundingUp)
	if i != 210 {
		t.Errorf("Format %s to Int fail, got %d", s, i)
	}

	i = ToInt(s, RoundingAuto)
	if i != 209 {
		t.Errorf("Format %s to Int fail, got %d", s, i)
	}

	s = "209.15"
	f := ToF64(s, 1, RoundingDown)
	if f != 209.10 {
		t.Errorf("Format %s to float fail, got %f", s, f)
	}

	i = ToInt(f, RoundingAuto)
	if i != 209 {
		t.Errorf("Format %f to int fail, got %d", f, i)
	}
}

func TestDateString(t *testing.T) {
	d := time.Date(2015, 2, 8, 0, 0, 0, 0, time.Now().Location())

	s := Date2String(d, "dd-MMM-YY")
	if s != "08-Feb-15" {
		t.Errorf("Format date %v to string fail. got %s", d, s)
	}
	s = Date2String(d, "dd-MMM-yyyy")
	if s != "08-Feb-2015" {
		t.Errorf("Format date %v to string fail. got %s", d, s)
	}
	s = Date2String(d, "d-M-yyyy")
	if s != "8-2-2015" {
		t.Errorf("Format date %v to string fail. got %s", d, s)
	}

	d, _ = time.Parse(time.UnixDate, "Sat Mar  7 09:06:39 PST 2015")
	s = Date2String(d, "dd-MMM-YY H:m:s")
	if s != "07-Mar-15 9:6:39" {
		t.Errorf("Format date %v to string fail. got %s", d, s)
	}

	d, _ = time.Parse(time.UnixDate, "Sat Mar  7 19:06:39 PST 2015")
	s = Date2String(d, "dd-MMM-YY H:m:s")
	if s != "07-Mar-15 19:6:39" {
		t.Errorf("Format date %v to string fail. got %s", d, s)
	}

}
