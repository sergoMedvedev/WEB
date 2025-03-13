package controller

import (
	"encoding/json"
	"io"
	"log"

	"lab_4/server/src/model/coach"
	"lab_4/server/src/model/costumer"
	"lab_4/server/src/model/footballer"
	"lab_4/server/src/model/system"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//____________AUTH________________________________

func Login(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	request := costumer.Costumer{}

	err = json.Unmarshal(jsonData, &request)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	user, err := costumer.LoginCostumer(c, &request)
	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid login or password"})
	}
	log.Println(user)

	if user.Role == "" {
		c.JSON(404, gin.H{"error": "Invalid login or password"})
	}
	session := sessions.Default(c)
	session.Set("role", user.Role)
	session.Save()

	c.JSON(200, "{status : ok}")
}

func Registration(c *gin.Context) {
	c.HTML(200, "registration.html", nil)
}

func RegistrationUser(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	log.Println(string(jsonData))

	request := costumer.Costumer{}

	err = json.Unmarshal(jsonData, &request)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	//создание тренера, футболиста если они регистрируются
	if request.Role == "coach" {
		coachDTO := coach.Coach{
			FirstName: request.FirstName,
			LastName:  request.LastName,
		}

		err := coach.CreateCoach(coachDTO)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error server"})
		}
	}

	if request.Role == "player" {
		footballerDTO := footballer.Footballer{
			FirstName: request.FirstName,
			LastNane:  request.LastName,
			Position:  "NUL",
			Rating:    0,
		}

		err := footballer.CreateFootballer(footballerDTO)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error server"})
		}
	}

	if request.Role == "coach" {
		coachDTO := coach.Coach{
			FirstName: request.FirstName,
			LastName:  request.LastName,
		}

		err := coach.CreateCoach(coachDTO)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error server"})
		}
	}

	err = costumer.RegistCostumer(c, &request)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error server"})
	}

	c.JSON(200, "")
}

//____________SYSTEM________________________________

func GetSystemPage(c *gin.Context) {
	role, _ := c.Get("role")
	log.Println(role)
	c.HTML(200, "index.html", nil)
}

func GetPageCreateFootballClub(c *gin.Context) {
	session := sessions.Default(c)
	role := session.Get("role")
	log.Println(role)
	if role != "coach" {
		c.JSON(403, gin.H{"status": "There is no access"})
	}
	c.HTML(200, "create-foolball-club.html", nil)
}

func GetFootballClubs(c *gin.Context) {
	c.HTML(200, "football-club.html", nil)
}

func GetFootballClub(c *gin.Context) {
	footballClubList, err := system.GetFootballClub()
	if err != nil {
		c.JSON(500, gin.H{"status": "Error get football club"})
	}

	c.JSON(200, footballClubList)
}

func GetPageSquad(c *gin.Context) {
	c.HTML(200, "squad.html", nil)
}

func GetSquad(c *gin.Context) {
	squadList, err := system.GetSquadList()
	if err != nil {
		c.JSON(500, gin.H{"status": "Error get squad"})
	}

	c.JSON(200, squadList)
}

func GetPageCoach(c *gin.Context) {
	c.HTML(200, "coaches.html", nil)
}

func GetPageFootballer(c *gin.Context) {
	c.HTML(200, "footballer.html", nil)
}

//____________COACH________________________________

func GetCoaches(c *gin.Context) {

	arrayCoach, err := coach.GetCoaches()
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, arrayCoach)
}

//____________FOOTBALLER_______________________________

func GetFootballers(c *gin.Context) {
	arrayCoach, err := footballer.GetFootballers()
	if err != nil {
		c.JSON(500, err.Error())
	}

	c.JSON(200, arrayCoach)
}

func CreateFootbalClubAndSquad(c *gin.Context) {
	session := sessions.Default(c)
	role := session.Get("role")
	log.Println(role)
	if role != "coach" {
		c.JSON(403, gin.H{"status": "There is no access"})
	}

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	request := system.CreateFootbalClubAndSquadRequest{}

	err = json.Unmarshal(jsonData, &request)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	log.Println(request)

	response, err := system.CreateFootballClubAndSquad(request)
	if err != nil {
		c.JSON(500, gin.H{"status": "Error"})
	}

	c.JSON(200, response)
}
