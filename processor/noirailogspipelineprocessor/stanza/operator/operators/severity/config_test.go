// Brought in as is from opentelemetry-collector-contrib
package severity

import (
	"path/filepath"
	"testing"

	noiraistanzaentry "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/entry"
	"github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operatortest"
)

func TestUnmarshal(t *testing.T) {
	operatortest.ConfigUnmarshalTests{
		DefaultConfig: NewConfig(),
		TestsFile:     filepath.Join(".", "testdata", "config.yaml"),
		Tests: []operatortest.ConfigUnmarshalTest{
			{
				Name:   "default",
				Expect: NewConfig(),
			},
			{
				Name: "on_error_drop",
				Expect: func() *Config {
					cfg := NewConfig()
					cfg.OnError = "drop"
					return cfg
				}(),
			},
			{
				Name: "parse_from_simple",
				Expect: func() *Config {
					cfg := NewConfig()
					from := noiraistanzaentry.Field{FieldInterface: noiraistanzaentry.NewBodyField("from")}
					cfg.ParseFrom = &from
					return cfg
				}(),
			},
			{
				Name: "parse_with_preset",
				Expect: func() *Config {
					cfg := NewConfig()
					from := noiraistanzaentry.Field{FieldInterface: noiraistanzaentry.NewBodyField("from")}
					cfg.ParseFrom = &from
					cfg.Preset = "http"
					return cfg
				}(),
			},
		},
	}.Run(t)
}
