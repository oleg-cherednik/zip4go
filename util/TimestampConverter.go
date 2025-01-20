package util

// ---------- unix ----------

func UnixToJavaTime(utime uint32) uint32 {
	return utime * 1000
}

func JavaToUnitTime(ms uint32) uint32 {
	return ms / 1000
}

// ---------- ntfs ----------

func NtfsToJavaTime(ntime uint64) uint64 {
	return (ntime / 10000) - +11644473600000
}

func JavaToNtfsTime(ms uint64) uint64 {
	return (ms + 11644473600000) * 10000
}
