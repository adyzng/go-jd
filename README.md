# go-AutoBuy

最近在学习Golang，虽然有些毛病吧，但是简洁高效，简直是Web开发利器。

那就拿Golang来重写Python版京东自动下单工具[jd-autobuy](https://github.com/Adyzng/jd-autobuy)来练手吧。


## 登陆

京东二维码扫码登陆，保存cookie，无需二次登陆。


## 第三方库

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
go run autobuy.go -goods 531065:2 -order

2017/06/27 16:55:13 [ INFO] ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
2017/06/27 16:55:13 [ INFO] 无需重新登录
2017/06/27 16:55:13 [ INFO] ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
2017/06/27 16:55:13 [ INFO] 购物车详情>
2017/06/27 16:55:14 [ INFO] 购买  数量  价格      总价      商品
2017/06/27 16:55:14 [ INFO]  +    2     ¥7.90     ¥15.80    【京东超市】绿箭（DOUBLEMINT）无糖薄荷糖原味薄荷味35粒23.8g单...
2017/06/27 16:55:14 [ INFO]  -    1     ¥1199.00  ¥1199.00  联想（ThinkVision）X24 23.8英寸纤薄超窄边框IPS屏显示器
2017/06/27 16:55:14 [ INFO]  +    1     ¥21.00    ¥21.00    王小波：沉默的大多数（2014版）
2017/06/27 16:55:14 [ INFO]  -    1     ¥59.00    ¥59.00    心匠 NanoFixit纳诺 纳米液态膜 创意手机膜 液体手机膜 通用于iPh...
2017/06/27 16:55:14 [ INFO] 总数: 3
2017/06/27 16:55:14 [ INFO] 总额: ¥36.80
2017/06/27 16:55:16 [ INFO] 编号: 531065, 库存: 有货, 价格: 7.90, 链接: https://cart.jd.com/gate.action?pid=531065&pcount=1&ptype=1
2017/06/27 16:55:16 [ INFO] ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
2017/06/27 16:55:16 [ INFO] 购买商品: 531065
2017/06/27 16:55:18 [ INFO] 成功加入进购物车 2 个 【京东超市】绿箭（DOUBLEMINT）无糖薄荷糖原味薄荷味35粒23.8g单...
2017/06/27 16:55:18 [ INFO] ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
2017/06/27 16:55:18 [ INFO] 订单详情>
2017/06/27 16:55:19 [ INFO] 总金额: ￥36.80
2017/06/27 16:55:19 [ INFO] 　运费: ￥6.00
2017/06/27 16:55:19 [ INFO] 应付款: ￥42.80
2017/06/27 16:55:19 [ INFO] 收货人： **** 188********
2017/06/27 16:55:19 [ INFO] 寄送至： 北京 **********************
2017/06/27 16:55:19 [ INFO] ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
2017/06/27 16:55:19 [ INFO] 提交订单>
2017/06/27 16:55:20 [ INFO] 下单成功，订单号：593*******16
```



[1]: https://github.com/go-clog/clog
[2]: https://github.com/PuerkitoBio/goquery
[3]: https://github.com/axgle/mahonia
[4]: https://github.com/bitly/go-simplejson

