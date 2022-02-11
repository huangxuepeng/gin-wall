kill -9 $(lsof -i:8082 -t)
go build -o main
BUILD_ID=DONTKILLME
source /var/lib/jenkins/workspace/git-wall
nohup ./main &