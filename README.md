# doha - Docker virtual environment for Habitat.sh

## Installation via golang package manager:

```
go get -u github.com/guskovd/doha
```


## Install via homebrew:

```
brew tap guskovd/doha
brew install doha
```

## Installation via binary:

Save the corresponding binary in the directory included in the PATH environment variable (/usr/bin, /usr/local/bin, etc ...)

e.g. /usr/local/bin/doha


## Usage:

### doha shell

```bash
doha shell
```

### doha exec

```bash
doha exec -- which hab
```

### doha start

Start doha daemon:

```bash
doha start
```

### doha stop

Stop doha daemon:

```bash
doha stop
```
