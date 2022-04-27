# go-jd

[2021-11] JD buy logic changed, not working now. TBD


Golang version of [jd-autobuy](https://github.com/Adyzng/jd-autobuy)。


## Login

Login by QRCode using JD app

## 3rd-party

+ [clog][1]: Clog is a channel-based logging package for Go.
+ [goquery][2]: A little like that j-thing, only in Go.
+ [mahonia][3]: Character-set conversion library implemented in Go.
+ [go-simplejson][4]: A Go package to interact with arbitrary JSON.


## Example

``` cmd
Usage 
  -area string                                                                      
        ship location string, default to Beijing (default "1_72_2799_0")            
  -goods string                                                                     
        the goods you want to by, find it from JD website.                          
        Single Goods:                                                               
          2567304(:1)                                                               
        Multiple Goods:                                                             
          2567304(:1),3133851(:2)                                                   
  -order                                                                            
        submit the order to JingDong when get the Goods.                            
  -period int                                                                       
        the refresh period when out of stock, unit: ms. (default 500)               
  -rush                                                                             
        continue to refresh when out of stock.                                      
```

``` cmd
# example
go run autobuy.go -goods 531065:2 -order
``` 



[1]: https://github.com/go-clog/clog
[2]: https://github.com/PuerkitoBio/goquery
[3]: https://github.com/axgle/mahonia
[4]: https://github.com/bitly/go-simplejson

