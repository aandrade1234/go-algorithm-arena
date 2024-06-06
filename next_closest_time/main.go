package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(nextClosestTime1("19:34"))
	fmt.Println(nextClosestTime1("23:59"))

	fmt.Println(nextClosestTime2("19:34"))
	fmt.Println(nextClosestTime2("23:59"))
}

func nextClosestTime2(time string) string {
	digits := make([]int, 0, 4)
	for _, c := range time {
		if c != ':' {
			digits = append(digits, int(c-'0'))
		}
	}

	validTimes := generateValidTimes(digits)

	currentTime := toMinutes(time)
	nextTime := currentTime
	for _, t := range validTimes {
		if t > currentTime {
			nextTime = t
			break
		}
	}

	if nextTime == currentTime {
		nextTime = validTimes[0]
	}

	return fromMinutes(nextTime)
}

func generateValidTimes(digits []int) []int {
	var validTimes []int
	for _, h1 := range digits {
		for _, h2 := range digits {
			if h1*10+h2 < 24 {
				for _, m1 := range digits {
					for _, m2 := range digits {
						if m1*10+m2 < 60 {
							validTimes = append(validTimes, (h1*10+h2)*60+m1*10+m2)
						}
					}
				}
			}
		}
	}
	sort.Ints(validTimes)
	return validTimes
}

func toMinutes(time string) int {
	parts := strings.Split(time, ":")
	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])
	return hours*60 + minutes
}

func fromMinutes(minutes int) string {
	hours := minutes / 60
	minutes = minutes % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func nextClosestTime1(time string) string {
	digits := make(map[rune]bool)
	for _, r := range time {
		if r >= '0' && r <= '9' {
			digits[r] = true
		}
	}

	hours, _ := strconv.Atoi(time[:2])
	minutes, _ := strconv.Atoi(time[3:])
	totalMinutes := (hours * 60) + minutes

	isValidTime := func(h, m int) bool {
		timeString := fmt.Sprintf("%02d%02d", h, m)
		for _, char := range timeString {
			if !digits[char] {
				return false
			}
		}

		return true
	}

	for {
		totalMinutes = (totalMinutes + 1) % (24 * 60)
		hours = totalMinutes / 60
		minutes = totalMinutes % 60

		if isValidTime(hours, minutes) {
			return fmt.Sprintf("%02d:%02d", hours, minutes)
		}
	}
}
