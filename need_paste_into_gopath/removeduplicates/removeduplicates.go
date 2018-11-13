package removeduplicates

func RemoveDuplicate(ss []string) []string {
	checked := make(map[string]struct{})
	var result []string
	for _, s := range ss {
		if _, ok := checked[s]; !ok {
			result = append(result, s)
			checked[s] = struct{}{}
		}
	}

	return result
}
