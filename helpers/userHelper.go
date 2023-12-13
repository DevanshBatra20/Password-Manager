package helpers

import "strconv"

func ValidateId(id string) (int, error) {
	res, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return res, nil
}
