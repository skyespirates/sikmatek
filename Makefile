migrations_path=./migrations
dsn="mysql://root:secret@tcp(localhost:3306)/sikmatek"

.PHONY: api/start
api/start:
	@go run .\cmd

.PHONY: api/dev
api/dev:
	@air

.PHONY: db/migrations/new
db/migrations/new:
	@migrate create -ext sql -dir $(migrations_path) -seq $(name)

.PHONY: db/migrations/up
db/migrations/up:
	@migrate -database $(dsn) -path $(migrations_path) up

.PHONY: db/migrations/down
db/migrations/down:
	@migrate -database $(dsn) -path $(migrations_path) down

.PHONY: db/migrations/version
db/migrations/version:
	@migrate -database $(dsn) -path $(migrations_path) version

.PHONY: db/migrations/force
db/migrations/force:
	@migrate -database $(dsn) -path $(migrations_path) force $(version)