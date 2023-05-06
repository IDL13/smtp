package sm

import (
	"context"
	"log"
	"net/smtp"

	"smtp_server/pkg/api"

	"smtp_server/internal/config"
)

//GRPCServer
type GRPCServer struct {
	api.UnimplementedAdderServer
}

//Add
func (GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {

	cfg := config.GetConf()

	auth := smtp.PlainAuth("", cfg.Gmail, cfg.GmailKey, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, cfg.Gmail, []string{req.GetEmail()}, []byte(req.GetMsg()))
	if err != nil {
		log.Fatal(err)
	}

	return &api.AddResponse{Result: 1}, nil
}

// mustEmbedUnimplementedAdderServer ...
func (GRPCServer) mustEmbedUnimplementedAdderServer() {}
