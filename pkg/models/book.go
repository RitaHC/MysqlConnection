package models
import (
	"github.com/jinzhu/gorm"
	// Brings in db connection
	"github.com/Rita376/MysqlConnection/pkg/config"
)

var db *gorm.DB

type Book struct{
	// Suggesting that the 'Book' struct shall be used with ORM(Object-Relational-Mapping)
	gorm.model
	// This line has 2 declarations
	// 1 'gorm:""' -> which means it may contain additional configration options
	// 2nd 'json:"name"' -> when converting the Book struct to JSON format, the Name field should be serialized as "name"
	Name string `gorm:""json:"name"`
	Author string `json: "author"`
	Publication string `json: "publication"`
}

func init(){
	config.Connect()
	// Calling the data stored in the DB function in config file and saving it
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func(b *Book) CreateBook() *Book{
	// Creates a new record 
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}