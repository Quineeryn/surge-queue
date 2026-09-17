package route

import (
	"myAPI/database"
	"myAPI/pkg/module/activity"
	"myAPI/pkg/module/auth"
	"myAPI/pkg/module/user"
	"myAPI/pkg/security"
	"myAPI/pkg/worker"
	"net/http"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/mongo"
)

func InitRouter(mux *http.ServeMux, db *gorm.DB, mongoDB *mongo.Database, emailChan chan<- worker.EmailJob, redis *redis.Client, jwtManager *security.JWTManager) {

	userRepo := user.NewRepository(db)
	txManager := database.NewTxManager(db)
	activityRepo := activity.NewRepository(mongoDB)

	activitySvc := activity.NewService(activityRepo, redis)
	userService := user.NewService(userRepo, txManager, activitySvc, redis)
	authService := auth.NewService(userService, jwtManager)

	UserRoute(mux, userService, emailChan, jwtManager, redis)
	AuthRoute(mux, authService)
	ActivityLogRoute(mux, activitySvc, jwtManager)
}
