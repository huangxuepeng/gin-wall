kill -9 $(lsof -i:8082 -t)
swag init
go build -o main
BUILD_ID=DONTKILLME
nohup ./main &