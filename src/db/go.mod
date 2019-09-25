module db

go 1.12

require (
	github.com/boltdb/bolt v1.3.1
	github.com/stretchr/testify v1.4.0
	model v0.0.0
)

replace model v0.0.0 => ../model
