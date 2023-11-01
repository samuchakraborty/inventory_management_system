createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path database/migration -database "mysql://root:new_password@tcp(127.0.0.1:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local" -verbose up

migratedown:
	migrate -path database/migration -database "mysql://root:new_password@tcp(127.0.0.1:3306)/inventory?charset=utf8mb4&parseTime=True&loc=Local" -verbose down

generate: 
	sqlboiler mysql -c conf.d/sqlboiler.toml -o models --no-tests --no-hooks

