package route

import (
	"myAPI/database"
	"myAPI/pkg/module/activity"
	"myAPI/pkg/module/user"
	"myAPI/pkg/worker"
	"net/http"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouter(mux *http.ServeMux, db *gorm.DB, mongoDB *mongo.Database, emailChan chan<- worker.EmailJob, redis *redis.Client) {

	userRepo := user.NewRepository(db)
	txManager := database.NewTxManager(db)
	activityRepo := activity.NewRepository(mongoDB)
	activitySvc := activity.NewService(activityRepo, redis)

	userService := user.NewService(userRepo, txManager, activitySvc)
	UserRoute(mux, userService, emailChan)
	ActivityLogRoute(mux, activitySvc)
}
