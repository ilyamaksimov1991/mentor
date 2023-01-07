rup:
	docker-compose  rm -f # удаление остановленных контейнеров
	docker rmi my_gmb_good_morning_bot # удаление строго образа productfilter-worker если он был
	docker-compose up # запуск всех образов

up:
	docker-compose up --remove-orphans
