package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/config"
	appGrpc "gitlab.ozon.dev/cyxrop/homework-2/internal/app/grpc/user"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/mail"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/mw"
	mbRepository "gitlab.ozon.dev/cyxrop/homework-2/internal/app/repository/mailbox"
	uRepository "gitlab.ozon.dev/cyxrop/homework-2/internal/app/repository/user"
	service "gitlab.ozon.dev/cyxrop/homework-2/internal/app/service/user"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/worker"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/db"
	"gitlab.ozon.dev/cyxrop/homework-2/pkg/api"
	"gitlab.ozon.dev/cyxrop/homework-2/pkg/cryptographer"
	"gitlab.ozon.dev/cyxrop/homework-2/pkg/telegram"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	EnvTgToken  = "TG_TOKEN"
	EnvCryptKey = "CRYPT_KEY"
	EnvDbConn   = "DB_CONN"
)

func main() {
	tgToken := os.Getenv(EnvTgToken)
	if tgToken == "" {
		log.Printf("Environment variable %q not specified\n", EnvTgToken)
		return
	}

	cryptKey := os.Getenv(EnvCryptKey)
	if cryptKey == "" {
		log.Printf("Environment variable %q not specified\n", EnvCryptKey)
		return
	}

	dbConnStr := os.Getenv(EnvDbConn)
	if dbConnStr == "" {
		log.Printf("Environment variable %q not specified\n", EnvDbConn)
		return
	}

	cfgFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Printf("read config file: %s\n", err)
		return
	}

	cfg, err := config.Parse(cfgFile)
	if err != nil {
		log.Printf("parse config file: %s\n", err)
		return
	}

	imapHosts := map[mail.Type]string{
		mail.TypeGoogle: cfg.Imap.Hosts.Google,
		mail.TypeMailRu: cfg.Imap.Hosts.MailRu,
		mail.TypeYandex: cfg.Imap.Hosts.Yandex,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbPool, err := db.New(ctx, dbConnStr)
	if err != nil {
		log.Printf("db connect: %s\n", err)
		return
	}
	defer dbPool.Close()

	poller := mail.NewPoller(imapHosts)
	crypt := cryptographer.NewAES(cryptKey)
	ur := uRepository.NewUserRepository(dbPool)
	mr := mbRepository.NewMailboxRepository(dbPool)

	tgBot, err := telegram.NewBot(tgToken, cfg.Telegram.Timeout, cfg.Telegram.Debug)
	if err != nil {
		log.Printf("create tg bot: %s\n", err)
		return
	}

	us := service.NewUserService(ur, mr, crypt, tgBot, poller)

	// Run background workers
	tgListener := worker.NewTgListener(tgBot, us)
	tgListener.Run(ctx)

	notifier := worker.NewNotifier(us, time.Minute*time.Duration(cfg.Notifier.Every))
	notifier.Run(ctx)

	// Run grpc server
	userServer := appGrpc.NewUserServiceServer(us)
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(mw.LogInterceptor))
	api.RegisterUserServiceServer(grpcServer, userServer)

	grpcEndpoint := fmt.Sprintf(":%d", cfg.App.GRPC.Port)
	lis, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		log.Printf("listen: %v\n", err)
		return
	}

	// Register grpc gateway server
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err = api.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts); err != nil {
		log.Fatalf("register grpc gateway server: %s", err)
	}

	errCh := make(chan error)
	signals := make(chan os.Signal)

	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Println("Run GRPC server...")
		if err = grpcServer.Serve(lis); err != nil {
			log.Printf("serve grpc: %s\n", err)
			errCh <- err
		}
	}()

	go func() {
		log.Println("Run GRPC Gateway server...")
		if err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.App.HTTP.Port), mux); err != nil {
			log.Printf("serve grpc gateway: %s\n", err)
			errCh <- err
		}
	}()

	select {
	case err = <-errCh:
		log.Printf("Graceful shutdown on error: %s\n", err)
		gracefulShutdown(cancel, grpcServer)
	case <-signals:
		log.Println("Graceful shutdown on syscall")
		gracefulShutdown(cancel, grpcServer)
	}
}

func gracefulShutdown(cancel context.CancelFunc, server *grpc.Server) {
	cancel()
	server.GracefulStop()
}
