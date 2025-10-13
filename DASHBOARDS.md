# 📊 Reign Dashboards

Beautiful, information-dense dashboards for two distinct user personas in the Sovereyn network.

## 🎯 Design Philosophy

Reign serves **two types of users** with distinct needs:

### 👨‍💻 AI Developers
**Focus**: Consumption, experimentation, cost optimization

They care about:
- ✅ Latency and performance
- ✅ Cost per request
- ✅ Model availability
- ✅ Request success rates
- ✅ Queue times

### 🖥️ Node Operators
**Focus**: Earning, efficiency, uptime, ROI

They care about:
- ✅ Credits earned
- ✅ Hardware efficiency
- ✅ Uptime & reputation
- ✅ Model demand
- ✅ Cost vs earnings

## 👨‍💻 AI Developer Dashboard

```bash
reign dev status
```

### What You See

**🤖 Inference Metrics**
- Per-model usage statistics (requests, average latency, credits spent)
- 7-day trends to understand usage patterns
- Model comparison at a glance

**💰 Credits & Usage**
- Current balance with today's activity
- Burn rate and runway calculations
- Cost trends (up/down vs last week)

**⚡ Performance**
- Average latency with percentiles (p50, p95, p99)
- Success rate tracking
- Local vs network preference split

**🎯 Smart Insights**
- AI-powered optimization suggestions
- Peak usage time recommendations
- Cost-saving opportunities

**📊 Network Health**
- Peer count and model availability
- Queue depth with visual progress bar
- Estimated wait times

### Example Output

```
╭────────────────────────────────────────────────────────────────────────╮
│  👑 REIGN - AI Developer Dashboard     v0.2.0                          │
│                                                                        │
│ 🤖 INFERENCE METRICS                                                   │
│   Model           Today      7d Avg       Latency    Credits           │
│   llama3.2:3b     47 req     12.3/day     127ms      23.5c             │
│   codellama       12 req     8.1/day      245ms      9.2c              │
│   llama3.2:70b    2 req      0.3/day      1200ms     8.1c              │
│                                                                        │
│ 💰 CREDITS & USAGE                                                     │
│   Balance:  847 credits      (-52c today, +15c earned)                 │
│   Burn Rate:     ~12.3 credits/day                                     │
│   Runway:       68 days at current usage                               │
│   Cost Trend:       ▼ 8% vs last week                                  │
│                                                                        │
│ ⚡ PERFORMANCE                                                         │
│   Avg Latency:  156ms  (p50: 127ms, p95: 342ms, p99: 1200ms)           │
│   Success Rate: 98.7%  (2 failures in 150 requests)                    │
│   Preferred:    Local: 82% | Network: 18%                              │
│                                                                        │
│ 🎯 SMART INSIGHTS                                                      │
│   → 89% of requests use llama3.2:3b - consider pulling locally         │
│   → Peak usage: 2-4pm EST - queue times +45% vs off-peak               │
│   → codellama: 4 pending requests - high network load                  │
│                                                                        │
│ 📊 NETWORK HEALTH                                                      │
│   Available Peers: 47 online | Models: 23                              │
│   Queue Depth:     [██░░░░░░░░] | Est Wait: 2.3s                       │
╰────────────────────────────────────────────────────────────────────────╯
```

## 🖥️ Node Operator Dashboard

```bash
reign node status
```

### What You See

**💰 Earnings & Contribution**
- Today, this week, and all-time earnings
- Pending settlements
- Global rank among all nodes
- USD conversion for today's earnings

**📈 Revenue Breakdown**
- Earnings by source (inference, bandwidth, storage, bonuses)
- Percentage contribution of each revenue stream

**🔥 Workload (Last 24h)**
- Requests served and success rate
- Average latency (with target benchmarks)
- Peak hour analysis

**🖥️ Hardware Utilization**
- GPU, CPU, RAM, Disk usage with progress bars
- Temperature monitoring with color-coded status
- Power consumption and cost estimates

**📦 Models Served**
- Per-model statistics (requests, latency, revenue)
- Status indicators (active, idle, warning)
- Performance at a glance

**🌍 Network Participation**
- Connected peers
- Data relayed (with earnings)
- Job queue status
- Reputation score with star rating

**⚠️ Alerts & Optimization**
- Real-time alerts (temperature, unused models)
- Actionable optimization suggestions
- Uptime streak tracking for bonuses

### Example Output

```
╭────────────────────────────────────────────────────────────────────────╮
│  🏛️  THRONE - Node Operator Dashboard     Uptime: 47d 3h 12m           │
│                                                                        │
│ 💰 EARNINGS & CONTRIBUTION                                             │
│   Today:        +127.3 credits    (Est: $2.14 USD)                     │
│   This Week:    +892.1 credits    ▲12% vs last week                    │
│   All Time:     15847 credits    (since Jan 3, 2025)                   │
│   Pending:      23.1 credits      (settlement in ~2h)                  │
│   Rank:         #482 / 4891 nodes globally                             │
│                                                                        │
│ 📈 REVENUE BREAKDOWN                                                   │
│   Inference:    94.2 credits   (74% of earnings)                       │
│   Bandwidth:    21.3 credits   (17% of earnings)                       │
│   Storage:      11.8 credits   (9% of earnings)                        │
│   Bonus:        0.0 credits   (0% of earnings)                         │
│                                                                        │
│ 🔥 WORKLOAD (Last 24h)                                                 │
│   Requests Served:  847 inferences                                     │
│   Success Rate:     99.4%  (5 failures)                                │
│   Avg Latency:      112ms (target: <150ms for premium tier)            │
│   Peak Hour:        2pm EST - 94 req/hr                                │
│                                                                        │
│ 🖥️  HARDWARE UTILIZATION                                               │
│   GPU:   [███████░░░] 78% (NVIDIA RTX 4090, 24GB)                      │
│   CPU:   [███░░░░░░░] 32% (AMD Ryzen 9 5950X, 16c/32t)                 │
│   RAM:   [████░░░░░░] 48% (18.2GB / 64GB)                              │
│   Disk:  [██░░░░░░░░] 23% (1.2TB / 4TB)                                │
│   Temp:  GPU: 67°C  |  CPU: 54°C  (healthy)                            │
│   Power: ~340W avg  (est $0.82/day @ $0.12/kWh)                        │
│                                                                        │
│ 📦 MODELS SERVED                                                       │
│   ✓ llama3.2:3b     712 reqs  |  Avg: 89ms   |  Rev: 89.2c             │
│   ✓ codellama       98 reqs  |  Avg: 134ms   |  Rev: 24.1c             │
│   ✓ llama3.2:70b    37 reqs  |  Avg: 1100ms   |  Rev: 14.0c            │
│   ○ mistral:7b      0 reqs  |  Avg: 0ms   |  Rev: 0.0c                 │
│                                                                        │
│ 🌍 NETWORK PARTICIPATION                                               │
│   Peers Connected:  47 nodes                                           │
│   Data Relayed:     2.30 GB (earning +8.0c)                            │
│   Jobs Queued:      12 in network | 3 assigned to you                  │
│   Reputation:       ★★★★☆ 4.7/5 (387 ratings)                          │
│                                                                        │
│ ⚠️  ALERTS & OPTIMIZATION                                              │
│   → GPU running hot (67°C) - optimize cooling for +3% efficiency       │
│   ✓ mistral:7b unused for 7d - free 4.1GB by removing                  │
│   ✓ Peak efficiency: 2-6pm EST - consider prioritizing availability    │
│   ✓ Uptime streak: 47 days - maintain for +15% bonus next week         │
╰────────────────────────────────────────────────────────────────────────╯
```

## 🎨 Visual Design Elements

### Progress Bars
```
[████████░░] 78%  - Green (healthy, < 50%)
[███████░░░] 68%  - Yellow (warning, 50-80%)
[█████████░] 92%  - Red (critical, > 80%)
```

### Star Ratings
```
★★★★★ 5.0  - Perfect
★★★★☆ 4.7  - Excellent
★★★☆☆ 3.2  - Good
```

### Status Indicators
```
✓ Active   - Green
⚠ Warning  - Yellow
✗ Error    - Red
○ Idle     - Gray
```

### Trend Indicators
```
▲ 12%  - Up (green for earnings, red for costs)
▼ 8%   - Down (red for earnings, green for costs)
```

## 🚀 Command Structure

### Smart Auto-Detection
```bash
reign status
# Automatically shows the right dashboard based on your activity
```

### Explicit Dashboards
```bash
reign dev status    # Force developer view
reign node status   # Force operator view
```

### Coming Soon
```bash
# Developer commands
reign dev history       # Request history with replay
reign dev optimize      # Cost optimization tips
reign dev playground    # Interactive model testing
reign dev limits        # Rate limits and quotas

# Operator commands
reign node earnings     # Detailed revenue analytics
reign node optimize     # Hardware tuning recommendations
reign node models       # Model management based on demand
reign node peers        # Network connection health
reign node logs         # Real-time inference logs
```

## 🔧 Technical Implementation

### Data Structure
The dashboards consume the `/stats/dashboard` endpoint which returns:

```json
{
  "role": "developer|operator|both",
  "uptime": "47d 3h 12m",
  "developer": { /* DeveloperStats */ },
  "operator": { /* OperatorStats */ },
  "network": { /* NetworkStats */ },
  "version": { /* VersionInfo */ }
}
```

### Rendering
- Built with [Lipgloss](https://github.com/charmbracelet/lipgloss) for beautiful styling
- Responsive layout that adapts to terminal width
- Color-coded information (success=green, warning=yellow, error=red)
- Information density optimized for 80+ column terminals

### Role Detection
Throne daemon automatically detects user role based on:
1. **Has GPU + serving models** → Node Operator
2. **Only making requests** → AI Developer
3. **Both** → Show combined view or let user choose

## 🎯 Information Density

Both dashboards are designed to be **scannable** and **information-dense**:

- ✅ No unnecessary whitespace
- ✅ Tables for multi-row data
- ✅ Visual progress bars for percentages
- ✅ Color coding for quick status recognition
- ✅ Trends and comparisons (vs last week, vs targets)
- ✅ Actionable insights, not just raw data

## 📚 Learn More

- See [README.md](README.md) for installation and basic usage
- Run `reign --help` for all available commands
- Use `reign demo dev` or `reign demo node` to see sample dashboards

---

**Made with ❤️ for the Sovereyn community**
