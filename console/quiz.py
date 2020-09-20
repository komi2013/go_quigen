
import time
from selenium import webdriver
import postgresql
import json
import csv
import os
import urllib.request

db = postgresql.open('pq://exam_8099:Zk1CGsBK@localhost:5432/programming')
# ps = db.prepare("SELECT * from c_fe2")
# for d in ps():
#     print(d[0])
#     x = d[0].split("/")
#     print(x)
#     os.makedirs("/Project/go_quigen/public/data/fe/" + x[4], exist_ok=True)
#     os.chdir("/Project/go_quigen/public/data/fe/" + x[4])
#     url = d[0]
#     with urllib.request.urlopen(url) as u:
#         with open(x[6], 'bw') as o:
#             o.write(u.read())


# JSONファイルのロード
arr = json.load(open('/Project/go_quigen/log/e-words.json', 'r'))
# list of dictの抽出

for d in arr:
    print(d[1])
    ps = db.prepare("INSERT INTO c_word (question, choice_0)"+
               "VALUES($1,$2)")
    ps(d[1],d[0])

