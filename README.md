# duration

This is a copy of this [`duration`](https://git.maze.io/go/duration) package with updated `go.mod` file.

Parse standard units of time based on go's stdlib time implementation, extended to use days, weeks and years with
approximations.

Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h", "d", "w", "y".

Furthermore, custom marshaller and un-marshaller are defined for yaml and json documents, thus marshalling fields with
duration such as "1w" are possible:

```go
var jsonFile = `
{
    "some_duration": "1w"
}
`

type output struct {
    SomeDuration Duration `json:"some_duration"`
}

func main() {
    var out output
    json.NewDecoder(strings.NewReader(jsonFile)).Decode(&out)
}
```