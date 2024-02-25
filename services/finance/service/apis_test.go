package service

import (
	"context"
	"testing"
)

func TestSampleFunc(t *testing.T) {
	ctx := context.Background()
	_, err := InitializeFinanceService(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// t.Run("bad - sample test", func(t *testing.T) {
	// 	ctx := context.Background()

	// 	_, err := svc.CreateRates(ctx, nil)
	// 	assert.Equal(t, fmt.Errorf("the request was nil"), err)

	// })

	// TODO: add clean up func to remove the resource created in the test (good way to test delete too)

}
