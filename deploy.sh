kill -9 $(lsof -i:8080 -t)
go build -o main
BUILD_ID=DONTKILLME
nohup ./main &