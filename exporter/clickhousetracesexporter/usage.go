package clickhousetracesexporter

import (
	"fmt"
	"strings"

	"github.com/NoirAI/noirai-otel-collector/usage"
	"github.com/google/uuid"
	"go.opencensus.io/metric/metricdata"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

const (
	NoirAISentSpansKey      = "singoz_sent_spans"
	NoirAISentSpansBytesKey = "singoz_sent_spans_bytes"
	NoirAISpansCount        = "noirai_spans_count"
	NoirAISpansBytes        = "noirai_spans_bytes"
)

var (
	// Measures for usage
	ExporterNoirAISentSpans = stats.Int64(
		NoirAISentSpansKey,
		"Number of noirai log records successfully sent to destination.",
		stats.UnitDimensionless)
	ExporterNoirAISentSpansBytes = stats.Int64(
		NoirAISentSpansBytesKey,
		"Total size of noirai log records successfully sent to destination.",
		stats.UnitDimensionless)

	// Views for usage
	SpansCountView = &view.View{
		Name:        NoirAISpansCount,
		Measure:     ExporterNoirAISentSpans,
		Description: "The number of spans exported to noirai",
		Aggregation: view.Sum(),
		TagKeys:     []tag.Key{usage.TagTenantKey, usage.TagExporterIdKey},
	}
	SpansCountBytesView = &view.View{
		Name:        NoirAISpansBytes,
		Measure:     ExporterNoirAISentSpansBytes,
		Description: "The size of spans exported to noirai",
		Aggregation: view.Sum(),
		TagKeys:     []tag.Key{usage.TagTenantKey, usage.TagExporterIdKey},
	}
)

func UsageExporter(metrics []*metricdata.Metric, id uuid.UUID) (map[string]usage.Usage, error) {
	data := map[string]usage.Usage{}
	for _, metric := range metrics {
		if !strings.Contains(metric.Descriptor.Name, NoirAISpansCount) && !strings.Contains(metric.Descriptor.Name, NoirAISpansBytes) {
			continue
		}
		exporterIndex := usage.GetIndexOfLabel(metric.Descriptor.LabelKeys, usage.ExporterIDKey)
		tenantIndex := usage.GetIndexOfLabel(metric.Descriptor.LabelKeys, usage.TenantKey)
		if exporterIndex == -1 || tenantIndex == -1 {
			return nil, fmt.Errorf("usage: failed to get index of labels")
		}
		if strings.Contains(metric.Descriptor.Name, NoirAISpansCount) {
			for _, v := range metric.TimeSeries {
				if v.LabelValues[exporterIndex].Value != id.String() {
					continue
				}
				tenant := v.LabelValues[tenantIndex].Value
				if d, ok := data[tenant]; ok {
					d.Count = v.Points[0].Value.(int64)
					data[tenant] = d
				} else {
					data[tenant] = usage.Usage{
						Count: v.Points[0].Value.(int64),
					}
				}
			}
		} else if strings.Contains(metric.Descriptor.Name, NoirAISpansBytes) {
			for _, v := range metric.TimeSeries {
				if v.LabelValues[exporterIndex].Value != id.String() {
					continue
				}
				tenant := v.LabelValues[tenantIndex].Value
				if d, ok := data[tenant]; ok {
					d.Size = v.Points[0].Value.(int64)
					data[tenant] = d
				} else {
					data[tenant] = usage.Usage{
						Size: v.Points[0].Value.(int64),
					}
				}
			}
		}
	}
	return data, nil
}
