package utils

func Assign(target, src map[string]string) {
	for key, value := range target {
		src[key] = value
	}
}