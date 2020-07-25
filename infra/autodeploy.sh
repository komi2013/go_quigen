#!/bin/bash -e

cd /Project/go
while :
do
FILES=`/usr/bin/find ./ -mmin 0.01`
KEY="/home/komatsu/.ssh/kagoya.key"
PATH="root@quigen.info:/var/www/go/"
# NOW=`/bin/date --date "-2 minutes" +%m%d%H%M`
# echo $FILES
for d in $FILES; do
#     # echo "${d}"
    if [ "`/bin/echo $d | /bin/grep -v .git`" ]; then
        F=`/bin/echo $d | /bin/sed 's/.\///'`
        /usr/local/go/bin/go build
        echo $NOW $F
        /usr/bin/rsync -ave "/usr/bin/ssh -i $KEY" \
            --delete \
            --exclude='log/' \
            $F $PATH$F
        # /bin/touch -t $NOW $F
    fi
#     # /usr/bin/rsync -avzPe "/usr/bin/ssh -i infra/autodeploy.sh root@quigen.info:/tmp/infra/autodeploy.sh
done
# /usr/bin/rsync -ave "/usr/bin/ssh -i $KEY" --delete \
#     --exclude='log/' \
#     ./ $PATH
# /bin/sleep 1

done