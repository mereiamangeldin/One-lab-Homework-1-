start:
	docker-compose up --build

stop:
	docker-compose down
	docker rmi homework1-golang

migrate-up:
	migrate -path ./migrations -database 'postgres://postgres:Merei04977773@@localhost:5434/onelab?sslmode=disable' up