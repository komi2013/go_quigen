
import time
from selenium import webdriver
import postgresql
import json
import csv
import os
import urllib.request

# driver = webdriver.Chrome(executable_path='/Project/go_quigen/log/chromedriver')
# driver.get('https://php-archive.net/php/essential-php-functions/')

# i = "7"
# print(driver.find_element_by_css_selector('#post-940 > div > p:nth-child(7)")').text)
# //*[@id="post-940"]/div/p[5]
# //*[@id="post-940"]/div/p[5]/strong
# //*[@id="post-940"]/div/p[6]
# //*[@id="post-940"]/div/p[101]


db = postgresql.open('pq://exam_8099:Zk1CGsBK@localhost:5432/programming')
ps = db.prepare("SELECT * from c_fe2")
for d in ps():
    print(d[0])
    x = d[0].split("/")
    print(x)
#     os.makedirs("/Project/go_quigen/public/data/fe/" + x[4], exist_ok=True)
#     os.chdir("/Project/go_quigen/public/data/fe/" + x[4])
#     url = d[0]
#     with urllib.request.urlopen(url) as u:
#         with open(x[6], 'bw') as o:
#             o.write(u.read())


# word_id
# 1 ~500 e-words
# 500 ~ 1250 otsuka
# 1250 ~ 1345
# question_id 
# 2001 ~ from word

# # JSONファイルのロード
# arr = json.load(open('/Project/go_quigen/log/lpic-study.json', 'r'))
# # list of dictの抽出

# for d in arr:
#     print(d[0])
#     ps = db.prepare("""INSERT INTO c_resource 
#         (choice_0,choice_1,choice_2,choice_3,resource_txt,explanation,category_id)
#         VALUES($1,$2,$3,$4,$5,$6,$7)""")
#     ps(d[0],d[1],d[2],d[3],d[4],d[5],75)

