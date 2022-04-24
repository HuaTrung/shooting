package repositories

import (
	"gorm.io/gorm"
	"log"
	"shootingplane/entity/models"
)

//
//import (
//	"database/sql"
//	"gorm.io/gorm"
//	"log"
//	models "shootingplane/entity/models"
//)
//
//func AddUser(Db *gorm.DB,user models.User) {
//	stmt, err := repo.Db.Prepare("INSERT INTO user(id, name) values(?,?)")
//	checkErr(err)
//
//	_, err = stmt.Exec(user.GetId(), user.GetName())
//	checkErr(err)
//}
//
//func (repo *UserRepository) RemoveUser(user models.User) {
//	stmt, err := repo.Db.Prepare("DELETE FROM user WHERE id = ?")
//	checkErr(err)
//
//	_, err = stmt.Exec(user.GetId())
//	checkErr(err)
//}
//
func FindUserByName(Db *gorm.DB,Name string) *models.User {
	tb := Db.Table("personal.users")
	var user models.User
	row := tb.Unscoped().First(&user,"Name = ?",Name)
	if row.Error!=nil{
		log.Print(row.Error)
		return nil
	} else {
		if user.GetId()==0{
			return nil
		} else {
			return &user
		}
	}

}
//
//func (repo *UserRepository) GetAllUsers() []models.User {
//
//	rows, err := repo.Db.Query("SELECT id, name FROM user")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	var users []models.User
//	defer rows.Close()
//	for rows.Next() {
//		var user User
//		rows.Scan(&user.Id, &user.Name)
//		users = append(users, &user)
//	}
//
//	return users
//}
//
//func (repo *UserRepository) FindUserByUsername(username string) *User {
//
//	row := repo.Db.QueryRow("SELECT id, name, username, password FROM user where username = ? LIMIT 1", username)
//
//	var user User
//
//	if err := row.Scan(&user.Id, &user.Name, &user.Username, &user.Password); err != nil {
//		if err == sql.ErrNoRows {
//			return nil
//		}
//		panic(err)
//	}
//
//	return &user
//
//}