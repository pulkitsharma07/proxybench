package benchmark

func StressHTTP(name string) Stress {
	return NewStress(name, false)
}
