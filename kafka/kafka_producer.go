/*
 * Copyright (c) 2017 TFG Co
 * Author: TFG Co <backend@tfgco.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package kafka

import (
	log "github.com/Sirupsen/logrus"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	raven "github.com/getsentry/raven-go"
	"github.com/spf13/viper"
	"github.com/topfreegames/extensions/kafka/interfaces"
	"github.com/topfreegames/extensions/util"
)

// Producer for producing push feedbacks to a kafka queue
type Producer struct {
	Brokers  string
	Config   *viper.Viper
	Producer interfaces.KafkaProducerClient
	Logger   *log.Logger
}

// NewProducer for creating a new Producer instance
func NewProducer(config *viper.Viper, logger *log.Logger, clientOrNil ...interfaces.KafkaProducerClient) (*Producer, error) {
	q := &Producer{
		Config: config,
		Logger: logger,
	}
	var producer interfaces.KafkaProducerClient
	if len(clientOrNil) == 1 {
		producer = clientOrNil[0]
	}
	err := q.configure(producer)
	return q, err
}

func (q *Producer) loadConfigurationDefaults() {
	q.Config.SetDefault("extensions.kafkaproducer.brokers", "localhost:9941")
}

func (q *Producer) configure(producer interfaces.KafkaProducerClient) error {
	q.loadConfigurationDefaults()
	q.Brokers = q.Config.GetString("extensions.kafkaproducer.brokers")
	c := &kafka.ConfigMap{
		"bootstrap.servers": q.Brokers,
	}
	l := q.Logger.WithFields(log.Fields{
		"brokers": q.Brokers,
	})
	l.Debug("configuring kafka producer")

	if producer == nil {
		p, err := kafka.NewProducer(c)
		q.Producer = p
		if err != nil {
			l.WithError(err).Error("error configuring kafka producer client")
			return err
		}
	} else {
		q.Producer = producer
	}
	go q.listenForKafkaResponses()
	l.Info("kafka producer initialized")
	return nil
}

func (q *Producer) listenForKafkaResponses() {
	l := q.Logger.WithFields(log.Fields{
		"method": "listenForKafkaResponses",
	})
	for e := range q.Producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			m := ev
			if m.TopicPartition.Error != nil {
				raven.CaptureError(m.TopicPartition.Error, map[string]string{
					"version":   util.Version,
					"extension": "kafka-producer",
				})
				l.WithError(m.TopicPartition.Error).Error("error sending feedback to kafka")
			} else {
				l.WithFields(log.Fields{
					"topic":     *m.TopicPartition.Topic,
					"partition": m.TopicPartition.Partition,
					"offset":    m.TopicPartition.Offset,
				}).Debug("delivered feedback to topic")
			}
			break
		default:
			l.WithField("event", ev).Warn("ignored kafka response event")
		}
	}
}

// SendMessage sends the message to a topic of kafka Queue
func (q *Producer) SendMessage(message []byte, topic string) {
	m := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: message,
	}
	q.Producer.ProduceChannel() <- m
}