// 派生摘要纯函数（phase15-03 三派生算法逐字实现）。
//
// 输入约定：events 为 repository 层单一查询返回的三键链 DESC 全序切片
// （occurred_at DESC, created_at DESC, id DESC；索引 i 越小 = 事件越晚）。
// Go 侧不重排，直接消费 SQL 返回序；排序唯一执行位在 repository 层 ORDER BY。
//
// 本函数零 I/O、零时间函数、零全局状态——输入即全部信息，可纯内存单元测试。
//
// 三派生项（N=10 冻结）：
//   - RecentEvents        events[0:min(10,len)]；空集 → 空切片（恒非 nil）
//   - LatestTaskCompleted 自 i=0 起首个 task_completed（三轨同序取最新）；不存在 → nil
//   - CurrentPhase 三态   无 phase_started → 空值（从未开始）；
//                         最新 phase_started 之后存在同 task_key 的
//                         phase_completed → 空值（全部完结）；
//                         否则 key = 该事件 task_key、label = 该事件 title
//
// 两种 current_phase 空值情形输出同型零值（空字符串），不引入第二套
// phase 状态字段（phase15-02 冻结）。
//
// 文件落点：backend/internal/progress/derive.go
package progress

// recentEventsLimit 派生摘要 recent_events 条数上限（N=10，phase15-03 冻结）。
const recentEventsLimit = 10

// DeriveProgressSummary 从三键链 DESC 全序事件切片计算派生摘要（纯函数）。
func DeriveProgressSummary(events []ProgressEventReadResult) ProgressSummary {
	// recent_events：最近 N=10 条三轨混合事件；空集恒构造空切片（非 nil，
	// 空态恒构造约束在 domain 层的承接，brief 装配侧与前端无双套判空）。
	n := min(recentEventsLimit, len(events))
	recent := make([]ProgressEventReadResult, n)
	copy(recent, events[:n])

	// latest_task_completed：自 i=0 起首个 task_completed（索引越小越晚）。
	// 复制后取址：指针指向副本，不受入参切片后续修改影响。
	var latestTaskCompleted *ProgressEventReadResult
	for i := range events {
		if events[i].EventKind == EventKindTaskCompleted {
			event := events[i]
			latestTaskCompleted = &event
			break
		}
	}

	summary := ProgressSummary{
		CurrentPhaseKey:     "",
		CurrentPhaseLabel:   "",
		LatestTaskCompleted: latestTaskCompleted,
		RecentEvents:        recent,
	}

	// current_phase 三态：找最新 phase_started（DESC 序第一个 = 三键链序最新）。
	latestStartedIdx := -1
	for i := range events {
		if events[i].EventKind == EventKindPhaseStarted {
			latestStartedIdx = i
			break
		}
	}
	if latestStartedIdx == -1 {
		return summary // 从未开始：空值
	}

	// 全部完结判定：存在 j < latestStartedIdx（序更晚）的同 task_key phase_completed。
	for j := 0; j < latestStartedIdx; j++ {
		if events[j].EventKind == EventKindPhaseCompleted &&
			events[j].TaskKey == events[latestStartedIdx].TaskKey {
			return summary // 全部完结：空值（与从未开始同型零值）
		}
	}

	// 进行中：key = 该最新 phase_started 的 task_key；label = 其 title。
	summary.CurrentPhaseKey = events[latestStartedIdx].TaskKey
	summary.CurrentPhaseLabel = events[latestStartedIdx].Title
	return summary
}
