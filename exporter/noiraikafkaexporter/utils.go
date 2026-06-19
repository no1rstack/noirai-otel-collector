package noiraikafkaexporter

import (
	"go.opentelemetry.io/collector/client"
)

const (
	DefaultKafkaTopicPrefix = "default"
)

// getKafkaTopicFromClientMetadata returns the kafka topic from client metadata
func getKafkaTopicPrefixFromClientMetadata(md client.Metadata) string {
	// return default topic if no tenant id is found in client metadata
	noiraiTenantId := md.Get("noirai_tenant_id")
	if len(noiraiTenantId) != 0 {
		return noiraiTenantId[0]
	}

	return DefaultKafkaTopicPrefix
}
