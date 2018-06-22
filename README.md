# onRead在线阅读(服务端)
- 使用golang语言
- 基于Beego RestFul API前后端分离模式
- 使用mysql数据库

## 特性
- RestFul API
- Access Token, User Auth


## 项目启动
- 项目根目录下执行命令： bee run -gendoc=true -downdoc=true
- -gendoc=true 表示每次自动化的 build 文档，-downdoc=true 就会自动的下载 swagger 文档查看器
- 浏览器登录  http://127.0.0.1:8080/swagger/#!/customer  或  http://localhost:8080/swagger/#!/customer
