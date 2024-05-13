package utils

import (
	"fmt"
	"time"
)

func DurationToFormatString(duration time.Duration) string {
	return fmt.Sprintf("%02d:%02d", int(duration.Hours()), int(duration.Minutes())%60)
}
