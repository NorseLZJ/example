FROM centos:7

RUN yum update -y \
    && yum upgrade -y \
    && yum install wget -y \
    && cd /etc/yum.repos.d && mv CentOS-Base.repo CentOS-Base.repo.bak \
    && cd /etc/yum.repos.d &&  wget http://mirrors.163.com/.help/CentOS6-Base-163.repo \
    && cd ~ && yum install mysql mysql-server http mod_dav_svn  php php-pear php-mysql subversion php-mbstring -y \
    && mkdir -p /data/svnroot/repos \
    && svnadmin create /data/svnroot/repos/test \
    && chown apache.apache /data/svnroot/ -R 




CMD ["export TERM=linux"]
CMD ["top -bn 1 &>> /dev/null"]

EXPOSE 80 443 