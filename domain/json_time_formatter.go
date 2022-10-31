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
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006/01/02 15:04:05 Z07:00"))
	return []byte(stamp), nil
}

func (t TimeFormatter) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("15:04:05 WIB"))
	return []byte(stamp), nil
}
func (t DateFormatter) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006/01/02"))
	return []byte(stamp), nil
}
