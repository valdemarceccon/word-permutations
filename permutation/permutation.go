package permutation

func Calculate(word string) []string {
	return per(word, "")
}

func per(a,b string) []string {
	var result []string
	if len(a) == 0 {
		return append(result, b)
	}

	for i, l := range a {
		aux := a[0:i] + a[i+1:]

		result = append(result, per(aux, b + string(l))...)
	}

	return result
}

