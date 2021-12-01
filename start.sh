projectDir=$(dirname $0)
cd $projectDir
projectDir=$(pwd)
mergedFile=$projectDir/data/aof.merged
binFIle=$projectDir/bin/parseAOF
step=0
version="0.5.0"
splitLines=1000
header="--------------------parseAOF | version="$version"--------------------"

# delete the generated files before start
if [ -f "$mergedFile" ]; then
    rm $mergedFile
fi

if [ -f "$binFIle" ]; then
    rm $binFIle
fi

# make the aof file split into many sub files
if [ ! -n "$1" ] ;then
    echo "please input the aof file"
    exit
fi

((step++))
echo [$step] split the aof file: $1
split -l $splitLines ${1} $projectDir/data/aof.split_

if [ $? -ne 0 ]; then
    echo -e "\033[31;4mfailed\033[0m"
    exit
else
    echo -e "\033[32msuccess\033[0m"
fi

# build project
((step++))
if [ ! -f "$binFIle" ]; then
  echo [$step] the bin file not exists, building ...
  cd $projectDir/src

  if [ "$(uname)" == "Darwin" ]; then
     CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o parseAOF .
  elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
     CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o parseAOF .
  else
     go build -o parseAOF .
  fi

  mv parseAOF $projectDir/bin/

  if [ $? -ne 0 ]; then
    echo -e "\033[31;4mfailed\033[0m"
    exit
  else
    echo -e "\033[32msuccess\033[0m"
  fi

else
  echo [$step] the bin file exists, skip build
fi

# run project
((step++))
echo [$step] run the bin file
$binFIle
if [ $? -ne 0 ]; then
  echo -e "\033[31;4mfailed\033[0m"
  exit
else
  echo -e "\033[32msuccess\033[0m"
fi

# merge the parsed files
((step++))
echo [$step] merge the parsed file
cd $projectDir/data/
i=0
for file in $(ls aof.split_*.parsed)
do
    ((i++))
    if [ "$i" -eq "1" ];then
      echo $header >> $mergedFile
    fi
    cat $file >> $mergedFile
    # echo "\n" >> $mergedFile
    rm $file
done
echo -e "merged file: " $mergedFile
if [ $? -ne 0 ]; then
  echo -e "\033[31;4mfailed\033[0m"
  exit
else
  echo -e "\033[32msuccess\033[0m"
fi
