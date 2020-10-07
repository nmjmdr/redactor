# Redactor

Masks properties from an input JSON string and writes the masked output to a output stream

Example:

```
    keys := map[string]bool{
		"licenses.number": true,
	}
	const jsonStream = `{"licenses":[{"name":"Savarkar","number":"123456347"},{"name":" Hegdewar","number":"135862286"}]}`

	input := strings.NewReader(jsonStream)
	err := Redact(keys, input, os.Stdout)
```
```
Output:
{"licenses":[{"name":"Savarkar","number":"******"},{"name":" Hegdewar","number":"*****"}]}
```

Benchmark:
```
goos: darwin
goarch: amd64
pkg: redactor/redactor
Benchmark_Simple-4   	 4850139	       256 ns/op	     880 B/op	       3 allocs/op
Benchmark_Large-4    	 4736006	       251 ns/op	     880 B/op	       3 allocs/op
PASS
ok  	redactor/redactor	2.958s
```