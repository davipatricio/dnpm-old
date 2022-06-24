package utils

func Contains(s []string, e any) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}