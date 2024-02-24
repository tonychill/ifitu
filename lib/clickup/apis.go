package clickup

import (
	"context"

	"github.com/tonychill/ifitu/apis/pb/go/global"
)

func (c *clientImpl) DeleteJourney(ctx context.Context, id string) error {
	panic("implement me")
}

func (c *clientImpl) AddExperienceToJourney(ctx context.Context, journeyId string, experience *global.Experience) error {
	panic("implement me")
}

func (c *clientImpl) UpdateExperience(ctx context.Context, journeyId string, experience *global.Experience) error {
	panic("implement me")
}

func (c *clientImpl) RemoveExperienceFromJourney(ctx context.Context, journeyId, experienceId string) error {
	panic("implement me")
}

func (c *clientImpl) UpdatePaymentPaymentStatus(ctx context.Context, journeyId, experienceId string, status PaymentStatus) error {
	panic("implement me")
}
