package main

import (
	"log"
	"mqtt_agv_simulator/config"
	"mqtt_agv_simulator/mqtt"
	"mqtt_agv_simulator/services"
	"mqtt_agv_simulator/state"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 1. 고정 사양 정보 초기화
	services.InitFactsheet()

	// 2. 설정 로드
	cfg := config.LoadConfig()

	// 3. MQTT 클라이언트 생성 및 연결
	client := mqtt.NewClient(cfg)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 4. 핵심 로직 고루틴 실행
	go state.PublishingLoop(client)        // 주기적인 State 발행 루프
	go services.AgvLogicController(client) // 주문/액션 처리 제어 루프

	// 5. 정상 종료 처리
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("프로그램을 종료합니다. 'OFFLINE' 상태를 발행합니다.")
	services.PublishConnectionState(client, "OFFLINE", false) // 정상 종료 메시지 발행
	client.Disconnect(250)
}
