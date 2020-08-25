
import time
from selenium import webdriver

# driver = webdriver.Firefox(executable_path='/Project/go_quigen/infra/geckodriver')
driver = webdriver.Chrome(executable_path='/Project/go_quigen/infra/chromedriver')
driver.get('https://translate.google.com/')
sources = [
    '鰻',
    'エイ',
    '鮪',
    '蟹',
    'パンダ',
    'ジュゴン',
    '狼',
    '蜥',
    'キングコブラ',
    'カメレオン',
    'エリマキトカゲ'
    # 'とびうお',
    # 'なまこ',
    # 'かつお',
    # 'かれい',
    # 'きす',
    # 'さけ・ます',
    # 'さば',
    # 'さより',
    # 'さんま',
    # 'たい',
    # 'かに',
    # 'こんぶ',
    # 'さざえ',
    # 'たこ',
    # 'アボカド',
    # 'アスパラガス',
    # 'メロン',
    # 'もも',
    # 'ようなし',
    # 'いちじく',
    # 'オレンジ',
    # 'いちご',
    # 'えだまめ',
    # 'わさび',
    # '牛乳',
    # 'らっきょう',
    # 'ラディッシュ',
    # 'レタス・サラダな',
    # 'レモン',
    # 'サザエ',
    # 'カニ',
    # 'クリオネ',
    # 'ホタテ',
    # 'ヤドカリ',
    # 'タコ',
    # 'ミミズ',
    # 'イカ',
    # 'カンガルー',
    # '狐',
    # 'アザラシ',
    # 'スイギュウ',
    # 'モグラ',
    # 'アフリカゾウ',
    # 'ビーバー',
    # 'マナティー',
    # 'アライグマ',
    # 'ネコ',
    # 'イヌ',
    # 'イノシシ',
    # 'サイ',
    # 'ヤマアラシ',
    # 'りんご',
    # 'れんこん',
    # '梅',
    # 'オクラ',
    # 'かぼちゃ',
    # 'カリフラワー',
    # 'きゃべつ',
    # 'カバ',
    # 'バター',
    # 'チーズ',
    # '小豆',
    # 'いんげん',
    # 'こまつな',
    # 'しいたけ',
    # 'じゃがいも',
    # 'セロリ',
    # 'たけのこ',
    # 'たまねぎ',
    # 'カピバラ',
    # 'カマイルカ',
    # 'カモノハシ',
    # 'オットセイ',
    # 'キツネ',
    # 'リス',
    # 'サイ',
    # 'シマウマ',
    # 'グリズリー',
    # 'コアラ',
    # 'コウモリ',
    # 'アルマジロ',
    # 'ネズミ',
    # 'ゴールデンハムスター',
    # 'クジラ',
    # '麒麟',
    # 'かえる',
    # 'だいこん',
    # '柿',
    # 'キーウィフルーツ',
    # 'グレープフルーツ',
    # 'すいか',
    # 'なし',
    # 'バナナ',
    # 'にわとり',
    # 'にわとりの卵',
    # 'ひつじ',
    # 'うし',
    # 'あなご',
    # 'あゆ',
    # 'いわし',
    # 'パイナップル',
    # 'パパイア',
    # 'ぶどう',
    # 'マンゴー',
    # 'いわな・やまめ',
    # 'そば',
    # 'ぶた',
    # 'うなぎ',
    # 'あじ',
    # '米',
    # '小麦',
    # 'あかがい',
    # 'あさり',
    # 'いか',
    # 'うに',
    # 'えび',
    # 'かき',
    # 'あおのり・のり',
    # 'えんどう',
    # 'ごま',
    # 'そら豆',
    # 'だいず',
    # 'かぶ',
    # 'パセリ',
    # 'ピーマン',
    # 'ブロッコリー',
    # 'ほうれんそう',
    # 'ほんしめじ',
    # 'マッシュルーム',
    # 'タヌキ',
    # 'チンパンジー',
    # 'トナカイ',
    # '猿',
    # '鹿',
    # 'モモンガ',
    # 'モグラ',
    # 'イルカ',
    # 'ヒツジ',
    # 'ラクダ',
    # '豚',
    # 'ハイエナ',
    # 'トラ',
    # 'ホッキョクグマ',
    # 'ライオン',
    # 'バク',
    # 'マントヒヒ',
    # 'ウォンバット',
    # 'アザラシ',
    # 'コウモリ',
    # 'モルモット',
    # 'ヤギ',
    # 'ラッコ',
    # 'ロバ',
    # 'アメーバ',
    # 'ゾウリムシ',
    # 'ミドリムシ',
    # 'ウミガメ',
    # '蛇',
    # '亀',
    # 'ワニ'
]
time.sleep(10)
for d in sources:
    id = driver.find_element_by_id("source")
    id.send_keys(d)
    # print("before click")
    time.sleep(2)
    driver.find_elements_by_class_name("res-tts")[0].click()

    print(driver.find_element_by_xpath('/html/body/div[2]/div[2]/div[1]/div[2]/div[1]/div[1]/div[2]/div[3]/div[1]/div[2]/div/span[1]/span').text)
    id.clear()


time.sleep(100000000)