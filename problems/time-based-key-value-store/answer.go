package main

import "fmt"

// Ref: https://leetcode.com/problems/time-based-key-value-store/solutions/1690481/

type entry struct {
	value string
	time  int
}

type TimeMap struct {
	entries map[string][]entry
}

func Constructor() TimeMap {
	return TimeMap{
		entries: map[string][]entry{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.entries[key] = append(this.entries[key], entry{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	timeIndex := this.findCloestTimeIndex(key, timestamp)
	if timeIndex < 0 {
		return ""
	} else {
		return this.entries[key][timeIndex].value
	}
}

func (this *TimeMap) findCloestTimeIndex(key string, timestamp int) int {
	l := 0
	r := len(this.entries[key]) - 1
	for l <= r {
		m := (l + r) / 2
		if this.entries[key][m].time < timestamp {
			l = m + 1
		} else if this.entries[key][m].time > timestamp {
			r = m - 1
		} else {
			return m
		}
	}
	return r
}

func main() {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)
	fmt.Println(timeMap.Get("foo", 1) == "bar")
	fmt.Println(timeMap.Get("foo", 3) == "bar")
	timeMap.Set("foo", "bar2", 4)
	fmt.Println(timeMap.Get("foo", 4) == "bar2")
	fmt.Println(timeMap.Get("foo", 5) == "bar2")
}
