package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	scheduleForm := "1/2/2006 15:04:05"
	dateNew, _ := time.Parse(scheduleForm, date)
	return dateNew
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	AppointmentForm := "January 2, 2006 15:04:05"
	dateAppointment, _ := time.Parse(AppointmentForm, date)
	dateNow := time.Now()
	return dateNow.After(dateAppointment)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	AppointmentForm := "Monday, January 2, 2006 15:04:05"
	dateAppointment, _ := time.Parse(AppointmentForm, date)
	dateHour := dateAppointment.Hour()
	if dateHour >= 12 && dateHour < 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	dateNew := Schedule(date)
	year, month, day := dateNew.Date()
	hour, min, _ := dateNew.Clock()
	weekday := dateNew.Weekday()
	desc := fmt.Sprintf("You have an appointment on %s, %s %v, %v, at %v:%v.", weekday, month, day, year, hour, min)
	return desc

}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
