#https://qiita.com/katoon081/items/ab805e7a6ffce855217e

from selenium.webdriver.common.by import By
from selenium import webdriver
import pandas as pd 
import csv
import pickle

USER ='b1021204'
PASS = 'EPa6ouQ2'


browser = webdriver.Chrome()
# login page
browser.get('https://sso.fun.ac.jp')

# sent ID and pass
'''
    htmlを除いた結果、label_input_1,2がそれぞれユーザーネームとパスを
    格納するもののnameらしい
    
    <label for="input_2" id="label_input_2">パスワード</label>　より

    <tr id="submit_row">
        <td class="credentials_table_unified_cell"><input type="submit" class="credentials_input_submit" value="ログイン"></td>
        
    </tr>
'''
elem_username = browser.find_element(By.NAME, "username")
elem_password = browser.find_element(By.NAME, "password")
browser_form = browser.find_element(By.CLASS_NAME, "credentials_input_submit")
#press submit
elem_username.send_keys(USER)
elem_password.send_keys(PASS)


browser_form.click()
'''
cookies = driver.get_cookies('https://sso.fun.ac.jp/my.policy') # クッキーを取得する
pickle.dump(cookies,open(cookies_file,'wb')) # クッキーを保存する
'''
print("ログイン成功です!")
print("\n\n")
#取得したい文字列の要素
elm = browser.find_element(By.CLASS_NAME, "webtop2")
# 要素のテキスト内容を取得
num = [elm.text]
header = [""]
df = pd.DataFrame({
    '成果':num
}, index = header)
df.to_csv("test.csv", encoding="UTF-8")

browser.close()