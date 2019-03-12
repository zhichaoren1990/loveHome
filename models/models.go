package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)


/**
	用户表
 */
type User struct {
	Id         int `json:"user_id"`
	Name       string `orm:"size(100);unique" json:"name"`
	Password   string `orm:"size(128)" json:"password"`
	Mobile     string `orm:"size(11);unique" json:"mobile"`
	Real_name  string `orm:"size(32)" json:"real_name"`
	Id_card    string `orm:"size(20)" json:"id_card"`
	Avatar_url string `orm:"size(256)" json:"avatar_url"`
	Reg_time   time.Time `orm:"auto_now_add;type(dateTime)"json:"Reg_time"`

}

/**
	房屋信息表
 */

type House struct {
	Id   int `json:"id"`
	Title string `orm:"size(64)"json:"title"`                                          //房屋标题
	Price int `orm:"default(0)" json:"price"`         	                               //房屋价格
	Address string `orm:"size(512)" orm:"default("")" json:"address"`                  //房屋地址
	Room_count int `orm:"default(1)"json:"room_count"`                                 //房间数量
	Acreage   int `orm:"default(0)" json:"acreage"`                                    //房间总面积
	Unit  string `orm:"size(32)" orm:"default" json:"Unit"`                            //房间单元 几室几厅
	Capacity  string  `orm:"default(1)" json:"capacity"`                               //房间容纳的总人数
	Beds   string  `orm:"size(64)" orm:"default("")" json:"beds"`                      //房屋床铺的配置
	Deposit int `orm:"default(0)"json:"deposit"`                                       //押金
	Min_days int `orm:"default(1)" json:"min_days"`                                    //最好入住的天数
	Max_days int `orm:"default(0)" json:"max_days"`                                    //最多入住的天数 0表示不限制
	Index_image_url string `orm:"size(256)" orm:"default("")" json:"index_image_url"`  //房间主图
	Ctime time.Time `orm:"auto_now_add;type(datetime)"json:"ctime"`

	User *User `orm:"rel(fk)"json:"user"`
	Area *Area `orm:"rel(fk)"json:"area"`
}


/**
	区域信息表
 */

type Area struct {
	Id int `json:"id"`
	Name string `orm:"size(32)"json:"name"`
}


/**
	房间设施表
 */
type Facility struct {
	Id int `json:"id"`
	Name string `orm:"size(32)"json:"name"`
	House []*House `orm:"rel(m2m)"json:"house"`
}

/**
	房屋图片
 */

type HouseImage struct {
	Id int `json:"id"`
	Url string `orm:"size(256)"json:"url"`
	House *House `orm:"rel(fk)"json:"house"`
}

/**
	订单表
 */

type OrderHouse struct {
	Id int `json:"id"`
	Days int `orm:"size(11)" orm:"default(1)"json:"days"`    //预定天数
	House_price int `orm:"default("")"`                      //房屋价格
	Amount  int  `orm:"size(11)"json:"amount"`				 //订单总价
	Status string `orm:"default(WAIT_ACCEPT)"`               //订单状态
	Comment string `orm:"size(512)"json:"comment"`
	Ctime time.Time `orm:"auto_now_add;type(dateTime)"json:"ctime"`
	House *House `orm:"rel(fk)"json:"house"`
	User *User `orm:"rel(fk)"json:"user"`

}
 
 

func init()  {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/loveHome?charset=utf8&loc=Asia%2FShanghai", 30)
	orm.RegisterModel(new(User),new(Area),new(House),new(Facility),new(HouseImage),new(OrderHouse))
	//orm.RunSyncdb("default", false, true)

}