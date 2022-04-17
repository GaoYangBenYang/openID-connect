create table user_info
(
	id int auto_increment primary key comment 'beego orm 创建表自带字段',
	u_u_i_d varchar(255) not null comment 'UUID',
	user_name varchar(255) not null comment '昵称',
	email varchar(255) not null comment '电子邮箱',
	mobile_number int not null comment '电话号码',
	password varchar(255) not null comment '密码',
	image_url varchar(255) null comment '头像在对象存储中的url',
	permission int default 1 null comment '账号权限',
	is_violation int default 1 not null comment '账号是否异常'
);

