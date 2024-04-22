#https://qiita.com/katoon081/items/ab805e7a6ffce855217e

from selenium.webdriver.common.by import By
from selenium import webdriver
import pandas as pd 
import csv

USER ='b1021204'
PASS = 'EPa6ouQ2'

browser = webdriver.Chrome()
# login page
browser.get('https://sso.fun.ac.jp/my.policy')

# sent ID and pass
'''
    htmlを除いた結果、label_input_1,2がそれぞれユーザーネームとパスを
    格納するもののnameらしい
    
    <label for="input_2" id="label_input_2">パスワード</label>　より

    <tr id="submit_row">
        <td class="credentials_table_unified_cell"><input type="submit" class="credentials_input_submit" value="ログイン"></td>
        
    </tr>
'''
elem_username = browser.find_element(By.CSS_SELECTOR, "label_input_1")
elem_password = browser.find_element(By.CSS_SELECTOR, "label_input_2")
browser_form = browser.find_element(By.CSS_SELECTOR, "credentials_input_submit")
#press submit
elem_username.send_keys(USER)
elem_username.Send_keys(PASS)
browser_from.click()
print("ログイン成功です!")
print("\n\n")
'''
#取得したい文字列の要素
elm = browser.find_element(By.CSS_SELECTOR, "")
# 要素のテキスト内容を取得
num = [elm.text]
header = ["タイトル"]
df = pd.DataFrame({
    '成果':num
}, index = header)
df.to_csv("test.csv", encoding="UTF-8")

'''
#browser.close()