package service

import (
	"context"

	"github.com/rs/zerolog/log"
	// idSvc "github.com/tonychill/ifitu/apis/pb/go/identity_service
	// contSvc "github.com/tonychill/ifitu/apis/pb/go/content_service"
	// idSvc "github.com/tonychill/ifitu/apis/pb/go/identity_service"
)

// func (s *ServiceImpl) CreateGuest(ctx context.Context, req *idSvc.CreateGuestRequest) (*idSvc.CreateGuestResponse, error) {
// 	resp, err := s.idClient.CreateGuest(ctx, req)
// 	if err != nil {
// 		log.Error().Err(err).Msgf("failed to create guest id %s", req.Guest.Id)
// 	}

// 	// TODO: implement some retries here.
// 	if err = s.repo.CreateGuest(ctx, req.Guest); err != nil {
// 		log.Error().Err(err).Msgf("failed to add guest id %s to the graph", req.Guest.Id)
// 	}

// 	return resp, nil
// }

// Calls the content service to upload content. Future implementations will allow for this
// method to send the content in chunks.
// func (s *ServiceImpl) UploadContent(ctx context.Context, content *global.Content) (global.Content, error) {
// 	stream, err := s.contentClient.UploadContent(ctx)
// 	if err != nil {
// 		log.Error().Err(err).Msg("Error creating upload stream")
// 		return global.Content{}, status.Errorf(codes.Internal, "Error creating upload stream: %v", err)
// 	}

// 	// TODO: pass the stream connection to a chunker that will send the content in chunks
// 	// as we're currently running into the 4MB limit of gRPC

// 	_ /*final*/, err = s.chunkContent(ctx, stream, content)
// 	if err != nil {
// 		log.Error().Err(err).Msg("Error chunking content")
// 		return global.Content{}, status.Errorf(codes.Internal, "Error chunking content: %v", err)
// 	}

// 	if err := stream.Send(&contSvc.UploadContentRequest{
// 		Name:        content.Name,
// 		Description: content.Description,
// 		CreatorType: content.CreatorType,
// 		CreatorId:   content.CreatorId,
// 		Metadata:    content.Metadata,
// 		ContentType: content.ContentType,
// 		MimeType:    content.MimeType,
// 		Data:        content.Data,
// 		// Size:        content.Size,

// 	}); err != nil {
// 		log.Error().Err(err).Msg("Error sending upload chunk to content service")
// 		return global.Content{}, status.Errorf(codes.Internal, "Error sending chunk to content service: %v", err)
// 	}

// 	resp, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Error().Err(err).Msg("Error closing upload stream to content service")
// 		return global.Content{}, status.Errorf(codes.Internal, "Error closing upload stream to content service: %v", err)
// 	}

// 	log.Debug().Msgf("content uploaded with url: %s", resp.ContentUrl)

// 	return enrichContentWithResp(resp, content), nil
// 	// for {

// 	// 	if err == io.EOF {
// 	// 		if err := stream.CloseSend(); err != nil {
// 	// 			log.Error().Err(err).Msg("Error closing upload stream")
// 	// 			return status.Errorf(codes.Internal, "Error closing upload stream: %v", err)
// 	// 		}
// 	// 		log.Info().Msg("End of stream")
// 	// 		break
// 	// 	}
// 	// 	if err != nil {
// 	// 		if err := stream.CloseSend(); err != nil {
// 	// 			log.Error().Err(err).Msg("Error closing upload stream to content service")
// 	// 			return status.Errorf(codes.Internal, "Error closing upload stream to content service after non-EOF error: %v", err)
// 	// 		}
// 	// 		log.Error().Err(err).Msg("received error from upload stream when recieving from the caller; non-EOF")
// 	// 		return status.Errorf(codes.Internal, "Error receiving upload chunk from caller; non-EOF: %v", err)
// 	// 	}
// 	// 	if err := stream.Send(req); err != nil {
// 	// 		log.Error().Err(err).Msg("Error sending upload chunk to content service")
// 	// 		return status.Errorf(codes.Internal, "Error sending chunk to content service: %v", err)
// 	// 	}

// 	// }
// 	// resp, err := stream.CloseAndRecv()
// 	// if err != nil {
// 	// 	log.Error().Err(err).Msg("Error closing upload stream to content service")
// 	// 	return status.Errorf(codes.Internal, "Error closing upload stream to content service: %v", err)
// 	// }
// 	// return svr.SendAndClose(resp)
// }

// func enrichContentWithResp(resp *contSvc.UploadContentResponse, content *global.Content) global.Content {
// 	return global.Content{
// 		Id:          resp.ContentId,
// 		Metadata:    resp.Metadata,
// 		ContentType: resp.ContentType,
// 		MimeType:    resp.MimeType,
// 		Name:        resp.Name,
// 		Description: content.Description,
// 		Url:         resp.ContentUrl,
// 		Size:        content.Size,
// 		CreatorId:   resp.CreatorId,
// 		CreatorType: resp.CreatorType,
// 		CreatedAt:   content.CreatedAt,
// 	}

// }

// TODO: Not yet implemented
// func (s *ServiceImpl) chunkContent(ctx context.Context, stream contSvc.ContentService_UploadContentClient, content *global.Content) (global.Content, error) {

// 	return global.Content{}, nil
// 	// Calculate the number of chunks based on the content size and desired chunk size (4MB)
// 	// Ref: https://jbrandhorst.com/post/grpc-binary-blob-stream/
// 	chunkSize := 4 * 1 << 20 // 4MB
// 	numChunks := int(math.Ceil(float64(content.Size) / float64(chunkSize)))

// 	// Create a buffer to hold each chunk
// 	buffer := make([]byte, chunkSize)

// 	// Read and send each chunk separately
// 	req := &contSvc.UploadContentRequest{
// 		Name:         content.Name,
// 		Description:  content.Description,
// 		CreatorType:  content.CreatorType,
// 		CreatorId:    content.CreatorId,
// 		Metadata:     content.Metadata,
// 		ContentType:  content.ContentType,
// 		MimeType:     content.MimeType,
// 		InitialChunk: true,
// 		// Size:        content.Size,

// 	}

// 	for i := 0; i < numChunks; i++ {
// 		// Calculate the start and end positions of the current chunk
// 		if i == numChunks-1 {
// 			req.FinalChunk = true
// 		}
// 		start := i * chunkSize
// 		end := int(math.Min(float64(start+chunkSize), float64(content.Size)))

// 		// Read the chunk from the content
// 		// _, err := content.Data.ReadAt(buffer, int64(start))
// 		// if err != nil && err != io.EOF {
// 		// 	return err
// 		// }
// 		req.Data = buffer[:end-start]

// 		// Send the chunk through the stream
// 		if err := stream.Send(req); err != nil {
// 			log.Error().Err(err).Msg("Error sending upload chunk to content service")
// 			return global.Content{}, status.Errorf(codes.Internal, "Error sending chunk to content service: %v", err)
// 		}
// 		// TODO: deprecate. the first chunk is always sent with the initial request; redundant
// 		req.InitialChunk = false
// 	}

// 	resp, err := stream.CloseAndRecv()
// 	if err != nil {
// 		log.Error().Err(err).Msg("Error closing upload stream to content service")
// 		return global.Content{}, status.Errorf(codes.Internal, "Error closing upload stream to content service: %v", err)
// 	}

// 	log.Debug().Msgf("content uploaded with url when chunking: %s", resp.ContentUrl)

// 	return enrichContentWithResp(resp, content), nil

// }

func (s *ServiceImpl) Shutdown(ctx context.Context) error {
	log.Info().Msg("shutting down concierge service...")
	close(s.shutdownCh)
	return nil
}
