package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
	"github.com/lehoangthienan/marvel-heroes-backend/endpoints"
	repo "github.com/lehoangthienan/marvel-heroes-backend/repository"
	groupRepo "github.com/lehoangthienan/marvel-heroes-backend/repository/group"
	heroRepo "github.com/lehoangthienan/marvel-heroes-backend/repository/hero"
	userRepo "github.com/lehoangthienan/marvel-heroes-backend/repository/user"
	"github.com/lehoangthienan/marvel-heroes-backend/service"
	authSvc "github.com/lehoangthienan/marvel-heroes-backend/service/auth"
	groupSvc "github.com/lehoangthienan/marvel-heroes-backend/service/group"
	heroSvc "github.com/lehoangthienan/marvel-heroes-backend/service/hero"
	imageSvc "github.com/lehoangthienan/marvel-heroes-backend/service/image"
	userSvc "github.com/lehoangthienan/marvel-heroes-backend/service/user"
	serviceGrpc "github.com/lehoangthienan/marvel-heroes-backend/transport/grpc"
	serviceHttp "github.com/lehoangthienan/marvel-heroes-backend/transport/http"
	"github.com/lehoangthienan/marvel-heroes-backend/transport/http/middlewares"
	"github.com/lehoangthienan/marvel-heroes-backend/util/config/db/pg"
	envConfig "github.com/lehoangthienan/marvel-heroes-backend/util/config/env"

	// "github.com/lehoangthienan/marvel-heroes-backend/util/config/redis"
	tx "github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var isProduction = envConfig.GetENV() == "production"
	if !isProduction {
		err := godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("failed to load .env by error: %v", err))
		}
	}

	// Setup addr
	port := "4002"
	if envConfig.GetPortEnv() != "" {
		port = envConfig.GetPortEnv()
	}

	httpAddr := fmt.Sprintf(":%v", port)

	// Setup log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Setup locale
	{
		loc, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			logger.Log("error", err)
			os.Exit(1)
		}
		time.Local = loc
	}

	// Setup service
	var (
		pgDB, closeDB = pg.New(envConfig.GetPGDataSourceEnv())
		// redisClient, closeRedisConn = redis.NewRedisClient(envConfig.GetRedisAddr())

		userRepo  = userRepo.NewRepo(pgDB)
		heroRepo  = heroRepo.NewRepo(pgDB)
		groupRepo = groupRepo.NewRepo(pgDB)

		repo = repo.Repository{
			UserRepository:  userRepo,
			HeroRepository:  heroRepo,
			GroupRepository: groupRepo,
		}

		txSvc   = tx.NewTransactionService(tx.NewConfig(pgDB))
		userSvc = service.Compose(
			userSvc.NewService(repo, txSvc),
			userSvc.ValidatingMiddleware(),
		).(userSvc.Service)
		heroSvc = service.Compose(
			heroSvc.NewService(repo, txSvc),
			heroSvc.ValidatingMiddleware(),
		).(heroSvc.Service)
		groupSvc = service.Compose(
			groupSvc.NewService(repo, txSvc),
			groupSvc.ValidatingMiddleware(),
		).(groupSvc.Service)
		authSvc = service.Compose(
			authSvc.NewAuthService(pgDB),
			authSvc.ValidationMiddleware(),
		).(authSvc.Service)
		imageSvc = service.Compose(
			imageSvc.NewService(),
			imageSvc.ValidatingMiddleware(),
		).(imageSvc.Service)

		s = service.Service{
			UserService:  userSvc,
			HeroService:  heroSvc,
			AuthService:  authSvc,
			GroupService: groupSvc,
			ImageService: imageSvc,
		}
	)

	defer closeDB()
	// defer closeRedisConn()

	endpoints := endpoints.MakeServerEndpoints(s)
	middlewares := middlewares.MakeHTTPpMiddleware(s)

	var h http.Handler
	{
		h = serviceHttp.NewHTTPHandler(
			middlewares,
			endpoints,
			logger,
		)
	}

	errs := make(chan error)

	go func() {
		logger.Log("transport", "HTTP", "addr", httpAddr)
		errs <- http.ListenAndServe(httpAddr, h)
	}()
	// grpc
	portGRPC := "4004"
	if envConfig.GetGRPCPortEnv() != "" {
		portGRPC = envConfig.GetGRPCPortEnv()
	}

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{}

	if envConfig.GetENV() == "secure-grpc" || isProduction {
		// Create the TLS credentials
		creds, err := tls.X509KeyPair([]byte(envConfig.GetServerCRT()), []byte(envConfig.GetServerKey()))
		if err != nil {
			logger.Log("could not load TLS keys", err)
		}
		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&creds)))
	}

	var (
		grpcServer = grpc.NewServer(opts...)
		grpcAddr   = fmt.Sprintf(":%v", portGRPC)
	)
	serviceGrpc.NewGRPCHandler(
		endpoints,
		logger,
		grpcServer,
	)

	go func() {
		lis, err := net.Listen("tcp", grpcAddr)
		defer lis.Close()
		if err != nil {
			errs <- err
		}

		logger.Log("transport", "GRPC", "addr", grpcAddr)
		errs <- grpcServer.Serve(lis)
	}()

	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-ch)
	}()

	logger.Log("exit", <-errs)
}
