package controllers

import (
	"fmt"

	"github.com/ganggas95/trawangan/app"
	"github.com/ganggas95/trawangan/app/models"
	"github.com/ganggas95/trawangan/app/routes"
	"github.com/revel/revel"
	"strings"
	"time"
)

type Persons struct {
	App
}

func (c Persons) checkUser() revel.Result {
	if user := c.connected(); &user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	}
	return nil
}
func (c Persons) Index() revel.Result {
	if user := c.connected(); &user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	}
	return c.Render()
}
func (c Persons) List(search string) revel.Result {
	user := c.connected()
	if user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	} else if !user.Verify {
		c.Flash.Error("Your account not verified. Check Your email to ferify your account!")
		c.RenderArgs["user"] = user.Username
		return c.Redirect(routes.Persons.UnverifyAcc())
	}
	sr := ""
	if search != "" {
		sr = strings.TrimSpace(search)
	} else {
		sr = strings.ToLower(search)
	}
	var person []models.Person
	db := app.GORM.Debug().Where("nama LIKE ?", "%"+sr+"%").Find(&person)
	if db.RecordNotFound() {
		c.Redirect(routes.Persons.Index())
	}
	for i := 0; i < len(person); i++ {
		person[i].TanggalLahir.Format("Jan _2, 2006")
	}
	fmt.Println(person)
	c.RenderArgs["persons"] = person
	return c.Render(&person)
}
func (c Persons) Show(id int64) revel.Result {
	user := c.connected()
	if user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	} else if !user.Verify {
		c.Flash.Error("Your account not verified. Check Your email to ferify your account!")
		c.RenderArgs["user"] = user.Username
		return c.Redirect(routes.Persons.UnverifyAcc())
	}
	persons := c.getPerson(id)
	return c.Render(persons)
}
func (c Persons) UnverifyAcc() revel.Result {
	user := c.connected()
	fmt.Println(user)
	if user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	}
	if user.Verify {
		return c.Redirect(routes.Persons.List(""))
	}
	return c.Render()
}
func (c Persons) Tambah() revel.Result {
	user := c.connected()
	if user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	} else if !user.Verify {
		c.Flash.Error("Your account not verified. Check Your email to ferify your account!")
		c.RenderArgs["user"] = user.Username
		return c.Redirect(routes.Persons.UnverifyAcc())
	}
	return c.Render()
}
func (c Persons) AddData(person *models.Person) revel.Result {
	person.TanggalLahir.Format("2006-01-02")
	c.Validation.Required(person.Nama)
	c.Validation.Required(person.Alamat)
	c.Validation.Required(person.TempatLahir)
	c.Validation.Required(person.TanggalLahir)
	c.Validation.Required(person.Pekerjaan)
	person.Validation(c.Validation)
	user := c.connected()
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Persons.Tambah())
	} else if user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	} else if !user.Verify {
		c.Flash.Error("Your account not verified. Check Your email to ferify your account!")
		c.RenderArgs["user"] = user.Username
		return c.Redirect(routes.Persons.UnverifyAcc())
	}

	app.GORM.Create(&person)
	return c.Redirect(routes.Persons.List(""))
}

func (c Persons) Delete(id int64) revel.Result {
	user := c.connected()
	if user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	} else if !user.Verify {
		c.Flash.Error("Your account not verified. Check Your email to ferify your account!")
		c.RenderArgs["user"] = user.Username
		return c.Redirect(routes.Persons.UnverifyAcc())
	}
	person := models.Person{IdPerson: id}
	err := app.GORM.Delete(&person)
	if err.Error != nil {
		panic(err.Error)
	}
	return c.Redirect(routes.Persons.List(""))
}

func (c Persons) Ubah(nama, alamat, tempatlahir, pekerjaan string, tanggallahir time.Time, id int64) revel.Result {
	user := c.connected()
	if user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	} else if !user.Verify {
		c.Flash.Error("Your account not verified. Check Your email to ferify your account!")
		c.RenderArgs["user"] = user.Username
		return c.Redirect(routes.Persons.UnverifyAcc())
	}
	person := c.getPerson(id)
	person.Nama = nama
	person.Alamat = alamat
	person.TempatLahir = tempatlahir
	person.TanggalLahir = tanggallahir
	person.Pekerjaan = pekerjaan
	err := app.GORM.Save(&person)
	if err.Error != nil {
		panic(err.Error)
	}
	return c.Redirect(routes.Persons.List(""))
}

func (c Persons) GetData(id int64) revel.Result {
	user := c.connected()
	if user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	} else if !user.Verify {
		c.Flash.Error("Your account not verified. Check Your email to ferify your account!")
		c.RenderArgs["user"] = user.Username
		return c.Redirect(routes.Persons.UnverifyAcc())
	}
	person := c.getPerson(id)
	return c.Render(person)
}

func (c Persons) getPerson(id int64) *models.Person {
	var persons models.Person
	person := app.GORM.First(&persons, id)
	if person.Error != nil {
		panic(person.Error)
	}
	return &persons
}

func (c Persons) Logout() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Login())
	}

	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Login())
}
