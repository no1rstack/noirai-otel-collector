package noiraiclickhousemetrics

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/no1rstack/noirai-otel-collector/usage"
	"go.opencensus.io/metric/metricdata"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

const (
	NoirAISentMetricPointsKey      = "noirai_sent_metric_points"
	NoirAISentMetricPointsBytesKey = "noirai_sent_metric_points_bytes"
	NoirAIMetricPointsCount        = "noirai_metric_points_count"
	NoirAIMetricPointsBytes        = "noirai_metric_points_bytes"
)

var (
	// Measures for usage
	ExporterNoirAISentMetricPoints = stats.Int64(
		NoirAISentMetricPointsKey,
		"Number of noirai metric points successfully sent to destination.",
		stats.UnitDimensionless)
	ExporterNoirAISentMetricPointsBytes = stats.Int64(
		NoirAISentMetricPointsBytesKey,
		"Total size of noirai metric points successfully sent to destination.",
		stats.UnitDimensionless)

	// Views for usage
	MetricPointsCountView = &view.View{
		Name:        NoirAIMetricPointsCount,
		Measure:     ExporterNoirAISentMetricPoints,
		Description: "The number of metric points exported to noirai",
		Aggregation: view.Sum(),
		TagKeys:     []tag.Key{usage.TagTenantKey, usage.TagExporterIdKey},
	}
	MetricPointsBytesView = &view.View{
		Name:        NoirAIMetricPointsBytes,
		Measure:     ExporterNoirAISentMetricPointsBytes,
		Description: "The size of metric points exported to noirai",
		Aggregation: view.Sum(),
		TagKeys:     []tag.Key{usage.TagTenantKey, usage.TagExporterIdKey},
	}
)

func UsageExporter(metrics []*metricdata.Metric, id uuid.UUID) (map[string]usage.Usage, error) {
	data := map[string]usage.Usage{}
	for _, metric := range metrics {
		if !strings.Contains(metric.Descriptor.Name, NoirAIMetricPointsCount) && !strings.Contains(metric.Descriptor.Name, NoirAIMetricPointsBytes) {
			continue
		}
		exporterIndex := usage.GetIndexOfLabel(metric.Descriptor.LabelKeys, usage.ExporterIDKey)
		tenantIndex := usage.GetIndexOfLabel(metric.Descriptor.LabelKeys, usage.TenantKey)
		if exporterIndex == -1 || tenantIndex == -1 {
			return nil, fmt.Errorf("usage: failed to get index of labels")
		}
		if strings.Contains(metric.Descriptor.Name, NoirAIMetricPointsCount) {
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
		} else if strings.Contains(metric.Descriptor.Name, NoirAIMetricPointsBytes) {
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
