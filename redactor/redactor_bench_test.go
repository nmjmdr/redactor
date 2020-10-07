package redactor

import (
	"bytes"
	"strings"
	"testing"
)

func benchRedact(input string, keys map[string]bool, b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	inReader := strings.NewReader(input)
	var outString string
	buffer := bytes.NewBufferString(outString)
	for n := 0; n < b.N; n++ {
		Redact(keys, inReader, buffer)
	}
}

func Benchmark_Simple(b *testing.B) {
	benchRedact(`{"name": true, "license": "2397973"}`, map[string]bool{"license": true}, b)
}

func Benchmark_Large(b *testing.B) {
	benchRedact(`{"customers":[{"name":{"firstName":"Barnett","lastName":"Henson"},"email":"barnetthenson@furnitech.com","license":"3r873r7r3983r73r7"},{"name":{"firstName":"Richards","lastName":"Bernard"},"email":"richardsbernard@furnitech.com","license":"3r873r7r3983r73r7"},{"name":{"firstName":"Holman","lastName":"Frank"},"email":"holmanfrank@furnitech.com","license":"3r873r7r3983r73r7"},{"name":{"firstName":"Lorena","lastName":"Monroe"},"email":"lorenamonroe@furnitech.com","license":"3r873r7r3983r73r7"},{"name":{"firstName":"Romero","lastName":"Hines"},"email":"romerohines@furnitech.com","license":"3r873r7r3983r73r7"},{"name":{"firstName":"Salas","lastName":"Mcleod"},"email":"salasmcleod@furnitech.com","license":"3r873r7r3983r73r7"},{"name":{"firstName":"Fisher","lastName":"Kim"},"email":"fisherkim@furnitech.com","license":"3r873r7r3983r73r7"},{"name":{"firstName":"Bernice","lastName":"Bryan"},"email":"bernicebryan@furnitech.com","license":"3r873r7r3983r73r7"}]}`,
		map[string]bool{"customers.name.lastName": true, "customers.email": true, "customers.license": true},
		b,
	)
}
