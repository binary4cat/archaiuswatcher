package archaiuswatcher

import (
	"github.com/go-chassis/go-archaius"
	"testing"
)

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

func TestNewWithWatcher(t *testing.T) {
	_ = archaius.Init(archaius.WithRequiredFiles([]string{"./test.yaml"}))
	type args struct {
		i      interface{}
		prefix string
	}
	tests := []struct {
		name string
		args args
	}{
		{"无嵌套struct，指定前缀",args{i:&Person{},prefix:"xian.city"}},
		{"嵌套struct，指定前缀",args{i:&Person1{},prefix:"xian.city"}},
		{"嵌套struct，不指定前缀",args{i:&Person2{},prefix:""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewWithWatcher(tt.args.i, tt.args.prefix)
		})
	}
}

func Test_changeValue(t *testing.T) {
	_ = archaius.Init(archaius.WithRequiredFiles([]string{"./test.yaml"}))
	var person Person1
	NewWithWatcher(&person,"xian.city")
	type args struct {
		yml string
		val interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{name:"测试字段Name",args:args{yml:"xian.city.name",val:"xxx"}},
		{name:"测试字段Age",args:args{yml:"xian.city.name",val:"111"}},
		{name:"测试字段NumOfMembers",args:args{yml:"xian.city.family.numOfMembers",val:"10"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			changeValue(tt.args.yml, tt.args.val)
		})
	}
}
