.PHONY: build-db
build-db:
	podman build --force-rm=true -f db/Containerfile -t waffle_db db

.PHONY: db-up
db-up: build-db
	podman run --replace -dt -e POSTGRES_HOST_AUTH_METHOD=trust --name waffle_db -p 5432:5432 localhost/waffle_db

.PHONY: dump-schema
dump-schema:
	pg_dump -c -s -U postgres -h localhost waffle > db/structure.sql
	pg_dump -a -t migrations -U postgres -h localhost waffle > db/migrations.sql

# Run using 'make new-migration NAME=<name>'
.PHONY: new-migration
new-migration:
	sql-migrate new -config=dbconfig.yml -env=development $(NAME)

.PHONY: migrate
migrate: apply-migrations dump-schema

.PHONE: apply-migrations
apply-migrations:
	sql-migrate up -config=dbconfig.yml -env=development