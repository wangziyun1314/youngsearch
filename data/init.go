package data

import (
	//"github.com/opesun/goquery"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

//Photo 图片的信息
type Photo struct {
	Index   int    `gorm:"primary_key"` //图片在mysql里面的id
	Url     string //图片的打开地址
	Content string //图片的描述文字
}

//init 初始化db
func init() {
	db, err := gorm.Open(mysql.Open("root:wangziyun@tcp(localhost:3306)/young_search?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Page{})
	db.AutoMigrate(&Photo{})
	Db = db
}

//InsertPage 将网页信息入库
func InsertPage(p *Page) {
	Db.Create(p)
}

// InsertPages 将多个网页入库
func InsertPages(pages []*Page) {
	Db.Create(pages)
}

// InsertPhoto 将图片信息入库
func InsertPhoto(photo *Photo) {
	Db.Create(photo)
}

// InsertPhotos 将多个图片入库
func InsertPhotos(photos []*Photo) {
	Db.Create(photos)
}

// FindPhotoById 由id找到对应的图片信息
func FindPhotoById(id int) Photo {
	var p Photo
	p = Photo{Index: id}
	Db.Find(&p)
	return p
}
