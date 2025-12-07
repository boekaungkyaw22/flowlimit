███████╗██╗      ██████╗ ██╗    ██╗██╗     ██╗███╗   ███╗██╗████████╗
██╔════╝██║     ██╔═══██╗██║    ██║██║     ██║████╗ ████║██║╚══██╔══╝
█████╗  ██║     ██║   ██║██║ █╗ ██║██║     ██║██╔████╔██║██║   ██║   
██╔══╝  ██║     ██║   ██║██║███╗██║██║     ██║██║╚██╔╝██║██║   ██║   
██║     ███████╗╚██████╔╝╚███╔███╔╝███████╗██║██║ ╚═╝ ██║██║   ██║   
╚═╝     ╚══════╝ ╚═════╝  ╚══╝╚══╝ ╚══════╝╚═╝╚═╝     ╚═╝╚═╝   ╚═╝   
                                                                     
High-performance, extensible, and framework-agnostic **rate limiting library for Go**.  
Supports multiple algorithms, clean middleware adapters, in-memory storage, and a modern, testable architecture.

`flowlimit` is designed to be the **core building block** for API gateways, backend services, and production Go applications that need precise and efficient traffic control.

---

## ✨ Features

- ⚡ **High-performance in-memory rate limiting**
- 🧠 **Multiple algorithms supported**
  - Token Bucket  
  - Leaky Bucket  
  - Sliding Window
- 🔌 **Framework integrations**
  - net/http  
  - Gin  
  - Fiber  
  - Echo  
  - Chi
- 🧱 **Clean and extensible architecture**
  - Pluggable algorithms  
  - Pluggable storage engine  
  - Customizable key extractors  
  - Custom rejection handlers  
- 🧪 **Fully testable with deterministic time provider**
- 🧹 **TTL-based automatic cleanup for expired keys**
- 🧭 **Designed for future Redis / distributed support**

---

## 🚀 Installation

```bash
go get github.com/yourname/flowlimit
