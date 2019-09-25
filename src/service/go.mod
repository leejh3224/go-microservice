module service

go 1.12

require (
	db v0.0.0
	github.com/gorilla/mux v1.7.3
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
	model v0.0.0
)

replace (
	db v0.0.0 => ../db
	model v0.0.0 => ../model
)
