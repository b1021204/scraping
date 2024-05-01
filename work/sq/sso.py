#https://qiita.com/katoon081/items/ab805e7a6ffce855217e

from selenium.webdriver.common.by import By
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
import pandas as pd 
import csv
import pickle
import time

USER ='b1021204'
PASS = 'EPa6ouQ2'

"""
ブラウザを表示させないようにするコマンド。
デバックのため、コメントアウト中
必要ならばwebdriver.Chromeの引数にoptionsを追記
options = Options()
options.add_argument('--headless')
"""

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

# go to page of vm
elem_vm = browser.find_element(By.ID, "/Common/Manage_P")
elem_vm.click()

# go to  page of vm info 
browser.switch_to.window(browser.window_handles[1])
elem_select = browser.find_element(By.XPATH, "/html/body/div/div/main/div/form/div[2]/div/span")
elem_select.click()

elem_ip = browser.find_element(By.ID, "copiable-ip_address-0")
print("adress : " + elem_ip.text)
"""
#取得したい文字列の要素
elm = browser.find_element(By.CLASS_NAME, "webtop2")
# 要素のテキスト内容を取得
num = [elm.text]
print(elm)
header = [""]
df = pd.DataFrame({
    '成果':num
}, index = header)
df.to_csv("test.csv", encoding="UTF-8")

"""
print("終了")
browser.close()