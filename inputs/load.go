package inputs

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func Load(day int) string {
	return load(generateFilename(day))
}

func LoadExample(day, n int) string {
	return load(fmt.Sprintf("%s_example_%d", generateFilename(day), n))
}

func load(f string) string {
	contents, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(contents))
}

func generateFilename(day int) string {
	_, d, _, ok := runtime.Caller(0)
	if !ok {
		panic("could not determine current caller")
	}
	return fmt.Sprintf("%s/day_%d", filepath.Dir(d), day)
}
