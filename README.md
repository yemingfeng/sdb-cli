### [SDB](https://github.com/yemingfeng/sdb) cli

#### 启动 [sdb server](https://github.com/yemingfeng/sdb)

SDB 默认监听端口为 10000

#### 启动

```shell
sh ./scripts/start_cli.sh
```

##### shell 启动

<font color="red">由于使用了 protobuf，该项目并没有将 protobuf 生成的 go 文件上传到 github。 需要手动触发编译 protobuf 文件</font>

```shell
sh ./scripts/build_protobuf.sh
```

```shell
sh ./scripts/start_cli.sh
```

#### 支持的命令列表

命令规则与 [sdb protobuf](https://github.com/yemingfeng/sdb-protobuf) 保持一致

```shell
Commands:
  bfadd               bfadd key values
  bfcreate            bfcreate key n p
  bfdel               bfdel key
  bfexist             bfexist key values
  bscountrange        bscountrange key start end
  bscreate            bscreate key size
  bsdel               bsdel key
  bsgetrange          bsgetrange key start end
  bscount             bscount key
  bsmget              bsmget key bits
  bsmset              bsmset key bits value
  bssetrange          bssetrange key start end value
  cinfo               cinfo 
  clear               clear the screen
  del                 del key
  exit                exit the program
  get                 get key
  ghadd               ghadd key id0 latitude0 longitude0 id1 latitude1 longitude1......
  ghcount             ghcount key
  ghcreate            ghcreate key precision
  ghdel               ghdel key
  ghgetboxes          ghgetboxes key latitude longitude
  ghgetneighbors      ghgetneighbors key latitude longitude
  ghmembers           ghmembers key
  ghpop               ghpop key ids
  help                display help
  hlladd              hlladd key values
  hllcount            hllcount key
  hllcreate           hllcreate key
  hlldel              hlldel key
  incr                incr key delta
  lcount              lcount key
  ldel                ldel key
  lexist              lexist key values
  llpush              llpush key values
  lmembers            lmembers key
  lpop                lpop key values
  lrange              lrange key offset limit
  lrpush              lrpush key values
  mcount              mcount key
  mdel                mdel key
  mexist              mexist key keys
  mget                mget keys
  mmembers            mmembers key
  mpop                mpop key keys
  mpush               mpush key key0 value0 key1 value1......
  mset                mset key0 value0 key1 value1 ......
  plist               plist dataType key offset limit
  scount              scount key
  sdel                sdel key
  set                 set key value
  setnx               setnx key value
  sexist              sexist key values
  smembers            smembers key
  spop                spop key values
  spush               spush key values
  zcount              zcount key
  zdel                zdel key
  zexist              zexist key values
  zmembers            zmembers key
  zpop                zpop key values
  zpush               zpush key value0 score0 key1 score1......
  zrange              zrange key offset limit

```

#### 使用

<img alt="cli" src="https://github.com/yemingfeng/sdb-cli/raw/master/docs/cli.png" width="50%" height="50%"/>
