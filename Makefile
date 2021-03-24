run:
	go run cmd/app/main.go

migrate:
	migrate -path migrations -database "mysql://magdv_test:q1w2e3r4@tcp(localhost:3306)/magdv_test" up