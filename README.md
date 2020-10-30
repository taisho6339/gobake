# gobake

## Overview

gobake is a cli tool to create directories for application and a file to export GO environment values.

### Feature   

- make directories for your application
- sniff your own project's GOPATH automatically and then create a file to export it.

## Build

```
make build
```

## Usage

### Make directories for your application and Create a file to export your project's GOPATH

```
## You can run anywhere
gobake app [author name] [app name] <flags>
```

### Only generate an env file for your project

```
## Run in the same folder as go.mod file
gobake up
```

Output:

.gobake
```
export GOPATH=/Users/taisho6339/git/taisho6339/gobake
export PATH=$PATH:$GOPATH/bin
export GOENV_DISABLE_GOPATH=1
```

### Configure your environment values

```
source .gobake
```

## LICENSE

gobake is released under the MIT license. See LICENSE.