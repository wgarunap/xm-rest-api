package producer

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/tryfix/kstream/data"
	"github.com/tryfix/kstream/producer"
	"github.com/tryfix/log"
	"github.com/wgarunap/xm-rest-api/config"
	"github.com/wgarunap/xm-rest-api/domain"
	"time"
)

var _ domain.Producer = (*prod)(nil)

type prod struct {
	producer producer.Producer
}

func (p *prod) Produce(topic string, key, value []byte) error {
	_, _, err := p.producer.Produce(context.Background(), &data.Record{
		Key:       key,
		Value:     value,
		Topic:     topic,
		Timestamp: time.Now(),
	})
	if err != nil {
		return err
	}
	//fmt.Println(fmt.Sprintf(`topic:%v key:%v value:%v`, topic, string(key), string(value)))
	return nil
}

func NewProducer(conf *config.Conf) domain.Producer {
	pcon := producer.NewConfig()
	pcon.Logger = log.StdLogger
	pcon.BootstrapServers = conf.KafkaBootstrapServers
	pcon.ClientID = `xm-api-client-ip`
	pcon.Version = sarama.V2_4_0_0
	pcon.RequiredAcks = producer.WaitForAll
	pcon.Partitioner = producer.HashBased
	//pcon.MetricsReporter = c.Resolve(application.ModuleStreamReporter).(metrics.Reporter)

	p, err := producer.NewProducer(pcon)
	if err != nil {
		panic(`error initializing kafka producer`)
	}
	return &prod{producer: p}
}
