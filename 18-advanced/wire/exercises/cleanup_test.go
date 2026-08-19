package exercises

import (
	"reflect"
	"testing"
)

func TestGeneratedCleanupUsesReverseConstructionOrder(t *testing.T) {
	log := NewCleanupLog()
	app, cleanup := InitializeExerciseApp(log)
	if !app.Ready() {
		t.Fatal("generated app is not ready")
	}

	// 上层资源必须先于它依赖的底层资源释放。
	cleanup()
	if got, want := log.Entries(), []string{"resource-b", "resource-a"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("cleanup entries = %v, want %v", got, want)
	}
}
