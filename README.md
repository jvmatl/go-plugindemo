**github.com/jvmatl/go-plugindemo**

This is a simple example project that shows one way to structure a go project that uses plugins.

This was originally written to respond to a StackOverflow question (https://stackoverflow.com/q/56693941/3117035) but may get enhanced over time if there is interest from the community.

**To Build and run**

First build the plugin: (on linux)
```
cd processors/shout
go build -buildmode=plugin -o shout.so
```

The come back to this directory and:
```
go run main.go
```

-(-- John
