package plans

import "time"

// Prorate returns the unused part of an annual premium, in pence, from cancelledOn to the end of
// the policy year. Days are counted exactly; a leap year has 366 days.
//
// rating-service has its own copy of this calculation (ProRata.kt) for mid-term changes. The two
// round differently: this one truncates, rating-service rounds half up.
func Prorate(annualPence int64, start, cancelledOn time.Time) int64 {
	end := start.AddDate(1, 0, 0)
	if !cancelledOn.Before(end) {
		return 0
	}
	total := end.Sub(start).Hours() / 24
	unused := end.Sub(cancelledOn).Hours() / 24
	return int64(float64(annualPence) * unused / total)
}
