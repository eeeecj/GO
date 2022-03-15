package rpckv

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// KV数据库
type KVStoreService struct {
	m      map[string]string           //存储数据
	filter map[string]func(key string) //筛选符合条件的数据
	mu     sync.Mutex                  //同步锁
}

//创建实例
func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (k *KVStoreService) Get(key string, value *string) error {
	k.mu.Lock()
	defer k.mu.Unlock()

	if v, ok := k.m[key]; ok {
		*value = v
		return nil
	}
	return fmt.Errorf("not found")
}

func (k *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	key, value := kv[0], kv[1]

	// 判断是否更改值，若更改，则触发筛选条件，进而引起watch方法输出
	if oldvalue := k.m[key]; oldvalue != value {
		// 遍历筛选函数，若存在多个筛选函数，分别执行对应的watch语句
		for _, fn := range k.filter {
			fn(key)
		}
	}
	// 若未更改值，则不处罚watch
	k.m[key] = value
	return nil
}

func (k *KVStoreService) Watch(timeout int, keyChanged *string) error {
	// 注册id
	id := fmt.Sprintf("Watch-%s-%03d", time.Now(), rand.Intn(100))
	ch := make(chan string, 10)

	k.mu.Lock()
	// 向对象中添加筛选函数
	k.filter[id] = func(key string) { ch <- key }
	k.mu.Unlock()
	// 阻塞函数进程，若检测到有值更改或超时，则运行退出
	select {
	case <-time.After(time.Second * time.Duration(timeout)):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
	return nil
}
