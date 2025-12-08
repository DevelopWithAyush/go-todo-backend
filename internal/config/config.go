package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	Env        string
	MongoURI   string
	MongoDB    string
	RedisURI   string
	RedisPass  string
	RedisDB    int
	JWTSecret  string
	CookieName string
	SMTPHost   string
	SMTPPort   string
	SMTPUser   string
	SMTPPass   string
}

func Load() *Config {
	return &Config{
		Port:       get("PORT", "5000"),
		Env:        get("ENV", "development"),
		MongoURI:   get("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:    get("MONGO_DB", "todo_app"),
		RedisURI:   get("REDIS_URI", "localhost:6379"),
		RedisPass:  get("REDIS_PASS", ""),
		RedisDB:    getInt("REDIS_DB", 0),
		JWTSecret:  get("JWT_SECRET", "secret"),
		CookieName: get("COOKIE_NAME", "todo_app"),
		SMTPHost:   get("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:   get("SMTP_PORT", "587"),
		SMTPUser:   get("SMTP_USER", ""),
		SMTPPass:   get("SMTP_PASS", ""),
	}
}

func get(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

func getInt(key string, def int) int {
	if v, ok := os.LookupEnv(key); ok {
		var n int
		_, err := fmt.Sscanf(v, "%d", &n)
		if err == nil {
			return n
		}
	}
	return def
}
