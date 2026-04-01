# TVBox Spider

[![Go Version](https://img.shields.io/badge/Go-1.26.1-blue)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

TVBox Spider 是一个基于 Go 语言开发的 TVBox 视频爬虫框架，提供标准化的接口定义和工具函数，用于从各种视频源获取内容。

## 📖 项目简介

本项目为 TVBox 应用提供了一套完整的爬虫接口规范，定义了如何获取首页内容、分类内容、视频详情、搜索、播放等核心功能。通过实现标准接口，开发者可以快速构建自己的视频源爬虫。

## ✨ 特性

- 🚀 **标准化接口**：定义了完整的 Spider 接口，涵盖所有 TVBox 核心功能
- 🔐 **加密工具**：提供 AES、Base64、MD5 等常用加密解密工具
- 📦 **类型安全**：完整的数据模型定义，支持 JSON 序列化
- 🎯 **易于扩展**：清晰的接口设计，便于实现自定义爬虫
- 🛠️ **工具完善**：包含单元测试，确保工具函数正确性

## 🏗️ 项目结构

```
tvbox-spider/
├── internal/
│   ├── models/          # 数据模型定义
│   │   └── models.go    # Vod, Result, Class 等数据结构
│   ├── spider/          # 爬虫接口定义
│   │   └── spider.go    # Spider 标准接口
│   └── utils/           # 工具函数
│       ├── crypto.go    # 加密解密工具（AES/Base64/MD5）
│       └── crypto_test.go  # 单元测试
├── go.mod               # Go 模块配置
└── README.md            # 项目文档
```

## 📦 核心组件

### 1. Spider 接口 (`internal/spider/spider.go`)

定义了 TVBox 爬虫的标准接口，包含以下核心方法：

| 方法 | 说明 |
|------|------|
| `Init(ext SpiderExt)` | 初始化爬虫实例 |
| `HomeContent(filter bool)` | 获取首页内容 |
| `HomeVideoContent()` | 获取首页推荐视频 |
| `CategoryContent(tid, page, filter, ext)` | 获取分类内容（支持分页和筛选） |
| `DetailContent(ids)` | 获取视频详情信息 |
| `SearchContent(keyword, quick)` | 搜索视频内容 |
| `SearchContentWithPage(keyword, quick, page)` | 带分页的搜索 |
| `PlayerContent(flag, id, vipFlags)` | 获取播放地址 |
| `LocalProxy(params)` | 本地代理处理 |
| `LiveContent(url)` | 获取直播内容 |
| `ManualVideoCheck()` | 手动视频检查 |
| `IsVideoFormat(url)` | 判断 URL 是否为视频格式 |
| `Action(action)` | 执行自定义操作 |
| `Destroy()` | 销毁爬虫实例，清理资源 |

### 2. 数据模型 (`internal/models/models.go`)

#### Vod - 视频对象
```go
type Vod struct {
    VodID       any    // 视频 ID
    VodName     string // 视频名称
    VodPic      string // 封面图片
    VodRemarks  string // 备注信息（集数、更新状态等）
    VodYear     string // 年份
    VodArea     string // 地区
    VodActor    string // 演员
    VodDirector string // 导演
    VodContent  string // 简介
    VodPlayFrom string // 播放来源
    VodPlayURL  string // 播放地址
}
```

#### Result - 响应结果
```go
type Result struct {
    Code      int                 // 状态码
    Message   string              // 消息
    List      []Vod               // 视频列表
    Class     []Class             // 分类列表
    Filters   map[string][]Filter // 筛选条件
    Page      int                 // 当前页码
    PageCount int                 // 总页数
    Total     int                 // 总记录数
}
```

### 3. 加密工具 (`internal/utils/crypto.go`)

提供常用的加密解密函数：

- **AESEncrypt/AESDecrypt**: AES-CBC 模式加解密（使用 PKCS7 填充）
- **Base64Encode/Base64Decode**: Base64 编码解码
- **MD5Hash**: MD5 哈希计算（32 位小写十六进制）

## 💻 使用示例

### 实现一个自定义爬虫

```go
package main

import (
    "github.com/bzw1204/tvbox-spider/internal/spider"
    "github.com/bzw1204/tvbox-spider/internal/models"
)

type MySpider struct {
    config string
}

// 实现 Spider 接口
func (s *MySpider) Init(ext spider.SpiderExt) error {
    s.config = string(ext)
    return nil
}

func (s *MySpider) HomeContent(filter bool) (*models.Result, error) {
    // 实现获取首页内容的逻辑
    return &models.Result{
        Code: 1,
        List: []models.Vod{},
    }, nil
}

func (s *MySpider) CategoryContent(tid string, page int, filter bool, ext spider.SpiderExt) (*models.Result, error) {
    // 实现获取分类内容的逻辑
    return &models.Result{
        Code: 1,
        Page: page,
        List: []models.Vod{},
    }, nil
}

func (s *MySpider) DetailContent(ids []string) (*models.Result, error) {
    // 实现获取视频详情的逻辑
    return &models.Result{
        Code: 1,
        List: []models.Vod{},
    }, nil
}

func (s *MySpider) SearchContent(keyword string, quick bool) (*models.Result, error) {
    // 实现搜索的逻辑
    return &models.Result{
        Code: 1,
        List: []models.Vod{},
    }, nil
}

func (s *MySpider) PlayerContent(flag, id string, vipFlags []string) (*models.Result, error) {
    // 实现获取播放地址的逻辑
    return &models.Result{
        Code: 1,
        Url: "https://example.com/video.mp4",
    }, nil
}

// 实现其他必需的方法...
func (s *MySpider) HomeVideoContent() (*models.Result, error) { return nil, nil }
func (s *MySpider) SearchContentWithPage(keyword string, quick bool, page int) (*models.Result, error) { return nil, nil }
func (s *MySpider) LocalProxy(params map[string]any) ([]*models.Result, error) { return nil, nil }
func (s *MySpider) LiveContent(url string) (*models.Result, error) { return nil, nil }
func (s *MySpider) ManualVideoCheck() (bool, error) { return false, nil }
func (s *MySpider) IsVideoFormat(url string) (bool, error) { return false, nil }
func (s *MySpider) Action(action string) string { return "" }
func (s *MySpider) Destroy() {}

func main() {
    spider := &MySpider{}
    spider.Init("{}")

    result, _ := spider.HomeContent(false)
    // 处理结果...
}
```

### 使用加密工具

```go
package main

import (
    "fmt"
    "github.com/bzw1204/tvbox-spider/internal/utils"
)

func main() {
    // AES 加密解密
    key := "my-secret-key-123"
    plainText := "Hello, TVBox!"

    encrypted, _ := utils.AESEncrypt(plainText, key)
    fmt.Println("加密后:", encrypted)

    decrypted, _ := utils.AESDecrypt(encrypted, key)
    fmt.Println("解密后:", decrypted)

    // Base64 编码解码
    encoded := utils.Base64Encode([]byte("test"))
    decoded, _ := utils.Base64Decode(encoded)

    // MD5 哈希
    hash, _ := utils.MD5Hash("password123")
    fmt.Println("MD5:", hash)
}
```

## 🧪 运行测试

```bash
# 运行所有测试
go test ./...

# 运行加密工具测试
go test ./internal/utils -v

# 运行特定测试
go test ./internal/utils -run TestAESDecrypt
```

## 📋 依赖要求

- Go 1.26.1 或更高版本
- 无第三方依赖（仅使用 Go 标准库）

## 🔧 开发指南

### 1. 创建新的爬虫实现

1. 导入 `spider` 包
2. 创建结构体并实现 `Spider` 接口的所有方法
3. 在 `Init` 方法中解析配置参数
4. 在各个方法中实现具体的爬取逻辑

### 2. 数据格式说明

所有响应都遵循统一的 `Result` 格式：

- `Code`: 1 表示成功，其他表示失败
- `List`: 视频列表，用于分类、搜索等场景
- `Class`: 分类信息，用于首页展示
- `Filters`: 筛选条件，支持多级筛选

### 3. 错误处理建议

- 所有方法都应该返回 `error`
- 发生错误时设置合适的 `Code` 值
- 使用 `Message` 字段描述错误信息

## 📝 注意事项

1. **接口完整性**：实现 Spider 接口时必须实现所有方法
2. **线程安全**：爬虫实例可能在多个 goroutine 中使用，注意并发安全
3. **资源管理**：在 `Destroy` 方法中及时释放资源（关闭连接、清理缓存等）
4. **性能优化**：对于频繁访问的数据建议使用缓存机制

## 🤝 贡献指南

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

## 👥 作者

GitHub: [@bzw1204](https://github.com/bzw1204)

## 🙏 致谢

感谢所有为 TVBox 生态做出贡献的开发者！
