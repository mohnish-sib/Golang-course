# Golang-course

### Differnt branches contains different modules

#### Delete all folders and files to reset:

`git rm -r .`

#### To create new independent branch(i.e. orphan branches):

`git checkout --orphan BRANCHNAME`

##### After that git needs a bit of cleanup after an orphan checkout

`rm .git/index`
`rm -r *`
