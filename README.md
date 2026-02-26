# ME Frp Python SDK

ME Frp 5.0 的官方 Python SDK，支持同步和异步两种调用模式。

## 安装

推荐使用 [uv](https://github.com/astral-sh/uv) 进行安装和管理。

### 仅同步支持 (requests)

```bash
uv add MEFrpLib --optional sync
# 或者使用 pip
pip install MEFrpLib[sync]
```

### 仅异步支持 (aiohttp)

```bash
uv add MEFrpLib --optional async
# 或者使用 pip
pip install MEFrpLib[async]
```

### 全部安装

```bash
uv add MEFrpLib --optional all
# 或者使用 pip
pip install MEFrpLib[sync,async]
```

## 快速开始

### 同步模式

```python
from MEFrpLib import MEFrpClient

# 初始化客户端
client = MEFrpClient()

# 登录
client.login("username", "password", "captcha_token")

# 获取用户信息
user_info = client.get_user_info()
print(f"欢迎回来, {user_info.username}! 剩余流量: {user_info.traffic} MB")

# 获取隧道列表
proxies_resp = client.get_proxy_list()
for proxy in proxies_resp.proxies:
    print(f"隧道: {proxy.proxyName} -> {proxy.domain}:{proxy.remotePort}")
```

### 异步模式

```python
import asyncio
from MEFrpLib import AsyncMEFrpClient

async def main():
    # 使用上下文管理器自动管理 session
    async with AsyncMEFrpClient() as client:
        # 登录
        await client.login("username", "password", "captcha_token")
        
        # 获取用户信息
        user_info = await client.get_user_info()
        print(f"用户: {user_info.username}")
        
        # 获取节点列表
        nodes = await client.get_node_list()
        for node in nodes:
            print(f"节点: {node.name} ({node.region})")

if __name__ == "__main__":
    asyncio.run(main())
```

## 异常处理

SDK 定义了以下异常：

- `MEFrpError`: 所有 SDK 异常的基类
- `APIError`: 当 API 返回非 200 状态码或业务错误时抛出
- `AuthError`: 身份验证失败（如 Token 无效）时抛出
- `NetworkError`: 网络请求失败或超时时抛出

```python
from MEFrpLib import MEFrpClient, APIError

client = MEFrpClient(token="your_token")
try:
    info = client.get_user_info()
except APIError as e:
    print(f"API 错误: {e.message} (代码: {e.code})")
```

## 开发指南

如果您想参与开发、构建或对代码进行格式化，请确保已安装 `uv`。

### 初始化环境

```bash
# 克隆仓库后
uv sync --all-extras
```

### 代码检查与格式化

```bash
# 检查并自动修复
uvx ruff check . --fix

# 格式化代码
uvx ruff format .
```

### 构建与发布

项目使用 `uv` 进行构建和发布。

1. **构建分发包**：

   ```bash
   uv build
   ```

   构建完成后，`dist/` 目录下会生成 `.tar.gz` (sdist) 和 `.whl` (wheel) 文件。

2. **发布到 PyPI**（需要权限）：

   ```bash
   uv publish
   ```

> **提示**：GitHub Actions 已配置为在发布新 Release 时自动使用 `uv` 构建并发布到 PyPI。

## 开源协议

AGPL v3
