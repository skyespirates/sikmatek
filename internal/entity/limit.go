package entity

type Limit struct {
	Id         int     `json:"id"`
	Requested  int     `json:"requested_limit"`
	Status     string  `json:"status"`
	ApprovedBy *int    `json:"approved_by,omitempty"`
	ApprovedAt *string `json:"approved_at,omitempty"`
	ConsumerId int     `json:"consumer_id"`
}

type LimitDetail struct {
	Id         int     `json:"id"`
	Requested  int     `json:"requested_limit"`
	Used       int     `json:"used_limit"`
	Remaining  int     `json:"remaining_limit"`
	Status     string  `json:"status"`
	ApprovedBy *int    `json:"approved_by,omitempty"`
	ApprovedAt *string `json:"approved_at,omitempty"`
	ConsumerId int     `json:"consumer_id"`
}

type CreateLimitPayload struct {
	Requested  int `json:"requested_limit"`
	ConsumerId int
}

type UpdateLimitPayload struct {
	LimitId int    `json:"limit_id"`
	Action  string `json:"action"`
	AdminId int    `json:"admin_id"`
}

type LimitListPayload struct {
	RoleId     int
	ConsumerId int
}

type ApprovedLimit struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
