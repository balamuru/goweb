package time

import "time"

func NowTime() string{
	return time.Now().String()
}