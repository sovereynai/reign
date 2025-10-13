package client

// MockDeveloperStats generates sample developer dashboard data
func MockDeveloperStats() *DashboardStats {
	return &DashboardStats{
		Role:   "developer",
		Uptime: "47d 3h 12m",
		Version: VersionInfo{
			Version:   "v0.2.0",
			Commit:    "ed3c5bd1234567890abcdef",
			BuildTime: "2025-10-13T10:16:00Z",
		},
		Developer: &DeveloperStats{
			Credits: CreditStats{
				Balance:      847,
				TodaySpent:   52,
				TodayEarned:  15,
				BurnRate:     12.3,
				RunwayDays:   68,
				TrendPercent: -8.0,
			},
			Inference: InferenceStats{
				Today:       47,
				WeekAvg:     12.3,
				Total:       150,
				SuccessRate: 98.7,
				Failures:    2,
			},
			Performance: PerformanceStats{
				AvgLatencyMs:   156,
				P50LatencyMs:   127,
				P95LatencyMs:   342,
				P99LatencyMs:   1200,
				LocalPercent:   82,
				NetworkPercent: 18,
			},
			Models: []ModelUsage{
				{
					Name:          "llama3.2:3b",
					RequestsToday: 47,
					WeekAvg:       12.3,
					AvgLatencyMs:  127,
					CreditsSpent:  23.5,
				},
				{
					Name:          "codellama",
					RequestsToday: 12,
					WeekAvg:       8.1,
					AvgLatencyMs:  245,
					CreditsSpent:  9.2,
				},
				{
					Name:          "llama3.2:70b",
					RequestsToday: 2,
					WeekAvg:       0.3,
					AvgLatencyMs:  1200,
					CreditsSpent:  8.1,
				},
			},
			Insights: []string{
				"89% of requests use llama3.2:3b - consider pulling locally",
				"Peak usage: 2-4pm EST - queue times +45% vs off-peak",
				"codellama: 4 pending requests - high network load",
			},
		},
		Network: NetworkStats{
			PeersConnected:  47,
			ModelsAvailable: 23,
			QueueDepth:      12,
			EstWaitSec:      2.3,
			DataRelayedGB:   2.3,
		},
	}
}

// MockOperatorStats generates sample operator dashboard data
func MockOperatorStats() *DashboardStats {
	return &DashboardStats{
		Role:   "operator",
		Uptime: "47d 3h 12m",
		Version: VersionInfo{
			Version:   "v0.2.0",
			Commit:    "ed3c5bd1234567890abcdef",
			BuildTime: "2025-10-13T10:16:00Z",
		},
		Operator: &OperatorStats{
			Earnings: EarningsStats{
				Today:     127.3,
				TodayUSD:  2.14,
				ThisWeek:  892.1,
				WeekTrend: 12.0,
				AllTime:   15847,
				Pending:   23.1,
				Rank:      482,
				TotalNodes: 4891,
				Breakdown: EarningsBreakdown{
					Inference: 94.2,
					Bandwidth: 21.3,
					Storage:   11.8,
					Bonus:     0.0,
				},
			},
			Workload: WorkloadStats{
				RequestsServed: 847,
				SuccessRate:    99.4,
				Failures:       5,
				AvgLatencyMs:   112,
				PeakHour:       "2pm EST",
				PeakRequests:   94,
			},
			Hardware: HardwareStats{
				GPU: ResourceUsage{
					Percent: 78,
					Used:    "18.7GB",
					Total:   "24GB",
					Details: "NVIDIA RTX 4090, 24GB",
				},
				CPU: ResourceUsage{
					Percent: 32,
					Used:    "10c",
					Total:   "32t",
					Details: "AMD Ryzen 9 5950X, 16c/32t",
				},
				RAM: ResourceUsage{
					Percent: 48,
					Used:    "18.2GB",
					Total:   "64GB",
					Details: "DDR4-3600",
				},
				Disk: ResourceUsage{
					Percent: 23,
					Used:    "1.2TB",
					Total:   "4TB",
					Details: "NVMe SSD",
				},
				Temperature: TempStats{
					GPU:    67,
					CPU:    54,
					Status: "healthy",
				},
				PowerWatts:     340,
				PowerCostDaily: 0.82,
			},
			ModelsServed: []ModelServed{
				{
					Name:         "llama3.2:3b",
					Requests:     712,
					AvgLatencyMs: 89,
					Revenue:      89.2,
					Status:       "active",
				},
				{
					Name:         "codellama",
					Requests:     98,
					AvgLatencyMs: 134,
					Revenue:      24.1,
					Status:       "active",
				},
				{
					Name:         "llama3.2:70b",
					Requests:     37,
					AvgLatencyMs: 1100,
					Revenue:      14.0,
					Status:       "active",
				},
				{
					Name:         "mistral:7b",
					Requests:     0,
					AvgLatencyMs: 0,
					Revenue:      0,
					Status:       "idle",
				},
			},
			Reputation: ReputationStats{
				Score:        4.7,
				MaxScore:     5.0,
				Ratings:      387,
				UptimeStreak: 47,
			},
			Alerts: []Alert{
				{
					Level:   "warning",
					Message: "GPU running hot (67°C) - optimize cooling for +3% efficiency",
				},
				{
					Level:   "info",
					Message: "mistral:7b unused for 7d - free 4.1GB by removing",
				},
				{
					Level:   "info",
					Message: "Peak efficiency: 2-6pm EST - consider prioritizing availability",
				},
				{
					Level:   "info",
					Message: "Uptime streak: 47 days - maintain for +15% bonus next week",
				},
			},
		},
		Network: NetworkStats{
			PeersConnected:  47,
			ModelsAvailable: 23,
			QueueDepth:      12,
			EstWaitSec:      2.3,
			DataRelayedGB:   2.3,
		},
	}
}

// MockBothStats generates sample data for a user who is both dev and operator
func MockBothStats() *DashboardStats {
	devStats := MockDeveloperStats()
	opStats := MockOperatorStats()

	return &DashboardStats{
		Role:      "both",
		Uptime:    "47d 3h 12m",
		Version:   devStats.Version,
		Developer: devStats.Developer,
		Operator:  opStats.Operator,
		Network:   devStats.Network,
	}
}
