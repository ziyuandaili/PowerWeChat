package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseOperationGetJsErrList struct {
	*response.ResponseMiniProgram
	Success    bool           `json:"success"`
	OpenID     string         `json:"openid"`
	Data       *power.HashMap `json:"data"`
	TotalCount int64          `json:"totalCount"`
}
