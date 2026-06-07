// Brought in as is from opentelemetry-collector-contrib
package trace

import (
	"path/filepath"
	"testing"

	noiraistanzaentry "github.com/NoirAI/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/entry"
	noiraistanzahelper "github.com/NoirAI/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/helper"
	"github.com/NoirAI/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operatortest"
)

func TestConfig(t *testing.T) {
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
				Name: "spanid",
				Expect: func() *Config {
					parseFrom := noiraistanzaentry.Field{FieldInterface: noiraistanzaentry.NewBodyField("app_span_id")}
					cfg := noiraistanzahelper.SpanIDConfig{}
					cfg.ParseFrom = &parseFrom

					c := NewConfig()
					c.SpanID = &cfg
					return c
				}(),
			},
			{
				Name: "traceid",
				Expect: func() *Config {
					parseFrom := noiraistanzaentry.Field{FieldInterface: noiraistanzaentry.NewBodyField("app_trace_id")}
					cfg := noiraistanzahelper.TraceIDConfig{}
					cfg.ParseFrom = &parseFrom

					c := NewConfig()
					c.TraceID = &cfg
					return c
				}(),
			},
			{
				Name: "trace_flags",
				Expect: func() *Config {
					parseFrom := noiraistanzaentry.Field{FieldInterface: noiraistanzaentry.NewBodyField("app_trace_flags_id")}
					cfg := noiraistanzahelper.TraceFlagsConfig{}
					cfg.ParseFrom = &parseFrom

					c := NewConfig()
					c.TraceFlags = &cfg
					return c
				}(),
			},
		},
	}.Run(t)
}
