package stringslice

func Contains(elem string, slice []string) bool {
	for _, s := range slice {
		if s == elem {
			return true
		}
	}

	return false
}
