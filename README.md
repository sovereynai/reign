# 👑 Reign

**Your terminal gateway to distributed AI.**

Reign lets you run AI inference across a network of GPU nodes—without managing infrastructure, paying for idle time, or vendor lock-in. Get instant access to language models and vision AI from your command line.

## Why Reign?

### For AI Developers

**Stop overpaying for idle GPU time.** Traditional cloud AI services charge you whether you're using them or not. With Reign, you only pay for actual inference—no monthly minimums, no idle charges.

**Access models anywhere.** Browse and use AI models across the entire Sovereyn network. From powerful language models to specialized vision AI, find the right model for your task.

**No vendor lock-in.** Open source CLI means you control your tools. Fork it, extend it, or swap it out. Your choice.

**Simple as `reign chat`.** No complex SDKs, no API key juggling, no infrastructure management. Just install and run.

## ✨ What You Get

- 🤖 **Instant AI Access** - Chat with language models from your terminal
- 🌍 **Network-Wide Discovery** - Browse 10+ models across all connected nodes
- 📊 **Usage Insights** - Track your credits, burn rate, and performance metrics
- 💡 **Smart Optimization** - Get recommendations to reduce costs
- 📦 **Model Flexibility** - Use Ollama LLMs or ONNX vision models
- ⚡ **Zero Setup** - Auto-discovers available models and services

## 🆕 What's New

- **Network Model Discovery** - See all available models across the network
- **Vision AI Support** - Access computer vision models for image classification and detection
- **Zero Configuration** - No environment variables or complex setup required
- **Model Locations** - Find which nodes host the models you need

## 🚀 Quick Start

### Prerequisites

**Install Ollama** (optional - for using local models):
```bash
# macOS
brew install ollama
ollama serve
ollama pull llama3.2:3b

# Linux
curl -fsSL https://ollama.com/install.sh | sh

# Windows
# Download from https://ollama.com/download
```

**Install Throne** (connects you to the Sovereyn network):
```bash
# macOS/Linux
brew install sovereynai/tap/throne
throne serve

# Windows
scoop bucket add sovereyn https://github.com/sovereynai/scoop-bucket
scoop install throne
```

That's it! Throne auto-configures and connects you to the network.

### Install Reign

**Homebrew (macOS/Linux):**
```bash
brew install sovereynai/tap/reign
```

**Scoop (Windows):**
```powershell
scoop bucket add sovereyn https://github.com/sovereynai/scoop-bucket
scoop install reign
```

**Or download from [Releases](https://github.com/sovereynai/reign/releases)**

### Start Using AI

```bash
# Check connection
reign status

# See available models
reign models network

# Chat with AI
reign chat "Explain machine learning in simple terms"
```

## 📖 Common Commands

### Chat with AI

```bash
# Simple chat
reign chat "Explain quantum computing in one sentence"

# Specify a model
reign chat -m llama3.2:latest "Write a haiku about recursion"
```

### Browse Models

```bash
# See local models
reign models

# See all models across the network
reign models network
```

### Monitor Usage

```bash
# Developer dashboard - see your credits and usage
reign dev status

# Node operator dashboard - see your earnings
reign node status
```

The dashboard shows you what matters:
- **For Developers:** Credit balance, burn rate, per-model costs, latency insights
- **For Operators:** Earnings, hardware utilization, model performance

## 🛠️ For Developers

Reign is open source and built with Go. Want to extend it or build your own tools?

```bash
# Clone and build
git clone https://github.com/sovereynai/reign.git
cd reign
go build -o reign ./cmd/reign
```

Reign communicates with the Throne daemon over HTTP, making it easy to understand and extend.

## 🌍 Contributing

We welcome contributions! Fork the repo, create a feature branch, and submit a pull request.

**Popular contribution ideas:**
- Streaming chat output
- Session history
- Interactive model playground
- Shell completions
- Export metrics to JSON/CSV

## 📚 Documentation

- **[AI Developer Guide](https://github.com/sovereynai/ai-developer-guide)** - Get started using Sovereyn Network
- **[Node Operator Guide](https://github.com/sovereynai/node-operator-guide)** - Run your own inference node
- **[Architecture Docs](https://github.com/sovereynai/architecture-docs)** - Technical deep dives & design docs

## 📄 License

MIT License - See LICENSE file for details

---

👑 Rule Your AI Destiny!
