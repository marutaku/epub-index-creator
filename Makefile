.PHONY: migrate ent_create gen_migration migrate goa_generate run test

goa_generate:
	goa gen github.com/marutaku/epub-index-creator/design

create_model:
	go run -mod=mod entgo.io/ent/cmd/ent new $(name)

ent_generate:
	go generate ./ent

gen_migration:
	go run -mod=mod ent/migrate/main.go ${migration_name}

migrate:
	atlas migrate apply --dir "file://ent/migrate/migrations" --url "sqlite://${db_name}"

dev:
	air -c .air.toml

test:
	go test -v ./...
