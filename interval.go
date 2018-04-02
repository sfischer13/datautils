package datautils

import "math"
import "regexp"
import "strconv"
import "strings"

var reRanges = regexp.MustCompile(`` +
	`^([0-9]*|[0-9]*:[0-9]*|[0-9]*:[0-9]*:[0-9]*)` +
	`(,([0-9]*|[0-9]*:[0-9]*|[0-9]*:[0-9]*:[0-9]*))*$`)
var reRangeS = regexp.MustCompile(`^(\d*)$`)
var reRangeM = regexp.MustCompile(`^(\d*):(\d*)$`)
var reRangeL = regexp.MustCompile(`^(\d*):(\d*):(\d*)$`)

// Interval TODO
type Interval struct {
	Low  int64
	High int64
	Step int64
}

// NewInterval TODO
func NewInterval(low, high, step string) Interval {
	if low == "" {
		low = "1"
	}
	if high == "" {
		high = strconv.FormatInt(math.MaxInt64, 10)
	}
	if step == "" {
		step = "1"
	}

	l, errL := strconv.ParseInt(low, 10, 64)
	h, errH := strconv.ParseInt(high, 10, 64)
	s, errS := strconv.ParseInt(step, 10, 64)

	if (errL != nil) || (errH != nil) || (errS != nil) {
		panic("Parameters could not be parsed!")
	}

	return Interval{
		Low:  l,
		High: h,
		Step: s,
	}
}

// Contains TODO
func (r *Interval) Contains(i int64) bool {
	return (i >= r.Low) && (i <= r.High) && (((i - r.Low) % r.Step) == 0)
}

// String2Intervals TODO
func String2Intervals(s string) []Interval {
	if reRanges.MatchString(s) {
		result := []Interval{}
		for _, r := range strings.Split(s, ",") {
			low := ""
			high := ""
			step := ""
			if m := reRangeS.FindStringSubmatch(r); len(m) != 0 {
				low = m[1]
				high = m[1]
			} else if m := reRangeM.FindStringSubmatch(r); len(m) != 0 {
				low = m[1]
				high = m[2]
			} else if m := reRangeL.FindStringSubmatch(r); len(m) != 0 {
				low = m[1]
				high = m[2]
				step = m[3]
			} else {
				return nil
			}
			result = append(result, NewInterval(low, high, step))
		}
		return result
	}
	return nil
}

// Intervals2Func TODO
func Intervals2Func(rs []Interval) func(int64) bool {
	return func(i int64) bool {
		result := false
		for _, r := range rs {
			result = result || r.Contains(i)
		}
		return result
	}
}
