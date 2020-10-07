package redactor

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Redact_Array(t *testing.T) {
	keys := map[string]bool{
		"licenses.number": true,
	}
	const jsonStream = `{"licenses":[{"name":"Savarkar","number":"123456347"},{"name":" Hegdewar","number":"135862286"}]}`

	input := strings.NewReader(jsonStream)
	var outString string
	buffer := bytes.NewBufferString(outString)
	err := Redact(keys, input, buffer)
	assert.NoError(t, err)

	decoder := json.NewDecoder(buffer)
	type License struct {
		Name   string `json:"name"`
		Number string `json:"number"`
	}
	type Doc struct {
		Licenses []License `json:"licenses"`
	}
	var m Doc
	err = decoder.Decode(&m)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(m.Licenses))
	for _, l := range m.Licenses {
		assert.Equal(t, Mask, l.Number)
	}
}

func Test_Redact_Complex_Object(t *testing.T) {
	keys := map[string]bool{
		"customer.license.number": true,
	}
	const jsonStream = `{
		"customer": {
			"license": {
				"number": "12936782947237"
			}
		}
	}
	`

	input := strings.NewReader(jsonStream)
	var outString string
	buffer := bytes.NewBufferString(outString)
	err := Redact(keys, input, buffer)
	assert.NoError(t, err)

	decoder := json.NewDecoder(buffer)
	type License struct {
		Name   string `json:"name"`
		Number string `json:"number"`
	}
	type Customer struct {
		License License `json:"license"`
	}
	type Doc struct {
		Customer Customer `json:"customer"`
	}
	var m Doc
	err = decoder.Decode(&m)
	assert.NoError(t, err)
	assert.Equal(t, Mask, m.Customer.License.Number)
}

func Test_Redact_Simple_Object(t *testing.T) {
	keys := map[string]bool{
		"number": true,
	}
	const jsonStream = `{
				"number": "12936782947237"
	}
	`

	input := strings.NewReader(jsonStream)
	var outString string
	buffer := bytes.NewBufferString(outString)
	err := Redact(keys, input, buffer)
	assert.NoError(t, err)

	decoder := json.NewDecoder(buffer)
	type Doc struct {
		Number string `json:"number"`
	}
	var m Doc
	err = decoder.Decode(&m)
	assert.NoError(t, err)
	assert.Equal(t, Mask, m.Number)
}
