package shortening

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/logrusorgru/aurora"
	pb "github.com/seed95/shortening/api/proto/shortening"
	"github.com/seed95/shortening/config"
	"github.com/seed95/shortening/pkg/derrors"
	"github.com/seed95/shortening/pkg/log"
	"github.com/seed95/shortening/pkg/translate"
	"github.com/seed95/shortening/server/rpc"
	"github.com/seed95/shortening/service"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"net"
	"net/http"
)

type (
	handler struct {
		pb.UnimplementedShorteningServer
		shorteningService service.Shortening
		logger            log.Logger
		translator        translate.Translator
	}

	Option struct {
		Cfg               *config.Server
		ShorteningService service.Shortening
		Logger            log.Logger
		Translator        translate.Translator
	}
)

func Start(opt *Option) error {

	grpcAddress := net.JoinHostPort(opt.Cfg.GRPC.Host, opt.Cfg.GRPC.Port)
	restAddress := net.JoinHostPort(opt.Cfg.Rest.Host, opt.Cfg.Rest.Port)

	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	server := &handler{
		shorteningService: opt.ShorteningService,
		logger:            opt.Logger,
		translator:        opt.Translator,
	}
	pb.RegisterShorteningServer(s, server)

	err = server.startHttp(grpcAddress, restAddress)
	if err != nil {
		return err
	}

	fmt.Printf("gRPC Server start: %v\n", aurora.Green(grpcAddress))

	reflection.Register(s)
	return s.Serve(listener)

}

func (s *handler) startHttp(grpcAddress, restAddress string) error {

	conn, err := grpc.DialContext(context.Background(), grpcAddress, grpc.WithInsecure())
	if err != nil {
		return err
	}
	router := runtime.NewServeMux(runtime.WithForwardResponseOption(httpRedirectModifier))
	if err = pb.RegisterShorteningHandler(context.Background(), router, conn); err != nil {
		return err
	}

	go func() {
		fmt.Printf("Http Server start: %v\n", aurora.Green(restAddress))
		_ = http.ListenAndServe(restAddress, wsproxy.WebsocketProxy(router))
	}()

	return nil
}

func httpRedirectModifier(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {

	headers := w.Header()
	location := headers.Get("Grpc-Metadata-Location")
	if len(location) != 0 {
		w.Header().Set("Location", location)
		w.WriteHeader(http.StatusFound)
	}
	return nil
}

func (s *handler) GenerateShort(ctx context.Context, in *pb.GenerateShortRequest) (*pb.GenerateShortResponse, error) {

	lang := rpc.GetLang(ctx)

	shortLink, err := s.shorteningService.GenerateShort(in.GetOriginalLink(), in.GetAlias(), in.GetExpiration())
	if err != nil {
		msg, code := derrors.GRPCError(err)
		return nil, status.Error(code, s.translator.Translate(msg, lang...))
	}

	res := &pb.GenerateShortResponse{
		OriginalLink: in.OriginalLink,
		ShortLink:    shortLink,
		Expiration:   in.Expiration,
	}
	return res, nil
}

func (s *handler) GetOriginal(ctx context.Context, in *pb.GetOriginalRequest) (*pb.GetOriginalResponse, error) {

	lang := rpc.GetLang(ctx)
	key := in.GetKey()

	originalLink, err := s.shorteningService.GetOriginalLink(key)
	if err != nil {
		msg, code := derrors.GRPCError(err)
		return nil, status.Error(code, s.translator.Translate(msg, lang...))
	}

	res := &pb.GetOriginalResponse{
		OriginalLink: originalLink,
		ShortLink:    s.shorteningService.GetShortLink(key),
	}

	return res, nil
}

func (s *handler) Redirect(ctx context.Context, in *pb.RedirectRequest) (*pb.RedirectResponse, error) {

	lang := rpc.GetLang(ctx)

	originalLink, err := s.shorteningService.GetOriginalLink(in.GetKey())
	if err != nil {
		msg, code := derrors.GRPCError(err)
		return nil, status.Error(code, s.translator.Translate(msg, lang...))
	}

	header := metadata.Pairs("Location", originalLink)

	_ = grpc.SendHeader(ctx, header)

	return &pb.RedirectResponse{}, nil

}
