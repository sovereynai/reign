# 👑 Reign

**The official CLI for Sovereyn's distributed intelligence network.**

Reign is the user-friendly command-line interface for interacting with the Sovereyn network. Submit AI inference jobs, manage models, and monitor network status—all from your terminal.

## ✨ Features

- 🤖 **Chat with AI models** - Submit inference jobs with beautiful output
- 📊 **Dual Dashboards** - Rich, information-dense views for AI Devs and Node Operators
- 📈 **Real-time Metrics** - Track credits, earnings, performance, and hardware
- 💡 **Smart Insights** - AI-powered optimization recommendations
- 📦 **Model management** - List and pull Ollama models
- 🎨 **Beautiful TUI** - Rich terminal UI with colors and formatting
- ⚡ **Zero configuration** - Auto-discovers local throne daemon

## 🚀 Quick Start

### Prerequisites

1. **Throne daemon running**:
   ```bash
   throne serve
   ```

2. **Ollama installed** (for LLM models):
   ```bash
   # macOS
   brew install ollama
   ollama serve

   # Pull a model
   ollama pull llama3.2:3b
   ```

### Installation

```bash
# Build from source
go build -o reign ./cmd/reign

# Run it
./reign --help
```

## 📖 Usage

### 👨‍💻 For AI Developers

```bash
# View your developer dashboard
reign dev status

# Smart dashboard auto-detects your role
reign status

# Chat with AI models
reign chat "Explain quantum computing in one sentence"
reign chat -m llama3.2:latest "Write a haiku about recursion"

# Coming soon
reign dev history     # View request history & replay
reign dev optimize    # Get cost reduction suggestions
reign dev playground  # Interactive model testing
```

**Developer Dashboard shows:**
- 💰 Credit balance, burn rate, and runway
- 🤖 Per-model inference metrics
- ⚡ Latency percentiles (p50, p95, p99)
- 📊 Success rates and failures
- 🎯 Smart optimization insights
- 🌍 Network health and queue depth

### 🖥️ For Node Operators

```bash
# View your operator dashboard
reign node status

# See earnings and hardware stats
reign status

# Coming soon
reign node earnings   # Detailed revenue breakdown
reign node optimize   # Hardware tuning recommendations
reign node models     # Add/remove models based on demand
reign node peers      # Network connections & health
reign node logs       # Real-time inference log stream
```

**Operator Dashboard shows:**
- 💰 Earnings (today, week, all-time) with trends
- 📈 Revenue breakdown (inference, bandwidth, storage, bonuses)
- 🔥 Workload stats (requests, success rate, peak hours)
- 🖥️ Hardware utilization (GPU, CPU, RAM, disk, temps)
- 📦 Model performance (requests, latency, revenue per model)
- 🌍 Network participation and reputation score
- ⚠️ Alerts and optimization suggestions

### Model Management

```bash
# List available models
reign models

# Pull a new model (coming soon)
reign models pull llama3.2:latest
```

### Status & Version

```bash
# Smart status (auto-detects role)
reign status

# Show version info
reign version
```

## 🏗️ Architecture

Reign is a **pure HTTP client** that communicates with the throne daemon:

```
┌─────────┐         ┌─────────┐         ┌────────┐
│  Reign  │  HTTP   │ Throne  │  API    │ Ollama │
│   CLI   ├────────>│ Daemon  ├────────>│  LLM   │
└─────────┘         └─────────┘         └────────┘
                          │
                          ├──> Ledger (logging)
                          ├──> Credits (tracking)
                          └──> Network (p2p)
```

**Key Design:**
- **Open Source** - Community-driven, forkable
- **Zero Secrets** - No proprietary logic (that's in throne)
- **Pure Client** - Just HTTP calls to throne's API
- **Extensible** - Easy to add new commands

## 🛠️ Development

### Project Structure

```
reign/
├── cmd/reign/           # Main CLI entry point
├── internal/
│   ├── client/          # HTTP client for throne API
│   ├── config/          # Config & daemon discovery
│   ├── ui/              # Rich dashboard rendering
│   └── bootstrap/       # First-run setup
└── README.md
```

### Building

```bash
go mod tidy
go build -o reign ./cmd/reign
```

### Testing

```bash
# Start throne daemon first
throne serve

# Then test reign commands
./reign status
./reign models
./reign chat "Hello world"
```

## 🌍 Contributing

Reign is open source! We welcome contributions:

1. Fork the repo
2. Create a feature branch
3. Submit a pull request

**Ideas for contributions:**
- Streaming chat output with real-time progress
- Session history and conversation replay
- Interactive model playground
- Hardware monitoring enhancements
- Shell completions
- Config file support
- Export metrics to JSON/CSV

## 📚 API Reference

Reign communicates with throne's HTTP API:

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/healthz` | GET | Health check |
| `/version` | GET | Version info |
| `/chat` | POST | Chat completion |
| `/generate` | POST | Text generation |
| `/ollama/models` | GET | List models |
| `/stats/dashboard` | GET | Comprehensive dashboard statistics |

See throne's API documentation for full details.

## 🔗 Links

- **Sovereyn Website**: https://sovereyn.ai (coming soon)
- **Throne Daemon**: (proprietary - binary releases only)
- **Documentation**: https://docs.sovereyn.ai (coming soon)

## 📄 License

MIT License - See LICENSE file for details

## 🙏 Acknowledgments

Built with:
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling

---

**Made with ❤️ by the Sovereyn community**
