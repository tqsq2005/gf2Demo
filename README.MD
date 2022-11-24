# GoFrame Template For SingleRepo

Project Makefile Commands: 
- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.

## GoFrame V2开发流程：

1. 设计表结构，初始化项目，修改配置文件
2. 使用gf gen dao 生成对应的 dao/do/model
3. 编写api层：定义业务侧数据结构，提供对外接口的输入/输出数据结构
4. 编写model层：定义数据侧数据结构，提供对内的数据处理的输入/输出数据结构
5. 编写logic层，自动生成service层代码。（通过配置goland File Watcher自动生成，也可以通过gf gen service手动执行脚本生成，建议前者）
6. 在service层代码生成RegisterXX()方法后，在对应的logic模块注册服务(每个模块只需要写一次)
7. 编写controller层，接收/解析用户输入的参数，调用service层的服务。
8. 注册路由，对外暴露接口，比如这个项目是编写cmd.go文件。
9. 在main.go中 加入一行 _ "project-name/internal/logic" （只需写一次）
10. 在main.go中加入一行 _ "github.com/gogf/gf/contrib/drivers/mysql/v2" （如果你使用的是mysql；只需写一次）

## 教程
- [【准备工作】用“抄作业”的思路带你快速上手电商实战项目](https://mp.weixin.qq.com/s?__biz=MzIyNjM0MzQyNg==&mid=2247485145&idx=1&sn=f830c9257f65a01ac232277708aa597d&chksm=e870a9b4df0720a248ec42467b71a5f8f0baee02a1e6e39224f0f4da37dc337050107db61424&scene=178&cur_album_id=2640964083592265729&rd2werd=1&key=5aeb35d4c256315d6c68d1dffbb1853abfa9dab224127378ca8d7d9f97a27c6e8bd1387690f9c9bac542ae297f8e09d5c2aa9614ca76e2c6733965d257dd4f32a982876d191eb1f91838f8230cda636e3188ced9b000f4e0a150ab0cc94388376a5a665930e7c580cba80decef8b6c373e1d43a458d501aa55b0a54de006d4d3&ascene=15&uin=OTE4MDk4NDIx&devicetype=Windows+7+x64&version=63080021&lang=zh_CN&session_us=gh_3e948720ff0c&exportkey=n_ChQIAhIQAqtsLumCEORBMuacF92XfBLtAQIE97dBBAEAAAAAANquEpE9gGUAAAAOpnltbLcz9gKNyK89dVj0pBbsnCWfKABf1ntzm4zKbOFLKc9XbxSqsqFDvxJ%2BXLWNnCxix0WkynX8XbCAxq4TXtdvvUt4ZhG%2FDufxPz%2Fz6TNpD7YCxsdp6Yo2%2BDeBab1RzH7cqgRMeZIPSnIHHiMarKOJyO2%2BfXXzc9bVZBngfcG0mkcR01etnp1oMvyP2B3r099hb%2BntZ%2FXwExqKAODbJNRak9M3157%2FG7x1XZ20JSRwgTqzGz0%2Bx%2BOupK4vUTuTBn9k%2F00%2FdnSC9Gmc57uZs%2FxzU0B5IQ%3D%3D&acctmode=0&pass_ticket=bRj12tW2x2zz%2FK7CJGrGiwratH601F%2BzhctS2yLvTtqvZBWLrK5caAZN8zmd12Nx&wx_header=0&fontgear=2)