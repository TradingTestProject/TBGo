package main

import (
	_ "TBGo/proto/generated"
	investapi "TBGo/proto/generated"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-co-op/gocron"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var cfg Config
var logger *zap.Logger

func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}

func runCronJobs() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Seconds().Do(func() {
		hello("John Doe")
	})
	s.StartBlocking()
}

func main() {
	runCronJobs()

	baseClientOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{Time: 6 * time.Minute, Timeout: 1 * time.Second}),
		grpc.WithBlock(),
	}

	baseClientOpts = appendMetadataGrpcOption(baseClientOpts, "x-app-name", "tbgo")
	prodClientOpts := appendMetadataGrpcOption(baseClientOpts, "Authorization", "Bearer "+cfg.BrokerAPIToken)

	investCtx, investCancel := context.WithTimeout(context.Background(), time.Second*10)
	defer investCancel()

	prodInvestApiConn, err := grpc.DialContext(
		investCtx,
		"sandbox-invest-public-api.tinkoff.ru:443",
		prodClientOpts...,
	)
	if err != nil {
		logger.Fatal("can't connect to SandBox invest api", zap.Error(err))
	}

	marketApi := investapi.NewMarketDataServiceClient(prodInvestApiConn)
	candles, err := marketApi.GetCandles(context.Background(), &investapi.GetCandlesRequest{
		Figi:     "",
		From:     timestamppb.New(time.Now()),
		To:       timestamppb.New(time.Now().Add(-1)),
		Interval: MapToCandleInterval(CandleInterval(1)),
	})
	println(candles)
}
