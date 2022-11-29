
# centos7.9 install nginx 


r=$(which apt)
if [[ $? == 0 ]];then
        # ubuntu 
        apt-get install gcc g++ make libpcre3 libpcre3-dev zlib1g zlib1g-dev libssl-dev -y
fi

yum install -y pcre zlib openssl  
yum install -y pcre-devel openssl-devel zlib-devel

./configure
make
make install
