
### 使用GODEV 跑标准库的test

>1、设置GODEV 为自己的源码目录

$  GODEV=/data/app/go/src/go

>2、进入源码目录
$  cd $GODEV/src

>3、设置引导安装到当前安装目录
$ GOROOT_BOOTSTRAP=$(go env GOROOT) ./make.bash

>4、使用新编译的工具链运行所有测试
$GODEV/bin/go test -v go/types 

>5、进入某个标准包下面运行单个测试

 $GODEV/bin/go test -v -test.run="TestCloseRead"
=== RUN   TestCloseRead
--- PASS: TestCloseRead (0.00s)
    net_test.go:26: skipping unixpacket test
PASS
Socket statistical information:
(inet4, stream, default): opened=2 connected=1 listened=1 accepted=0 closed=2 openfailed=0 connectfailed=1 listenfailed=0 acceptfailed=0 closefailed=0
(local, stream, default): opened=2 connected=1 listened=1 accepted=0 closed=2 openfailed=0 connectfailed=0 listenfailed=0 acceptfailed=0 closefailed=0

ok      net    0.007s