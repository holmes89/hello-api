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
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Database{
		conn: rdb,
	}
}

func (s *Database) Translate(word string, language string) string {
	return s.conn.Get(context.Background(), fmt.Sprintf("%s:%s", word, language)).Val()
}
