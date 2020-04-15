package dates

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// GetNow return now() at UTC Timezone
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString return now() at string at UTC Timezone
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
