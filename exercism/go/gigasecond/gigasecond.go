/*
Package gigasecond adds 1 giga second to the given input and return the new date
*/
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond accepts time as input and adds 1 gigasecond to it and returns it as output
func AddGigasecond(t time.Time) time.Time {
	gigaVal := 1000_000_000
	return t.Add(time.Second * time.Duration(gigaVal))
}
