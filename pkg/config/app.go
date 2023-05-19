// MAIN AIM OF THIS PACKAGE
// It helps creating connection with db.
// The db conncetion created here helps all other files to interact with db

package config

import(
	// Both Mysql packages
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Initiation variable
var(
	db * gorm.DB
)

// Opens connection with db
func Connect(){
	// UserName : Password / TableName ? 
	d, err := gorm.Open("mysql", "Rita376:Hameer376!/testTable1?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	// Saving all the data (d) to db (database)
	db = d
}

// Return db(database) with the help of the package
func GetDB() *gorm.DB{
	return db
}