package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"
	res, _ := time.Parse(layout, date)
	return res
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	const layout = "January 2, 2006 15:04:05"
	res, _ := time.Parse(layout, date)

	return res.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January 2, 2006 15:04:05"
	res, _ := time.Parse(layout, date)
	hour := res.Hour()

	return (hour >= 12 && hour < 18)
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	const layout = "1/2/2006 15:04:05"
	res, _ := time.Parse(layout, date)

	return fmt.Sprintf("You have an appointment on %s", res.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	t, _ := time.Parse("2006-01-2", fmt.Sprintf("%d-09-15", time.Now().Year()))
	return t
}
