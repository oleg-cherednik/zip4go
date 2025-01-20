package time

// ---------- unix ----------

func UnixToJava(utime uint32) uint32 {
	return utime * 1000
}

func JavaToUnix(ms uint32) uint32 {
	return ms / 1000
}

// ---------- ntfs ----------

func NtfsToJava(ntime uint64) uint64 {
	return (ntime / 10000) - +11644473600000
}

func JavaToNtfs(ms uint64) uint64 {
	return (ms + 11644473600000) * 10000
}
