package translation

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	_ "embed"

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
		Password: cfg.DatabasePassword,
		DB:       0, // use default DB
	})
	return &Database{
		conn: rdb,
	}
}

func (s *Database) Translate(word string, language string) string {
	out := s.conn.Get(context.Background(), fmt.Sprintf("%s:%s", word, language))
	return out.Val()
}

//go:embed translations.json
var migrations []byte

func (s *Database) LoadData() error {
	type translation struct {
		Language string `json:"language"`
		Hello    string `json:"hello"`
	}
	var result []translation
	if err := json.Unmarshal(migrations, &result); err != nil {
		return err
	}
	ctx := context.Background()
	for _, t := range result {
		err := s.conn.Set(ctx, fmt.Sprintf("hello:%s", strings.ToLower(t.Language)), t.Hello, 0).Err()
		if err != nil {
			panic(err)
		}
	}
	return nil
}
