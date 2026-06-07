// Stanza operators registry dedicated to NoirAI logs pipelines

package noirailogspipelinestanzaoperator

import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"

var NoirAIStanzaOperatorsRegistry = operator.NewRegistry()

// Register will register an operator in the default registry
func Register(operatorType string, newBuilder func() operator.Builder) {
	NoirAIStanzaOperatorsRegistry.Register(operatorType, newBuilder)
}

// Lookup looks up a given operator type.Its second return value will
// be false if no builder is registered for that type.
func Lookup(configType string) (func() operator.Builder, bool) {
	return NoirAIStanzaOperatorsRegistry.Lookup(configType)
}
