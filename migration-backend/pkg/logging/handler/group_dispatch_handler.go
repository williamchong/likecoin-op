package handler

import (
	"context"
	"log/slog"
)

type groupDispatchHandler struct {
	groupMap       map[string]slog.Handler
	initialHandler slog.Handler
}

func NewGroupDispatchHandler(
	groupMap map[string]slog.Handler,
	initialHandler slog.Handler,
) slog.Handler {
	return &groupDispatchHandler{
		groupMap,
		initialHandler,
	}
}

func (h *groupDispatchHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.initialHandler.Enabled(ctx, level)
}

func (h *groupDispatchHandler) Handle(ctx context.Context, record slog.Record) error {
	return h.initialHandler.Handle(ctx, record)
}

func (h *groupDispatchHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	initialHandler := h.initialHandler.WithAttrs(attrs)
	return &groupDispatchHandler{
		h.groupMap,
		initialHandler,
	}
}

func (h *groupDispatchHandler) WithGroup(group string) slog.Handler {
	selectedHandler, ok := h.groupMap[group]
	if ok {
		return selectedHandler.WithGroup(group)
	}
	initialHandler := h.initialHandler.WithGroup(group)
	return &groupDispatchHandler{
		h.groupMap,
		initialHandler,
	}
}
