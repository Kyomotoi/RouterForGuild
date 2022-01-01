package global

import "time"

type Timestamp string

func (t Timestamp) ToTimeStamp() (time.Time, error) {
	return time.Parse(time.RFC3339, string(t))
}
