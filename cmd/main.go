package main

import (
	"context"
	"encoding/json"
	"net/http"

	pb "proto-test/service"
)

type Messages struct{}

func (h *Messages) GetMessage(context.Context, *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	return &pb.GetMessageResponse{
		Message: &pb.Message{
			Type:    "asdSERVER",
			Payload: "nekaPoruka",
		},
	}, nil
}

func main() {
	messagesHandler := pb.NewMessagesServer(&Messages{}, nil)
	http.ListenAndServe(":3000", messagesHandler)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	resp, _ := json.Marshal(struct {
		Ping string `json:"ping"`
	}{Ping: "pong"})
	_, _ = w.Write(resp)
}
