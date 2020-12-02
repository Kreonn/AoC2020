package tools

import "strconv"

// Str2Int converts a list of strings to a list of integers
func Str2Int(base []string) ([]int, error) {
	res := make([]int, 0, len(base))
	for _, str := range base {
		i, err := strconv.Atoi(str)
		if err != nil {
			return res, err
		}
		res = append(res, i)
	}
	return res, nil
}
