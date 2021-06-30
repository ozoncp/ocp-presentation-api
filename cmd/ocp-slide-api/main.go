package main

import (
	"fmt"
	"net"
	"os"

	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/api"
	"github.com/ozoncp/ocp-presentation-api/internal/ocp-slide-api/repo"
	desc "github.com/ozoncp/ocp-presentation-api/pkg/ocp-slide-api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/caarlos0/env/v6"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const serviceName = "ocp-slide-api"

type config struct {
	Address   string `env:"ADDRESS" envDefault:"0.0.0.0:8001"`
	Database  string `env:"POSTGRES_DB,unset,notEmpty"`
	User      string `env:"POSTGRES_USER,unset,notEmpty"`
	Password  string `env:"POSTGRES_PASSWORD,unset,notEmpty"`
	Host      string `env:"POSTGRES_HOST,unset,notEmpty"`
	Port      int    `env:"POSTGRES_PORT,unset,notEmpty"`
	ChunkSize uint   `env:"CHUNK_SIZE" envDefault:"256"`
	Debug     bool   `env:"DEBUG" envDefault:"false"`
}

func runGRPC(cfg *config) error {
	// sslmode=verify-full
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable&connect_timeout=10",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database)
	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Error().Msgf("Failed to connect to the database: %v", err)
		return err
	}
	log.Info().Msgf("A connection was successfully established with the database")

	repo := repo.NewRepo(db)

	server := grpc.NewServer()
	reflection.Register(server)
	desc.RegisterSlideAPIServer(server, api.NewSlideAPI(repo, cfg.ChunkSize))

	listen, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Error().Msgf("Failed to listen: %v", err)
		return err
	}
	log.Info().Msgf("The server listening on the %s gRPC server endpoint", cfg.Address)

	return server.Serve(listen)
}

func main() {
	log.Logger = zerolog.New(os.Stdout).With().
		Timestamp().
		Str("role", serviceName).
		Caller().
		Logger()
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal().Err(err).Msg("Read failed configuration parameters")
	}

	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if err := runGRPC(&cfg); err != nil {
		log.Fatal().Err(err).Msg("Run failed the gRPC server")
	}
}
