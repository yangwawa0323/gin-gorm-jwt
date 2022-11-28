package utils

import "log"

func ErrorDebug(err error) error {
	if err != nil {
		log.Printf("[DEBUG]: %s", err.Error())
		return err
	}
	return nil
}
