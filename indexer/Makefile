.PHONY: migrate ent_create gen_migration migrate

ent_create:
	go run -mod=mod entgo.io/ent/cmd/ent new $(name)

ent_generate:
	go generate ./ent

gen_migration:
	atlas migrate diff ${migration_name} --dir "file://ent/migrate/migrations" --to "file://schema.hcl" --dev-url "sqlite://file?mode=memory"

migrate:
	atlas migrate apply --dir "file://ent/migrate/migrations" --url "sqlite://${db_name}"
