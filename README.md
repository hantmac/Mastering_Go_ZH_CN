# Mastering_Go_ZH_CN
### 《Mastering GO》

![](https://ws3.sinaimg.cn/large/006tNbRwly1fyma67vtssj30830a074b.jpg)

------


本书适用于Golang程序员。您之前应该阅读有关Go的介绍性书籍，或者已经完成了[Go By Example](https://books.studygolang.com/gobyexample/)。本书的内容包括但不限于并发、网络编程、垃圾回收、组合、GO UNIX系统编程、基本数据类型（Array,Slice,Map）、GO源码、反射，接口，类型方法等高级概念。阅读本书需要一定的编程经验。如果你在工作中使用Go或者业余时间爱好GO，那么这本书一定会让你对GO的理解更上一层楼。

### 翻译进度

*持续更新中。。。。*
- [目录](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/TOC.md)
- [chapter 0 前言](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter0)
  - [00.1 前言](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter0/00.1.md)
  - [00.2 面向读者](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter0/00.2.md)
  - [00.3 章节概览](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter0/00.3.md)
  - [00.4 更多信息](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter0/00.4.md)
    - [00.4.1 代码规范约定](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter0/00.4.1.md)
- [chapter 1 Go与操作系统](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter1)
  - [01.1 本书结构](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter1/01.1.md)
  - [01.2 Go的历史](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter1/01.2.md)
  - [01.3 为什么是Go](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter1/01.3.md)
  - [01.4 Go的优点](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter1/01.4.md)
    - [01.4.1 Go是完美的么](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter1/01.4.1.md)
- [chapter 2 Go内部机制](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter2)
  - [02.1 本章概述](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter2/02.1.md)
  - [02.2 编译器](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter2/02.2.md)
  - [02.5 C中调用Go函数](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter2/02.5.md)
    - [02.5.1 Go Package](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter2/02.5.1.md)
    - [02.5.2 C代码](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter2/02.5.2.md)
  - [02.6 defer关键字](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter2/02.6.md)
  - [02.7 Panic和Recover](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter2/02.7.md)
- [chapter 3 Go基本数据类型](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter3)
  - [03.3 Go切片](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.3.md)
    - [03.3.1 切片基本操作](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.3.1.md)
    - [03.3.2 切片的扩容](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.3.2.md)
    - [03.3.3 字节切片](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.3.3.md)
    - [03.3.4 copy()函数](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.3.4.md)
    - [03.3.5 多维切片](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.3.5.md)
    - [03.3.6 使用切片的代码示例](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.3.6.md)
    - [03.3.7 使用sort.Slice()排序](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.3.7.md)
  - [03.4 Go 映射(map)](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.4.0.md)
    - [03.4.1 Map值为nil的坑](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.4.1.md)
    - [03.4.2 何时该使用Map?](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.4.2.md)
  - [03.5 Go 常量](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.5.md)
    - [03.5.1 常量生成器：iota](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.5.1.md)
  - [03.6 Go 指针](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.6.md)
  - [03.7 时间与日期的处理技巧](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.7.md)
    - [03.7.1 解析时间](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.7.1.md)
    - [03.7.2 解析时间的代码示例](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.7.2.md)
    - [03.7.3 解析日期](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.7.3.md)
    - [03.7.4 解析日期的代码示例](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.7.4.md)
    - [03.7.5 格式化时间与日期](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.7.5.md)
  - [03.8 延展阅读](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.8.md)
  - [03.9 练习](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.9.md)
  - [03.10 本章小结](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter3/03.10.md)
- [chapter4 组合类型的使用](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.0.md)
  - [04.1 关于组合类型](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.1.md)
  - [04.2 结构体](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.2.md)
    - [04.2.1 结构体指针](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.2.1.md)
    - [04.2.2 使用new关键字](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.2.2.md)
  - [04.3 元组](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.3.md)
  - [04.4 正则表达式与模式匹配](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.4.md)
    - [04.4.1 理论知识](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.4.1.md)
    - [04.4.2 简单的正则表达式示例](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.4.2.md)
    - [04.4.3 高级的正则表达式示例](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.4.3.md)
    - [04.4.4 正则匹配IPv4地址](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.4.4.md)
  - [04.5 字符串](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.5.md)
    - [04.5.1 rune是什么？](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.5.1.md)
    - [04.5.2 关于Unicode的包](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.5.2.md)
    - [04.5.3 关于字符串处理的包](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.5.3.md)
  - [04.6 switch语句](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.6.md)
  - [04.7 计算Pi的精确值](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.7.md)
  - [04.8 实现简单的K-V存储](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.8.md)
  - [04.9 延展阅读](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.9.md)
  - [04.10 练习](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.10.md)
  - [04.11 本章小结](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter4/04.11.md)
- [chapter5 数据结构](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter5/05.0.md)
  - [05.1 图和节点](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter5/05.1.md)
  - [05.2 算法复杂度](https://github.com/hantmac/Mastering_Go_ZH_CN/blob/master/eBook/chapter5/05.2.md)
- [chapter 7 反射和接口](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7)
  - [07.1 类型方法](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.1.md)
  - [07.2 Go的接口](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.2.md)
  - [07.3 类型断言](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.3.md)
  - [07.4 设计接口](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.4.md)
    - [07.4.1 接口的使用](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.4.1.md)
    - [07.4.2 Switch用于类型判断](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.4.2.md)
  - [07.5 反射](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.5.md)
    - [07.5.1 使用反射的简单示例](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.5.1.md)
    - [07.5.2 反射进阶](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.5.2.md)
    - [07.5.3 反射的三个缺点](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.5.3.md)
  - [07.6 Go的OOP思想](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.6.md)
  - [07.7 延展阅读](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.7.md)
  - [07.8 练习](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.8.md)
  - [07.9 本章小结](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter7/07.9.md)

- [chapter 9 并发-Goroutines,Channel和Pipeline](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter9/09.0.md)
  - [09.1 关于进程，线程和Go协程](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter9/09.1.md)
      - [09.1.1 Go调度器](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter9/09.1.1.md)
      - [09.1.2 并发与并行](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter9/09.1.2.md)
  - [09.2 Goroutines](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter9/09.2.md)
      - [09.2.1 创建一个Goroutine](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter9/09.2.1.md)
      - [09.2.2 创建多个Goroutine](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter9/09.2.2.md)
  - [09.3 优雅地结束goroutines](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter9/09.3.md)
      - [09.3.1 当Add()和Done()的数量不匹配时会发生什么？](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter9/09.3.1.md)
  - [09.4 Channel(通道)](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter9/09.4.md)

- [chapter 12 Go网络编程基础](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.0.md)
  - [12.1 关于net/http,net和http.RoundTripper](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.1.md)
    - [12.1.1 http.Response类型](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.1.1.md)
    - [12.1.2 http.Request类型](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.1.2.md)
    - [12.1.3 http.Transport类型](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.1.3.md)
  - [12.2 关于TCP/IP](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.2.md)
  - [12.3 关于IPv4和IPv6](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.3.md)
  - [12.4 命令行工具netcat](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.4.md)
  - [12.5 读取网络接口的配置文件](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.5.md)
  - [12.6 实现DNS查询](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.6.md)
    - [12.6.1 获取域名的 NS 记录](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.6.1.md)
    - [12.6.2 获取域名的 MX 记录](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.6.2.md)
  - [12.7 Go实现web服务器](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.7.md)
    - [12.7.1 分析HTTP服务](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.7.1.md)
    - [12.7.2 用Go创建网站](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.7.2.md)
  - [12.8 追踪 HTTP](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.8.md)
    - [12.8.1 测试 HTTP handler](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.8.1.md)
  - [12.9 Go实现web客户端](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.9.md)
    - [12.9.1 Go web客户端进阶](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter12/12.9.1.md)
- [chapter 13 网络编程 - 构建服务器与客户端](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.0.md)
  - [13.1 Go 标准库-net](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.1.md)
  - [13.2 TCP 客户端](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.2.md)
    - [13.2.1 另一个版本的 TCP 客户端](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.2.1.md)
  - [13.3 TCP 服务器](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.3.md)
    - [13.3.1 另一个版本的 TCP 服务器](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.3.1.md)  
  - [13.4 UDP 客户端](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.4.md)
  - [13.5 UDP 服务器](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.5.md)
  - [13.6 并发 TCP 服务器](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.6.md)
    - [13.6.1 简洁的并发TCP服务器](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.6.1.md)
  - [13.7 远程调用（RPC）](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.7.md)
    - [13.7.1 RPC 客户端](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.7.1.md)
    - [13.7.2 RPC 服务器](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.7.2.md)
  - [13.8 底层网络编程](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.8.md)
    - [13.8.1 获取ICMP数据](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.8.1.md)
  - [13.9 接下来的任务](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.9.md)
  - [13.10 延展阅读](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.10.md)
  - [13.11 练习](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.11.md)   
  - [13.12 本章小节](https://github.com/hantmac/Mastering_Go_ZH_CN/tree/master/eBook/chapter13/13.12.md)

-------
### GitBook
[Mastering_Go_ZH_CN](https://wskdsgcf.gitbook.io/mastering-go-zh-cn)

根据翻译进度实时更新。
=======

-------
### 支持本书

如果你喜欢本书 《玩转 Go》，你可以参与到本书的翻译或纠正工作中来，具体请联系【Jack E-mail：hantmac@outlook.com】，一同完善本书并帮助壮大 Go 语言在国内的学习群体，给大家提供更好的学习资源。

-------
### 规则&&Fork&&PR

- 章节命名规则：举例，第一章第一节，01.1.md,如果第一节下面还有分支，01.1.1,依次类推;
- 联系邮箱，取得电子版，获得安排的翻译章节，Fork分支，提交PR;
- 由多人审核后，合并

-------
### 交流社区

微信群：

![](https://ws3.sinaimg.cn/large/006tNc79ly1fyuq9sie4nj305i0a23yi.jpg)

若以上二维码已过期，请加我微信，拉入群组。

![](https://ws4.sinaimg.cn/large/006tNc79ly1fyus99uw3bj306b06bjrq.jpg)

--------
### 致谢

- 本书原作者：Mihalis Tsoukalos
- 参与翻译人员
  - [Jack](https://github.com/hantmac)
  - [xionghui](https://github.com/xionghui)
  - [leeweir](https://github.com/leeweir)
  - [songxuexian](https://github.com/songxuexian)
  - [tangjun1990](https://github.com/tangjun1990)
  - [glbrtchen](https://github.com/glbrtchen)
  - [wskdsgcf](https://github.com/wskdsgcf)
  - [mark1995](https://github.com/mark1995)
  - [themoonbear](https://github.com/themoonbear)

---------
### 授权许可

除特别声明外，本书中的内容使用 [CC BY-SA 3.0 License](http://creativecommons.org/licenses/by-sa/3.0/)（创作共用 署名-相同方式共享3.0 许可协议）授权，代码遵循 [BSD 3-Clause License](https://github.com/astaxie/build-web-application-with-golang/blob/master/LICENSE.md)（3 项条款的 BSD 许可协议）。

--------
### 开始阅读

------

### Go学习资料及社区（持续更新中。。。）
- [Go By Example 英文网站](https://gobyexample.com/)
- [Go By Example 中文网站](https://books.studygolang.com/gobyexample/)
- [GOCN Forum](https://gocn.vip/)
- [Go语言中文网](https://studygolang.com/)
- [Go walker 强大的Go在线API文档](https://gowalker.org/)
- [jsonTOGo 好用的json转go struct工具](https://mholt.github.io/json-to-go/)
- [Go web框架beego](https://beego.me/)
- [官方代码规范指导](https://github.com/golang/go/wiki/CodeReviewComments)
- [xorm](https://github.com/go-xorm/xorm)支持 MySQL、PostgreSQL、SQLite3 以及 MsSQL
- [mgo](http://labix.org/mgo)MongoDB 官方推荐驱动
- [gorm](https://github.com/jinzhu/gorm)全功能 ORM (无限接近) 支持 MySQL、PostgreSQL、SQLite3 以及 MsSQL

