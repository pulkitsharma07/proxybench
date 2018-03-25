package benchmark

func StressHTTPS() *Stress {
	return NewStress("HTTPS Stress", true)
}
