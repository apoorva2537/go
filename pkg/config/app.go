package config
import(
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)
var {
	db* gorm.DB
}
func connect(){
	d,err := gorm.Open("mysql","apoorva:apoorva@123/simpleres?charset=utf8&parserTime=True&;oc=Local")
	if err!=nil{
		panic(err)
	}
	db=d
}
func getDB() *gorm.DB(
	return db
)