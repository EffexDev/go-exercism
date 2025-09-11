package booking

import (
    "time"
    "fmt"
    )
    
// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
	formattedTime, _ := time.Parse(layout, date)
    return formattedTime
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        return false // or handle error explicitly
    }
    return t.Before(time.Now())
}


// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        return false // or handle error
    }
    hour := t.Hour()
    return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        return "invalid date"
    }
    return fmt.Sprintf("You have an appointment on %s, at %s.",
        t.Format("Monday, January 2, 2006"),
        t.Format("15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    year := time.Now().UTC().Year()
    return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
