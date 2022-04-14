package db

import (
	"github.com/astaxie/beego/orm"
	//导入mysql驱动
	_ "github.com/go-sql-driver/mysql"
	"openid_connect/models"
)

func init() {
	// 参数1   driverName
	// 参数2   数据库类型
	// 这个用来设置 driverName 对应的数据库类型
	// mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)

	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	_ = orm.RegisterDataBase("default", "mysql", "root:123456@tcp(101.35.90.54:3306)/dbTest?charset=utf8", 30, 30)

	//注册表
	orm.RegisterModel(new(models.UserInfo))

	//使用表名前缀,创建后的表名为 prefix_user
	//orm.RegisterModelWithPrefix("prefix_", new(User))

	//RunSyncdb 运行 syncdb 命令行。
	//name 表示表的别名。默认为“默认”。
	//force 表示如果当前错误则运行下一个 sql。
	//verbose 意味着在运行命令时显示所有信息。
	_ = orm.RunSyncdb("default", false, true)
}
