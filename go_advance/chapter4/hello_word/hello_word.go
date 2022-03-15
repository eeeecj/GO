package helloword

type HelloServiece struct{}

// 必须满足 Go 语言的 RPC 规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，
// 并且返回一个 error 类型，同时必须是公开的方法
func (p *HelloServiece) Hello(request string, replay *string) error {
	*replay = "hello:" + request
	return nil
}
