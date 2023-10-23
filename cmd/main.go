// main.go
package main

import (
    "log"
    "google.golang.org/protobuf/proto"
    clientanalytics "metrics/log/event"
)

func encodeLogRequest(extension string) ([]byte, error) {
	clientInfo := &clientanalytics.ClientInfo{
		ClientType: proto.Int32(1),
	}

	logEvent := &clientanalytics.LogEvent{
		EventTimeMs:      proto.Int64(1625140123456),
		SourceExtension:  []byte(extension),
	}

    req := &clientanalytics.LogRequest{
		ClientInfo:    clientInfo,
		LogSource:     proto.Int32(2),
		RequestTimeMs: proto.Int64(1625140123456),
		LogEvent:      []*clientanalytics.LogEvent{logEvent},
		//TODO: add --> LogSourceName: proto.String("log_source_name"),
	}

    return proto.Marshal(req)
}

func main() {
	log.Printf("Welcome to the Cuttlefish Metrics!")
    extension := "source_extension"
    data, err := encodeLogRequest(extension)
    if err != nil {
        log.Fatal("Marshaling error: ", err)
    }
    log.Printf("Encoded data: %v", data)
}
