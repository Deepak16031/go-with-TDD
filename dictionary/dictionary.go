package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", errors.New("this is an unknown key")
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
