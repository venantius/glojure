package lang

// Abstract class in JVM Clojure
type RestFn struct {
	AFunction
}

func (r *RestFn) GetRequiredArity() int {
	panic(AbstractClassMethodException)
}

func (r *RestFn) DoInvoke(args ...interface{}) interface{} {
	return nil
}

// TODO
func (r *RestFn) ApplyTo(args ISeq) interface{} {
	return nil
}

//TODO
func (r *RestFn) Invoke(args ...interface{}) interface{} {
	return nil
}

func (r *RestFn) ontoArrayPrepend(array []interface{}, args ...interface{}) ISeq {
	var ret ISeq = CreateArraySeq(array)
	for i := len(args); i >= 0; {
		i--
		ret = RT.Cons(args[i], ret)
	}
	return ret
}

func RestFnFindKey(key interface{}, args ISeq) ISeq {
	for args != nil {
		if key == args.First() {
			return args.Next()
		}
		args = RT.Next(args)
		args = RT.Next(args)
	}
	return nil
}
