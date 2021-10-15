package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseCalendarGet struct {
	*response.ResponseWork

	CalendarList []*power.HashMap `json:"calendar_list"`

}