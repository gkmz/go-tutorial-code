//go:build wireinject

package exercises

import "github.com/google/wire"

// InitializeExerciseApp 声明带两个清理函数的练习依赖图。
func InitializeExerciseApp(log *CleanupLog) (*ExerciseApp, func()) {
	wire.Build(NewResourceA, NewResourceB, NewExerciseApp)
	return nil, nil
}
