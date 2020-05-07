package env

import (
	"os"
)

func GetENV() string {
	return os.Getenv("ENV")
}

// GetPortEnv
func GetPortEnv() string {
	return os.Getenv("PORT")
}

// GetPGDataSourceEnv
func GetPGDataSourceEnv() string {
	return os.Getenv("PG_DATASOURCE")
}

func GetJWTSerectKeyEnv() string {
	return os.Getenv("JWT_SECRET_KEY")
}

func GetRedisAddr() string {
	return os.Getenv("REDIS_ADDR")
}

func GetGRPCPortEnv() string {
	return os.Getenv("GRPC_PORT")
}

func GetServerKey() string {
	return os.Getenv("SERVER_KEY")
}

func GetServerCRT() string {
	return os.Getenv("SERVER_CRT")
}
