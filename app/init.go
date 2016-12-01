package app

import (
	"fmt"
	"github.com/ganggas95/trawangan/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
)

var GORM *gorm.DB

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}
	// register startup functions with OnAppStart
	// ( order dependent )
	revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

func InitDB() {
	var err error
	DB_USER := revel.Config.StringDefault("db.user", "")
	DB_PASS := revel.Config.StringDefault("db.pass", "")
	DB_NAME := revel.Config.StringDefault("db.name", "")
	DB_HOST := revel.Config.StringDefault("db.host", "")

	GORM, err = gorm.Open(revel.Config.StringDefault("db.driver", ""), fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PASS, DB_NAME))
	if err != nil {
		revel.INFO.Println("Db Error ", err)
	}
	GORM.SingularTable(true)
	GORM.AutoMigrate(&models.AgentTravel{},
		&models.User{},
		&models.UserToken{},
		&models.Person{},
		&models.AgentService{},
		&models.UserFoto{},
		&models.FotoService{},
		&models.AddOnService{},
	)
	GORM.LogMode(true)
	revel.INFO.Println("Db COnnected")
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
