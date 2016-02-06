2016-02-06
----------
* I'm going to mark places where JVM Clojure has features that I'm punting on for now with `ROADMAP`.

2016-02-05
----------
* I've been wondering for a long time how `clojure.core` actually gets loaded, since I couldn't
 find anywhere obvious in `clojure.main` where it happens. The magical secret sauce is in
 RT.doInit(), where we see `load("clojure.core")` -- RT.doInit() is in turn triggered by RT's
 `static` section. All of this could probably just go in RT's `init()` method.

2016-02-04
----------
* Additional notes to follow up on last night's - after extensive thought I don't see a better path
 forward than using the context package or rolling something similar myself. I think the various
 thread-local storage packages that exist in the OSS world are generally frowned upon and I'm wary
 of going down any path that wouldn't be likely to have compiler and runtime support in the long run. 
 
 This brings us to the question of how actually to implement things. I think the place to start here
 will be to begin by adding an additional `ctx` argument to all functions that will be a pointer to
 a `context`. If the pointer is nil, then that should serve as an appropriate proxy for the JVM's 
 AtomicBoolean `threadBound` field.
  
 The obvious place to begin adding this is any function that otherwise would have relied upon
 `threadBound` or `dvals`. We can propogate the appropriate function signature upwards from there. 
 I have no idea how much refactoring is going to be involved in this, but my suspicion is that it
 will be significant.
 
 On the up side, thread locals don't get much use in the compiler otherwise (Agents are the next
 thing I have to worry about, and I'm not going to pay too much attention to them for now). 

2016-02-03
----------
* This evening I really spent a few hours considering the different concurrency models of JVM Clojure
 that I'm going to need to adapt or modify to get to work idiomatically (slash, at all) in Go. For
 a while I was worried about getting namespaces to work well, but I think those will work fine in the
 short term with a simple mutex on a map. More difficult, I think, is the question of vars, which 
 depend in JVM Clojure on ThreadLocals. 
 
 I'm still reasoning my way through the necessary properties of vars and their storage patterns. At
 the moment the two "obvious" paths I could go down are ones in which subsequent storage contexts are
 propogated and then locally stored in structs or channels, or in which I use the x/net/context
 package to do something similar. Either way, the functional interface will be very different, I think.

2016-02-02
----------
* Henceforth I'm going to tag parts of the codebase that seem like un-necessary JVM cruft with
 a hashtag note, `#VESTIGIAL`

2016-01-31
----------
* Sentinel values are a pain in the ass and just bit me hard. I might end up creating a special
 utility type for them just so I can easily go on about my day. For the moment it's sufficient to be
 (very) lazy and just use `rand.Int()`.
 
* The pattern I was copying from JVM Clojure where I created a public variable, e.g. EMPTY_BLAH...
 and then built on top of that in various functions left me open to a shared memory bug that I've
 identified. Henceforth the pattern for these will have to involve helper functions and shallow
 cloning, which I think is a better pattern anyways. I'm not even sure I should have it as anything
 other than a local variable, but I can return to that later.
 
* My laziness has gotten the better of me again. In the future, when writing up function signatures
 without implementations, I should encode a `panic(NotYetImplementedError)` within the empty function
 so that I don't unintentionally end up depending on a large stack of functions that just return their
 respective zero values blindly. Frankly, it's astonishing that I didn't do anything about this sooner.

2016-01-29
----------
* The initialization methods for PersistentHashMaps vs PersistentArrayMaps are different (the former
 takes varargs, the latter takes interface slices). This sort of difference in constructor function
 interfaces is not good, and these two should be brought in line.

* Right now there are various interface methods that I've chosen to have default to the least-specific
 common denominator. For instance, ITransientAssociative forces implementing structs to default to always
 provide ITransientCollection, even though individual sub-interfaces in JVM Clojure require the return
 value to be, e.g. ITransientMap or ITransientVector. I think this approach is fine for now and can 
 always be refactored later. `Conj`, `Assoc`, `Persistent` are particularly guilty culprits here.

2016-01-24
----------
* I've made an active decision not to attempt to implement functions in JVM clojure that are
 tied to the java.util.Collection interface. These have been a major thorn in my side to stub
 out so far and as there is no analog in Go I don't see much need in pretending. Note that there
 are some that are likely to prove valuable later (i.e. compareTo, for instance, as an explicit
 non-direct-equality comparison), but I'm okay on punting on that decision for now. 

2016-01-23
----------
* The existing work on abstract classes is going to have to go. It doesn't fit with
 Go's inheritance model at all, which has turned out to be a real pain in the neck.
 Specifically: there's no way to have an abstract method for a struct that gets
 embedded in another struct such that the calling the method as the embedding struct
 retains the relevant type information of the embedding struct. My guess is that the
 ultimate takeaway here will involve moving the current abstract classes to interfaces
 and refactoring so that the individual classes each get the concrete implementations.
 These concrete implementations might be linked to private static methods that they can
 all share that are tied to the new "abstract interfaces" in question.

2016-01-18
----------
* I've finally found a simple answer on StackOverflow that explains how I could do
 simple type checking for, e.g. maps, slices, etc. Link here:
 http://stackoverflow.com/questions/25772347/go-lang-generic-map-value

2016-01-15
----------
* I've left off on making a decision about "List" for a long time. Henceforth I
 am going to use array/slice builtins as a substitute for List in all cases
 where it's not clearly a PersistentList. I think there might be an exception for
 pendingForms in the LispReader, which should maybe be a linked list based on the
 JVM inplementation
    
* From what I understand of Java's `volatile` keyword, a fairly simple locking strategy
 should suffice as a Go analog.

2016-01-11
----------
* The INode interface in JVM Clojure has an overloaded Assoc method that takes
 some additional arguments when dealing with transient nodes. At the moment I've
 named these `AssocWithEdit`, but something more like `TransientAssoc` might make
 more sense.

* I've started to experiment with working on a LispReader.go file. There's a
 tremendous amount of work to be done there so I'm taking things one step at a
 time. I've got a ways to go with familiarizing myself with Go's IO paradigm.
 
* I haven't figured out yet how exception handling is going to be implemented
 here. The Exception / error paradigm in Go is extremly different from that
 of Java.

2016-01-06
----------
* I've been embedding structs incorrectly. Embedding with pointers is acceptable 
 for some use cases, but not the ones I wanted (where the "concrete class" passes 
 along its full struct to the "abstract class" pointer receiver). I've started to 
 refactor this a bit but it's possible I've missed something - TL;DR, any struct 
 embedding should not be by pointer.

* I haven't yet figured out what the right constructor / initialization pattern is.
 Many of these data structures start out with default values, which isn't something
 Go supports by design. I'll need to figure something out at some point - maybe 
 calling Initialize() on new structs, or using a constructor function like 
 CreatePersistentArrayMap() (which is what I ended up doing with vectors).

2016-01-03
----------
* So far I've been making different versions of Cons based on the various
 interfaces entailed. I think this is probably the wrong approach - so far 
 everything that supports Cons seems to extend IPersistentCollection, so maybe 
 instead we can just defined Cons at the IPersistentCollection interface level 
 once and not for any of the sub-interfaces.

* Clojure is heavily dependent on murmurhash3 for hashing equality. I think mmh3
 is a great choice for a hashing function but unfortunately Go doesn't have an 
 implementation in the standard library, so I'm going to need to find an acceptable
 dependency and use that.

2016-01-02
----------
* I've decided to return pointer types for any structs I'm working with. It would 
 appear that specifying a return value of an interface expects a pointer type anyways, 
 so this is actually a move in the direction of consistency.
 
* Embedded fields in Go are...not great when compared to Java. They're fine for 
 accessing, but if I want to set them I have to be specific about where exactly the 
 original field was set any time I do a new object construction. It's possible that's
 a more memory-efficient way of doing things than adding a new field outright, but 
 the code is pretty ugly so I'm going to just encode a new field for now. This is 
 primarily coming up with regard to inheritance of a `_meta` field from `Obj`
 
* The lack of union types in Go presents a significant problem when it comes to 
 interface management. Consider the case of ISeq and IPersistentCollection - two 
 interfaces that specify a different function signature for the same function, `cons`.
 By itself this isn't necessarily a problem, but it is when we come to the case of the
 EmptyList class/struct, which must implement both (it inherits the IPersisentCollection
 implementation from PersistentList). I believe the best choice for the moment is to 
 refactor both interfaces to have independent functions (e.g. ConsISeq and 
 ConsIPersistentCollection) and target those directly. In the future I expect a general
 `Cons` will be required with type switching at runtime but this will have to do for now.
 
* Something I wasn't aware of about Clojure is the fact that various classes implement 
 their own versious of Seq by way of nested classes. I'll need to be careful when working 
 on this that I create a useful and functional interface here.

* `ToArray` methods should probably be `ToSlice` for Go.

* Nested interfaces are a recipe for disaster. I think ultimately all of the core
 interfaces are going to have to be flattened.

2016-01-01
----------
* I'm confused about the model for return types in Go. Critically, when returning a
 struct, it's not clear to me if the better pattern is to return the struct itself or
 to return a pointer to the struct. Feels like the latter might be the better option
 but I haven't had any clear guidance on that yet.
 
* I think part of the test suite should probably do explicit interface checking since 
 Go doesn't have a built-in mechanism that enforces this until you try to write code 
 that relies on a given interface.