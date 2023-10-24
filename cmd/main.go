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
	"net/http"
	"bytes"
	"io/ioutil"
	"errors"
)

const (
	//Refernce google3/wireless/android/play/playlog/proto/clientanalytics.proto
	ClientTypeValue   = 16
	LogSourceValue    = 971
	LogSourceNameValue = "CUTTLEFISH_METRICS"
)

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func encodeLogRequest(extension []byte) ([]byte, error) {
	currentTimeMs := currentTimeMillis()
	log.Printf("current time = %v", currentTimeMs)

	clientInfo := &clientanalytics.ClientInfo{
		ClientType: proto.Int32(ClientTypeValue),
	}

	logEvent := &clientanalytics.LogEvent{
		EventTimeMs:     proto.Int64(currentTimeMs),
		SourceExtension: extension,
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


func createAndEncodeAtestLogEventInternal(commandLine string) ([]byte, error) {
	// Generate UUIDs for UserKey and RunID
	userKey := uuid.New().String()
	runID := uuid.New().String()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}
	log.Printf("cwd = %v", cwd)


	atestStartEvent := &internal_user_log.AtestLogEventInternal_AtestStartEvent{
		CommandLine:    proto.String(commandLine),
		TestReferences: TestReferences,
		Cwd:            proto.String(cwd),
		Os:             proto.String(runtime.GOOS),
	}

	logEvent := &internal_user_log.AtestLogEventInternal{
		UserKey:         proto.String(userKey),
		RunId:           proto.String(runID),
		UserType:        &UserType,
		ToolName:        proto.String(ToolName),
		Event:           &internal_user_log.AtestLogEventInternal_AtestStartEvent_{AtestStartEvent: atestStartEvent},
	}
	return proto.Marshal(logEvent)
}

// ClearcutServer represents the Clearcut server.
type ClearcutServer int

// Enum values for ClearcutServer
const (
	Local ClearcutServer = iota + 1
	Staging
	Prod
)

// clearcutServerUrl returns the URL of the Clearcut server based on the given server type.
func clearcutServerUrl(server ClearcutServer) string {
	switch server {
	case Local:
		return "http://localhost:27910/log"
	case Staging:
		return "https://play.googleapis.com:443/staging/log"
	case Prod:
		return "https://play.googleapis.com:443/log"
	default:
		log.Println("Invalid host configuration")
		return ""
	}
}

func postRequest(output []byte) (error) {
	clearcutURL := clearcutServerUrl(Prod)

	resp, err := http.Post(clearcutURL, "application/json", bytes.NewBuffer(output))
	if err != nil {
		log.Printf("Failed to initialize HTTP request: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Metrics message failed: %s", string(output))
		log.Printf("HTTP error code: %d", resp.StatusCode)
		log.Printf("HTTP response body: %s", string(body))
		return errors.New("metrics message failed")
	}

	log.Println("Metrics posted to ClearCut")
	return nil
}

func sendLaunchCommand(commandLine string) (error) {
	log.Printf("Command Line: %v", commandLine)
	encoded, err := createAndEncodeAtestLogEventInternal(commandLine)
	if err != nil {
		log.Fatalf("Failed to encode log event: %v", err)
	}
	//Pass the encoded log event to the encodeLogRequest function
    data, err := encodeLogRequest(encoded)
    if err != nil {
        log.Println("Marshaling error: ", err)
    }
	//log.Printf("Encoded Log Event: %v", encoded)
    //log.Printf("Encoded data: %v", data)
	
	//Pass the encoded data to the postRequest function
	err = postRequest(data)

	if err != nil {
		log.Println("Failed to post request: ", err)
		return err
	}
	return nil
}

func main() {
	log.Printf("-----------------------------------")
	log.Printf("Welcome to the Cuttlefish Metrics!")
	commandLine := strings.Join(os.Args, " ")
	err := sendLaunchCommand(commandLine)
	if err != nil {
		log.Println("Failed to send launch command to metrics server: ", err)
	}

}
