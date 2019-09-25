module github.com/leejh3224/micro-go

go 1.12

require (
	db v0.0.0
	model v0.0.0
	server v0.0.0
	service v0.0.0
)

replace (
	db v0.0.0 => ./src/db
	model v0.0.0 => ./src/model
	server v0.0.0 => ./src/server
	service v0.0.0 => ./src/service
)
