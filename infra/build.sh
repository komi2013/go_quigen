#!/bin/bash -e
cd /var/www/go
while :
do

FILES=`find ./app/ -mmin 0.01`
# NOW=`/bin/date --date "-2 minutes" +%m%d%H%M`
# echo $FILES
MODIFIED=0
for d in $FILES; do
    MODIFIED=1
    # echo $MODIFIED
    # if [ "`/bin/echo $d | /bin/grep -v .git`" ]; then
    #     F=`/bin/echo $d | /bin/sed 's/.\///'`
    #     echo $F $NOW
    #     # touch -t $NOW $F
    # fi
done
if [ $MODIFIED -eq 1 ]; then
    echo "restart"
    PS=`ps aux | grep main | grep -v grep | awk 'NR==1{print $2}'`
    if [ -n "$PS" ]; then
        kill $PS
    fi
    PS=`ps aux | grep main | grep -v grep | awk 'NR==1{print $2}'`
    if [ -n "$PS" ]; then
        kill $PS
    fi
    # pkill main
    /usr/local/go/bin/go run main.go >> log/main.log &
    sleep 1
fi

done