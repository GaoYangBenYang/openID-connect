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
	//健康检查
	beego.Router("/op/healthy",&controllers.Provider{})
	//rp注册
	beego.Router("/op/register",&controllers.Provider{},"Post:RegisterOp")
	//授权
	// beego.Router("/op/register",&controllers.Provider{},"Post:RegisterOp")
	//授权
}
