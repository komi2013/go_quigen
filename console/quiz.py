
import time
from selenium import webdriver
import postgresql
import json
import csv

# JSONファイルのロード
json_dict = json.load(open('/Project/go_quigen/console/basicInfo.json', 'r'))
# list of dictの抽出
target_dicts = json_dict
# print(target_dicts)
i = 0
for d in json_dict:
    print(d[1])
    print(i)
    i = i + 1

# with open('/Project/go_quigen/console/basicInfo.csv', 'w') as f:
#     # dialectの登録
#     csv.register_dialect('dialect01', doublequote=True, quoting=csv.QUOTE_ALL)
#     # DictWriter作成
#     writer = csv.DictWriter(f, fieldnames=target_dicts[0].keys(), dialect='dialect01')
#     # CSVへの書き込み
#     writer.writeheader()
#     for target_dict in target_dicts:
#         writer.writerow(target_dict)


time.sleep(200000)
db = postgresql.open('pq://exam_8099:Zk1CGsBK@194.135.95.163:5432/shikaku')

# driver = webdriver.Firefox(executable_path='/Project/go_quigen/infra/geckodriver')
driver = webdriver.Chrome(executable_path='/Project/go_quigen/infra/chromedriver')
driver.get('https://www.fe-siken.com/fekakomon.php')

driver.find_element_by_xpath('//*[@id="modeWrap"]/fieldset[2]/legend/h4/label/input').click()
driver.find_element_by_xpath('//*[@id="optionWrap"]/label[5]/input').click()
driver.find_element_by_xpath('//*[@id="fs2"]/div[4]/button[2]').click()

time.sleep(2)

driver.find_elements_by_class_name('submit')[0].click()

while True:

    r = {}
    choice_0 = ''
    # print(driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[1]').get_attribute('innerHTML') )
    txt = driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[1]').get_attribute('innerHTML')
    print(driver.find_element_by_id('answerChar').get_attribute('innerHTML'))
    answer = driver.find_element_by_id('answerChar').get_attribute('innerHTML')
    if driver.find_element_by_id('select_a') :
        print(driver.find_element_by_id('select_a').get_attribute('innerHTML') )
        arr = []
        if answer == 'ア' :
            choice_0 = driver.find_element_by_id('select_a').get_attribute('innerHTML')
        else :
            arr.append(driver.find_element_by_id('select_a').get_attribute('innerHTML'))
        if answer == 'イ' :
            choice_0 = driver.find_element_by_id('select_i').get_attribute('innerHTML')
        else :
            arr.append(driver.find_element_by_id('select_i').get_attribute('innerHTML'))
        if answer == 'ウ' :
            choice_0 = driver.find_element_by_id('select_u').get_attribute('innerHTML')
        else :
            arr.append(driver.find_element_by_id('select_u').get_attribute('innerHTML'))
        if answer == 'エ' :
            choice_0 = driver.find_element_by_id('select_e').get_attribute('innerHTML')
        else :
            arr.append(driver.find_element_by_id('select_e').get_attribute('innerHTML'))
    else :
        print(driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[4]/ul/li/ul/li[1]/a/button').get_attribute('innerHTML'))
        
        print(driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[4]/ul/li/ul/li[2]/a/button').get_attribute('innerHTML'))
        
        print(driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[4]/ul/li/ul/li[3]/a/button').get_attribute('innerHTML'))
        
        print(driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[4]/ul/li/ul/li[4]/a/button').get_attribute('innerHTML'))
      
    # print(driver.find_elements_by_class_name('ansbg')[1].get_attribute('innerHTML') )
    explain = driver.find_elements_by_class_name('ansbg')[1].get_attribute('innerHTML')
    
    ps = db.prepare("INSERT INTO c_resource (resource_txt, choice_0, choice_1, choice_2, choice_3, resource_sound, explanation)"+
               "VALUES($1,$2,$3,$4,$5,$6,$7)")
    ps(txt,choice_0,arr[0],arr[1],arr[2],answer,explain)
    # //*[@id="mainCol"]/div[2]/div[4]/ul/li[1]/a
    # driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[3]/ul/li[1]/a').click()
    clickable = 0
    try:
        driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[3]/ul/li[1]/a').click()
    except:
        clickable = 1
        print('clickable = 1')
    if clickable > 0 :
        try:
            driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[4]/ul/li[1]/a').click()
        except:
            clickable = 2
            print('clickable = 2')

    # answer1 = driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[3]/ul/li[1]/a')
    # answer2 = driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[4]/ul/li[1]/a')
    # if len(answer1) > 0:
    #     answer1.click()
    # elif len(answer2) > 0:
    #     answer2.click()
    #configform > div.bottomBtns > button.submit
    # time.sleep(1)
    # driver.find_element_by_xpath('//*[@id="configform"]/div[1]/button[1]').click()
    driver.find_element_by_css_selector('#configform > div.bottomBtns > button.submit').click()
    # driver.find_element_by_class_name('submit').click()
    # driver.get('https://www.fe-siken.com/fekakomon.php')
    time.sleep(2)

time.sleep(100000000)



# # const DB_CONNECT = "user=exam_8099 password=Zk1CGsBK dbname=shikaku sslmode=disable port=5432 host=194.135.95.163"
# for d in db.prepare("SELECT * FROM m_category_name"):
#     print(d["category_name"])