package sessionUtils

import (
	"github.com/aaronchen2k/openstc/internal/server/biz/domain"
	"github.com/aaronchen2k/openstc/internal/server/cfg"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	SessionID = "OpenSTC_SessionID"
	CredKey   = "OpenSTC_Credentials"
	session   = sessions.New(sessions.Config{Cookie: SessionID})
)

func GetCredentials(ctx iris.Context) (cred *domain.UserCredentials) {
	if serverConf.Config.Redis.Enable {
		credObj := ctx.Values().Get("sess")
		if credObj == nil {
			return
		}
		cred = credObj.(*domain.UserCredentials)
	} else {
		sess := session.Start(ctx)
		credObj := sess.Get(CredKey)
		if credObj == nil {
			return
		}

		cred = credObj.(*domain.UserCredentials)
	}

	return
}

func SaveCredentials(ctx iris.Context, cred *domain.UserCredentials) {
	sess := session.Start(ctx)
	sess.Set(CredKey, cred)
}
func RemoveCredentials(ctx iris.Context) {
	sess := session.Start(ctx)
	sess.Delete(CredKey)
}
