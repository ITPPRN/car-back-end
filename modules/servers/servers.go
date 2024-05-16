package servers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"testBackend/configs"
	"testBackend/logs"
	"testBackend/pkg/utils"
)

type server struct {
	App *fiber.App
	Db  *gorm.DB
	Cfg *configs.Config
}

func NewServer(cfg *configs.Config,
	db *gorm.DB,

) *server {
	return &server{
		App: fiber.New(),
		Db:  db,
		Cfg: cfg,
	}
}

func (s *server) Start() {
	if err := s.Handlers(); err != nil {
		logs.Error(err)
		panic(err.Error())
	}

	fiberConnURL, err := utils.UrlBuilder("fiber", s.Cfg)
	if err != nil {
		logs.Error(err)
		panic(err.Error())
	}

	port := s.Cfg.App.Port
	logs.Info("server started on localhost:", zap.String("port", port))

	if err := s.App.Listen(fiberConnURL); err != nil {
		logs.Error(err)
		panic(err.Error())
	}
}
