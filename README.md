# ipvsGo
查找ip运营商
#example
```
  go get https://github.com/sevenelevenlee/ipvsGo.git
  
  ip := "218.244.0.0"
  re := ipvsGo.IPString2Long(ip)
  ipvsGo.SearchByBinary(re) #when result is "" 未在本库中查到
```
