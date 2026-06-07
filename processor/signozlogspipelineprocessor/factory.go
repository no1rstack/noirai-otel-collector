// Async processor: ConsumeLogs returns immediately after enqueuing to
// FromPdataConverter. The factory therefore bypasses processorhelper.NewLogs
// (which only supports sync ProcessLogs callbacks) and returns the processor
// directly as a processor.Logs implementation.
package noirailogspipelineprocessor

import (
	"context"
	"errors"
	"fmt"

	noirailogspipelinestanzaadapter "github.com/NoirAI/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/adapter"
	noirailogspipelinestanzaoperator "github.com/NoirAI/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"

	"github.com/NoirAI/noirai-otel-collector/processor/noirailogspipelineprocessor/internal/metadata"
)

func NewFactory() processor.Factory {
	return processor.NewFactory(
		metadata.Type,
		createDefaultConfig,
		processor.WithLogs(createLogsProcessor, metadata.LogsStability))
}

// Note: This isn't a valid configuration (no operators would lead to no work being done)
func createDefaultConfig() component.Config {
	return &Config{
		BaseConfig: noirailogspipelinestanzaadapter.BaseConfig{
			Operators: []noirailogspipelinestanzaoperator.Config{},
		},
	}
}

func createLogsProcessor(
	_ context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	pCfg, ok := cfg.(*Config)
	if !ok {
		return nil, errors.New("could not initialize noirailogspipeline processor")
	}
	if len(pCfg.BaseConfig.Operators) == 0 {
		return nil, errors.New("no operators were configured for noirailogspipeline processor")
	}

	proc, err := newLogsPipelineProcessor(pCfg, set, nextConsumer)
	if err != nil {
		return nil, fmt.Errorf("couldn't build \"noirailogspipeline\" processor %w", err)
	}

	return proc, nil
}
