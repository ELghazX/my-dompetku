package connection

import (
	"fmt"
	"log"

	"github.com/elghazx/my-dompetku/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase(conf config.Database) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable Timezone=%s",
		conf.Host,
		conf.Port,
		conf.Password,
		conf.User,
		conf.Name,
		conf.Tz,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB from gorm.DB:", err.Error())
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("failed to ping database:", err.Error())
	}
	return db
}
