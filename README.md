gowf
====
golang web framework open source

need config go:

   #!/bin/bash

   export GOROOT=$yourGoroot      #$HOME/milo/build/go
   export PATH=$PATH:$GOROOT/bin

   export GOPATH=$currentProjectPath:$thridpartyPath      #$HOME/milo/web/gowf:$GOROOT/libs

   go "$@"


runing command: ./go run src/server.go [dev|prod|xxxx]#[dev|prod|xxxx] are different modes, can read different config file.
  eg: dev->cfg/dev.cfg, prod->cfg/prod.cfg or xxxx->cfg/xxxx.cfg

mode dev: dynamic loading page, if you change it.
other modes: cache page in memory, cannot change until your restart server.

config port & template path in cfg/*.cfg:

   ;my config
   [section]
   port=8080
   [template]
   path=assets/page/ ;with '/' end;
