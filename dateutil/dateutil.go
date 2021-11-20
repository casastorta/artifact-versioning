package dateutil

import (
	"fmt"
	"time"
)

func GetCurrentYearWeek(loc *time.Location) string {
	t := time.Now()
	return GetYearWeek(loc, &t)
}

func GetYearWeek(loc *time.Location, t *time.Time) string {
	tl := t.In(loc)
	year, week := tl.ISOWeek()
	return fmt.Sprintf("%d.%d", year, week)
}
