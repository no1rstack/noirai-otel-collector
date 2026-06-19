// Brought in as is from opentelemetry-collector-contrib

package severity

import (
	"go.opentelemetry.io/collector/component"

	noirailogspipelinestanzaoperator "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator"
	noiraistanzahelper "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

const operatorType = "severity_parser"

func init() {
	noirailogspipelinestanzaoperator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new severity parser config with default values
func NewConfig() *Config {
	return NewConfigWithID(operatorType)
}

// NewConfigWithID creates a new severity parser config with default values
func NewConfigWithID(operatorID string) *Config {
	return &Config{
		TransformerConfig: noiraistanzahelper.NewTransformerConfig(operatorID, operatorType),
		SeverityConfig:    noiraistanzahelper.NewSeverityConfig(),
	}
}

// Config is the configuration of a severity parser operator.
type Config struct {
	noiraistanzahelper.TransformerConfig `mapstructure:",squash"`
	noiraistanzahelper.SeverityConfig    `mapstructure:",omitempty,squash"`
}

// Build will build a severity parser operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	transformerOperator, err := c.TransformerConfig.Build(set)
	if err != nil {
		return nil, err
	}

	severityParser, err := c.SeverityConfig.Build(set)
	if err != nil {
		return nil, err
	}

	return &Parser{
		TransformerOperator: transformerOperator,
		SeverityParser:      severityParser,
	}, nil
}
