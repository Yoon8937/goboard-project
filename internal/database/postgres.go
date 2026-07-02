package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB() (*gorm.DB, error) {
	// 1. .env 파일 로드 (로컬 환경에서 .env 파일을 읽어옵니다)
	// 만약 배포 환경(도커, 쿠버네티스 등)에서 시스템 환경 변수를 직접 주입한다면
	// 이 로드 과정에서 에러가 날 수 있으므로 에러 처리는 생략하거나 로그만 남깁니다.
	_ = godotenv.Load()

	// 2. 환경 변수에서 값 읽어오기 (값이 없을 때를 대비한 기본값 세팅)
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "cherry")
	password := getEnv("DB_PASSWORD", "password123")
	dbname := getEnv("DB_NAME", "board_db")
	port := getEnv("DB_PORT", "5432")
	timezone := getEnv("DB_TIMEZONE", "Asia/Seoul")

	// 3. 안전하게 DSN 문자열 조립
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host, user, password, dbname, port, timezone,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// 환경 변수가 비어있을 때 기본값을 반환해주는 헬퍼 함수
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
