package teehee

import (
	"io"
	"strings"
	"testing"
)

var TestConfigTable = []struct {
	File       string
	ShouldFail bool
}{
	{
		File:       "non-existent.toml",
		ShouldFail: true,
	},
	{
		File:       "example_config.toml",
		ShouldFail: false,
	},
}

func TestConfig(t *testing.T) {
	for _, tt := range TestConfigTable {
		err := Config(tt.File)
		if tt.ShouldFail {
			if err == nil {
				t.Errorf("Expected a failed result for %s\n", tt.File)
			}
		} else {
			if err != nil {
				t.Errorf("Expected a non-failed result for %s\n", tt.File)
			}
		}
	}
}

var TestConfigFromReaderTable = []struct {
	Config     *io.reader
	ShouldFail bool
}{
	{
		Config: strings.NewReader(`
		[http]
		ip = "0.0.0.0"
		port = "8090"
		`),
		ShouldFail: false,
	},
}

func TestConfigFromReader(t *testing.T) {

}
