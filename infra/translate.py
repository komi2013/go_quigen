
import time
from selenium import webdriver

# driver = webdriver.Firefox(executable_path='/Project/go_quigen/infra/geckodriver')
driver = webdriver.Chrome(executable_path='/Project/go_quigen/infra/chromedriver')
driver.get('https://translate.google.com/')
sources = [
'education quiz for 3 years old'
,'creature (3 years old)'
,'reptile (3 years old)'
,'bird (3 years old)'
,'creature under the water'
,'Insect (3 years old)'
,'mammalian(3 years old)'
,'food (3 years old)'
,'dairy (3 years old)'
,'carbohydrate (3 years old)'
,'meat (3 years old)'
,'fish food (3 years old)'
,'sea food'
,'beans'
,'vegetable'
,'fruits'
,'fish (3 years old)'
,'microorganism'
,'green vegetable'
,'non green vegetable'
,'pet'
,'livestock animal'
,'big animal'
,'small animal'
,'4 legs animal'
,'easy and simple picture quiz especially for 3 years old'
,'it is common creature quiz for 3 years old kids'
,'to understand reptile for 3 years old kids'
,'some birds which 3 years old can recognize'
,'creature which live under the water such as crab or clay fish'
,'insects such as worm'
,'mammalian such as lion, cat, dog'
,'food quiz which is easy for 3 years old kids'
,'dairy such as milk butter'
,'carbohydrate such as rice bread'
,'meat quiz about pork beef for 3 years old to memorize easily'
,'ocean fish river fish quiz which 3 years old recognize'
,'easy quiz about sea food for 3 years old kids'
,'beans such as soy beans edamame'
,'common vegetable and understandable for 3 years old kids'
,'fruits such as apple banana'
,'this category is about some kinds of fishes'
,'for 3 years old, easy to answer quiz about microorganism'
,'in vegetable, take all green color vegetable'
,'divide green color and non green color '
,'it is common pet such as cat dog'
,'animal which is livestock'
,'big animal and animal that doesn’t walk with 4 legs'
,'relatively small animal'
,'divide animal that walk with 4 legs not walk with 4 legs within big animal'
]
time.sleep(15)
for d in sources:
    id = driver.find_element_by_id("source")
    id.send_keys(d)
    # print("before click")
    time.sleep(2)
    # driver.find_elements_by_class_name("res-tts")[0].click()

    print(driver.find_element_by_xpath('/html/body/div[2]/div[2]/div[1]/div[2]/div[1]/div[1]/div[2]/div[3]/div[1]/div[2]/div/span[1]/span').text)
    time.sleep(5)
    id.clear()


time.sleep(100000000)