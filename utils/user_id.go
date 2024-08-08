package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetUserID(r *http.Request) (id int, err error) {
	id, err = strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return -1, fmt.Errorf("malformed user_id passed: %v", err)
	}
	return id, nil
}
