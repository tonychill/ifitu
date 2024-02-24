package service

import (
	"context"
	"testing"
	// "github.com/tonychill/ifitu/apis/pb/go/finance"
	// "github.com/tonychill/ifitu/apis/pb/go/global"
)

func TestInitiateCreateExperienceFlow(t *testing.T) {
	// ctx := context.Background()
	// svc, err := InitializeConciergeService(ctx)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	t.Run("bad - the request was nil", func(t *testing.T) {
		// ctx := context.Background()
		// expectedErr := fmt.Errorf("the request was nil")
		// _, err := svc.InitiateFlow(ctx, nil)
		// assert.Equal(t, expectedErr, err)

	})
	t.Run("bad - the experience flow was nil", func(t *testing.T) {
		// ctx := context.Background()

		// _, err := svc.InitiateFlow(ctx, nil)
		// assert.Equal(t, globals.ErrRequestNil, err)

	})

}
func TestCreateExperienceFlow(t *testing.T) {
	ctx := context.Background()
	_, err := InitializeConciergeService(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// t.Run("bad - the request was nil", func(t *testing.T) {
	// 	ctx := context.Background()
	// 	expectedErr := fmt.Errorf("the request was nil")
	// 	_, err := svc.InitiateFlow(ctx, nil)
	// 	assert.Equal(t, expectedErr, err)

	// })

	// t.Run("bad - the flow in the request was nil", func(t *testing.T) {
	// 	ctx := context.Background()
	// 	expectedErr := fmt.Errorf("the request was nil")
	// 	_, err := svc.InitiateFlow(ctx, &conSvc.InitiateFlowRequest{})
	// 	assert.Equal(t, expectedErr, err)

	// })

	// t.Run("bad - the create experience flow in the request was nil", func(t *testing.T) {
	// 	ctx := context.Background()
	// 	expectedErr := fmt.Errorf("the request was nil")
	// 	_, err := svc.InitiateFlow(ctx, &conSvc.InitiateFlowRequest{
	// 		Flow: &conSvc.Flow{
	// 			Data: &conSvc.Flow_CreateExperienceFlow{
	// 				CreateExperienceFlow: nil,
	// 			},
	// 		},
	// 	})
	// 	assert.Equal(t, expectedErr, err)

	// })

	// t.Run("bad - no experience data provied in the create experience flow request", func(t *testing.T) {
	// 	ctx := context.Background()
	// 	expectedErr := fmt.Errorf("no experience data provied in the create experience flow request")
	// 	_, err := svc.InitiateFlow(ctx, &conSvc.InitiateFlowRequest{
	// 		Flow: &conSvc.Flow{
	// 			Data: &conSvc.Flow_CreateExperienceFlow{
	// 				CreateExperienceFlow: &conSvc.CreateExperienceFlow{
	// 					Experience: nil,
	// 				},
	// 			},
	// 		},
	// 	})
	// 	assert.Equal(t, expectedErr, err)

	// })

	t.Run("bad - no resources were provided in the request", func(t *testing.T) {
		// TODO: mod this test to handle requests with no resources; already created.
		// ctx := context.Background()
		// expectedErr := fmt.Errorf("no resources were provided in the request")
		// _, err := svc.InitiateFlow(ctx, &conSvc.InitiateFlowRequest{
		// 	Flow: &conSvc.Flow{
		// 		Kind: &conSvc.Flow_CreateExperienceFlow{
		// 			CreateExperienceFlow: &conSvc.CreateExperienceFlow{
		// 				Experience: nil,
		// 			},
		// 		},
		// 	},
		// })
		// assert.Equal(t, expectedErr, err)

	})

	// t.Run("good - experience and resources created", func(t *testing.T) {
	// 	ctx := context.Background()
	// 	expectedErr := fmt.Errorf("no experience data provied in the create experience flow request")
	// 	resp, err := svc.InitiateFlow(ctx, &conSvc.InitiateFlowRequest{
	// 		Flow: &conSvc.Flow{
	// 			Data: &conSvc.Flow_CreateExperienceFlow{
	// 				CreateExperienceFlow: &conSvc.CreateExperienceFlow{
	// 					Experience: &global.Experience{
	// 						Name:        "test name",
	// 						Description: "test description",
	// 						Media: []*global.Content{
	// 							{
	// 								Name:        "test media name",
	// 								Description: "test media description",
	// 								Url:         "www.example.com/test-media",
	// 								MimeType:    "dono",
	// 								Size:        0,
	// 								Metadata:    map[string]string{"test_key": "test_value"},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	})

	// 	assert.Equal(t, expectedErr, err)
	// 	assert.NotNil(t, resp)
	// 	assert.Equal(t, "", "") // flow id
	// 	// flowStatus := resp.GetFlowStatus()
	// 	// assert.Equal(t, 1, len(flowStatus.ResourceIds)) // Get the resource id from the flow metadata

	// })

	/*
		- Client will make a request to the coordinator containing the resource and experience info in the payload.
		- the coordinator will return to the client a flow id and the status of the flow
		- the operations service will validate and persist the resource data sent by the client
		- the journey service will save the experinece data such as pricing and availability (scheduler mod will handle availability)
		- the coordinator, via its content module, will save the images associated with the experience/resource(s)
		- in the event that the job was not completed the client will be able to use the requeest id to check the sstatus of the job.
		- The client should be able to query the coordinator for jobs.
	*/

}
