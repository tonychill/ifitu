package clickup

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func NewClient(ctx context.Context) (c Client, err error) {
	ci := &clientImpl{
		repo:    &repo{},
		baseUrl: "https://api.clickup.com/api/v2/",
		app: fiber.New(fiber.Config{
			// Prefork:       true,
			// CaseSensitive: true,
			// StrictRouting: true,
			// BodyLimit:    100 * MB, // 100MB
			ServerHeader: "juvae/clickup-client",
			AppName:      "clickup_client_v1.0.0",
		}),
	}
	if err = ci.repo.initializeRepo(ctx); err != nil {
		return
	}
	return ci, err
}

func (cl *clientImpl) makePostRequest(req interface{}) (err error) {
	// TODO: need to implement
	log.Debug().Msgf("*** TESTING_CLICKUP: making dummy post request: %+v", req)
	return nil
	listId := "901700234454"
	taskId := "something"
	reqUrl := "https://api.clickup.com/api/v2/list/" + listId + "/task/" + taskId
	a := fiber.Post(reqUrl)
	defer fiber.ReleaseAgent(a)
	reqBody, err := json.Marshal(req)
	if err != nil {
		log.Error().Err(err).Msg("error marshalling request")
		return err
	}

	a = a.Body(reqBody)
	a = a.Set("Authorization", os.Getenv("CLICKUP_API_KEY"))
	code, body, errs := a.Bytes()
	for _, _err := range errs {
		err = errors.Join(err, _err)
		log.Error().Msgf("error making post request: %v", errs)
	}
	if err != nil {
		log.Error().Err(err).Msg("got some errors after making post request")
		return err
	}

	fmt.Printf("respose from fiber: %d\n%s\n%+v", code, string(body), err)
	return nil
}
