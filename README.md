# Golang-course

### Differnt branches contains different modules

#### Delete all folders and files to reset:

```
$ git rm -r .
```

#### To create new independent branch(i.e. orphan branches):

```
$ git checkout --orphan BRANCHNAME
```

##### After that git needs a bit of cleanup after an orphan checkout

```
$ rm .git/index
```
```
$ rm -r *
```

### Useful links

- [ ] [Arithmetic Operators](https://www.golangprograms.com/arithmetic-operators-in-go-programming-language.html)
- [ ]  [Arithmetic Operators](https://go.dev/ref/spec#Arithmetic_operators)
- [ ] [Operators Precedence](https://www.tutorialspoint.com/go/go_operators_precedence.htm)
