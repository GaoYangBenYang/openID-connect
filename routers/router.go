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
	"openid_connect_op/controllers"
)

func init() {
	//Authorize api
	beego.Router("/op/authorize",&controllers.Authorize{},"Get:Authorize")
	//token api
	//beego.Router("/op/token",&controllers.Provider{},"Post:token")
	//userInfo api
	//beego.Router("/op/userInfo",&controllers.Provider{},"Post:userInfo")
}
