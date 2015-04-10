package gigasecond

import "time"

const TestVersion = 2

var Birthday = time.Date(1989, time.August, 28, 11, 58, 0, 0, time.Local)
var gigasecond time.Duration = 1e9 * time.Second

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
