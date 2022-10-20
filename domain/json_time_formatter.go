package domain

import (
	"fmt"
	"time"
)

type (
	LastUpdatedTimeFormatter time.Time
	TimeFormatter            time.Time
	DateFormatter            time.Time
)

func (t LastUpdatedTimeFormatter) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("20060102150405"))
	return []byte(stamp), nil
}

func (t TimeFormatter) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("150405"))
	return []byte(stamp), nil
}
func (t DateFormatter) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("20060102"))
	return []byte(stamp), nil
}
