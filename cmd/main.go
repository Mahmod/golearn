// main.go
package main

import (
	"fmt"
	"log"
	"google.golang.org/protobuf/proto"
	clientanalytics "metrics/log/event"
)

func main() {
	fmt.Println("Welcome to the Cuttlefish Metrics!")
	clientInfo := &clientanalytics.ClientInfo{
		ClientType: proto.Int32(1),
	}

	logEvent := &clientanalytics.LogEvent{
		EventTimeMs:      proto.Int64(1234567890),
		SourceExtension:  []byte("some source extension"),
	}

	logRequest := &clientanalytics.LogRequest{
		ClientInfo:    clientInfo,
		LogSource:     proto.Int32(2),
		RequestTimeMs: proto.Int64(1234567890123),
		LogEvent:      []*clientanalytics.LogEvent{logEvent},
	}
	
	data, err := proto.Marshal(logRequest)
	if err != nil {
		log.Fatalf("Failed to encode LogRequest: %v", err)
	}

	log.Printf("Encoded LogRequest: %v", data)
}
