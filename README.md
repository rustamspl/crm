# crm
B-APPS CRM

This is First Readme File

#Dependences

github.com/beego/bee<br /> 
github.com/astaxie/beego<br /> 
github.com/go-sql-driver/mysql<br /> 
github.com/astaxie/beego/orm<br /> 
golang.org/x/crypto<br /> 
github.com/Shopify/go-lua<br /> 
github.com/tealeg/xlsx<br /> 



#Quick install

\#mysql -u root -p<br />

CREATE USER 'crm'@'localhost' IDENTIFIED BY 'crm';GRANT ALL PRIVILEGES ON *.* TO 'crm'@'localhost' IDENTIFIED BY 'crm' WITH GRANT OPTION MAX_QUERIES_PER_HOUR 0 MAX_CONNECTIONS_PER_HOUR 0 MAX_UPDATES_PER_HOUR 0 MAX_USER_CONNECTIONS 0;GRANT ALL PRIVILEGES ON `golang`.* TO 'crm'@'localhost';<br />

Ctrl+D

cd /usr/local/go/src/github.com/<br />
mkdir beego<br />
cd beego<br />
git clone http://github.com/beego/bee.git<br />
cd ..<br />

mkdir astaxie<br />
cd astaxie<br />
git clone github.com/astaxie/beego.git<br />
cd ..<br />

mkdir go-sql-driver<br />
cd go-sql-driver<br />
git clone http://github.com/go-sql-driver/mysql.git<br />
cd ..<br />

mkdir x<br />
cd x<br />
git clone http://golang.org/x/crypto.git<br />
cd ..<br />

mkdir Shopify<br />
cd Shopify<br />
git clone github.com/Shopify/go-lua.git<br />
cd ..<br />

mkdir tealeg<br />
cd tealeg<br />
git clone github.com/tealeg/xlsx.git<br />
cd ..<br />




