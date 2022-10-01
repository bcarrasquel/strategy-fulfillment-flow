package kafkaclient

import (
	"io/ioutil"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const ( // enums
	KAFKA_VERSION_2_4 = "v2.4"
	KAFKA_VERSION_2_8 = "v2.8"
)

func GetKafkConfigByVersion(version string) kafka.ConfigMap {
	certPassword, _ := ioutil.ReadFile(os.Getenv("SSL_KEY_PASSWORD"))
	connections := map[string]kafka.ConfigMap{
		KAFKA_VERSION_2_4: {
			"bootstrap.servers":        os.Getenv("KAFKA_BROKERS_SSL"),
			"security.protocol":        "SSL",
			"ssl.ca.location":          os.Getenv("SSL_CA_LOCATION"),
			"ssl.key.location":         os.Getenv("SSL_KEY_LOCATION"),
			"ssl.certificate.location": os.Getenv("SSL_CERTIFICATE_LOCATION"),
			"ssl.key.password":         string(certPassword),
		},
		KAFKA_VERSION_2_8: {
			"bootstrap.servers": os.Getenv("KAFKA_BROKERS"),
		},
	}
	return connections[version]
}
