gowf
====
golang web framework open source




Run: go run src/server.go

PLEASE ignore below words.














===============================中文============================

先更新下配置go文件:

   #!/bin/bash
   
   export GOROOT=$yourGoroot      #go根目录

   export PATH=$PATH:$GOROOT/bin

   export GOPATH=$currentProjectPath:$thridpartyPath      #当前项目路径:第三方资源路径

   go "$@"

项目启动命令：

  ./go run src/server.go [dev|prod|xxxx]

关于模式Mode：    

  [dev|prod|xxxx] 是运行模式，通过指定模式读取不同的配置文件，

  eg：dev->cfg/dev.cfg, prod->cfg/prod.cfg or xxxx->cfg/xxxx.cfg

  dev模式：动态加载页面，若页面有更改会及时呈现（开发用该模式）；

  其他模式：页面第一次加载后会cache到memory中，除非重启server，否则不会改变（发布时候用该模式）；


通配置文件配置端口及template路径cfg/*.cfg，

可以根据自己的需求自定义config架构，参见src/util/config.go->struct Config：

   ;my config

   [section]

   port=8080

   [template]

   path=assets/page/ ;with '/' end;


=======================================EN============================


need config go:

   #!/bin/bash
   
   export GOROOT=$yourGoroot      #$HOME/milo/build/go；

   export PATH=$PATH:$GOROOT/bin

   export GOPATH=$currentProjectPath:$thridpartyPath      #$HOME/milo/web/gowf:$GOROOT/libs；

   go "$@"


runing command: 

 ./go run src/server.go [dev|prod|xxxx]


Abort Mode:

  #[dev|prod|xxxx] are different modes, can read different config file.

  eg: dev->cfg/dev.cfg, prod->cfg/prod.cfg or xxxx->cfg/xxxx.cfg

  mode dev: dynamic loading page, if you change it.

  other modes: cache page in memory, cannot change until your restart server.



config port & template path in cfg/*.cfg，

you can define its struct as your like, reference to src/util/config.go->struct Config:

   ;my config

   [section]

   port=8080

   [template]

   path=assets/page/ ;with '/' end;

