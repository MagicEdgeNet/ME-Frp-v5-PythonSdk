# ME Frp Golang SDK

ME Frp 5.0 官方 Go 语言 SDK，提供完整的 API 封装。

## 安装

```bash
go get github.com/MagicEdgeNet/MEFrp-v5-GolangSdk
```

## 快速开始

### 1. 注册与登录

```go
package main

import (
    "fmt"
    "log"

    "github.com/MagicEdgeNet/MEFrp-v5-GolangSdk"
)

func main() {
    // 创建客户端（无需 token 用于注册/登录）
    client := mefrp.NewClient("")

    // 获取注册验证码（需人机验证 token）
    err := client.GetRegisterEmailCode("your@email.com", "captcha_token_here")
    if err != nil {
        log.Fatal(err)
    }

    // 注册账户
    err = client.Register(mefrp.RegisterRequest{
        Username:  "myusername",
        Email:     "your@email.com",
        EmailCode: "123456",
        Password:  "MyPass123",
    })
    if err != nil {
        log.Fatal(err)
    }

    // 登录获取 token（需人机验证 token）
    token, err := client.Login(mefrp.LoginRequest{
        Username:     "myusername",
        Password:     "MyPass123",
        CaptchaToken: "captcha_token_here",
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Login successful, token: %s\n", token)
}
```

### 2. 用户信息与签到

```go
package main

import (
    "fmt"
    "log"

    "github.com/MagicEdgeNet/MEFrp-v5-GolangSdk"
)

func main() {
    client := mefrp.NewClient("YOUR_API_TOKEN")

    // 获取用户信息
    user, err := client.GetUserInfo()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("User: %s, Traffic: %d MB\n", user.Username, user.Traffic/1024)

    // 每日签到（需人机验证 token）
    err = client.Sign("captcha_token_here")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Sign-in successful!")

    // 获取用户组信息
    groups, err := client.GetUserGroups()
    if err != nil {
        log.Fatal(err)
    }
    for _, g := range groups {
        fmt.Printf("Group: %s, Max Proxies: %d\n", g.FriendlyName, g.MaxProxies)
    }
}
```

### 3. 隧道管理

```go
package main

import (
    "fmt"
    "log"

    "github.com/MagicEdgeNet/MEFrp-v5-GolangSdk"
)

func main() {
    client := mefrp.NewClient("YOUR_API_TOKEN")

    // 获取隧道列表
    resp, err := client.GetProxyList()
    if err != nil {
        log.Fatal(err)
    }
    for _, t := range resp.Proxies {
        status := "Offline"
        if t.IsOnline {
            status = "Online"
        }
        fmt.Printf("- [%s] %s (ID: %d, Type: %s)\n", status, t.ProxyName, t.ProxyID, t.ProxyType)
    }

    // 创建隧道
    err = client.CreateProxy(mefrp.CreateProxyRequest{
        ProxyName:      "MyProxy",
        ProxyType:      "tcp",
        LocalIP:        "127.0.0.1",
        LocalPort:      8080,
        RemotePort:     10080,
        NodeID:         1,
        UseEncryption:  false,
        UseCompression: false,
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Proxy created successfully!")

    // 获取隧道配置（用于启动 frpc）
    config, err := client.GetProxyConfig(12345, "toml")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Config:\n%s\n", config.Config)

    // 启用/禁用隧道
    err = client.ToggleProxy(12345, true) // true = 禁用
    if err != nil {
        log.Fatal(err)
    }

    // 删除隧道
    err = client.DeleteProxy(12345)
    if err != nil {
        log.Fatal(err)
    }
}
```

### 4. 节点信息

```go
package main

import (
    "fmt"
    "log"

    "github.com/MagicEdgeNet/MEFrp-v5-GolangSdk"
)

func main() {
    client := mefrp.NewClient("YOUR_API_TOKEN")

    // 获取节点列表
    nodes, err := client.GetNodeList()
    if err != nil {
        log.Fatal(err)
    }
    for _, n := range nodes {
        fmt.Printf("Node: %s (ID: %d, Region: %s, Online: %v)\n",
            n.Name, n.NodeID, n.Region, n.IsOnline)
    }

    // 获取节点状态
    statuses, err := client.GetNodeStatus()
    if err != nil {
        log.Fatal(err)
    }
    for _, s := range statuses {
        fmt.Printf("Node %s: %d clients, %d proxies, Load: %d%%\n",
            s.Name, s.OnlineClient, s.OnlineProxy, s.LoadPercent)
    }
}
```

### 5. 公共信息

```go
package main

import (
    "fmt"
    "log"

    "github.com/MagicEdgeNet/MEFrp-v5-GolangSdk"
)

func main() {
    client := mefrp.NewClient("")

    // 获取公共统计
    stats, err := client.GetStatistics()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Users: %d, Nodes: %d, Proxies: %d\n",
        stats.Users, stats.Nodes, stats.Proxies)

    // 获取商城商品
    items, err := client.GetStoreItems()
    if err != nil {
        log.Fatal(err)
    }
    for _, item := range items {
        fmt.Printf("%s: %.2f CNY/%s\n", item.Name, item.CurrentPrice, item.Unit)
    }
}
```

## API 列表

### 公共接口（无需认证）

- `GetStatistics()` - 获取平台统计信息
- `GetStoreItems()` - 获取商城商品列表
- `GetHolidayData(year)` - 获取节假日数据
- `CheckUpdate(req)` - 检查软件更新
- `GetRegisterEmailCode(email, captchaToken)` - 获取注册验证码
- `Register(req)` - 注册账户
- `Login(req)` - 密码登录
- `GenerateMagicLink(req)` - 生成免密登录链接
- `VerifyMagicLink(mid)` - 验证免密登录链接
- `RequestIForgotEmailCode(req)` - 请求找回密码验证码
- `IForgot(req)` - 找回密码
- `GetPublicAdsByPlacement(placement)` - 获取公共广告

### 用户相关

- `GetUserInfo()` - 获取用户信息
- `Sign(captchaToken)` - 每日签到
- `GetUserFrpToken()` - 获取 FRP 启动 Token
- `GetUserGroups()` - 获取用户组信息
- `ResetAccessKey(captchaToken)` - 重置访问密钥
- `ChangePassword(req)` - 修改密码
- `GetUserLogs(filter)` - 获取操作日志
- `GetUserLogStats()` - 获取日志统计
- `GetRealnameInfo()` - 获取实名认证信息
- `PerformRealnameLegacy(req)` - 执行实名认证
- `GetUserTrafficStats(datePeriod)` - 获取流量统计
- `GetUserIcpDomain()` - 获取已备案域名列表
- `AddIcpDomain(domain)` - 添加域名备案
- `DeleteIcpDomain(domain)` - 删除域名备案
- `KickAllProxies()` - 强制下线所有隧道
- `GetPurchaseStatus()` - 获取购买状态
- `GetOperationLogCategories()` - 获取操作日志分类

### 隧道相关

- `GetProxyList()` - 获取隧道列表
- `CreateProxy(req)` - 创建隧道
- `UpdateProxy(req)` - 更新隧道
- `DeleteProxy(proxyID)` - 删除隧道
- `KickProxy(proxyID)` - 强制下线隧道
- `ToggleProxy(proxyID, isDisabled)` - 启用/禁用隧道
- `GetProxyConfig(proxyID, format)` - 获取单一隧道配置
- `GetMultipleProxyConfigs(proxyIDs, format)` - 获取多个隧道配置
- `GetEasyStartupConfig(proxyID)` - 获取快速启动配置
- `GetCreateProxyData()` - 获取创建隧道所需数据

### 节点相关

- `GetNodeList()` - 获取节点列表
- `GetNodeFreePort(nodeID, protocol)` - 获取节点空闲端口
- `GetNodeStatus()` - 获取节点状态
- `GetNodeToken(nodeID)` - 获取节点 Token
- `GetNodeConnectionList()` - 获取节点连接地址

### 广告系统

- `GetUserAds()` - 获取我的广告
- `GetAdsByPlacement(placement, slotID)` - 查询广告
- `AddAd(ad)` - 添加广告
- `UpdateAd(ad)` - 更新广告
- `DeleteAd(adsID)` - 删除广告
- `RenewAd(adsID)` - 续费广告
- `TrackAdClick(adsID)` - 追踪广告点击
- `GetAvailableAdSlots()` - 获取可用广告位
- `GetUserAdCredits()` - 获取广告额度
- `ApplyAd(ad)` - 申请广告
- `GetUserAdsStats()` - 获取广告统计
- `GetAdSlotByPlacement(placement)` - 获取广告位信息
- `ValidateCoupon(code, productType, orderAmount)` - 验证优惠券
- `PurchaseAdCredits(slotID, amount)` - 购买广告额度

### 财务与 CDK

- `GetOrders(page, pageSize, status)` - 获取订单列表
- `RepayOrder(orderID, payMethod, force)` - 重新支付订单
- `SubmitOrder(req)` - 提交订单
- `QueryOrder(orderID)` - 查询订单状态
- `Proceed(orderID)` - 支付后处理
- `RedeemCDK(code, captchaToken)` - 兑换 CDK
- `GetMyCDKUsage(page, pageSize)` - 获取 CDK 使用记录

### 节点捐赠

- `ApplyNodeDonate(donate)` - 申请节点捐赠
- `GetUserNodeDonates()` - 获取我的捐赠列表
- `ApplyNodeDelete(nodeID, reason)` - 申请删除捐赠节点
- `GetUserNodeDeleteRequests()` - 获取删除申请列表
- `ApplyNodeEdit(req)` - 申请编辑捐赠节点
- `GetUserNodeEditRequests()` - 获取编辑申请列表
- `GetInstallScript(nodeID, system, arch, nodeType)` - 获取安装脚本

### 系统相关

- `GetSystemStatus()` - 获取系统状态
- `GetPopupNotice()` - 获取重要公告
- `GetNotice()` - 获取系统公告
- `GetDownloadSources()` - 获取下载源列表
- `GetProducts()` - 获取产品列表

## 配置选项

### 初始化选项

```go
import "time"

client := mefrp.NewClient("token",
    mefrp.WithTimeout(30 * time.Second),
    mefrp.WithBaseURL("https://api.mefrp.com/api"),
    mefrp.WithUserAgent("MyApp/1.0"),
)
```

### 动态修改配置

您可以在客户端初始化后随时修改配置：

```go
// 修改 Endpoint
client.SetBaseURL("https://your-proxy-api.com/api")

// 修改 Token
client.SetToken("NEW_API_TOKEN")

// 修改 User-Agent
client.SetUserAgent("NewAgent/2.0")
```

## 注意事项

### 人机验证

以下接口需要人机验证 Token（`captchaToken`）：
- `GetRegisterEmailCode` - 获取注册验证码
- `Login` - 密码登录
- `Sign` - 每日签到
- `ResetAccessKey` - 重置访问密钥

您需要在前端或其他地方完成人机验证后获取 Token 并传入。

### User-Agent 要求

根据官方文档，建议使用自定义 User-Agent，格式为：

```
客户端名称/版本号 联系方式
```

例如：`MyFrpClient/1.0.0 admin@example.com`

不要使用 `MEFrp-Client` 开头的 User-Agent。

### 密码修改警告

调用 `ChangePassword` 会**重置启动令牌和访问密钥**，请谨慎操作。

## 目录结构

```
.
├── README.md           # 本文档
├── ads.go              # 广告系统接口
├── auth.go             # 注册、登录、密码管理
├── cash.go             # 财务、订单与 CDK 接口
├── client.go           # 核心 HTTP 客户端
├── donate.go           # 节点捐赠接口
├── examples/           # 使用示例
│   └── demo/
│       └── main.go
├── go.mod              # Go 模块定义
├── node.go             # 节点相关接口
├── public.go           # 公共信息接口
├── system.go           # 系统状态接口
├── proxy.go           # 隧道管理接口
├── types.go            # 数据结构定义
└── user.go             # 用户信息接口
└── version.go          # SDK 版本定义
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License
