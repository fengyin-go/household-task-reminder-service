package cleanup

type Retry struct{ repo *Repository }

func NewRetry(repo *Repository) *Retry { return &Retry{repo: repo} }

func (r *Retry) Run(id string) error {
	// 执行且仅执行一次清理。部分失败的清理不能自动重试：底层操作可能已经
	// 改动了状态（例如已经清理了一部分标签），盲目重试会导致重复清理。
	// 将错误上抛，由调用方决定如何恢复。
	return r.repo.Delete(id)
}
