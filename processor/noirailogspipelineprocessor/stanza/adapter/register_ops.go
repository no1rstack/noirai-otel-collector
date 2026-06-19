// Register copies of stanza operators dedicated to noirai logs pipelines
package noirailogspipelinestanzaadapter

import (
	_ "github.com/no1rstack/noirai-otel-collector/pkg/parser/grok"
	_ "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operators/add"
	_ "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operators/copy"
	_ "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operators/json"
	_ "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operators/move"
	_ "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operators/noop"
	_ "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operators/normalize"
	_ "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operators/regex"
	_ "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operators/remove"
	_ "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operators/router"
	_ "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operators/severity"
	_ "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operators/time"
	_ "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/operators/trace"
)
