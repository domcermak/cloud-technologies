package client

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
)

func getString(options map[string]interface{}, name string) (string, error) {
	arg, ok := options[name]
	if !ok {
		return "", errors.Errorf("required argument `%s` not given", name)
	}

	return arg.(string), nil
}

func getInt(options map[string]interface{}, name string) (int, error) {
	arg, err := getString(options, name)
	if err != nil {
		return 0, err
	}

	integer, err := strconv.Atoi(arg)
	if err != nil {
		return 0, errors.Errorf("`%s` is not a number", name)
	}

	return integer, nil
}

func prettyJson(input interface{}) (string, error) {
	data, err := json.MarshalIndent(input, "", "\t")

	return string(data), err
}
