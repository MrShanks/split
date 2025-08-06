migrate-up:
	goose -dir db/migrations postgres "user=split password=split dbname=split sslmode=disable" up

migrate-down:
	goose -dir db/migrations postgres "user=split password=split dbname=split sslmode=disable" down
