# sup

`sup` is a command line utility that will show you outstanding TODO items in the specified path.
If the path belongs to a git repository, the result can further be filtered by the author who made the commit, so you can quickly track down your outstanding work.

## Usage
### `sup`
```
sup [command]

Available Commands:
  list        List outstanding TODO items in current or specified directory

Use "sup [command] --help" for more information about a command.
```

### `sup list`
```
sup list [path] [flags]

Flags:
  -e, --email string   Override author look up value (default git config --global user.email
  -a, --skip-author    Skip author lookup
  -g, --skip-git       Skip current git commit lookup

```

```
$ sup list

Looking up commits by dharrigan118@gmail.com...

> inner/search_test.py 2 results
inner/search_test.py [0] >>> # TODO
inner/search_test.py [2] >>> return "test"  # TODO

> simple.txt 1 results
simple.txt [0] >>> TODO hello


Found 3 outstanding TODOs!
```

## Install
```bash
go get github.com/davidharrigan/sup
```

Make sure your $GOPATH bin is in your $PATH:
```bash
export PATH=$PATH:$GOPATH/bin
```
