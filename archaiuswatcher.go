package archaiuswatcher

import (
	"fmt"
	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-archaius/core"
	"github.com/go-mesh/openlogging"
	"reflect"
	"sync"
)

var once sync.Once

var yaml2value = make(map[string]reflect.Value)

type Listener struct {
	Key string
}

func (e *Listener) Event(event *core.Event) {
	changeValue(event.Key, event.Value)
}

func NewWithWatcher(i interface{}, prefix string) {
	var err error
	once.Do(func() {
		listener := Listener{Key: "test"}
		if err = archaius.RegisterListener(&listener, "s*"); err != nil {
			openlogging.GetLogger().Errorf("Failed archaius registry listener : %s", err.Error())
			panic("Failed archaius registry listener")
		}
	})

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if v.Kind() != reflect.Struct {
		openlogging.GetLogger().Errorf("`NewWithFile` need struct type! %v", v.Kind())
		panic("config type need struct!")
	}
	for i := 0; i < t.NumField(); i++ {
		fv := v.Field(i)
		ft := t.Field(i)

		yml := ft.Tag.Get("yaml")
		key := fmt.Sprintf("%s.%s", prefix, yml)

		if fv.Kind() != reflect.Struct {
			yaml2value[key] = fv
		} else {
			NewWithWatcher(fv.Addr().Interface(), key)
		}
	}

	initVal := archaius.GetConfigs()
	for i, iv := range initVal {
		// init value
		for y, yv := range yaml2value {
			if y == i {
				yv.Set(reflect.ValueOf(iv))
			}
		}
		fmt.Printf("key:%s    value:%s\n", iv, i)
	}
}

func changeValue(yml string, val interface{}) {
	var refVal reflect.Value
	var ok bool
	if refVal, ok = yaml2value[yml]; !ok {
		return
	}
	v := reflect.ValueOf(val)
	refVal.Set(v)
}
