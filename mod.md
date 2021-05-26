# mod

## go mod
go mod init "包名" 初始化module

go mod tidy 移除不需要的包，添加需要引入的包

go mod verify 验证工程文件以及依赖的源代码

go mod why -m "github上引入的地址" 告诉我们包引入的使用位置

go mod edit 修改mod上的一些内容
go mod edit -require github.com/。。。。    修改引入一个包

go mod vendor 把所有依赖移到vendor目录，会产生依赖冗余

go list -m all 列出所有的引入

go get github.com/。。。 添加引入

go build 自动把引入的包添加下载