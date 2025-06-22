package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/peesaphanthavong/adapters/api"
	"github.com/peesaphanthavong/adapters/database"
	"github.com/peesaphanthavong/config"
	"github.com/peesaphanthavong/core/usecases"
)

func main() {
	var (
		cfg            = config.InitialConfig()
		app            = initFiber(cfg)
		postgresClient = newPostgresClient(context.Background(), cfg)
	)
	defer postgresClient.Close()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Ready!!")
	})

	userGroups := app.Group("api/v1/user")
	userHandel := api.NewUserHandler(*usecases.NewUserUseCase(
		database.NewUserRepositoryDB(postgresClient),
	))
	userGroups.Get("/all", userHandel.GetAllUser)
	userGroups.Post("/create", userHandel.CreateUser)

	log.Printf("Listening on port: %s", cfg.Server.Port)
	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", cfg.Server.Port)); err != nil {
			log.Fatal(err.Error())
		}
	}()

	gracefulShutdown(app)
}

func initFiber(cfg *config.Config) *fiber.App {
	app := fiber.New(
		fiber.Config{
			ReadTimeout:           cfg.Server.ReadTimeout,
			WriteTimeout:          cfg.Server.WriteTimeout,
			IdleTimeout:           cfg.Server.IdleTimeout,
			DisableStartupMessage: true,
			CaseSensitive:         true,
			StrictRouting:         true,
		},
	)
	return app
}

func gracefulShutdown(f *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := f.Shutdown(); err != nil {
		log.Fatal("shutdown server:", err)
	}
}
func newPostgresClient(ctx context.Context, cfg *config.Config) *pgxpool.Pool {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.DBname,
	)
	pgxpoolConfig, err := pgxpool.ParseConfig(psqlInfo)

	if err != nil {
		log.Fatalf("postgres client parse config error %v", err)
	}

	pgxpoolConfig.MaxConnLifetime = cfg.DB.MaxConnLifeTime
	pgxpoolConfig.MaxConnIdleTime = cfg.DB.MaxIdle
	pgxpoolConfig.MaxConns = cfg.DB.MaxOpenConn

	pool, err := pgxpool.NewWithConfig(ctx, pgxpoolConfig)
	if err != nil {
		log.Fatalf("postgres client connect error %v", err)
	}
	log.Printf("postgres client connection config %+v", pool.Config().ConnString())

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("postgres client Ping error %v", err)
	}
	return pool
}
