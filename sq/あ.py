from selenium import webdriver
from selenium.webdriver.common.keys import Keys
import chromedriver_binary
import pickle

#chromedriver
driver = None

def main():

   global driver

   driver = webdriver.Chrome()
   driver.get('https://sso.fun.ac.jp/my.policy') #対象ログインページ
   
   print("ログインしたら何か入力してください。")
   val = input()

   pickle.dump(driver.get_cookies() , open("cookies.pkl","wb"))

   driver.close()


if __name__ == '__main__':
   main()