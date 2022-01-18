
Just a test repo for [unused](https://github.com/bep/Unused).

Running `unused` in the current repo currently prints the listing below. This mostly works great, but there are some gotchas re interfaces:

* Implemented interfaces which are never used in, e.g., a return value is reported as Unused (`UnusedInterface`). This is not surprising.
* An implemented interface method is reported as unused in the interface itself (`UsedInterfaceMethodReturningInt`). This is also not surprising, but may be a bigger problem than the first; maybe we should also look for implementations.

```
unused "**.go"                                                                                                                                                                
firstpackage/code1.go:5:2 variable UnusedVar is unused (EU1002)
firstpackage/code1.go:10:2 constant UnusedConst is unused (EU1002)
firstpackage/code1.go:17:6 function UnusedFunction is unused (EU1002)
firstpackage/code1.go:23:2 field UnusedField is unused (EU1002)
firstpackage/code1.go:34:6 interface UnusedInterfaceWithUsedAndUnusedMethod is unused (EU1002)
firstpackage/code1.go:36:2 method UnusedInterfaceMethodReturningInt is unused (EU1002)
firstpackage/code1.go:35:2 method UsedInterfaceMethodReturningInt is unused (EU1002)
firstpackage/code1.go:39:6 interface UnusedInterface is unused (EU1002)
firstpackage/code1.go:40:2 method UnusedInterfaceReturningInt is unused (EU1002)
firstpackage/code1.go:43:6 interface UsedInterface is unused (EU1002)
```
