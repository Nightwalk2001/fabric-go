package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Uri       string
	RedisAddr string
	RedisPass string
}

func Load() Config {
	_ = godotenv.Load()
	uri := os.Getenv("URI")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPass := os.Getenv("REDIS_PASS")

	return Config{
		Uri:       uri,
		RedisAddr: redisAddr,
		RedisPass: redisPass,
	}
}

const Mail string = "wangwei@nibs.ac.cn"
