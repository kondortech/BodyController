package nutrition

import (
	"context"
	"github.com/kirvader/BodyController/services/nutrition/server/source"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := source.serve(ctx); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
