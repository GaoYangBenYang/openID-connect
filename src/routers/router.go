// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"openid_connect/internal/controllers"
)

func init() {

	
	beego.Router("/.well-known/openid-configuration",&controllers.Authorize{},"Get:Authorize")
	//Authorize api
	beego.Router("/op/authorize",&controllers.Authorize{},"Get:Authorize")
	//认证请求响应
	beego.Router("/op/callback",&controllers.Authorize{},"Get:CallBack")
	//Token EndPoint
	// beego.Router("/op/token",&controllers.Token{},"Post:token")
	//userInfo api
	//beego.Router("/op/userInfo",&controllers.Provider{},"Post:userInfo")
}
