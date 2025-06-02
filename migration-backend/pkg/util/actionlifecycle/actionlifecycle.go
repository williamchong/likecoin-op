package actionlifecycle

import "context"

type ActionLifecycle[SucResp any, ActionModel any] interface {
	Begin(ctx context.Context) (*ActionModel, error)
	Success(ctx context.Context, sucResp *SucResp) (*ActionModel, error)
	Failed(ctx context.Context, err error) (*ActionModel, error)
}

func WithActionLifecycle[SucResp any, ActionModel any, LC ActionLifecycle[SucResp, ActionModel]](
	ctx context.Context,
	lc LC,
	fn func(ctx context.Context) (*SucResp, error),
) (*ActionModel, error) {
	_, err := lc.Begin(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := fn(ctx)
	if err != nil {
		return lc.Failed(ctx, err)
	}
	return lc.Success(ctx, resp)
}
