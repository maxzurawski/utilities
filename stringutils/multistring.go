package stringutils

import (
	"encoding/json"
	"strings"
)

type MultiString string

func (c MultiString) ToString() string {
	return string(c)
}

func ToMultiString(value string) MultiString {
	return MultiString(value)
}

func (c *MultiString) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == `true` {
		*c = "true"
		return nil
	}

	if str == `false` {
		*c = "false"
		return nil
	}

	res := str

	if strings.Contains(res, "\"") {
		res = res[1 : len(res)-1]
	}

	*c = MultiString(res)
	return nil
}

func (c MultiString) MarshalJSON() (out []byte, err error) {
	value := c.ToString()
	err = nil

	if value == "true" || value == "false" {
		out = []byte(strings.Replace(value, "\"", "`", -1))
		return
	}
	return json.Marshal(value)
}
