DB_DSN=host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable

up:
	goose -dir migrations postgres "$(DB_DSN)" up

down:
	goose -dir migrations postgres "$(DB_DSN)" down

status:
	goose -dir migrations postgres "$(DB_DSN)" status

create:
	@read -p "Migration name: " name; \
	goose -dir migrations create $$name sql