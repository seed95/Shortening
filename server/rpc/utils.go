package rpc

import (
	"context"
	"espad_task/pkg/translate"
	"espad_task/server"
	"google.golang.org/grpc/metadata"
	"strings"
)

func GetLang(ctx context.Context) []translate.Language {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return []translate.Language{}
	}
	languages := md.Get("grpcgateway-accept-language")
	return server.GetLanguage(strings.Join(languages, ","))
}
