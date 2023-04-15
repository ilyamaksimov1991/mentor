rup:
	docker-compose  rm -f # удаление остановленных контейнеров
	docker rmi my_gmb_good_morning_bot # удаление строго образа productfilter-worker если он был
	docker-compose up # запуск всех образов
up:
	docker-compose up --remove-orphans
test:
	go test -coverprofile=coverage.out.tmp ./...
	cat coverage.out.tmp | grep -v _mock.go > coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
