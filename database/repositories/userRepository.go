package repositories
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
//func (repo *UserRepository) FindUserById(ID string) models.User {
//
//	row := repo.Db.QueryRow("SELECT id, name FROM user where id = ? LIMIT 1", ID)
//
//	var user User
//
//	if err := row.Scan(&user.Id, &user.Name); err != nil {
//		if err == sql.ErrNoRows {
//			return nil
//		}
//		panic(err)
//	}
//
//	return &user
//
//}
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