# go-macho-libre

A library for dealing with wrapping and unwrapping Mach-O multi-architecture binaries (a.k.a. Universal binaries). 

## Developing

When developing / running tests, note that there is data being referenced that is not checked into the source tree 
but instead residing in a separate store and is managed via [git-lfs](https://git-lfs.github.com/). You will need to 
install and initialize git-lfs to work with this data:

```
git lfs install
```

Note on testing: since `lipo` is used to verify that the resulting binaries are as expected, these test fixtures 
can only be generated on a Mac.
