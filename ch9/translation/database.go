package translation

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/holmes89/hello-api/config"
	"github.com/holmes89/hello-api/handlers/rest"
)

var _ rest.Translator = &Database{}

type Database struct {
	conn *redis.Client
}

func NewDatabaseService(cfg config.Configuration) *Database {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DatabaseURL, cfg.DatabasePort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Database{
		conn: rdb,
	}
}

func (s *Database) Translate(word string, language string) string {
	out := s.conn.Get(context.Background(), fmt.Sprintf("%s:%s", word, language))
	if out.Err() != nil {
		panic(out.Err())
	}
	return out.Val()
}
