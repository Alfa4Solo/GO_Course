package Pr3

func Num4(words []string) string {
	mi := len(words[0])
	w := words[0]
	counter := 0
	for _, val := range words {
		for i := 0; i < len(w); i++ {
			if w[i] == val[i] {
				counter++
			} else {
				break
			}
		}
		if mi > counter {
			mi = counter
		}
		counter = 0
	}
	return w[:mi]
}
