package gigasecond

import "time"

//Just takes whatever time you pass in and adds a a gigasecond
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second*1000000000)
	return t
}
