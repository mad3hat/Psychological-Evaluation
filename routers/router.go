// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/index",
			beego.NSInclude(
				&controllers.IndexController{},
			),
		),

		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
		),
		),

		beego.NSNamespace("/signup",
			beego.NSInclude(
				&controllers.SignupController{},
			),
		),

		beego.NSNamespace("/logout",
			beego.NSInclude(
				&controllers.LogoutController{},
			),
		),

		beego.NSNamespace("/articles",
			beego.NSInclude(
				&controllers.ArticleController{},
			),
		),

		beego.NSNamespace("/addarticle",
			beego.NSInclude(
				&controllers.AddArticleController{},
			),
		),


		beego.NSNamespace("/managers",
			beego.NSInclude(
				&controllers.ManagersController{},
			),
		),

		beego.NSNamespace("/posts",
			beego.NSInclude(
				&controllers.PostsController{},
			),
		),

		beego.NSNamespace("/questions",
			beego.NSInclude(
				&controllers.QuestionsController{},
			),
		),

		beego.NSNamespace("/addquestion",
			beego.NSInclude(
				&controllers.AddQuestionsController{},
			),
		),

		beego.NSNamespace("/resources",
			beego.NSInclude(
				&controllers.ResourceController{},
			),
		),

		beego.NSNamespace("/addresource",
			beego.NSInclude(
				&controllers.AddResourceController{},
			),
		),

		beego.NSNamespace("/sessions",
			beego.NSInclude(
				&controllers.SessionsController{},
			),
		),

		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersController{},
			),
		),

		beego.NSNamespace("/about",
			beego.NSInclude(
				&controllers.AboutController{},
			),
		),

		beego.NSNamespace("/test",
			beego.NSInclude(
				&controllers.TestController{},
			),
		),

		beego.NSNamespace("/myinfo",
			beego.NSInclude(
				&controllers.MyinfoController{},
			),
		),

		beego.NSNamespace("/change_info",
			beego.NSInclude(
				&controllers.ChangeinfoController{},
			),
		),

		beego.NSNamespace("/questions1",
			beego.NSInclude(
				&controllers.Questions1Controller{},
			),
		),

		beego.NSNamespace("/questions2",
			beego.NSInclude(
				&controllers.Questions2Controller{},
			),
		),

		beego.NSNamespace("/questions3",
			beego.NSInclude(
				&controllers.Questions3Controller{},
			),
		),

		beego.NSNamespace("/questions4",
			beego.NSInclude(
				&controllers.Questions4Controller{},
			),
		),

		beego.NSNamespace("/questions5",
			beego.NSInclude(
				&controllers.Questions5Controller{},
			),
		),

		beego.NSNamespace("/answer",
			beego.NSInclude(
				&controllers.AnswersController{},
			),
		),

		beego.NSNamespace("/posting",
			beego.NSInclude(
				&controllers.PostingController{},
			),
		),

		beego.NSNamespace("/reading",
			beego.NSInclude(
				&controllers.ReadingController{},
			),
		),

		beego.NSNamespace("/changepass",
			beego.NSInclude(
				&controllers.ChangePassController{},
			),
		),

		beego.NSNamespace("/changearticle",
			beego.NSInclude(
				&controllers.ChangeArticleController{},
			),
		),

		beego.NSNamespace("/changequestion",
			beego.NSInclude(
				&controllers.ChangeQuestionController{},
			),
		),

		beego.NSNamespace("/changeresource",
			beego.NSInclude(
				&controllers.ChangeResourceController{},
			),
		),

		beego.NSNamespace("/addrespond",
			beego.NSInclude(
				&controllers.AddRespondController{},
			),
		),
	)
	beego.AddNamespace(ns)
	
}
