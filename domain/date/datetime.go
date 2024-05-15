package date

import (
	"database/sql/driver"
	"time"
)

type JSONDateTime time.Time

// Marshal returns the JSON encoding of t.
func (t JSONDateTime) MarshalJSON() ([]byte, error) {
	thaiLocation, _ := time.LoadLocation("Asia/Bangkok")
	formatted := time.Time(t).In(thaiLocation).Format("2006-01-02T04:05")
	return []byte(`"` + formatted + `"`), nil
}

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by t.
func (t *JSONDateTime) UnmarshalJSON(b []byte) error {
	thaiLocation, _ := time.LoadLocation("Asia/Bangkok")
	s := string(b)
	s = s[1 : len(s)-1]
	tt, err := time.ParseInLocation("2006-01-02T04:05", s, thaiLocation)
	if err != nil {
		return err
	}
	*t = JSONDateTime(tt)

	return nil
}

func (t JSONDateTime) Value() (driver.Value, error) {
	thaiLocation, _ := time.LoadLocation("Asia/Bangkok")

	return time.Time(t).In(thaiLocation).Format("2006-01-02T15:04:05Z07:00"), nil
}

func (t JSONDateTime) Time() time.Time {
	thaiLocation, _ := time.LoadLocation("Asia/Bangkok")

	return time.Time(t).In(thaiLocation)
}

func (t *JSONDateTime) Set(tt time.Time) {
	*t = JSONDateTime(tt)
}
