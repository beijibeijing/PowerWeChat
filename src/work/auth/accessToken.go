package auth

// #reference: https://open.work.weixin.qq.com/api/doc/90000/90135/91039

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

type AccessToken struct {
	*kernel.AccessToken
}

func NewAccessToken(app *kernel.ApplicationInterface) *AccessToken {
	token := &AccessToken{
		kernel.NewAccessToken(app),
	}

	// Override fields and functions
	token.EndpointToGetToken = "cgi-bin/gettoken"
	token.OverrideGetCredentials()

	return token
}


// Override GetCredentials
func (component *AccessToken) OverrideGetCredentials() {
	config := (*component.App).GetContainer().GetConfig()
	component.GetCredentials = func() *object.StringMap {
		return &object.StringMap{
			"corpid":     (*config)["corp_id"].(string),
			"corpsecret": (*config)["secret"].(string),
		}
	}
}