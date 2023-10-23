// main.go
package main

import (
    "log"
    "github.com/golang/protobuf/proto"
    clientanalytics "metrics/log/event"
	internal_user_log "metrics/log/atest"
	"runtime"
	"os"
	"time"
	"github.com/google/uuid"
	"strings"
)

const (
	//TODO: change it to appropriate clearcut value
	ClientTypeValue   = 1
	LogSourceValue    = 2
	LogSourceNameValue = "cuttlefish"
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
		LogSourceName: proto.String(LogSourceNameValue),
	}

	return proto.Marshal(req)
}

const (
	ToolName      = "cvdr"
)

var (
	UserType        = internal_user_log.UserType_GOOGLE
	TestReferences  = []string{"test1", "test2"}
)


func encodeAtestLogEventInternal(logEvent *internal_user_log.AtestLogEventInternal) ([]byte, error) {
	return proto.Marshal(logEvent)
}

func createAtestLogEventInternal(commandLine string) *internal_user_log.AtestLogEventInternal {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}
	log.Printf("cwd = %v", cwd)
	// Generate UUIDs for UserKey and RunID
	userKey := uuid.New().String()
	runID := uuid.New().String()

	atestStartEvent := &internal_user_log.AtestLogEventInternal_AtestStartEvent{
		CommandLine:    proto.String(commandLine),
		TestReferences: TestReferences,
		Cwd:            proto.String(cwd),
		Os:             proto.String(runtime.GOOS),
	}

	return &internal_user_log.AtestLogEventInternal{
		UserKey:         proto.String(userKey),
		RunId:           proto.String(runID),
		UserType:        &UserType,
		ToolName:        proto.String(ToolName),
		Event:           &internal_user_log.AtestLogEventInternal_AtestStartEvent_{AtestStartEvent: atestStartEvent},
	}
}

func main() {
	log.Printf("-----------------------------------")
	log.Printf("Welcome to the Cuttlefish Metrics!")
	commandLine := strings.Join(os.Args, " ")
	logEvent := createAtestLogEventInternal(commandLine)
	encoded, err := encodeAtestLogEventInternal(logEvent)
	if err != nil {
		log.Fatalf("Failed to encode log event: %v", err)
	}
	log.Printf("Encoded Log Event: %v", encoded)

	log.Printf("-----------------------------------")

    extension := "source_extension"
    data, err := encodeLogRequest(extension)
    if err != nil {
        log.Fatal("Marshaling error: ", err)
    }
    log.Printf("Encoded data: %v", data)
}
