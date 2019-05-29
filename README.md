# 介绍

封装对`github.com/go-chassis/go-archaius`的动态使用，可以匹配自定义的配置前缀，自定义struct接收配置项，动态的修改配置全局字段的值。

# 使用

在archaius包初始化后调用`archaiuswatcher.NewWithWatcher()`，给指定的结构体赋值，并且随着配置文件的修改动态修改该结构体变量的值。

示例配置：

```yaml
xian:
  city:
    name: zhangsan
    age: 00
    family:
      numOfMembers: 3
      father: zhangsi
      mather: lisi
```

该配置有一个固定的前缀`xian.city`，其余的配置都是在这前缀之下的。

示例代码：

```golang
type Person struct {
	Name string `yaml:"name"`
	Age int `yaml:"age"`
}

type Person1 struct {
	Name string `yaml:"name"`
	Age int `yaml:"age"`
	Family Family `yaml:"family"`
}

type Family struct {
	NumOfMembers int `yaml:"numOfMembers"`
 	Father string `yaml:"father"`
	Mather string `yaml:"mather"`
}



type Person2 struct {
	Name string `yaml:"xian.city.name"`
	Age int `yaml:"xian.city.age"`
	Family Family2 `yaml:"xian.city.family"`
}
type Family2 struct {
	NumOfMembers int `yaml:"xian.city.family.numOfMembers"`
	Father string `yaml:"xian.city.family.father"`
	Mather string `yaml:"xian.city.family.mather"`
}

var person Person
var person1 Person1
var person2 Person2

// 没有嵌套的struct，并且需要拼接指定的前缀
archaiuswatcher.NewWithWatcher(&pereson, "xian.city")
// 有嵌套的struct，并且需要拼接指定的前缀
archaiuswatcher.NewWithWatcher(&pereson1, "xian.city")
// 没有嵌套的struct，不需要拼接前缀，因为字段的tag中已经写全了配置项的名称
archaiuswatcher.NewWithWatcher(&pereson2, "")
```

- 指定的前缀需要和定义的字段tag的`yaml`匹配起来，匹配实际配置文件的层级
- `yaml`字段tag必须和配置文件中的一致，区分大小写
- 可以Watcher多个全局变量