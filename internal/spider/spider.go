package spider

import "github.com/bzw1204/tvbox-spider/internal/models"

// SpiderExt 爬虫扩展配置，通常为 JSON 字符串或配置对象
type SpiderExt string

// Spider 定义了 TVBox 爬虫的标准接口
// 所有爬虫实现都需要实现此接口，用于从各种视频源获取内容
type Spider interface {
	// Init 初始化爬虫
	// ext: 扩展配置参数，可以是 JSON 字符串或其他配置格式
	// 在爬虫实例化后调用，用于设置爬虫的基本配置
	Init(ext SpiderExt) error

	// HomeContent 获取首页内容
	// filter: 是否启用筛选功能
	// 返回首页的分类、推荐视频等信息
	HomeContent(filter bool) (*models.Result, error)

	// HomeVideoContent 获取首页视频内容
	// 返回首页的视频列表，通常用于展示热门或推荐视频
	HomeVideoContent() (*models.Result, error)

	// CategoryContent 获取分类内容
	// tid: 分类ID（type id）
	// page: 页码，从1开始
	// filter: 是否启用筛选
	// ext: 扩展参数，通常包含筛选条件等额外信息
	// 返回指定分类下的视频列表
	CategoryContent(tid string, page int, filter bool, ext SpiderExt) (*models.Result, error)

	// DetailContent 获取详情内容
	// ids: 视频ID列表，通常只包含一个ID
	// 返回视频的详细信息，包括播放地址、简介、演员等
	DetailContent(ids []string) (*models.Result, error)

	// SearchContent 搜索内容
	// keyword: 搜索关键词
	// quick: 是否快速搜索模式
	// 返回搜索结果列表
	SearchContent(keyword string, quick bool) (*models.Result, error)

	// SearchContentWithPage 带分页的搜索内容
	// keyword: 搜索关键词
	// quick: 是否快速搜索模式
	// page: 页码，从1开始
	// 返回指定页码的搜索结果列表
	SearchContentWithPage(keyword string, quick bool, page int) (*models.Result, error)

	// PlayerContent 获取播放内容
	// flag: 播放标识，用于区分不同的播放源或解析器
	// id: 视频ID或播放地址
	// vipFlags: VIP标识列表，用于判断是否需要VIP权限
	// 返回视频的实际播放地址和相关信息
	PlayerContent(flag, id string, vipFlags []string) (*models.Result, error)

	// LocalProxy 本地代理处理
	// params: 代理参数，包含请求的URL、请求头等信息
	// 返回代理响应，格式为 [状态码, 响应头map, 响应体]
	// 用于处理需要代理的请求，如视频流、图片等
	LocalProxy(params map[string]any) ([]*models.Result, error)

	// LiveContent 获取直播内容
	// url: 直播源地址
	// 返回直播源的播放列表或配置信息
	LiveContent(url string) (*models.Result, error)

	// ManualVideoCheck 手动视频检查
	// 返回是否需要手动检查视频的有效性
	// 某些源可能需要手动验证视频是否可播放
	ManualVideoCheck() (bool, error)

	// IsVideoFormat 判断URL是否为视频格式
	// url: 待检查的URL地址
	// 返回true表示该URL直接指向视频文件，false表示需要进一步解析
	// 用于快速判断链接类型，决定是否需要调用解析器
	IsVideoFormat(url string) (bool, error)

	// Action 执行自定义操作
	// action: 操作名称，由具体实现定义
	// 返回操作结果字符串，如果操作不存在则返回空字符串
	// 用于扩展功能，如登录、获取验证码等自定义操作
	Action(action string) string

	// Destroy 销毁爬虫实例，清理资源
	// 在爬虫不再使用时调用，用于释放连接、关闭文件等清理工作
	Destroy()
}
