
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ] 


run/api:
	go run ./cmd/api


db/psql:
	psql ${GREENLIGHT_DB_DSN}


db/migrations/new:
	migrate create -seq -ext=.sql -dir=./migrations ${name}


db/migrations/up: confirm
	migrate -path ./migrations -database ${GREENLIGHT_DB_DSN} up
