# GO for FUN
##May2021 -> ?

##### 08/06/2021 -> 
build a site, tracking prices  
	1-08/06/2021- new prj: query tiki & unmarshal json to get url + price
	2-09/06/2021- logger instance + correct unit test

##### 07/06/2021
Bye Go!

##### 06/06/2021
REST CRUD

##### 19/05/2021
- interface, interface implements, <values-types(concrete type)>, nil receiver vs nil interface values, empty interface
- type assertion to get back concrete type: i.(T)
- type switches i.(type) <<<<< ok
- 
##### 18/05/2021
- ref/pointers, strings, format, struct, receiver, maps, range (wordcount), func as var, func closure (fibonacci)
- method: You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package
- value/ipointer intepret: convenient? no, too much implicit works behind the scene. not good for dev at all. Still prefer C
- all methods on a given type should have either value or pointer receivers, but not a mixture of both.

##### 14/05/2021
defer, panic, restore
go test

##### 11/05/2021
HTTP, web page & form submit

##### 10/05/2021
regex, channel

##### 09/05/2021
dependencies
> go mod init, go mod edit -replace=tuts/greetings=../greetings, go mod tidy  
> log, math, rand  
> slices, map  
No overloading?
