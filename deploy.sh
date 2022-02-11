kill -9 $(lsof -i:8082 -t)
go build -o main
BUILD_ID=DONTKILLME
cd /var/lib/jenkins/workspace/git-wall
nohup ./main &