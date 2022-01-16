
Just a test repo for [Unused](https://github.com/bep/Unused).

Running `Unused` in the current repo currently prints the listing below. This mostly works great, but there are some gotchas re interfaces:

* Implemented interfaces which are never used in, e.g., a return value is reported as Unused (`UnusedInterface`). This is not surprising.
* An implemented interface method is reported as unused in the interface itself (`UsedInterfaceMethodReturningInt`). This is also not surprising, but may be a bigger problem than the first; maybe we should also look for implementations.

```
Unused ./...                                                                                                                                                                
Scanning .
Unused: firstpackage/code1.go:5:2:Variable:UnusedVar
Unused: firstpackage/code1.go:10:2:Constant:UnusedConst
Unused: firstpackage/code1.go:17:6:Function:UnusedFunction
Unused: firstpackage/code1.go:23:2:Field:UnusedField
Unused: firstpackage/code1.go:34:6:Interface:UnusedInterfaceWithUsedAndUnusedMethod
Unused: firstpackage/code1.go:36:2:Method:UnusedInterfaceMethodReturningInt
Unused: firstpackage/code1.go:35:2:Method:UsedInterfaceMethodReturningInt
Unused: firstpackage/code1.go:39:6:Interface:UnusedInterface
Unused: firstpackage/code1.go:40:2:Method:UnusedInterfaceReturningInt
Unused: firstpackage/code1.go:43:6:Interface:UsedInterface
```