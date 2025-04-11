package event

import (
	"fmt"
	"strings"

	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos/model"
)

const HEADER = "_migration"

func MakeMemoFromEvent(events []model.Event) string {
	parts := make([]string, 1, len(events))
	parts[0] = HEADER

	for _, event := range events {
		if event.Memo != "" {
			parts = append(parts, fmt.Sprintf(`%s
Tx: %s`, event.Memo, event.TxHash))
		}
	}

	return strings.Join(parts, "\n\n")
}
