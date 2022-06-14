package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/cucumber/godog"
	resty "github.com/go-resty/resty/v2"
	"github.com/holmes89/hello-api/config"
	"github.com/holmes89/hello-api/handlers/rest"
)

type apiFeature struct {
	client   *resty.Client
	server   *httptest.Server
	word     string
	language string
}

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
	// if res.Language != api.language {
	// 	return fmt.Errorf("language should be set to %s but instead was %+v", api.language, res)
	// }

	if res.Translation != arg1 {
		return fmt.Errorf("translation should be set to %s", arg1)
	}

	return nil
}

func (api *apiFeature) theWord(arg1 string) error {
	api.word = arg1
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {

	client := resty.New()
	cfg := config.Configuration{}
	cfg.LoadFromEnv()

	mux := API(cfg)
	server := httptest.NewServer(mux)
	api := &apiFeature{
		client: client,
		server: server,
	}

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
