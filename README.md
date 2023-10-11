# azure-blob-demo

a small demo using the go-sdk for azure storage.

## usage

create 10 blobs with data

``` shell
go run main.go -a create -n 10 -s <STORAGEACCOUNT_NAME>
```

delete 10 blobs

``` shell
go run main.go -a delete -n 10 -s <STORAGEACCOUNT_NAME>
```


