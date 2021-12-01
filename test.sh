projectDir=$(dirname $0)
cd $projectDir
projectDir=$(pwd)

# cache disabled
go test --count=1 ./src/global
result=`go test --count=1 ./src/global 2>&1 | grep 'failed'`
if [ -n "$result" ]; then
    echo "==========You failed the test=========="
    exit 1
fi

go test --count=1 ./src/parser
result=`go test --count=1 ./src/parser 2>&1 | grep 'failed'`
if [ -n "$result" ]; then
    echo "==========You failed the test=========="
    exit 1
fi

go test --count=1 ./src/writer
result=`go test --count=1 ./src/writer 2>&1 | grep 'failed'`
if [ -n "$result" ]; then
    echo "==========You failed the test=========="
    exit 1
fi
