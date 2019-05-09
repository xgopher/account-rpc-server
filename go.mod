module app

go 1.12

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.1
	github.com/jinzhu/gorm v1.9.5
	github.com/joho/godotenv v1.3.0
	google.golang.org/genproto v0.0.0-20190425155659-357c62f0e4bb // indirect
	google.golang.org/grpc v1.20.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
