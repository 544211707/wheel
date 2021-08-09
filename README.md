初学者只实现了基础的功能，但是也感觉自己暴风成长了~

# 前端-client

基于Vue框架，利用anxios收发数据，转盘样式根据网上的图自己Photoshop制作。

实现了项目要求中的：

- 转盘中有6个奖品
- 用户点击中间抽奖按钮可以进行抽奖
- 5秒后抽奖结束
- 转盘的转速逐渐降低

运行方式：

```
npm run vue3build
npm run vue3serve
```



# 后端-server

利用Golang，mysql，主要实现数据库的增删改查。

实现了项目中的：

- 奖品通过Server端下发
- 搭建一个简单的运营后台，支持简单的增加奖品、删除奖品（未来得及实现）、设置中奖概率（前端实现）、设置库存等功能

各文件功能：

```
controller：
service：
dao：执行数据库的操作
datasource：数据库引擎初始化、新建数据库表格
model：数据库表
	-wheelPrize：抽奖记录（抽奖id，奖品名称）
	-prizeContent：奖品列表（奖品id，奖品名称，奖品数量）
route：路由
```
运行方式：

```
直接运行main.go即可
```
抽奖转盘效果如下：

![20210805_210942.gif](https://github.com/544211707/wheel/blob/master/img/20210805_210942.gif?raw=true)

数据库更新如下：

1.奖品表

![奖品列表.png](https://github.com/544211707/wheel/blob/master/img/%E5%A5%96%E5%93%81%E5%88%97%E8%A1%A8.png?raw=true)

2.抽奖记录表

![抽奖记录.png](https://github.com/544211707/wheel/blob/master/img/%E6%8A%BD%E5%A5%96%E8%AE%B0%E5%BD%95.png?raw=true)
