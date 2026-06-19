// Brought in as is from opentelemetry-collector-contrib

package severity // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/severity"

import (
	"context"

	noiraistanzahelper "github.com/no1rstack/noirai-otel-collector/processor/noirailogspipelineprocessor/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// Parser is an operator that parses severity from a field to an entry.
type Parser struct {
	noiraistanzahelper.TransformerOperator
	noiraistanzahelper.SeverityParser
}

// Process will parse severity from an entry.
func (p *Parser) Process(ctx context.Context, entry *entry.Entry) error {
	return p.ProcessWith(ctx, entry, p.Parse)
}

func (p *Parser) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	return p.ProcessBatchWith(ctx, entries, p.Process)
}
