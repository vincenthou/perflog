# perflog
Use go lang to implement simple service for logging frontend performance data

# Setup

## Set go path

Edit user bash profile.

```sh
vi ~/.bash_profile
```

Input the file content.

```sh
export GOPATH="perflog folder path here"
export GOBIN="$GOPATH/bin"
export PATH="$PATH:$GOBIN"
```

Make it work

```sh
source ~/.bash_profile
```

Check runable path

```sh
echo $Path
```

## Get dependecies

```sh
go get github.com/tools/godep
go get github.com/mitchellh/gox
go get github.com/ant0ine/go-json-rest/rest
go get github.com/ogier/pflag
```