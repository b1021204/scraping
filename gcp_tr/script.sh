#! /bin/bash 
yum install -y httpd
echo "<html><body>Hello World!</body></html> > /var/www/html/index.html
systemctl enable --now httpd