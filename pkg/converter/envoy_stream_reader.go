package converter

import (
	"encoding/json"
	"fmt"
	v2 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v2"
	"github.com/golang/protobuf/jsonpb"
	"io"
)

type server struct {
	marshaler jsonpb.Marshaler
}

var _ v2.AccessLogServiceServer = &server{}

func New() v2.AccessLogServiceServer {
	return &server{}
}

func (s *server) StreamAccessLogs(stream v2.AccessLogService_StreamAccessLogsServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		logEntries := transform(in.GetHttpLogs())
		for _, logEntry := range logEntries {
			var data []byte
			data, err = json.Marshal(logEntry)
			fmt.Println(string(data))
		}
		//_ = es.Close(context.Background())
	}
}
