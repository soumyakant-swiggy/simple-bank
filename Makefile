postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

.PHONY: postgres createdb dropdb