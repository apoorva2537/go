package models

import(
	"github.com/jinzhu/gorm"
	"github.com/apoorva2537/go-bookstore/pkg/config"
)
var db *gorm.DB
type Book struct{
	gorm.model
	Name string 'gorm:""json:"name"'
	Author string 'json:"author"'
	Publication string 'json:"piblication"'
}
func init(){
	config.connect()
	db=config.GetDB()
	db.AutoMigrate(&Book{})
}
func (b*book) Create1book() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBooks()[]Books{
	var Books[]Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int 64)(*Book, *gorm.DB){
	var getBook Book
	db:=db.Where("ID=?",Id).Find(&getBook)
	return &getBook,DB
}

func DeleteBook(ID int 64)Book{
	var book Book
	db.Where("IID=?",ID).Delete(book)
	retrun book
}