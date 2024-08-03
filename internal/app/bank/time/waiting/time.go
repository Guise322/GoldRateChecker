package waiting

import (
	"PriceWatcher/internal/app/bank/time/call"
	"PriceWatcher/internal/app/bank/time/randomizer"
	"time"
)

func GetWaitTimeWithRandomComp(now time.Time, callHours []int) time.Duration {
	variation := 1800
	randDur := randomizer.RandomSec(variation)
	callTime := call.GetCallTime(now, callHours)

	return getWaitDurWithProcessingTime(now, callTime, randDur)
}

func getWaitDurWithProcessingTime(now time.Time, callTime time.Time, randDur time.Duration) time.Duration {
	waitDur := callTime.Sub(now)

	if waitDur < 0 {
		var zeroDur time.Duration
		return zeroDur
	}
	processingTime := 3 * time.Minute
	randComp := randDur + processingTime

	if waitDur < randComp {
		return waitDur
	}

	return waitDur - randComp
}