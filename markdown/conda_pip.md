# anaconda 与 pip 加速配置

- anaconda,minconda

在家目录下新建文件 `.condarc`

将以下内容写入此文件,这里使用的是阿里云的加速

```
channels:
  - defaults
show_channel_urls: true
default_channels:
  - http://mirrors.aliyun.com/anaconda/pkgs/main
  - http://mirrors.aliyun.com/anaconda/pkgs/r
  - http://mirrors.aliyun.com/anaconda/pkgs/msys2
custom_channels:
  conda-forge: http://mirrors.aliyun.com/anaconda/cloud
  msys2: http://mirrors.aliyun.com/anaconda/cloud
  bioconda: http://mirrors.aliyun.com/anaconda/cloud
  menpo: http://mirrors.aliyun.com/anaconda/cloud
  pytorch: http://mirrors.aliyun.com/anaconda/cloud
  simpleitk: http://mirrors.aliyun.com/anaconda/cloud
```

执行命令

```
conda clean -i && conda update conda
```

[阿里云官网参考链接](https://developer.aliyun.com/mirror/anaconda?spm=a2c6h.13651102.0.0.5bfd1b110l0fwp)

- pip 加速

几个命令要熟悉下

```
# 设置pip加速为阿里云

pip config set global.index-url https://mirrors.aliyun.com/pypi/simple/

# 查看pip加速是谁

pip config list

```
