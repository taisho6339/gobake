# gobake

## Overview

gobake is a cli tool to generate a file to export environment values for Go.
gobake sniffs your own project's GOPATH automatically and then creates a file to export your GOPATH.

## Build

```
make build
```

## Usage

### Generate an Env file for your project

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