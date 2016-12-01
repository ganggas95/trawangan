// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) SetUser(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.SetUser", args).Url
}

func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Login", args).Url
}

func (_ tApp) Register(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Register", args).Url
}

func (_ tApp) AddUser(
		user interface{},
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("App.AddUser", args).Url
}

func (_ tApp) AddUserWithSosmed(
		user interface{},
		password string,
		verifyPassword string,
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "verifyPassword", verifyPassword)
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("App.AddUserWithSosmed", args).Url
}

func (_ tApp) AuthApp(
		username string,
		password string,
		remember bool,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	revel.Unbind(args, "remember", remember)
	return revel.MainRouter.Reverse("App.AuthApp", args).Url
}

func (_ tApp) SetUp(
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("App.SetUp", args).Url
}

func (_ tApp) RegisterWithFacebook(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.RegisterWithFacebook", args).Url
}

func (_ tApp) LoginFb(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.LoginFb", args).Url
}

func (_ tApp) LoginWithFb(
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("App.LoginWithFb", args).Url
}

func (_ tApp) AuthFb(
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("App.AuthFb", args).Url
}

func (_ tApp) RegisterWithGPlus(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.RegisterWithGPlus", args).Url
}

func (_ tApp) LoginGplus(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.LoginGplus", args).Url
}

func (_ tApp) LoginWithGplus(
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("App.LoginWithGplus", args).Url
}

func (_ tApp) GplusAuth(
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("App.GplusAuth", args).Url
}

func (_ tApp) VerifyAcoount(
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("App.VerifyAcoount", args).Url
}

func (_ tApp) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Logout", args).Url
}

func (_ tApp) UploadFoto(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.UploadFoto", args).Url
}

func (_ tApp) Upload(
		avatar []byte,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "avatar", avatar)
	return revel.MainRouter.Reverse("App.Upload", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tPersons struct {}
var Persons tPersons


func (_ tPersons) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Persons.Index", args).Url
}

func (_ tPersons) List(
		search string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "search", search)
	return revel.MainRouter.Reverse("Persons.List", args).Url
}

func (_ tPersons) Show(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Persons.Show", args).Url
}

func (_ tPersons) UnverifyAcc(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Persons.UnverifyAcc", args).Url
}

func (_ tPersons) Tambah(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Persons.Tambah", args).Url
}

func (_ tPersons) AddData(
		person interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "person", person)
	return revel.MainRouter.Reverse("Persons.AddData", args).Url
}

func (_ tPersons) Delete(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Persons.Delete", args).Url
}

func (_ tPersons) Ubah(
		nama string,
		alamat string,
		tempatlahir string,
		pekerjaan string,
		tanggallahir interface{},
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "nama", nama)
	revel.Unbind(args, "alamat", alamat)
	revel.Unbind(args, "tempatlahir", tempatlahir)
	revel.Unbind(args, "pekerjaan", pekerjaan)
	revel.Unbind(args, "tanggallahir", tanggallahir)
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Persons.Ubah", args).Url
}

func (_ tPersons) GetData(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("Persons.GetData", args).Url
}

func (_ tPersons) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Persons.Logout", args).Url
}


type tAgent struct {}
var Agent tAgent


func (_ tAgent) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agent.Index", args).Url
}

func (_ tAgent) ServiceAgent(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agent.ServiceAgent", args).Url
}

func (_ tAgent) OrderAgent(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agent.OrderAgent", args).Url
}

func (_ tAgent) ChatAgent(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agent.ChatAgent", args).Url
}

func (_ tAgent) MemberAgent(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agent.MemberAgent", args).Url
}

func (_ tAgent) RegisterAgent(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agent.RegisterAgent", args).Url
}

func (_ tAgent) AddAgentFromUser(
		travelAgent interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "travelAgent", travelAgent)
	return revel.MainRouter.Reverse("Agent.AddAgentFromUser", args).Url
}

func (_ tAgent) SetService(
		agentService interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "agentService", agentService)
	return revel.MainRouter.Reverse("Agent.SetService", args).Url
}

func (_ tAgent) DeleteService(
		idService int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "idService", idService)
	return revel.MainRouter.Reverse("Agent.DeleteService", args).Url
}

func (_ tAgent) ActiveService(
		idService int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "idService", idService)
	return revel.MainRouter.Reverse("Agent.ActiveService", args).Url
}

func (_ tAgent) DisableService(
		idService int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "idService", idService)
	return revel.MainRouter.Reverse("Agent.DisableService", args).Url
}

func (_ tAgent) AddOnService(
		idService int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "idService", idService)
	return revel.MainRouter.Reverse("Agent.AddOnService", args).Url
}

func (_ tAgent) SetAddOn(
		addOn interface{},
		idService int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "addOn", addOn)
	revel.Unbind(args, "idService", idService)
	return revel.MainRouter.Reverse("Agent.SetAddOn", args).Url
}

func (_ tAgent) RemoveAddOn(
		idAddOn int64,
		idService int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "idAddOn", idAddOn)
	revel.Unbind(args, "idService", idService)
	return revel.MainRouter.Reverse("Agent.RemoveAddOn", args).Url
}

func (_ tAgent) UploadPage(
		idService int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "idService", idService)
	return revel.MainRouter.Reverse("Agent.UploadPage", args).Url
}

func (_ tAgent) UploadFoto(
		foto []byte,
		idService int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "foto", foto)
	revel.Unbind(args, "idService", idService)
	return revel.MainRouter.Reverse("Agent.UploadFoto", args).Url
}

func (_ tAgent) EditDesc(
		desc string,
		idService int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "desc", desc)
	revel.Unbind(args, "idService", idService)
	return revel.MainRouter.Reverse("Agent.EditDesc", args).Url
}

func (_ tAgent) GaleryFoto(
		idService int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "idService", idService)
	return revel.MainRouter.Reverse("Agent.GaleryFoto", args).Url
}

func (_ tAgent) Service(
		idService int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "idService", idService)
	return revel.MainRouter.Reverse("Agent.Service", args).Url
}

func (_ tAgent) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Agent.Logout", args).Url
}


type tApi struct {}
var Api tApi


func (_ tApi) Auth(
		username string,
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "username", username)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Api.Auth", args).Url
}

func (_ tApi) AddUser(
		user interface{},
		password string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	revel.Unbind(args, "password", password)
	return revel.MainRouter.Reverse("Api.AddUser", args).Url
}


