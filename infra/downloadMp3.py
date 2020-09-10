
import time
from selenium import webdriver

# driver = webdriver.Firefox(executable_path='/Project/go_quigen/infra/geckodriver')
driver = webdriver.Chrome(executable_path='/Project/go_quigen/infra/chromedriver')
driver.get('https://soundoftext.com/')
sources = [
'жираф'
,'белка летяга'
,'крот'
,'Дельфин'
,'салат'
,'Эвглена зелёная'
,'Инфузория туфелька'
,'Скат'
]
time.sleep(10)
for d in sources:
    # id = driver.find_element_by_id("source")
    # id.send_keys(d)
    driver.find_elements_by_class_name("field__textarea")[0].send_keys(d)
    # print("before click")
    time.sleep(5)
    # time.sleep(2000)
    driver.find_elements_by_class_name("field__submit")[0].click()

    # print(driver.find_element_by_xpath('/html/body/div[2]/div[2]/div[1]/div[2]/div[1]/div[1]/div[2]/div[3]/div[1]/div[2]/div/span[1]/span').text)
    time.sleep(5)
    driver.find_elements_by_class_name("field__textarea")[0].clear()


time.sleep(100000000)