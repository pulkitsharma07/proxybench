package benchmark

func StressHTTP() *Stress {
	return NewStress("HTTP Stress", false)
}
