package date

import (
	"database/sql/driver"
	"time"
)

type JSONDate time.Time

// Marshal returns the JSON encoding of t.
func (t JSONDate) MarshalJSON() ([]byte, error) {
	thaiLocation, _ := time.LoadLocation("Asia/Bangkok")
	formatted := time.Time(t).In(thaiLocation).Format("2006-01-02")
	return []byte(`"` + formatted + `"`), nil
}

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by t.
func (t *JSONDate) UnmarshalJSON(b []byte) error {
	thaiLocation, _ := time.LoadLocation("Asia/Bangkok")
	s := string(b)
	s = s[1 : len(s)-1]
	tt, err := time.ParseInLocation("2006-01-02", s, thaiLocation)
	if err != nil {
		return err
	}
	*t = JSONDate(tt)

	return nil
}

func (t JSONDate) Value() (driver.Value, error) {
	thaiLocation, _ := time.LoadLocation("Asia/Bangkok")

	return time.Time(t).In(thaiLocation).Format("2006-01-02T15:04:05Z07:00"), nil
}

func (t JSONDate) Time() time.Time {
	thaiLocation, _ := time.LoadLocation("Asia/Bangkok")

	return time.Time(t).In(thaiLocation)
}

func (t *JSONDate) Set(tt time.Time) {
	*t = JSONDate(tt)
}
