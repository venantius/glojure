2016-01-02
----------
* I've decided to return pointer types for any structs I'm working with. It would appear that specifying a return value of an interface expects a pointer type anyways, so this is actually a move in the direction of consistency.

* Embeded fields in Go are...not great when compared to Java. They're fine for accessing, but if I want to set them I have to be specific about where exactly the original field was set any time I do a new object construction. It's possible that's a more memory-efficient way of doing things than adding a new field outright, but the code is pretty ugly so I'm going to just encode a new field for now. This is primarily coming up with regard to inheritance of a `_meta` field from `Obj`

* The lack of union types in Go presents a significant problem when it comes to interface management. Consider the case of ISeq and IPersistentCollection - two interfaces that specify a different function signature for the same function, `cons`. By itself this isn't necessarily a problem, but it is when we come to the case of the EmptyList class/struct, which must implement both (it inherits the IPersisentCollection implementation from PersistentList). I believe the best choice for the moment is to refactor both interfaces to have independent functions (e.g. ConsISeq and ConsIPersistentCollection) and target those directly. In the future I expect a general `Cons` will be required with type switching at runtime but this will have to do for now.

* Something I wasn't aware of about Clojure is the fact that various classes implement their own versious of Seq by way of nested classes. I'll need to be careful when working on this that I create a useful and functional interface here.

* `ToArray` methods should probably be `ToSlice` for Go.

2016-01-01
----------
* I'm confused about the model for return types in Go. Critically, when returning a struct, it's not clear to me if the better pattern is to return the struct itself or to return a pointer to the struct. Feels like the latter might be the better option but I haven't had any clear guidance on that yet.
* I think part of the test suite should probably do explicit interface checking since Go doesn't have a built-in mechanism that enforces this until you try to write code that relies on a given interface.
