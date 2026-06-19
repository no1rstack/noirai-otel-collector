package noirailogspipelinestanzaadapter

import (
	noirailogspipelinestanzaoperator "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator"
)

type BaseConfig struct {
	// Using our own version of Config allows using a dedicated registry of stanza ops for logs pipelines.
	Operators []noirailogspipelinestanzaoperator.Config `mapstructure:"operators"`
}
