# unzipgbk
#### Unzip UTF-8, GBK, BIG5 or SHIFT-JIS encoded zip files in UTF-8 environment.

the codec of filename inside zip file are detected automatically, or you can use `-c` to give specific codec, but forcing wrong codec will fail to extract the file.

### Install
```sh
go get github.com/meoow/unzipgbk
```

### Usage
```sh
Usage: unzipgbk [options] zipfile ...

Options:
  -c="": Forcing codec instead of auto detecting
```
