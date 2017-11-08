
## beego-self-increment
为了解决Beego Orm多条数据循环无法获取自增id的问题,自己写了这个小方法<br>
## 项目地址
github地址: https://github.com/kiwi520/beego-self-increment<br>
欢迎使用并提出不足

## 使用方法
   ### 首先安装此包<br>
    go get  github.com/kiwi520/beego-self-increment
   ### 引用此包
    self "github.com/kiwi520/beego-self-increment"<br>
    
``` go
    var ai *sd.AutoInc<br>
    ai = sd.New(NumberId,1)<br>
    for _,v :=range JsonData{
	Ai :=ai.Id() //自增id
        ...
    }
    
```
    
   
   
