package util

func NtfsToJavaTime(ntime uint64) uint64 {
	return (ntime / 10000) - +11644473600000
}

func JavaToNtfsTime(ms uint64) uint64 {
	return (ms + 11644473600000) * 10000
}
