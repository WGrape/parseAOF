projectDir=$(dirname $0)
cd $projectDir
projectDir=$(pwd)

# cache disabled
go test --count=1 ./src/global
go test --count=1 ./src/parser
go test --count=1 ./src/writer
