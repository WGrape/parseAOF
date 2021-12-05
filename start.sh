projectDir=$(dirname $0)
cd $projectDir
projectDir=$(pwd)
mergedFile=$projectDir/data/aof.merged
binFIle=$projectDir/bin/parseAOF
step=0
version="0.5.0"
splitLines=100000
splitSize=100kb
aofFileMaxSize=1 # 1G: allow max size of aof file
srt=$(date +%s)
header="--------------------parseAOF | version="$version"--------------------"

# delete the generated files before start
if [ -f "$mergedFile" ]; then
    rm $mergedFile
fi

if [ -f "$binFIle" ]; then
    rm $binFIle
fi

# check the aof file and split it into many sub files
if [ ! -n "$1" ] ;then
    echo "please input the aof file"
    exit
fi
aofFileSize=$(ls -l $1 | awk '{print $5}')
aofFileSize=`expr $aofFileSize / 1024 / 1024 / 1024`
if [ "$aofFileSize" -gt "$aofFileMaxSize" ];then
  echo "allow max size of aof file: "$aofFileMaxSize"G""(${aofFileSize}G)"
  exit
fi

((step++))
echo [$step] Start splitting ...
split -a 5 -b $splitSize ${1} $projectDir/data/aof.split_
#split -a 3 -l $splitLines ${1} $projectDir/data/aof.split_
if [ $? -ne 0 ]; then
    echo -e "\033[31;4mfailed\033[0m"
    exit
else
    echo -e "Input AOF file: "$1
    for file in $(ls $projectDir/data/aof.split_*)
    do
        ((i++))
        if [ "$i" -gt "10" ];then
          echo -e "... ..."
          break
        fi
        echo -e "Split file $i: "$file
    done
    echo -e "\033[32msuccess\033[0m"
fi

# start building
((step++))
if [ ! -f "$binFIle" ]; then
  echo [$step] Start building ...
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

# start running
((step++))
echo [$step] Start running ...
$binFIle
if [ $? -ne 0 ]; then
  echo -e "\033[31;4mfailed\033[0m"
  exit
else
  echo -e "\033[32msuccess\033[0m"
fi

# merge the parsed files
((step++))
echo [$step] Start merging ...
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
echo -e "The merged file: " $mergedFile
if [ $? -ne 0 ]; then
  echo -e "\033[31;4mfailed\033[0m"
  exit
else
  echo -e "\033[32msuccess\033[0m"
fi

end=$(date +%s)
echo "[5] End: $[$end-$srt]s"
