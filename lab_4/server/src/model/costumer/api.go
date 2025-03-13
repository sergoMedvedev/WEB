package costumer

import (
	"lab_4/server/src/settings"
	"log"

	"github.com/gin-gonic/gin"
)

func LoginCostumer(c *gin.Context, cosm *Costumer) (*Costumer, error) {

	user := Costumer{}

	err := settings.DBinst.QueryRow(
		"SELECT user_id, first_name, last_name, role FROM costumer WHERE login = $1 AND password = $2",
		cosm.Login, cosm.Password,
	).Scan(&user.CostumerId, &user.FirstName, &user.LastName, &user.Role)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func RegistCostumer(c *gin.Context, cosm *Costumer) (error) {

	log.Println(*cosm)
	_, err := settings.DBinst.Exec("INSERT INTO costumer (first_name, last_name, login, password, role) VALUES ($1, $2, $3, $4, $5)",
        cosm.FirstName, cosm.LastName, cosm.Login, cosm.Password, cosm.Role)
	if err != nil {
		return  err
	}

	return nil

}
