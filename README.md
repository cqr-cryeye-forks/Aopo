# Aopo - 这个项目准备推翻重构 代码太乱了 目前的版本先开源了 师傅们可以自己优化 我重构好了会覆盖上传

###  内网自动化快速打点工具｜资产探测｜漏洞扫描｜服务扫描｜弱口令爆破

----

![image-20220502164731997](https://tva1.sinaimg.cn/large/e6c9d24egy1h1u5nlze31j20ym0u0q8r.jpg)

### 解决问题 (遇到问题请在[Issues](https://github.com/ExpLangcn/Aopo/issues)按照[模版](https://github.com/ExpLangcn/Aopo/issues/1)提交问题！其他方式反馈问题不给予受理)

Linux 报错：`/lib/ld-musl-x86_64.so.1: bad ELF interpreter: 没有那个文件或目录`

```
curl https://forensics.cert.org/cert-forensics-tools-release-el7.rpm -o cert-forensics-tools-release-el7.rpm 
rpm -Uvh cert-forensics-tools-release*rpm
yum --enablerepo=forensics install -y musl-libc-static
```

----

### 当前功能

- [x] ICMP并发探测存活IP
- [x] 对A段只进行网关探测
- [x] 集成Xray POC扫描
- [x] 解析Xray POC V1模版
- [x] 并发探测主机端口服务
- [x] 并发探测主机服务弱口令
- [x] 服务弱口令探测模块 支持 SSH
- [x] 服务弱口令探测模块 支持 Mysql
- [x] 服务弱口令探测模块 支持 Mssql
- [x] 服务弱口令探测模块 支持 Mongodb
- [x] 服务弱口令探测模块 支持 Oracle
- [x] 服务弱口令探测模块 支持 FTP
- [x] 服务弱口令探测模块 支持 SMB
- [x] 服务弱口令探测模块 支持 Redis 未授权
- [x] 自定义Password字典 自动合并到程序内置

----

### 更新记录

* **2022.5.02 Aopo V1.0.2**
  * 新增 Redis 未授权扫描
  * 新增 自定义Password字典，自动合并去重
* **2022.4.29 Aopo V1.0.1**
  * 发布程序

----

### 优势

- [x] 速度快
- [x] 体积小
- [x] 功能够用

这些应该算吧...

----

### 使用教程

**扫描内网全部资产** （默认探测：10 A段｜172 A段｜192 B段）

> ./Aopo -all

**指定网段扫描** （支持格式：192.168.1.1｜192.168.1.1/24｜192.168.1.1/16｜192.168.1.1/8｜192.168.1.1,192.168.1.2）

> ./Aopo -ip 192.168.1.1/24

**从文件读取目标** （从文件读取目标，支持格式如上 注意：一行一个）

> ./Aopo -ipf target.txt

**自定义密码字典** （从文件读取目标，自动去重复合并到程序内置字典中）

> ./Aopo -addpass password.txt -ip xxx...

----

### Twitter

#### [@ExpLangcn](https://twitter.com/ExpLangcn)

----

### 联系微信

![img](https://tva1.sinaimg.cn/large/e6c9d24egy1h1u6xw8ygsj20u0123act.jpg)