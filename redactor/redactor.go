package redactor

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

const Mask = "***********"

func merge(path string, k string) string {
	if len(path) == 0 {
		return k
	}
	if strings.HasSuffix(path, ".") {
		return path + k
	} else {
		return fmt.Sprintf("%s.%s", path, k)
	}
}

func hide(in *interface{}, keys map[string]bool, path string) {
	hold := (*in)
	if m, ok := hold.(map[string]interface{}); ok {
		for k, v := range m {
			p := merge(path, k)
			if keys[p] {
				m[k] = Mask
				continue
			}
			hide(&v, keys, merge(path, k))
		}
	} else if m, ok := hold.([]interface{}); ok {
		for _, v := range m {
			hide(&v, keys, path)
		}
	}
}

func Redact(keys map[string]bool, r io.Reader, w io.Writer) error {
	dec := json.NewDecoder(r)

	var m interface{}
	enc := json.NewEncoder(w)

	err := dec.Decode(&m)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	hide(&m, keys, "")

	enc.Encode(m)
	return nil
}
