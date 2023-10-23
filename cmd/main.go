// main.go
package main

import (
    "log"
    "google.golang.org/protobuf/proto"
    clientanalytics "metrics/log/event"
	"time"
)

const (
	ClientTypeValue   = 1
	LogSourceValue    = 2
	//LogSourceNameValue = "log_source_name"
)

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func encodeLogRequest(extension string) ([]byte, error) {
	currentTimeMs := currentTimeMillis()
	log.Printf("current time = %v", currentTimeMs)

	clientInfo := &clientanalytics.ClientInfo{
		ClientType: proto.Int32(ClientTypeValue),
	}

	logEvent := &clientanalytics.LogEvent{
		EventTimeMs:     proto.Int64(currentTimeMs),
		SourceExtension: []byte(extension),
	}

	req := &clientanalytics.LogRequest{
		ClientInfo:    clientInfo,
		LogSource:     proto.Int32(LogSourceValue),
		RequestTimeMs: proto.Int64(currentTimeMs),
		LogEvent:      []*clientanalytics.LogEvent{logEvent},
		//LogSourceName: proto.String(LogSourceNameValue),
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
