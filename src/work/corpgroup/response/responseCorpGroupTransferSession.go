package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCorpGroupTransferSession struct {
	*response.ResponseWork
	Userid     string `json:"userid"`
	SessionKey string `json:"session_key"`
}
