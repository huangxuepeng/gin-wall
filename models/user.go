package models

import (
	"time"

	"gorm.io/gorm"
)

//基础
type Base struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"column:create_time"`
	UpdatedAt time.Time      `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt `gorm:"<-:update"`
}

//评论, 动态的公共部分
type DynamicInfo struct {
	Like        uint `gorm:"column:like; comment '点赞'; default: 0"`
	Unlike      uint `gorm:"column:un_like; comment '踩'; default: 0"`
	NumberFloor uint `gorm:"column:number_floor;comment '楼层的高度 按照点赞的数量来排序'"`
}

//用户注册
type UserRegister struct {
	Base
	Mobile              string `gorm:"index;type:char(11);column:mobile;unique;not null;comment '手机号码'"`
	NickName            string `gorm:"type:varchar(100);column:nick_name;not null;comment '昵称'"`
	Avatar              string `gorm:"type:varchar(100);column:avatar;comment '头像'"`
	StudentNumber       string `gorm:"index;type:char(9);unique;not null;comment '学号'"`
	Email               string `gorm:"type:varchar(50);column:email;not null;comment '邮箱'"`
	Password            string `gorm:"type:varchar(150);column:password;not null;cpment '密码'"`
	Sex                 uint8  `gorm:"column:sex;comment '1 代表男生 0 代表女生';default:0"`
	Constellation       string `gorm:"type:varchar(10);column:constellation;comment '星座'"`
	Role                uint8  `gorm:"type:varchar(5);column:role; comment '0 普通用户 1 老师 2 管理员';default:0"`
	AuthenticationToken string `gorm:"type:varchar(196);column:authentication_token;comment '认证token'"`
	EmailAuthentication string `gorm:"type:varchar(15);column:email_authentication;not null;comment '邮箱验证码'"`
	BinningTime         string `gorm:"type:varchar(15);column:binning_time;comment '拉黑时间'"`
	ISReal              uint8  `gorm:"column:is_real;not null;comment '0 未实名 1 已经实名';default:0"`
	AdminID             uint   `gorm:"column:admin_id;comment '被审核之后存入的审核人的ID'"`
	UserRealname        UserRealname
	OnlineLog           []OnlineLog
	InfoLog             []InfoLog
	LoginInfo           []LoginInfo
	FinshSum            []FinshSum
	DynamicInformation  []DynamicInformation
	FatherComment       []FatherComment
	SonComment          []SonComment
}

//用户实名
type UserRealname struct {
	Base
	StudentNumber        string `gorm:"index;type:char(9);unique;not null;comment '学号'"`
	School               string `gorm:"column:school;type:varchar(100);not null;comment '学校名称'"`
	RealName             string `gorm:"column:real_name;type:varchar(20);not null;comment '真实姓名'"`
	Academy              string `gorm:"column:academy;type:varchar(20);not null;comment '学院名称'"`
	Profession           string `gorm:"column:type:varchar(20);not null;comment '专业'"`
	Age                  uint8  `gorm:"column:age;type:varchar(4);not null;comment '年龄'"`
	TeacherName          string `gorm:"column:teacher_name;type:varchar(20);comment '导师姓名(允许为空)'"`
	DormitoryFloorNumber string `gorm:"column:dormitory_floor_number;type:varchar(5);comment '寝室楼号'"`
	DormitoryNumber      string `gorm:"column:dormitory_number;type:varchar(5);comment '寝室号'"`
	UserRegisterID       uint
	TeacherID            uint
}

//期末总结(先不写)
type FinshSum struct {
	Base
	UserRegisterID uint
}

//导师
type Teacher struct {
	Base
	Name          string `gorm:"type:varchar(20);column:name;not null;comment '导师姓名'"`
	TeacherNumber string `gorm:"type:char(9);unique;column:teacher_number;not null;comment '教工号'"`
	Mobile        string `gorm:"index;type:char(11);column:mobile;unique;not null;comment '手机号码'"`
	JobTitle      string `gorm:"type:varchar(20);column:job_title;not null;comment '导师职称'"`
	UserRealname  []UserRealname
}

// 在线信息 用户名 IP 浏览器 操作系统 登录地点
type OnlineLog struct {
	Base
	IP              string `gorm:"type:varchar(50);column:ip;not null;comment 'ip地址'"`
	Browser         string `gorm:"type:varchar(20);column:ip;not null;comment '浏览器'"`
	OperatingSystem string `gorm:"type:varchar(20);column:operating_system;not null;comment '操作系统'"`
	UserRegisterID  uint
}

//动态信息
type DynamicInformation struct {
	Base
	DynamicInfo
	Type           string `gorm:"type:varchar(20);column:type;comment '文章类型'"`
	ISReal         uint8  `gorm:"default:0;comment '0 未认证 1 已认证'"`
	AdminID        uint   `gorm:"column:admin_id;comment '用户的ID'"`
	Report         uint8  `gorm:"default:0;comment '0 未被举报 1 被举报'"`
	FatherComment  []FatherComment
	UserRegisterID uint
}

//父评论
type FatherComment struct {
	Base
	DynamicInfo
	Comment              string `gorm:"type:varchar(255);column:comment;comment '评论内容'"`
	Report               uint8  `gorm:"column:report;default:0;comment '0 未被举报 1 被举报'"`
	SonComment           []SonComment
	DynamicInformationID uint
	UserRegisterID       uint
}

//子评论
type SonComment struct {
	Base
	DynamicInfo
	Comment         string `gorm:"type:varchar(255);column:comment;comment '评论内容'"`
	Report          uint8  `gorm:"default:0;comment '0 未被举报 1 被举报'"`
	FatherCommentID uint
	UserRegisterID  uint
}

//日志信息  用户名 访问的url
type InfoLog struct {
	Base
	URL            string `gorm:"type:varchar(50);column:url;comment 'url'"`
	UserRegisterID uint
}

// 登录事件记录
type LoginInfo struct {
	Base
	UserRegisterID uint
}
type Forum struct {
	Base
}

//文章信息(有可能有些人发表的文章字数比较多, 所以将文章的内容重新拿出来一张表)
type Essay struct {
	Base
}

//情侣关系表
//聊天内容
//表白功能
//广告信息
