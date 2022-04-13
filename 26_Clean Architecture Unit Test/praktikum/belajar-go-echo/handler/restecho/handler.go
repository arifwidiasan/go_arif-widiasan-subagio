package restecho

import (
	"belajar-go-echo/config"
	"belajar-go-echo/database"
	"belajar-go-echo/domain"

	//m "belajar-go-echo/middleware"
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
	"belajar-go-echo/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.ConnectDB()
	_ = NewData()
	repo := repository.NewMysqlRepository(db)

	svc := service.NewServiceUser(repo, conf)

	cont := EchoController{
		svc: svc,
	}

	apiUser := e.Group("/users",
		middleware.Logger(),
		middleware.CORS(),
		//m.APIKEYMiddleware,
	)

	apiUser.GET("", cont.GetUsersController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.POST("/login", cont.LoginUserController)
	apiUser.POST("", cont.CreateUserController, middleware.JWT([]byte(conf.JWT_KEY)))
}

type Datas struct {
}

func (d *Datas) CreateUsers(user model.User) error {
	panic("impl")
}
func (d *Datas) GetAll() []model.User {
	panic("impl")
}
func (d *Datas) GetOneByEmail(email string) (user model.User, err error) {
	panic("impl")
}
func NewData() domain.AdapterRepository {
	return &Datas{}
}
