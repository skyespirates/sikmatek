package entity

type LimitUsage struct{}

type CreateLimitUsagePayload struct {
	UsedAmount    int
	InstallmentId int
	LimitId       int
}
