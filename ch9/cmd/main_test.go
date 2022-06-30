package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"

	"github.com/cucumber/godog"
	resty "github.com/go-resty/resty/v2"
	"github.com/holmes89/hello-api/config"
	"github.com/holmes89/hello-api/handlers/rest"
	"github.com/ory/dockertest"
)

type apiFeature struct {
	client   *resty.Client
	server   *httptest.Server
	word     string
	language string
}

var (
	pool     *dockertest.Pool
	database *dockertest.Resource
)

func (api *apiFeature) iTranslateItTo(arg1 string) error {
	api.language = arg1
	return nil
}

func (api *apiFeature) theResponseShouldBe(arg1 string) error {
	url := fmt.Sprintf("%s/translate/%s", api.server.URL, api.word)

	resp, err := api.client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			"language": api.language,
		}).
		SetResult(&rest.Resp{}).
		Get(url)

	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusNotFound {
		return fmt.Errorf("unable to call api: %d %s", resp.StatusCode(), url)
	}

	res := resp.Result().(*rest.Resp)
	if res.Translation != arg1 {
		return fmt.Errorf("translation should be set to %s", arg1)
	}

	return nil
}

func (api *apiFeature) theWord(arg1 string) error {
	api.word = arg1
	return nil
}

func InitializeTestSuite(sc *godog.TestSuiteContext) {

	var err error

	sc.BeforeSuite(func() {
		pool, err = dockertest.NewPool("")
		if err != nil {
			panic(fmt.Sprintf("unable to create connection pool %s", err))
		}

		wd, err := os.Getwd()
		if err != nil {
			panic(fmt.Sprintf("unable to get working directory %s", err))
		}

		mount := fmt.Sprintf("%s/data/:/data/", filepath.Dir(wd))
		fmt.Println(mount)
		redis, err := pool.RunWithOptions(&dockertest.RunOptions{
			Repository: "redis",
			Mounts:     []string{mount},
		})
		if err != nil {
			panic(fmt.Sprintf("unable to create container: %s", err))
		}
		if err := redis.Expire(600); err != nil {
			panic("unable to set expiration on container")
		} //Destroy container if it takes too long
		database = redis
	})

	sc.AfterSuite(func() {
		database.Close()
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	client := resty.New()
	api := &apiFeature{
		client: client,
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		cfg := config.Configuration{}
		cfg.LoadFromEnv()
		cfg.DatabasePort = database.GetPort("6379/tcp")
		cfg.DatabaseURL = "localhost"
		fmt.Printf("%+v\n", cfg)
		mux := API(cfg)
		server := httptest.NewServer(mux)
		api.server = server
		return ctx, nil
	})
	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		api.server.Close()
		return ctx, nil
	})

	ctx.Step(`^I translate it to "([^"]*)"$`, api.iTranslateItTo)
	ctx.Step(`^the response should be "([^"]*)"$`, api.theResponseShouldBe)
	ctx.Step(`^the word "([^"]*)"$`, api.theWord)
}

// var opt = godog.Options{Output: colors.Colored(os.Stdout), Format: "progress", Randomize: time.Now().UTC().UnixNano()}

// func TestMain(m *testing.M) {
// 	pflag.Parse()
// 	opt.Paths = pflag.Args()
// 	status := godog.TestSuite{Name: "godogs", ScenarioInitializer: InitializeScenario, Options: &opt}.Run()
// 	if st := m.Run(); st > status {
// 		status = st
// 	}

// 	os.Exit(status)
// }
