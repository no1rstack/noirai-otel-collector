// Brought in as is from opentelemetry-collector-contrib

package noop

import (
	"go.opentelemetry.io/collector/component"

	noirailogspipelinestanzaoperator "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator"
	noiraistanzahelper "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

const operatorType = "noop"

func init() {
	noirailogspipelinestanzaoperator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// NewConfig creates a new noop operator config with default values
func NewConfig() *Config {
	return NewConfigWithID(operatorType)
}

// NewConfigWithID creates a new noop operator config with default values
func NewConfigWithID(operatorID string) *Config {
	return &Config{
		TransformerConfig: noiraistanzahelper.NewTransformerConfig(operatorID, operatorType),
	}
}

// Config is the configuration of a noop operator.
type Config struct {
	noiraistanzahelper.TransformerConfig `mapstructure:",squash"`
}

// Build will build a noop operator.
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	transformerOperator, err := c.TransformerConfig.Build(set)
	if err != nil {
		return nil, err
	}

	return &Transformer{
		TransformerOperator: transformerOperator,
	}, nil
}
