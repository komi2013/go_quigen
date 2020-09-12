
import time
from selenium import webdriver

# driver = webdriver.Firefox(executable_path='/Project/go_quigen/infra/geckodriver')
driver = webdriver.Chrome(executable_path='/Project/go_quigen/infra/chromedriver')
driver.get('https://www.fe-siken.com/fekakomon.php')

txt = 'nothing'
# try:
#     while True:
#         time.sleep(100000000)
# except KeyboardInterrupt:
#     # driver.get('https://www.fe-siken.com/fekakomon.php')
#     print("what?")
time.sleep(10)

driver.find_elements_by_class_name('submit')[0].click()
time.sleep(2)

while True:

    print(driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[1]').get_attribute('innerHTML') )
    print('komatsu')
    if driver.find_element_by_id('select_a') :
        print(driver.find_element_by_id('select_a').get_attribute('innerHTML') )
        print('komatsu')
        print(driver.find_element_by_id('select_i').get_attribute('innerHTML') )
        print('komatsu')
        print(driver.find_element_by_id('select_u').get_attribute('innerHTML') )
        print('komatsu')
        print(driver.find_element_by_id('select_e').get_attribute('innerHTML') )
        print('komatsu')
    else :
        print(driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[4]/ul/li/ul/li[1]/a/button').get_attribute('innerHTML'))
        print('komatsu')
        print(driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[4]/ul/li/ul/li[2]/a/button').get_attribute('innerHTML'))
        print('komatsu')
        print(driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[4]/ul/li/ul/li[3]/a/button').get_attribute('innerHTML'))
        print('komatsu')
        print(driver.find_element_by_xpath('//*[@id="mainCol"]/div[2]/div[4]/ul/li/ul/li[4]/a/button').get_attribute('innerHTML'))
        print('komatsu')

    print(driver.find_element_by_id('answerChar').get_attribute('innerHTML') )
    print('komatsu')
    print(driver.find_elements_by_class_name('ansbg')[1].get_attribute('innerHTML') )

    driver.find_element_by_xpath('//*[@id="configform"]/div[1]/button[1]').click()
    time.sleep(2)

time.sleep(100000000)