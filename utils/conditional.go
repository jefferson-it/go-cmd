package utils

func OptStr(cond bool, v1 string, v2 string) string {
	if cond {
		return v1
	}

	return v2
}
