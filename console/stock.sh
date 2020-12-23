#!/bin/bash -e

/usr/bin/ssh -f -N -L 5432:quigen.info:5432 -i "~/.ssh/kagoya.key" root@quigen.info

cd /Project/go_quigen/
/usr/local/go/bin/go run console/console.go UpStock
