package buildlazyevaluator

type Any interface{}

type EvalFunc func(Any) (Any, Any)

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	//接受函数跟初始值
	retValChan := make(chan Any)
	loopFunc := func() {
		var actstate Any = initState
		var retval Any
		for {
			//将初始值传入函数并返回计算结果，更新初始值
			retval, actstate = evalFunc(actstate)
			retValChan <- retval
		}
	}
	retFunc := func() Any {
		//获取结果，当结果不获取时，程序堵塞
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}

func BuildIntLazyEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		//进行结果计算
		return ef().(int)
	}
}
