if [ "$1" == "start" ]; then
  nohup go run main.go >nohup.out &
elif [ "$1" == "stop" ]; then
  ps -ef | grep main | grep -v grep | awk '{print $1}' | xargs kill -9
elif [ "$1" == "restart" ]; then
  ps -ef | grep main | grep -v grep | awk '{print $1}' | xargs kill -9
  nohup go run main.go >nohuo.out &
fi

if [ $? -eq 0 ]; then
    echo "[$1]服务执行成功,干得漂亮(*￣︶￣)"
else
    echo "[$1]执行失败,请重试(╥╯^╰╥)"
fi



