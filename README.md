# ipvsGo
查找ip运营商
运营商数据截止到2019.7.18
#example
```
  go get https://github.com/sevenelevenlee/ipvsGo.git
  
  ip := "218.244.0.0"
  #将ip转换成uint类型
  re := ipvsGo.IPString2Long(ip)
  ipvsGo.SearchByBinary(re) #when result is "" 未在本库中查到
```
