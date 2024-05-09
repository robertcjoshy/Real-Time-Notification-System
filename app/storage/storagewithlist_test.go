package storage

import (
	"context"
	"testing"

	"github.com/robert/notification/app/entity"
	"github.com/robert/notification/app/pkg"
)

func BenchmarkMemorywithlist_Push(b *testing.B) {
	mem := Newmemorywithlist(1000) // Create a new memory with list instance

	ctx := context.Background()
	//clientID := 10
	//notification := entity.Unreadmessagenotification{Count: 1} // Initialize notification instance for the benchmark

	b.ResetTimer() // Reset the benchmark timer before running the benchmark loop

	for i := 0; i < b.N; i++ {
		// Run the Push method b.N times to benchmark its performance
		if err := mem.Push(ctx, i, entity.Unreadmessagenotification{Count: i}); err != nil {
			b.Fatalf("Push error: %v", err)
		}
	}

	// Stop the benchmark timer
	b.StopTimer()
	b.Log("for ", b.N, " notifications: ")
	pkg.PrintMemUsage()
}

// You can add more benchmark tests for other methods if needed
