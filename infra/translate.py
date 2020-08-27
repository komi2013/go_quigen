
import time
from selenium import webdriver

# driver = webdriver.Firefox(executable_path='/Project/go_quigen/infra/geckodriver')
driver = webdriver.Chrome(executable_path='/Project/go_quigen/infra/chromedriver')
driver.get('https://translate.google.com/')
sources = [
'カピバラ'
,'カマイルカ'
,'カモノハシ'
,'オットセイ'
,'キツネ'
,'リス'
,'サイ'
,'シマウマ'
,'グリズリー'
,'コアラ'
,'コウモリ'
,'アルマジロ'
,'ネズミ'
,'ゴールデンハムスター'
,'クジラ'
,'麒麟'
,'蛙'
,'タヌキ'
,'チンパンジー'
,'トナカイ'
,'猿'
,'鹿'
,'モモンガ'
,'イルカ'
,'ヒツジ'
,'ラクダ'
,'豚'
,'ハイエナ'
,'虎'
,'ホッキョクグマ'
,'ライオン'
,'バク'
,'マントヒヒ'
,'ウォンバット'
,'アザラシ'
,'コウモリ'
,'モルモット'
,'ヤギ'
,'ラッコ'
,'ロバ'
,'アメーバ'
,'ゾウリムシ'
,'ミドリムシ'
,'ウミガメ'
,'蛇'
,'亀'
,'ワニ'
,'カバ'
]
time.sleep(15)
for d in sources:
    id = driver.find_element_by_id("source")
    id.send_keys(d)
    # print("before click")
    time.sleep(2)
    driver.find_elements_by_class_name("res-tts")[0].click()

    print(driver.find_element_by_xpath('/html/body/div[2]/div[2]/div[1]/div[2]/div[1]/div[1]/div[2]/div[3]/div[1]/div[2]/div/span[1]/span').text)
    time.sleep(5)
    id.clear()


time.sleep(100000000)