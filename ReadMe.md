## for 与 select 配合使用是否会导致 cpu 占用过高 

`select`语句中存在`default`分支时`select`不会阻塞，导致`cpu`空转；如果不存在`default`分支，无论`case`是阻塞或非阻塞管道，`select`语句会阻塞直到有`case`满足条件。

```shell
$ go build select_loop.go
$ ./select_loop
```

