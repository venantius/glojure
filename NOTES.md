2016-01-01
----------
* I'm confused about the model for return types in Go. Critically, when returning a struct, it's not clear to me if the better pattern is to return the struct itself or to return a pointer to the struct. Feels like the latter might be the better option but I haven't had any clear guidance on that yet.
* I think part of the test suite should probably do explicit interface checking since Go doesn't have a built-in mechanism that enforces this until you try to write code that relies on a given interface.
