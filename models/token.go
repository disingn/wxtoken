package models

type SetToken struct {
	Name           string `json:"name"`
	RemainQuota    int    `json:"remain_quota"`
	ExpiredTime    int64  `json:"expired_time"`
	UnlimitedQuota bool   `json:"unlimited_quota"`
}

type GetToken struct {
	Data []struct {
		Id             int    `json:"id"`
		UserId         int    `json:"user_id"`
		Key            string `json:"key"`
		Status         int    `json:"status"`
		Name           string `json:"name"`
		CreatedTime    int    `json:"created_time"`
		AccessedTime   int    `json:"accessed_time"`
		ExpiredTime    int    `json:"expired_time"`
		RemainQuota    int    `json:"remain_quota"`
		UnlimitedQuota bool   `json:"unlimited_quota"`
		UsedQuota      int    `json:"used_quota"`
	} `json:"data"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}
