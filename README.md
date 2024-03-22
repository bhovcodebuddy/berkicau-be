# berkicau-be
twitter knockoff


#create migration file
migrate create -ext sql -dir migrations/ -seq {tableName} *plural
ex: migrate create -ext sql -dir migrations/ -seq users

#run migration
migrate -path migrations -database postgres://postgres:Passw0rd@localhost:5432/berkicau?sslmode=disable up