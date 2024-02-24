package router

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/tonychill/ifitu/apis/pb/go/global"
	"github.com/tonychill/ifitu/lib/utils"
)

func getAllContentFromCtx(c *fiber.Ctx, mimeType string) (all []*global.Content, err error) {
	mpfm, _err := c.MultipartForm()
	if _err != nil {
		err = errors.Join(err, _err)
		log.Error().Err(err).Msg("error getting multipart form")
		return nil, err
	}

	for _, mpfHeaders := range mpfm.File {
		for _, mpfh := range mpfHeaders {
			var content *global.Content
			content, _err = getContentFromCtx(c, mpfh, "image")
			if _err != nil {
				log.Error().Err(_err).Msg("error getting content from context while looping through multipart form")
				err = errors.Join(err, _err)
				// log.Error().Err(err).Msg("image upload error")
				// return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

			}
			all = append(all, content)
		}

	}

	return all, err
}

func getContentFromCtx(c *fiber.Ctx, mpfh *multipart.FileHeader, mimeType string) (cnt *global.Content, err error) {
	var (
		mpf multipart.File
		f   *os.File
	)

	mpf, err = mpfh.Open()
	if err != nil {
		return nil, err
	}

	defer func() {
		e := mpf.Close()
		if err == nil {
			err = e
		}
		if err != nil {
			log.Error().Err(err).Msgf("error closing file name %s while defering", f.Name())
		}
	}()
	// generate new uuid for image name

	fileExt := strings.Split(filepath.Ext(mpfh.Filename), ".")[1]
	mimeType = fmt.Sprintf("image/%s", fileExt)
	// generate image from filename and extension

	// s, err := f.Stat()
	// if err != nil {
	// 	log.Error().Err(err).Msg("error getting file stats")
	// }

	// log.Debug().Msgf("file size: %d", s.Size())
	// log.Debug().Msgf("file name: %s", s.Name())
	// log.Debug().Msgf("file mode string: %s", s.Mode().String())

	var ok bool
	if f, ok = mpf.(*os.File); ok {
		// Windows can't rename files that are opened.
		// if err = mpf.Close(); err != nil {
		// 	return nil, err
		// }
		// log.Debug().Msgf("f name: %s", f.Name())
		// If renaming fails we try the normal copying method.
		// Renaming could fail if the files are on diferent devices.
		// if err := os.Rename(f.Name(), newName); err != nil {
		// 	log.Error().Err(err).Msg("error renaming file")
		// 	return nil, fmt.Errorf("rerenaming from %s to %s failed: %v",
		// 		f.Name(), newName, err)
		// }

		// Reopen f for the code below.
		// if mpf, err = mpfh.Open(); err != nil {
		// 	log.Error().Err(err).Msgf("error opening file name %s after renaming", f.Name())
		// 	return nil, err
		// }
	} else if false {
		// if err = mpf.Close(); err != nil {
		// 	return nil, err
		// }
		// newMpf, err := mpfh.Open()
		// if err != nil {
		// 	log.Error().Err(err).Msg("error opening file name %s after renaming")
		// }
		// if f, ok = newMpf.(*os.File); !ok {
		// 	log.Error().Err(err).Msg("error casting multipart file to os file")
		// }
		// fmt.Printf("********** mpf file name: %+v", newMpf)

		// abs, err := filepath.Abs(mpfh.Filename)
		// if err != nil {
		// 	log.Error().Err(err).Msg("error getting absolute path")
		// }
		// log.Debug().Msgf("*** absol;ute path: %s", abs)
		// if err := os.Rename(mpfh.Filename, newName); err != nil {
		// 	log.Error().Err(err).Msg("error renaming file")
		// 	return nil, fmt.Errorf("rerenaming from %s to %s failed: %v",
		// 		mpfh.Filename, newName, err)
		// }
		// Reopen f for the code below.
		// if mpf, err = mpfh.Open(); err != nil {
		// 	log.Error().Err(err).Msgf("error opening file name %s after renaming", f.Name())
		// 	return nil, err
		// }
	}

	// if f, err = os.Create(newName); err != nil {
	// 	log.Error().Err(err).Msg("error creating file while getting content from context")
	// 	return nil, err
	// }

	// defer func() {
	// 	e := f.Close()
	// 	if err == nil {
	// 		err = e
	// 	}
	// 	if err != nil {
	// 		log.Error().Err(err).Msgf("error closing f file name %s while defering", f.Name())
	// 	}
	// }()

	// // save image to ./images dir
	// if err = c.SaveFile(mpfh, fmt.Sprintf("./temp_images/%s", imageName)); err != nil {
	// 	log.Error().Err(err).Msg("image save error")
	// 	// return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	// 	return nil, err
	// }

	// generate image url to serve to client using CDN

	// imageUrl := fmt.Sprintf("http://localhost:4000/temp_images/%s", imageName)

	// create meta data and send to client

	// data := map[string]interface{}{

	// 	"imageName": imageName,
	// 	"imageUrl":  imageUrl,
	// 	"header":    file.Header,
	// 	"size":      file.Size,
	// }
	var (
		creatorId   string
		creatorType = global.CreatorType_CREATOR_TYPE_PARTNER
	)

	if creatorId = c.Get("x-partner-id"); creatorId == "" {
		creatorId = c.Get("x-guest-id")
		creatorType = global.CreatorType_CREATOR_TYPE_GUEST
	}

	return &global.Content{
		Metadata: map[string]string{"ip_address": c.IP()},
		// TODO: get name & id from client & ctx
		CreatorId:   creatorId,
		Name:        mpfh.Filename,
		CreatorType: creatorType,
		MimeType:    mpfh.Header.Get("Content-Type"),
		// TODO: create helper func to analyze mime type and return content type
		ContentType: global.ContentType_CONTENT_TYPE_IMAGE,
		Size:        mpfh.Size,
		Data:        utils.StreamToByte(mpf),
	}, nil

}
