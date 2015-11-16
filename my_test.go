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

	s = "209.15"
	f := ToF64(s, 1, RoundingAuto)
	if f != 209.20 {
		t.Errorf("Format %s to float fail, got %f", s, f)
	}

	i = ToInt(f, RoundingAuto)
	if f != 209 {
		t.Errorf("Format %f to int fail, got %d", f, i)
	}
}

func TestDateString(t *testing.T) {
	d := time.Date(2015, 2, 8, 0, 0, 0, 0, time.Now().Location())
	s := Date2String(d, "dd-MMM-yy")
	if s != "08-Feb-15" {
		t.Errorf("Format date %v to string fail. got %s", d, s)
	}
}
