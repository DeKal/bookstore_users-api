package dates

import "time"

const (
	apiDateLayout   = "2006-01-02T15:04:05Z"
	apiDateDBLayout = "2006-01-02 15:04:05"
)

// GetNow return now() at UTC Timezone
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString return now() at string at UTC Timezone
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDBString return now() at string with DB format at UTC Timezone
func GetNowDBString() string {
	return GetNow().Format(apiDateDBLayout)
}
