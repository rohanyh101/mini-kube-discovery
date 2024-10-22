package main

import "net/http"

func helthCheck(addr string) (bool, error) {
	res, err := http.Get(addr)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}
