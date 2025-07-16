package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config 구조체는 MQTT 연결 설정을 저장합니다.
type Config struct {
	Broker   string
	Port     string
	ClientID string
	Username string
	Password string
}

// LoadConfig 함수는 .env 파일에서 환경 변수를 로드하여 Config 구조체를 반환합니다.
func LoadConfig() *Config {
	// .env 파일 로드
	if err := godotenv.Load(); err != nil {
		log.Println("경고: .env 파일을 찾을 수 없습니다.")
	}

	return &Config{
		Broker:   getEnv("MQTT_BROKER", "localhost"),
		Port:     getEnv("MQTT_PORT", "1883"),
		ClientID: getEnv("MQTT_CLIENT_ID", "go-mqtt-client"),
		Username: getEnv("MQTT_USERNAME", ""),
		Password: getEnv("MQTT_PASSWORD", ""),
	}
}

// getEnv 함수는 환경 변수를 읽어오고, 값이 없을 경우 기본값을 반환합니다.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
