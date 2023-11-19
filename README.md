# adbc-duckdb-example

Fails on Mac OS 13.4
Go 1.21
Duckdb libduckdb.dylib 0.9.2 

```
SIGSEGV: segmentation violation
PC=0x0 m=10 sigcode=1
signal arrived during cgo execution

goroutine 19 [syscall]:
runtime.cgocall(0x100428b90, 0xc000167c50)
	/usr/local/opt/go/libexec/src/runtime/cgocall.go:157 +0x4b fp=0xc000167c28 sp=0xc000167bf0 pc=0x10000b70b
github.com/apache/arrow-adbc/go/adbc/drivermgr._Cfunc_AdbcStatementRelease(0xc00011e390, 0xc000390db0)
	_cgo_gotypes.go:386 +0x48 fp=0xc000167c50 sp=0xc000167c28 pc=0x100424468
github.com/apache/arrow-adbc/go/adbc/drivermgr.(*stmt).Close.func1(0xc000167cc8?, 0x10041da25?)
	/Users/anton/go/pkg/mod/github.com/apache/arrow-adbc/go/adbc@v0.8.0/drivermgr/wrapper.go:270 +0x65 fp=0xc000167c90 sp=0xc000167c50 pc=0x100426765
github.com/apache/arrow-adbc/go/adbc/drivermgr.(*stmt).Close(0x1008e9ef0?)
	/Users/anton/go/pkg/mod/github.com/apache/arrow-adbc/go/adbc@v0.8.0/drivermgr/wrapper.go:270 +0x31 fp=0xc000167cc0 sp=0xc000167c90 pc=0x1004266b1
github.com/levakin/adbc-duckdb-example.selectAll.func1()
	/Users/anton/Git/github.com/levakin/adbc-duckdb-example/main.go:57 +0x25 fp=0xc000167cd8 sp=0xc000167cc0 pc=0x100428125
github.com/levakin/adbc-duckdb-example.selectAll({0x10059c120, 0x100928240}, {0x10059f0e0?, 0xc000120058?})
	/Users/anton/Git/github.com/levakin/adbc-duckdb-example/main.go:74 +0x249 fp=0xc000167d78 sp=0xc000167cd8 pc=0x100428029
github.com/levakin/adbc-duckdb-example.Run({0x10059c120, 0x100928240})
	/Users/anton/Git/github.com/levakin/adbc-duckdb-example/main.go:41 +0x212 fp=0xc000167f30 sp=0xc000167d78 pc=0x100427c72
github.com/levakin/adbc-duckdb-example.TestRun(0xc000102ea0)
	/Users/anton/Git/github.com/levakin/adbc-duckdb-example/main_test.go:9 +0x26 fp=0xc000167f70 sp=0xc000167f30 pc=0x100428526
testing.tRunner(0xc000102ea0, 0x100512760)
	/usr/local/opt/go/libexec/src/testing/testing.go:1595 +0xff fp=0xc000167fc0 sp=0xc000167f70 pc=0x100115bbf
testing.(*T).Run.func1()
	/usr/local/opt/go/libexec/src/testing/testing.go:1648 +0x25 fp=0xc000167fe0 sp=0xc000167fc0 pc=0x100116b45
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc000167fe8 sp=0xc000167fe0 pc=0x100074561
created by testing.(*T).Run in goroutine 1
	/usr/local/opt/go/libexec/src/testing/testing.go:1648 +0x3ad

goroutine 1 [chan receive]:
runtime.gopark(0xc0001699e8?, 0x100014885?, 0x68?, 0x61?, 0x18?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc000169980 sp=0xc000169960 pc=0x10004126e
runtime.chanrecv(0xc00017e070, 0xc000169a67, 0x1)
	/usr/local/opt/go/libexec/src/runtime/chan.go:583 +0x3cd fp=0xc0001699f8 sp=0xc000169980 pc=0x10000db4d
runtime.chanrecv1(0x1008f8820?, 0x10045d340?)
	/usr/local/opt/go/libexec/src/runtime/chan.go:442 +0x12 fp=0xc000169a20 sp=0xc0001699f8 pc=0x10000d752
testing.(*T).Run(0xc000102d00, {0x1004fa403?, 0x2117b2b062076?}, 0x100512760)
	/usr/local/opt/go/libexec/src/testing/testing.go:1649 +0x3c8 fp=0xc000169ae0 sp=0xc000169a20 pc=0x1001169e8
testing.runTests.func1(0x1008f9ae0?)
	/usr/local/opt/go/libexec/src/testing/testing.go:2054 +0x3e fp=0xc000169b30 sp=0xc000169ae0 pc=0x100118afe
testing.tRunner(0xc000102d00, 0xc000169c48)
	/usr/local/opt/go/libexec/src/testing/testing.go:1595 +0xff fp=0xc000169b80 sp=0xc000169b30 pc=0x100115bbf
testing.runTests(0xc00011b220?, {0x1008ea7b0, 0x1, 0x1}, {0xd0?, 0xc000169d08?, 0x0?})
	/usr/local/opt/go/libexec/src/testing/testing.go:2052 +0x445 fp=0xc000169c78 sp=0xc000169b80 pc=0x1001189e5
testing.(*M).Run(0xc00011b220)
	/usr/local/opt/go/libexec/src/testing/testing.go:1925 +0x636 fp=0xc000169ec0 sp=0xc000169c78 pc=0x1001173d6
main.main()
	_testmain.go:47 +0x19c fp=0xc000169f40 sp=0xc000169ec0 pc=0x10042877c
runtime.main()
	/usr/local/opt/go/libexec/src/runtime/proc.go:267 +0x2bb fp=0xc000169fe0 sp=0xc000169f40 pc=0x100040dfb
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc000169fe8 sp=0xc000169fe0 pc=0x100074561

goroutine 2 [force gc (idle)]:
runtime.gopark(0x0?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc00006afa8 sp=0xc00006af88 pc=0x10004126e
runtime.goparkunlock(...)
	/usr/local/opt/go/libexec/src/runtime/proc.go:404
runtime.forcegchelper()
	/usr/local/opt/go/libexec/src/runtime/proc.go:322 +0xb3 fp=0xc00006afe0 sp=0xc00006afa8 pc=0x1000410d3
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc00006afe8 sp=0xc00006afe0 pc=0x100074561
created by runtime.init.6 in goroutine 1
	/usr/local/opt/go/libexec/src/runtime/proc.go:310 +0x1a

goroutine 3 [GC sweep wait]:
runtime.gopark(0x1?, 0x0?, 0x0?, 0x0?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc00006b778 sp=0xc00006b758 pc=0x10004126e
runtime.goparkunlock(...)
	/usr/local/opt/go/libexec/src/runtime/proc.go:404
runtime.bgsweep(0x0?)
	/usr/local/opt/go/libexec/src/runtime/mgcsweep.go:321 +0xdf fp=0xc00006b7c8 sp=0xc00006b778 pc=0x10002bf3f
runtime.gcenable.func1()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:200 +0x25 fp=0xc00006b7e0 sp=0xc00006b7c8 pc=0x100021085
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc00006b7e8 sp=0xc00006b7e0 pc=0x100074561
created by runtime.gcenable in goroutine 1
	/usr/local/opt/go/libexec/src/runtime/mgc.go:200 +0x66

goroutine 4 [GC scavenge wait]:
runtime.gopark(0xc000036070?, 0x100596508?, 0x0?, 0x0?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc00006bf70 sp=0xc00006bf50 pc=0x10004126e
runtime.goparkunlock(...)
	/usr/local/opt/go/libexec/src/runtime/proc.go:404
runtime.(*scavengerState).park(0x1008f8f40)
	/usr/local/opt/go/libexec/src/runtime/mgcscavenge.go:425 +0x49 fp=0xc00006bfa0 sp=0xc00006bf70 pc=0x100029789
runtime.bgscavenge(0x0?)
	/usr/local/opt/go/libexec/src/runtime/mgcscavenge.go:658 +0x59 fp=0xc00006bfc8 sp=0xc00006bfa0 pc=0x100029d39
runtime.gcenable.func2()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:201 +0x25 fp=0xc00006bfe0 sp=0xc00006bfc8 pc=0x100021025
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc00006bfe8 sp=0xc00006bfe0 pc=0x100074561
created by runtime.gcenable in goroutine 1
	/usr/local/opt/go/libexec/src/runtime/mgc.go:201 +0xa5

goroutine 18 [finalizer wait]:
runtime.gopark(0x0?, 0xc000335540?, 0x30?, 0x50?, 0x1000000010?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc00006a628 sp=0xc00006a608 pc=0x10004126e
runtime.runfinq()
	/usr/local/opt/go/libexec/src/runtime/mfinal.go:193 +0x107 fp=0xc00006a7e0 sp=0xc00006a628 pc=0x100020107
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc00006a7e8 sp=0xc00006a7e0 pc=0x100074561
created by runtime.createfing in goroutine 1
	/usr/local/opt/go/libexec/src/runtime/mfinal.go:163 +0x3d

goroutine 20 [GC worker (idle)]:
runtime.gopark(0xc000066760?, 0x100093eef?, 0xb0?, 0x67?, 0x100115bbf?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc000066750 sp=0xc000066730 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc0000667e0 sp=0xc000066750 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc0000667e8 sp=0xc0000667e0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 5 [GC worker (idle)]:
runtime.gopark(0x2117b8621e476?, 0x3?, 0xa2?, 0x6b?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc00006c750 sp=0xc00006c730 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc00006c7e0 sp=0xc00006c750 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc00006c7e8 sp=0xc00006c7e0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 34 [GC worker (idle)]:
runtime.gopark(0x2117b86213e12?, 0x3?, 0x7e?, 0x2a?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc000508750 sp=0xc000508730 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc0005087e0 sp=0xc000508750 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc0005087e8 sp=0xc0005087e0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 35 [GC worker (idle)]:
runtime.gopark(0x2117b86217716?, 0x1?, 0x73?, 0x2c?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc000508f50 sp=0xc000508f30 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc000508fe0 sp=0xc000508f50 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc000508fe8 sp=0xc000508fe0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 36 [GC worker (idle)]:
runtime.gopark(0x2117b8621e6ce?, 0x3?, 0x1?, 0x11?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc000509750 sp=0xc000509730 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc0005097e0 sp=0xc000509750 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc0005097e8 sp=0xc0005097e0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 6 [GC worker (idle)]:
runtime.gopark(0x2117b8621e3da?, 0x3?, 0x78?, 0x5?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc00006cf50 sp=0xc00006cf30 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc00006cfe0 sp=0xc00006cf50 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc00006cfe8 sp=0xc00006cfe0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 7 [GC worker (idle)]:
runtime.gopark(0x100929240?, 0x1?, 0xf0?, 0x96?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc00006d750 sp=0xc00006d730 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc00006d7e0 sp=0xc00006d750 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc00006d7e8 sp=0xc00006d7e0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 37 [GC worker (idle)]:
runtime.gopark(0x2117b8624e891?, 0x1?, 0x47?, 0xd8?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc000509f50 sp=0xc000509f30 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc000509fe0 sp=0xc000509f50 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc000509fe8 sp=0xc000509fe0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 8 [GC worker (idle)]:
runtime.gopark(0x100929240?, 0x1?, 0x8f?, 0xdb?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc00006df50 sp=0xc00006df30 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc00006dfe0 sp=0xc00006df50 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc00006dfe8 sp=0xc00006dfe0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 21 [GC worker (idle)]:
runtime.gopark(0x2117b8624cdfd?, 0x3?, 0x84?, 0x44?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc000066f50 sp=0xc000066f30 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc000066fe0 sp=0xc000066f50 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc000066fe8 sp=0xc000066fe0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 9 [GC worker (idle)]:
runtime.gopark(0x2117b8621dec7?, 0x3?, 0x5c?, 0x3?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc000504750 sp=0xc000504730 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc0005047e0 sp=0xc000504750 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc0005047e8 sp=0xc0005047e0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 22 [GC worker (idle)]:
runtime.gopark(0x2117b86256997?, 0x3?, 0x9e?, 0xac?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc000067750 sp=0xc000067730 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc0000677e0 sp=0xc000067750 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc0000677e8 sp=0xc0000677e0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 38 [GC worker (idle)]:
runtime.gopark(0x2117b8621e33b?, 0xc000030140?, 0x1a?, 0x14?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc00050a750 sp=0xc00050a730 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc00050a7e0 sp=0xc00050a750 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc00050a7e8 sp=0xc00050a7e0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 10 [GC worker (idle)]:
runtime.gopark(0x2117b86254ab6?, 0x1?, 0xc7?, 0x71?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc000504f50 sp=0xc000504f30 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc000504fe0 sp=0xc000504f50 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc000504fe8 sp=0xc000504fe0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 23 [GC worker (idle)]:
runtime.gopark(0x2117b8624cf93?, 0x3?, 0xc2?, 0x24?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc000067f50 sp=0xc000067f30 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc000067fe0 sp=0xc000067f50 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc000067fe8 sp=0xc000067fe0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

goroutine 39 [GC worker (idle)]:
runtime.gopark(0x2117b8624a39a?, 0x1?, 0x92?, 0xee?, 0x0?)
	/usr/local/opt/go/libexec/src/runtime/proc.go:398 +0xce fp=0xc00050af50 sp=0xc00050af30 pc=0x10004126e
runtime.gcBgMarkWorker()
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1293 +0xe5 fp=0xc00050afe0 sp=0xc00050af50 pc=0x100022c05
runtime.goexit()
	/usr/local/opt/go/libexec/src/runtime/asm_amd64.s:1650 +0x1 fp=0xc00050afe8 sp=0xc00050afe0 pc=0x100074561
created by runtime.gcBgMarkStartWorkers in goroutine 19
	/usr/local/opt/go/libexec/src/runtime/mgc.go:1217 +0x1c

rax    0x100c04280
rbx    0xc00011e390
rcx    0xc000167c50
rdx    0xc000167be0
rdi    0xc00011e390
rsi    0xc000390db0
rbp    0x70001205ceb0
rsp    0x70001205ce98
r8     0xc000700000
r9     0x0
r10    0x100453820
r11    0x10
r12    0x40000000000000
r13    0xc02c2acb0ab2c2ac
r14    0xc000168000
r15    0x49
rip    0x0
rflags 0x10a87
cs     0x2b
fs     0x0
gs     0x0
```
