package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/ganggas95/trawangan/app"
	//"github.com/ganggas95/trawangan/app/job"
	"github.com/ganggas95/trawangan/app/models"
	"github.com/ganggas95/trawangan/app/routes"
	"github.com/revel/revel"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Agent struct {
	App
}

func (a Agent) CheckAgent() *models.AgentTravel {
	user := a.connected()
	if user == nil {
		return nil
	}
	var agent models.AgentTravel
	err := app.GORM.Where("user_id = ?", user.UID).Find(&agent)
	if err.RecordNotFound() {
		return nil
	}
	return &agent

}

func (a Agent) Index() revel.Result {
	agent := a.CheckAgent()
	if agent != nil {
		user := a.GetUserWitId(agent.UserId)
		if !user.Verify {
			a.Flash.Error("Your account not verified. Check Your email to ferify your account!")
			a.RenderArgs["user"] = user.Username
			return a.Redirect(routes.Persons.UnverifyAcc())
		} else {
			var foto models.UserFoto
			db := app.GORM.Where("user_id = ?", user.UID).Find(&foto)
			if db.RecordNotFound() {

			}
			dir := "/public/data/" + user.Username + "/profile/"
			fotodir := dir + foto.Foto
			dahsboard := "dashboard"
			a.RenderArgs["agent"] = agent
			a.RenderArgs["dashboard"] = "dashboard"
			a.RenderArgs["fotodir"] = fotodir

			return a.Render(agent, dahsboard, fotodir)
		}

	} else {
		a.FlashParams()
		a.Flash.Error("Please register your account to Agent")
		return a.Redirect(routes.App.Index())
	}
}

func (a Agent) ServiceAgent() revel.Result {
	agent := a.CheckAgent()
	if agent != nil {
		user := a.GetUserWitId(agent.UserId)
		if !user.Verify {
			a.Flash.Error("Your account not verified. Check Your email to ferify your account!")
			a.RenderArgs["user"] = user.Username
			return a.Redirect(routes.Persons.UnverifyAcc())
		} else {
			var service []models.AgentService
			db := app.GORM.Where("agent_id = ?", agent.IdAgent).Find(&service)
			if db.RecordNotFound() {
				a.FlashParams()
				a.Flash.Error("No Service")
				return a.Redirect(routes.Agent.ServiceAgent())
			}

			var fotoUser models.UserFoto
			db = app.GORM.Where("user_id = ?", user.UID).Find(&fotoUser)
			if db.RecordNotFound() {

			}
			dir := "/public/data/" + user.Username + "/profile/"
			fotodir := dir + fotoUser.Foto
			for i := 0; i < len(service); i++ {
				var foto models.FotoService
				db1 := app.GORM.Find(&foto, models.FotoService{IdService: service[i].IdService})
				if db1.RecordNotFound() {

				}
				service[i].Foto = foto

			}
			a.RenderArgs["service"] = service
			a.RenderArgs["agent"] = agent
			a.RenderArgs["fotodir"] = fotodir
			return a.Render()

		}

	} else {
		a.FlashParams()
		a.Flash.Error("Please register your account to Agent")
		return a.Redirect(routes.App.Index())
	}
}

func (a Agent) OrderAgent() revel.Result {
	agent := a.CheckAgent()
	if agent != nil {
		user := a.GetUserWitId(agent.UserId)
		if !user.Verify {
			a.Flash.Error("Your account not verified. Check Your email to ferify your account!")
			a.RenderArgs["user"] = user.Username
			return a.Redirect(routes.Persons.UnverifyAcc())
		} else {
			var foto models.UserFoto
			db := app.GORM.Where("user_id = ?", user.UID).Find(&foto)
			if db.RecordNotFound() {

			}
			order := "order"
			dir := "/public/data/" + user.Username + "/profile/"
			fotodir := dir + foto.Foto
			a.RenderArgs["agent"] = agent
			a.RenderArgs["fotodir"] = fotodir
			a.RenderArgs["order"] = order
			return a.Render(agent, order, fotodir)
		}

	} else {
		a.FlashParams()
		a.Flash.Error("Please register your account to Agent")
		return a.Redirect(routes.App.Index())
	}
}

func (a Agent) ChatAgent() revel.Result {
	agent := a.CheckAgent()
	if agent != nil {
		user := a.GetUserWitId(agent.UserId)
		if !user.Verify {
			a.Flash.Error("Your account not verified. Check Your email to ferify your account!")
			a.RenderArgs["user"] = user.Username
			return a.Redirect(routes.Persons.UnverifyAcc())
		} else {
			var foto models.UserFoto
			db := app.GORM.Where("user_id = ?", user.UID).Find(&foto)
			if db.RecordNotFound() {

			}
			chat := "chat"
			dir := "/public/data/" + user.Username + "/profile/"
			fotodir := dir + foto.Foto
			a.RenderArgs["agent"] = agent
			a.RenderArgs["fotodir"] = fotodir
			a.RenderArgs["chat"] = chat
			return a.Render()
		}

	} else {
		a.FlashParams()
		a.Flash.Error("Please register your account to Agent")
		return a.Redirect(routes.App.Index())
	}
}
func (a Agent) MemberAgent() revel.Result {
	agent := a.CheckAgent()
	if agent != nil {
		user := a.GetUserWitId(agent.UserId)
		if !user.Verify {
			a.Flash.Error("Your account not verified. Check Your email to ferify your account!")
			a.RenderArgs["user"] = user.Username
			return a.Redirect(routes.Persons.UnverifyAcc())
		} else {
			var foto models.UserFoto
			db := app.GORM.Where("user_id = ?", user.UID).Find(&foto)
			if db.RecordNotFound() {

			}
			member := "member"
			dir := "/public/data/" + user.Username + "/profile/"
			fotodir := dir + foto.Foto
			a.RenderArgs["agent"] = agent
			a.RenderArgs["fotodir"] = fotodir
			a.RenderArgs["member"] = member
			return a.Render(agent, member, fotodir)
		}

	} else {
		a.FlashParams()
		a.Flash.Error("Please register your account to Agent")
		return a.Redirect(routes.App.Index())
	}
}

func (a Agent) RegisterAgent() revel.Result {
	user := a.connected()
	if user == nil {
		return a.Redirect(routes.App.Index())
	}
	return a.Render(user)
}

func (a Agent) UniqueHandler(email, website string) bool {
	var travelAgent models.AgentTravel
	db := app.GORM.Debug().Where("website_agent = ? OR email_agent = ?", website, email).Find(&travelAgent)
	if db.RecordNotFound() {
		return false
	} else {
		return true
	}

}

func (a Agent) AddAgentFromUser(travelAgent models.AgentTravel) revel.Result {
	users := a.connected()
	agent := a.UniqueHandler(travelAgent.Email, travelAgent.Website)
	if agent {
		a.Validation.Keep()
		a.FlashParams()
		a.Flash.Error("Your email and website url have been registered!")
		return a.Redirect(routes.Agent.RegisterAgent())
	}
	travelAgent.Validation(a.Validation)
	if a.Validation.HasErrors() {
		a.Validation.Keep()
		a.FlashParams()
		return a.Redirect(routes.Agent.RegisterAgent())
	}
	root_path := revel.BasePath
	dst_path := "public/data/" + users.Username + "/agent/"
	path := filepath.Join(root_path, dst_path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err2 := os.Mkdir(path, os.ModePerm)
		if err2 != nil {
			panic(err2)
		}
	}
	travelAgent.UserId = users.UID
	err := app.GORM.Create(&travelAgent)
	if err.Error != nil {
		panic(err.Error)
	}
	return a.Redirect(routes.Agent.Index())
}

func (a Agent) GetAgent(userId int64) *models.AgentTravel {
	var agent models.AgentTravel
	err := app.GORM.Where("user_id = ?", userId).Find(&agent)
	if err.Error != nil {
		panic(err.Error)
	}

	return &agent
}

func (a Agent) ParseAgentService() (models.AgentService, error) {
	service := models.AgentService{}
	req := a.Request.Body
	data := map[string]string{}
	content, _ := ioutil.ReadAll(req)
	json.Unmarshal(content, &data)
	for k, v := range data {
		revel.INFO.Println(k, "\t : ", v)
	}
	err := json.NewDecoder(req).Decode(&service)
	return service, err
}

func (a Agent) SetService(agentService models.AgentService) revel.Result {

	agent := a.CheckAgent()
	if agent != nil {
		user := a.GetUserWitId(agent.UserId)

		if !user.Verify {
			a.Flash.Error("Your account not verified. Check Your email to ferify your account!")
			a.RenderArgs["user"] = user.Username
			return a.Redirect(routes.Persons.UnverifyAcc())
		} else {

			a.Validation.Required(agentService.Service).Message("Service is required")
			a.Validation.Required(agentService.Kategori).Message("Kategori is required")
			a.Validation.Required(agentService.Price).Message("Price is required")

			if a.Validation.HasErrors() {
				a.Validation.Keep()
				a.FlashParams()
				return a.Redirect(routes.Agent.ServiceAgent() + "#modalAddService")
			}
			root_path := revel.BasePath
			dst_path := "public/data/" + user.Username + "/agent/service"
			path := filepath.Join(root_path, dst_path)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				err2 := os.Mkdir(path, os.ModePerm)
				if err2 != nil {
					panic(err2)
				}
			}
			agentService.IdAgent = agent.IdAgent
			db := app.GORM.Create(&agentService)
			if db.Error != nil {
				panic(db.Error)
			}
			return a.Redirect(routes.Agent.ServiceAgent())
		}

	}
	a.FlashParams()
	a.Flash.Error("Please register your account to Agent")
	return a.Redirect(routes.App.Index())
}

func (a Agent) DeleteService(idService int64) revel.Result {
	agent := a.CheckAgent()
	if agent == nil {
		return a.Redirect(routes.Agent.ServiceAgent())
	}
	serviceAgent := models.AgentService{IdService: idService}
	var serviceFoto models.FotoService
	db := app.GORM.Where("service_id = ?", idService).Find(&serviceFoto)
	db = app.GORM.Debug().Delete(&serviceFoto)
	db = app.GORM.Debug().Delete(&serviceAgent)
	if db.Error != nil {
		panic(db.Error)
	}
	a.Flash.Success("Service success deleted!!")
	return a.Redirect(routes.Agent.ServiceAgent())
}

func (a Agent) ActiveService(idService int64) revel.Result {
	agent := a.CheckAgent()
	if agent == nil {
		return a.Redirect(routes.Agent.ServiceAgent())
	}
	var serviceAgent models.AgentService
	db := app.GORM.First(&serviceAgent, idService)
	if db.RecordNotFound() {
		panic(db.RecordNotFound())
	}
	serviceAgent.Status = true
	db = app.GORM.Save(&serviceAgent)
	if db.Error != nil {
		a.Flash.Error("Service failed to active")
		return a.Redirect(routes.Agent.ServiceAgent())
	}
	a.Flash.Success("Service activated")
	return a.Redirect(routes.Agent.ServiceAgent())
}
func (a Agent) DisableService(idService int64) revel.Result {
	agent := a.CheckAgent()
	if agent == nil {
		return a.Redirect(routes.Agent.ServiceAgent())
	}
	var serviceAgent models.AgentService
	db := app.GORM.First(&serviceAgent, idService)
	if db.RecordNotFound() {
		panic(db.RecordNotFound())
	}
	serviceAgent.Status = false
	db = app.GORM.Save(&serviceAgent)
	if db.Error != nil {
		a.Flash.Error("Service failed to active")
		return a.Redirect(routes.Agent.ServiceAgent())
	}
	a.Flash.Success("Service activated")
	return a.Redirect(routes.Agent.ServiceAgent())
}
func (c Agent) GetServiceWithId(idService int64) (*models.AgentService, error) {
	db := app.GORM
	var serviceAgent models.AgentService
	db.First(&serviceAgent, idService)
	if db.RecordNotFound() {
		return nil, db.Error
	}
	return &serviceAgent, nil
}

/*
func (c Agent) EditService(idService int64) revel.Result {
	agent := c.ChatAgent()

	if agent != nil {
		var service models.AgentService
		app.GORM.First(&service, idService)
		return c.RenderJson(&service)
	}
	return c.Redirect(routes.App.Index())
}
*/
func (c Agent) AddOnService(idService int64) revel.Result {
	agent := c.CheckAgent()
	if agent != nil {
		var addOn []models.AddOnService
		db := app.GORM.Find(&addOn).Where("service_id = ?", idService)
		if db.RecordNotFound() {
			c.Flash.Error("Layanan Tambahan Kosong")
			c.FlashParams()
		}
		c.RenderArgs["addOn"] = addOn
		return c.RenderJson(addOn)
	}
	return c.Redirect(routes.App.Index())
}

func (c Agent) SetAddOn(addOn []models.AddOnService, idService int64) revel.Result {
	agent := c.CheckAgent()
	if agent != nil {
		var service models.AgentService
		app.GORM.First(&service, idService)
		service.AddOn = addOn
		db := app.GORM.Save(&service)
		if db.Error != nil {
			panic(db.Error)
		}
		return c.Redirect(routes.Agent.Service(idService))
	}
	return c.Redirect(routes.App.Index())
}
func (c Agent) GetAddOn(idService int64) *[]models.AddOnService {
	var addOn []models.AddOnService
	db := app.GORM.Where("service_id = ?", idService).Find(&addOn)
	if db.RecordNotFound() {
		return nil
	}
	return &addOn
}

func (c Agent) RemoveAddOn(idAddOn, idService int64) revel.Result {
	agen := c.CheckAgent()
	if agen != nil {
		app.GORM.Delete(&models.AddOnService{IdAddOnService: idAddOn})
		return c.Redirect(routes.Agent.Service(idService))
	}
	return c.Redirect(routes.App.Index())
}

func (a Agent) UploadPage(idService int64) revel.Result {
	agent := a.CheckAgent()
	if agent != nil {
		service, err := a.GetServiceWithId(idService)
		if err != nil {
			panic(err)
		}
		if service == nil {
			return a.Redirect(routes.Agent.ServiceAgent())
		}
		a.RenderArgs["agent"] = agent
		a.RenderArgs["service"] = service
		return a.Render()
	}
	return a.Redirect(routes.App.Index())
}

func (a Agent) UploadFoto(foto []byte, idService int64) revel.Result {
	user := a.connected()
	if user != nil {
		agent := a.CheckAgent()
		if agent != nil {
			a.Validation.Required(foto).Message("foto is required")

			a.Validation.MinSize(foto, 2*KB).
				Message("Minimum a file size of 2KB expected")
			a.Validation.MaxSize(foto, 2*MB).
				Message("File cannot be larger than 2MB")

			// Check format of the file.
			conf, format, err1 := image.DecodeConfig(bytes.NewReader(foto))
			a.Validation.Required(err1 == nil).Key("foto").
				Message("Incorrect file format")
			a.Validation.Required(format == "jpeg" || format == "png").Key("foto").
				Message("JPEG or PNG file format is expected")

			// Check resolution.
			a.Validation.Required(conf.Height >= 150 && conf.Width >= 150).Key("foto").
				Message("Minimum allowed resolution is 150x150px")
			if a.Validation.HasErrors() {
				a.Validation.Keep()
				a.FlashParams()
				a.Flash.Error("Upload failed.")
				return a.Redirect(routes.Agent.GaleryFoto(idService))
			}
			var img models.FotoService
			m := a.Request.MultipartForm
			for fname, _ := range m.File {
				fheader := m.File[fname]
				for i, _ := range fheader {
					file, err := fheader[i].Open()
					defer file.Close()
					if err != nil {
						panic(err)
					}
					root_path := revel.BasePath
					dst_path := "public/data/" + user.Username + "/agent/service/"
					path := filepath.Join(root_path, dst_path)
					if _, err := os.Stat(path); os.IsNotExist(err) {
						err2 := os.Mkdir(path, os.ModePerm)
						if err2 != nil {
							panic(err2)
						}
					} else {
						dst, err := os.Create(path + "/" + fheader[i].Filename)
						if err != nil {
							panic(err)
						}
						if _, err := io.Copy(dst, file); err != nil {
							panic(err)
						}
						defer dst.Close()

						img.Foto = fheader[i].Filename
						img.Height = conf.Height
						img.Width = conf.Width
						img.Format = format
						img.Size = len(foto)
						img.Dir = "/" + dst_path + "/"
					}
					service, err9 := a.GetServiceWithId(idService)
					if err9 != nil {
						panic(err9)
					}
					if service == nil {
						return a.Redirect(routes.Agent.ServiceAgent())
					}
					img.IdService = service.IdService
					db := app.GORM.Create(&img)
					if db.Error != nil {
						panic(db.Error)
					}
					return a.Redirect(routes.Agent.GaleryFoto(idService))

				}

			}
		}
		return a.Redirect(routes.App.Index())
	}

	return a.Redirect(routes.App.Index())
}

func (a Agent) FotoService(idService int64) ([]*models.FotoService, error) {
	var foto []*models.FotoService
	db := app.GORM.Where("service_id = ?", idService).Find(&foto)
	if db.Error != nil {
		return nil, db.Error
	}
	return foto, nil
}

func (c Agent) EditDesc(desc string, idService int64) revel.Result {
	agent := c.CheckAgent()
	if agent != nil {
		user := c.connected()
		if user == nil {
			return c.Redirect(routes.App.Index())
		}
		service, err := c.GetServiceWithId(idService)
		if err != nil {
			panic(err)
		}
		service.Desc = desc
		app.GORM.Save(&service)
		return c.Redirect(routes.Agent.Service(idService))
	}
	return c.Redirect(routes.App.Index())
}

func (a Agent) GaleryFoto(idService int64) revel.Result {
	agent := a.CheckAgent()
	if agent != nil {
		var foto []models.FotoService
		db := app.GORM.Where("service_id = ?", idService).Find(&foto)
		if db.RecordNotFound() {
			a.Flash.Success("Nothing Foto in galery")
			a.RenderArgs["agent"] = agent
			a.RenderArgs["foto"] = foto
			return a.Render()
		}
		service, err := a.GetServiceWithId(idService)
		if err != nil {
			panic(err)
		}
		a.RenderArgs["service"] = service
		a.RenderArgs["agent"] = agent
		a.RenderArgs["foto"] = foto
		return a.Render()
	}
	return a.Redirect(routes.App.Index())
}

func (c Agent) Service(idService int64) revel.Result {
	agent := c.CheckAgent()
	if agent != nil {
		user := c.connected()
		if user == nil {
			return c.Redirect(routes.App.Index())
		}
		service, err := c.GetServiceWithId(idService)
		if err != nil {
			panic(err)
		}
		addOn := c.GetAddOn(idService)
		foto, err := c.FotoService(idService)
		if err != nil {
			panic(err)
		}
		c.RenderArgs["service"] = service
		c.RenderArgs["foto"] = foto
		c.RenderArgs["agent"] = agent
		c.RenderArgs["addOn"] = addOn
		return c.Render()
	}
	return c.Redirect(routes.App.Index())
}
func (a Agent) Logout() revel.Result {
	if user := a.connected(); user == nil {
		a.Flash.Error("Please log in first")
		return a.Redirect(routes.App.Login())
	}

	for k := range a.Session {
		delete(a.Session, k)
	}
	return a.Redirect(routes.App.Login())
}
