package benchmark

func StressHTTPS(name string) Stress {
	return NewStress(name, true)
}
