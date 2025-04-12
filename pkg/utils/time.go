package utils

import "time"

var (
	_ TimeNower = (*TimeNow)(nil)
	_ TimeNower = (*TimeNowTest)(nil)
)

type TimeNower interface {
	Now() time.Time
}

type TimeNow struct{}

func NewTimeNow() TimeNow {
	return TimeNow{}
}

func (t TimeNow) Now() time.Time {
	return time.Now()
}

type TimeNowTest struct {
	t time.Time
}

func NewTimeNowTest(t time.Time) TimeNowTest {
	return TimeNowTest{t: t}
}

func (t TimeNowTest) Now() time.Time {
	return t.t
}
