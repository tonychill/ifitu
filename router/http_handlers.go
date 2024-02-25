package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	// idSvc "github.com/tonychill/ifitu/apis/pb/go/identity_service"
	// "github.com/tonychill/ifitu/lib/oauth"
)

// FIXME: remove after authentication implementation and testing
// DO NOT PUSH TO PROD!!!
// func (r *routerImpl) handleGetSession(c *fiber.Ctx) error {
// 	val, err := oauth.GetValueFromSession(c, "google")
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	type _tempSesh struct {
// 		AuthURL      string
// 		AccessToken  string
// 		RefreshToken string
// 		ExpiresAt    string
// 		IDToken      string
// 	}
// 	cv := &_tempSesh{}
// 	if err := json.Unmarshal([]byte(val), cv); err != nil {
// 		log.Error().Err(err).Msg("error unmashalling value to Sesh")
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 		"session": cv,
// 	})

// }

// func (r *routerImpl) handleLogin(c *fiber.Ctx) error {
// 	url, err := oauth.BeginAuthHandler(oauth.HttpReqResp{
// 		ReqCtx: c,
// 	})
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}
// 	return c.Redirect(url, fiber.StatusTemporaryRedirect)
// }

// func (r *routerImpl) handleAuthCallback(c *fiber.Ctx) error {
// 	guest, err := oauth.CompleteGuestAuth(context.Background(), oauth.CompleteGuestAuthRequest{
// 		ReqCtx: c,
// 	})
// 	if err != nil {
// 		log.Error().Err(err).Msg("error handling authentication callback")
// 		return c.Status(fiber.StatusPreconditionFailed).JSON(fiber.Map{
// 			"error": fmt.Sprintf("Error handling authentication callback for provier... %s",
// 				err.Error()),
// 		})
// 	}
// 	// val, err := oauth.GetFromSession("google", c)
// 	// if err != nil {
// 	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 	// 		"error": err.Error(),
// 	// 	})
// 	// }

// 	// cv := &oauth.Sesh{}
// 	// if err := json.Unmarshal([]byte(val), cv); err != nil {
// 	// 	log.Error().Err(err).Msg("error unmashalling value to Sesh")
// 	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 	// 		"error": err.Error(),
// 	// 	})
// 	// }
// 	// cv.AccessToken = guest.AccessToken
// 	// cv.ExpiresAt = guest.ExpiresAt.String()
// 	// cv.IDToken = guest.IDToken
// 	// b, err := json.Marshal(cv)
// 	// if err != nil {
// 	// 	log.Error().Err(err).Msg("error mashaling json")
// 	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 	// 		"error": err.Error(),
// 	// 	})
// 	// }

// 	// if err := oauth.StoreInSession("google", string(b), c); err != nil {
// 	// 	log.Error().Err(err).Msg("error adding value to session")
// 	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 	// 		"error": err.Error(),
// 	// 	})
// 	// }

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"response":    guest,
// 		"raw_headers": string(c.Request().Header.RawHeaders()),
// 	})
// }
// func (r *routerImpl) handleLogout(c *fiber.Ctx) error {
// 	printCookies(c, "handleLogout")
// 	if err := oauth.Logout(oauth.LogoutRequest{
// 		ReqCtx: c,
// 	}); err != nil {
// 		log.Error().Err(err).Msg("error logging out")
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": fmt.Sprintf("Error logging out %s",
// 				err.Error()),
// 		})
// 	}
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message": "Adios!",
// 	})
// }
// func (r *routerImpl) handleCreateGuest(c *fiber.Ctx) error {
// 	ctx, reqId := setRequestId(c, nil)
// 	req := &idSvc.CreateGuestRequest{}
// 	if err := utils.DecodeFiberRequest(c, req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"request_id": reqId,
// 			"message":    fmt.Sprintf("Error decoding client's request: %s", err.Error()),
// 		})
// 	}

// 	resp, err := r.conImpl.CreateGuest(ctx, req)
// 	if err != nil {
// 		log.Error().Msgf("error creating and guest via the concierge: %v", err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"request_id": reqId,
// 			"message":    fmt.Sprintf("Error creating guest: %s", err.Error()),
// 		})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"request_id": reqId,
// 		"response":   resp,
// 	})
// }

// Pulls the content from the context and adds it to the request. The service is then responsible for
// uploading the content to cloudflar
// TODO: deprecate: uploading handled separately. I might bring this back into play later but for now
// I am going to leave it out.
func (r *routerImpl) addContentToRequest(c *fiber.Ctx, req any) error {

	ctnt, err := getAllContentFromCtx(c, "image")
	if err != nil {
		log.Error().Err(err).Msg("error getting all content from context when adding content to request")
		// log.Error().Err(err).Msg("image upload error")
		// return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	fmt.Printf("*** TESTING: content from upload: %+v\n", ctnt)

	// Upload contnet to cloa
	// for _, o := range req.Experience.Options {
	// 	for _, m := range o.Resource.Media {
	// 		for _, c := range content {
	// 			if c.Id == m.Id {
	// 				m = c
	// 			}
	// 		}
	// 	}
	// }
	return nil
}

// func (r *routerImpl) handleUpdateResource(c *fiber.Ctx) error {
// 	req := &pb.UpdateResourceRequest{}
// 	if err := c.BodyParser(req); err != nil {
// 		return err
// 	}
// 	ctx := context.Background()
// 	_, err := r.service.UpdateResource(ctx, req)
// 	if err != nil {
// 		// do something
// 		log.Error().Msgf("error updating a resource via the concierge: %v", err)
// 		return err
// 	}
// 	return c.SendString("Add experiences to resource!")
// }

// func (r *routerImpl) handleUploadContent(fc *fiber.Ctx) error {
// 	content, err := getAllContentFromCtx(fc, "image")
// 	if err != nil {
// 		log.Error().Err(err).Msg("error getting all content from context while handling the call to upload content")
// 		return fc.JSON(fiber.Map{
// 			"status":  500,
// 			"message": "Server error",
// 		})
// 	}

// 	if len(content) == 0 {
// 		log.Error().Msg("no content found in context")
// 		return fc.JSON(fiber.Map{
// 			"status":  400,
// 			"message": "No content was provided by the client",
// 		})
// 	}

// 	enrichedContent := make([]*global.Content, 0, len(content))
// 	for _, c := range content {
// 		resp, err := r.conImpl.UploadContent(fc.Context(), c)
// 		if err != nil {
// 			log.Error().Err(err).Msg("error uploading content")
// 			return fmt.Errorf("error uploading content: %w", err)
// 		}

// 		enrichedContent = append(enrichedContent, &resp)
// 	}
// 	return fc.JSON(enrichedContent)
// }

// func (r *routerImpl) uploadContent(ctx context.Context, content []*global.Content) ([]*global.Content, error) {
// 	var err error
// 	for _, c := range content {
// 		*c, err = r.conImpl.UploadContent(ctx, c)
// 		if err != nil {
// 			log.Error().Err(err).Msg("error uploading content")
// 			// TODO: use errors.Join
// 			return nil, fmt.Errorf("error uploading content: %w", err)
// 		}
// 	}

// 	return content, nil
// }

func printCookies(c *fiber.Ctx, method string) {
	fc := c.Context()
	valStr := ""
	fc.Request.Header.VisitAllCookie(func(key, b []byte) {
		valStr += string(key) + ":" + string(b) + "|"
	})
	fmt.Printf("*** TESTING: method: %s | cookie values: %s\n", method, valStr)
}
