package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
)
var (
    kafkaProducer *kafka.Writer
)

const(
    kafkaURL = "localhost:29092,localhost:29093,localhost:29094"
    kafkaTopic = "test-topic"
)
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
    brokers := strings.Split(kafkaURL, ",")
	return kafka.NewWriter(kafka.WriterConfig{
        Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

func getKafkaReader(kafkaURL, topic string, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		CommitInterval: time.Second, // commit every second
		StartOffset: kafka.FirstOffset, // start reading from the first message
	})
}

type StockInfo struct {
    Message string `json:"message"`
    Type string `json:"type"`
}

func newStock(msg, typeMsg string) *StockInfo{
    return &StockInfo{
        Message: msg,
        Type: typeMsg,
    }
}

func actionStock(c *gin.Context) {
    s := newStock(c.Query("msg"), c.Query("type"))
    body := make(map[string]any)
    body["action"] = "action"
    body["info"] = s
    jsonBody, _ := json.Marshal(body)
    // tạo message cho kafka
    msg := kafka.Message{
        Key: []byte("action"),
        Value: []byte(jsonBody),
    }
    // viết message cho producer
    err := kafkaProducer.WriteMessages(context.Background(), msg)
    if err != nil {
        c.JSON(500, gin.H{
            "error": err.Error(),
        })
        return
    }
    c.JSON(200, gin.H{
        "error": nil,
        "message": "Action sent successfully",
    })
}

// Consumer hóng mua ATC
func RegisterConsumerATC(id int){
    kafkaGroupID := "consumer-group-" + string(rune(id))
    kafkaReader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupID)
    defer kafkaReader.Close()
    fmt.Printf("Consumer %d started\n", id)
    for {
        m, err := kafkaReader.ReadMessage(context.Background())
        if err != nil {
            fmt.Printf("Consumer %d read message failed: %v\n", id, err)
            continue
        }
        fmt.Printf("Consumer %d read message from topic %s partition %d offset %d at %d key %s value %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))
    }
}

func main(){
    r := gin.Default()
    r.POST("/action/stock", actionStock)
    kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
    defer kafkaProducer.Close()
    go RegisterConsumerATC(1)
    go RegisterConsumerATC(2)
    go RegisterConsumerATC(3)

    r.Run(":8999")
}
