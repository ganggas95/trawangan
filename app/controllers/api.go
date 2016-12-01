package controllers

import (
	"github.com/ganggas95/trawangan/app"
	"github.com/ganggas95/trawangan/app/job"
	"github.com/ganggas95/trawangan/app/models"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type Api struct {
	App
}

type Error404 struct {
	TextError string
	TypeError string
}

func (api Api) Auth(username, password string) revel.Result {
	user := api.GetUser(username)

	data := make(map[string]interface{})
	if &user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		agent := api.DetectAgent(user.UID)
		if err == nil {
			user.Agent = agent
			data["error"] = nil
			data["user"] = user
			return api.RenderJson(data)
		} else {
			error404 := Error404{TextError: "Username And Password is invalid!", TypeError: "Valid User"}
			data["error"] = error404
			data["user"] = nil
			return api.RenderJson(data)
		}
	}
	data["user"] = nil
	return api.RenderJson(data)
}

func (c Api) AddUser(user models.User, password string) revel.Result {
	c.Validation.Required(password)
	user.Validation(c.Validation)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.RenderJson(c)
	}
	user.HashedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	tok := job.RandomToken(32)
	err3 := job.SendToken(user.Email, tok)
	if err3 != nil {
		panic(err3)
	}
	var token models.UserToken
	token.AccessToken = tok
	user.TokenUser = token
	db := app.GORM.Create(&user)
	if db.Error != nil {
		revel.INFO.Println(db.GetErrors())
	}
	return c.RenderJson(user)
}

func (api Api) DetectAgent(userId int64) models.AgentTravel {
	var agent models.AgentTravel
	db := app.GORM.Where("user_id = ?", userId).Find(&agent)
	if db.RecordNotFound() {

	}
	return agent
}
