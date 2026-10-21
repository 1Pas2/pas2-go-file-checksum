# pas2-go-file-checksum

Stream SHA-256 digests and compare expected checksums.

Part of **Pas2 Polyglot** by [1Pas2](https://github.com/1Pas2): independent small projects in PHP, Python, Go and C#. No runtime dependencies outside the language standard library. Initial version: 0.1.0.

## Requirements and installation

Go 1.23+. Use a local checkout with `go test ./...`. After GitHub publication, applications can use `go get github.com/1Pas2/pas2-go-file-checksum@main`; no version tag is claimed.

## Example

```go
package main

import (
	"fmt"
	lib "github.com/1Pas2/pas2-go-file-checksum"
	"strings"
)

func main() {
	digest, err := lib.Digest(strings.NewReader("hello"))
	fmt.Println(digest, err)
}
```

## API and limits

Digest(reader) streams SHA-256 as lowercase hex. Verify(reader,expected) validates a 64-character hexadecimal digest and compares the result. Reader errors propagate; readers are consumed and never closed. Checksums detect changes only relative to a trusted expected value, not provenance.

## Tests and example

From the repository root:

```sh
go test ./...
go vet ./...
go run ./examples/basic
```

Tests use the standard library or a dependency-free executable check harness. A failing assertion exits nonzero. The example is self-contained and uses local in-memory or temporary data. CI runs the same checks on Linux. No registry releases, downloads, users or production deployments are claimed.

## Contributing

Open a focused issue or pull request with reproduction steps and a behavioral test. Run the commands above before submitting. Do not include credentials or private data in reports.

## Ecosystem

[Browse Pas2 projects](https://github.com/1Pas2?tab=repositories&q=pas2-). Each repository is usable independently. The earlier TypeScript series and this polyglot series have distinct repository names.

## License

[MIT](LICENSE) © 2026 1Pas2.

1Pas2
