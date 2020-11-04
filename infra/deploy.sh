#!/bin/bash -e

cd /Project/go_quigen/
# arr=("shikaku" "programming" "kids" "english")
arr=("shikaku")
d=`date "+%Y%m%d-%H%M%S"`
for v in "${arr[@]}"
do
    echo "$v"
    cp -f common/conf.$v.env common/config.go
    sed -i -e "s/date/$d/" common/config.go
    /usr/local/go/bin/go build

    /usr/bin/rsync -avzPe "/usr/bin/ssh -i /home/komatsu/.ssh/kagoya.key" \
      go_quigen root@quigen.info:/var/www/$v/go_quigen
    /usr/bin/rsync -avzPe "/usr/bin/ssh -i /home/komatsu/.ssh/kagoya.key" \
      html/ root@quigen.info:/var/www/$v/html/
    /usr/bin/rsync -avzPe "/usr/bin/ssh -i /home/komatsu/.ssh/kagoya.key" \
      public/css/ root@quigen.info:/var/www/$v/public/css/
    /usr/bin/rsync -avzPe "/usr/bin/ssh -i /home/komatsu/.ssh/kagoya.key" \
      public/js/ root@quigen.info:/var/www/$v/public/js/
    /usr/bin/rsync -avzPe "/usr/bin/ssh -i /home/komatsu/.ssh/kagoya.key" \
      public/img/ root@quigen.info:/var/www/$v/public/img/
done

# while :
# do
# FILES=`/usr/bin/find ./ -mmin 0.01`
# KEY="/home/komatsu/.ssh/kagoya.key"
# PATH="root@quigen.info:/var/www/go/"
# # NOW=`/bin/date --date "-2 minutes" +%m%d%H%M`
# # echo $FILES
# for d in $FILES; do
# #     # echo "${d}"
#     if [ "`/bin/echo $d | /bin/grep -v .git`" ]; then
#         F=`/bin/echo $d | /bin/sed 's/.\///'`
#         /usr/local/go/bin/go build
#         echo $NOW $F
#         /usr/bin/rsync -ave "/usr/bin/ssh -i $KEY" \
#             --delete \
#             --exclude='log/' \
#             $F $PATH$F
#         # /bin/touch -t $NOW $F
#     fi
# #     # /usr/bin/rsync -avzPe "/usr/bin/ssh -i infra/autodeploy.sh root@quigen.info:/tmp/infra/autodeploy.sh
# done
# # /usr/bin/rsync -ave "/usr/bin/ssh -i $KEY" --delete \
# #     --exclude='log/' \
# #     ./ $PATH
# # /bin/sleep 1

# done
