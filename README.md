#### **some program by golang or cpp**

- cmd/fileserver 很简单的文件服务器，提供最简单的上传下载功能，
    - 需要通过 share.json 配置程序端口，（上传，下载）文件夹

- cmd/jjson 对单个 json 文件格式化，或者整个文件夹下的 json文件格式化
    - 通过命令行 -f xxx.json or -d dir 进行

- cmd/reptile 爬取链家房屋（在售，售出）信息，
    - 通过配置文件配置 爬取的 城市 以及 区县 和MySQL数据库，先存入 MySQL再导出到Excel 做数据分析
