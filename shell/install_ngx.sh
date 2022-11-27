
# centos7.9 install nginx 


yum install -y pcre zlib openssl  
yum install -y pcre-devel openssl-devel zlib-devel

./configure
make
make install
