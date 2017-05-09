# cracker-heroku

Deploy the server side of [cracker](https://github.com/lovedboy/cracker) on [heroku](https://heroku.com/)

需要 github 以及 Heroku 账户用来部署
基本步骤可以参照此链接 https://github.com/521xueweihan/shadowsocks-heroku/blob/master/README.md
除了最后我们是部署cracker而不是ss



####1. 注册Heroku账户和Github账户

####2.登陆Github 并且fork cracker-heroku项目 

https://github.com/lovedboy/cracker-heroku

![](https://github.com/wangwill/cracker-heroku/blob/master/fork.png)


####3.登陆Heroku创建新app并关联Github账户

![](https://github.com/wangwill/cracker-heroku/blob/master/new%20app.png)

可以部署在美国和欧洲

![](https://github.com/wangwill/cracker-heroku/blob/master/app%20name.png)

选depoly后在下方勾选github并关联自己的账户

![](https://github.com/wangwill/cracker-heroku/blob/master/deploy.png)

红圈项可以不勾选。

####4.点部署（deploy）

10s左右会有显示 部署结果

成功的话 会显示Heroku项目网址

#### !!!! 可以在这里设定自己的密码

Setting 页面 ——> Reveal Config Vars，添加参数：
SECRET= yourpassword
![](https://github.com/wangwill/cracker-heroku/blob/master/change%20secret.png)


####5.客户端使用方法

linux下使用方法 作者已明确描述

windows下 下载cracker最新的release：https://github.com/lovedboy/cracker/releases

保存exe文件（local）

打开cmd 输入exe存放地址 后跟 -raddr https://*****.herokuapp.com -secret ***** 默认本地1080端口

例 windows客户端我存放在E盘/WALL/Cracker 文件夹下 执行命令如下

![](https://github.com/wangwill/cracker-heroku/blob/master/cmd-instruction.png)

####6.浏览器以及本地支持代理的设置就不说了
