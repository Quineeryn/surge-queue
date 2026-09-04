package route

import (
	"myAPI/database"
	"myAPI/pkg/module/user"
	"myAPI/pkg/worker"
	"net/http"

	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouter(mux *http.ServeMux, db *gorm.DB, mongoDB *mongo.Database, emailChan chan<- worker.EmailJob) {

	userRepo := user.NewRepository(db)
	txManager := database.NewTxManager(db)
	userService := user.NewService(userRepo, txManager)
	UserRoute(mux, userService, emailChan)
}
