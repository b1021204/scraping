#https://magazine.techacademy.jp/magazine/26426
import urllib.request
from bs4 import BeautifulSoup

url = 'https://techacademy.jp/magazine/'
ua = 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) '
'AppleWebKit/537.36 (KHTML, like Gecko) '
'Chrome/55.0.2883.95 Safari/537.36 '

req = urllib.request.Request(url, headers={'User-Agent': ua})
html = urllib.request.urlopen(req)
soup = BeautifulSoup(html, "html.parser")
topicsindex = soup.find('div')
print(topicsindex)
print('\n\n')
topics = topicsindex.find_all('li')
for topic in topics:
  print(topic.find('a').contents[0])