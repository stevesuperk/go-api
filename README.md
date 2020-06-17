## go-api

基于 [Gin](https://github.com/gin-gonic/gin) 进行模块化设计的 API 框架，封装了常用的功能，使用简单，致力于进行快速的业务研发。

## 技术选型
1. web:[gin](https://github.com/gin-gonic/gin)
2. orm:[gorm](https://github.com/jinzhu/gorm)
3. database:[Mysql](https://github.com/mattn/go-sqlite3)
4. 配置文件 [go-yaml](https://github.com/go-yaml/yaml)
5. 定时任务

认证，鉴权，安全，流量管控，缓存，路由，熔断，灰度发布，监控报警

热更新，负载均衡，访问策略，IP黑白名单，扩展插件机制
版本控制，日志，

## 项目结构
```
go-api
    |-app 模块存放路径
        |-users 用户模块
            |-controller.go 控制器
            |-model.go 数据库访问目录
            |-view 模板文件目录
            |-route.go 路由
        ...
    |-config 配置文件目录
    |-logs 日志目录
    |-middleware 中间件
    |-helpders 公共方法目录
    |-static 静态资源目录
        |-css css文件目录
        |-images 图片目录
        |-js js文件目录
        |-libs js类库
    |-tests 测试目录
    |-vendor 项目依赖其他开源项目目录
    |-main.go 程序执行入口
```

支持多种数据库，配置文件实时加载更新

初始化，数据库，路由，日志，缓存