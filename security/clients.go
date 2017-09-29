/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package security

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var clientRegistryStore *bolt.DB

func AddClient(name string) (string, error) {
	token, err := generateRandomString(32)
	if err != nil {
		return "", err
	}
	err = put([]byte(token), []byte(name))
	if err != nil {
		return "", err
	}
	return token, nil
}

func IsAuthorizedClient(token string) (bool, error) {
	if !enabled {
		return true, nil
	}

	name, err := get([]byte(token))
	if err != nil {
		return false, err
	}
	namestr := string(name)
	if namestr == "" {
		return false, fmt.Errorf("Not a valid token")
	} else {
		appendLog(namestr, token)
		return true, nil
	}
}

func RemoveClient(token string) error {
	return nil
}
