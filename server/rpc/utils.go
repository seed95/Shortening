package rpc

import (
	"context"
	"github.com/seed95/shortening/pkg/translate"
	"github.com/seed95/shortening/server"
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
