package main

import "fmt"

// Permission 位标志枚举
type Permission uint

const (
	PermRead    Permission = 1 << iota // 1 (001)
	PermWrite                          // 2 (010)
	PermExecute                        // 4 (100)
)

// 组合权限
const (
	PermReadWrite    = PermRead | PermWrite               // 3 (011)
	PermReadExecute  = PermRead | PermExecute             // 5 (101)
	PermWriteExecute = PermWrite | PermExecute            // 6 (110)
	PermAll          = PermRead | PermWrite | PermExecute // 7 (111)
)

// 检查权限
func hasPermission(userPerm, checkPerm Permission) bool {
	return userPerm&checkPerm == checkPerm
}

// 添加权限
func addPermission(userPerm, addPerm Permission) Permission {
	return userPerm | addPerm
}

// 移除权限
func removePermission(userPerm, removePerm Permission) Permission {
	return userPerm &^ removePerm
}

// Permission 的 String 方法
func (p Permission) String() string {
	var names []string
	if p&PermRead != 0 {
		names = append(names, "Read")
	}
	if p&PermWrite != 0 {
		names = append(names, "Write")
	}
	if p&PermExecute != 0 {
		names = append(names, "Execute")
	}
	if len(names) == 0 {
		return "None"
	}
	result := ""
	for i, name := range names {
		if i > 0 {
			result += " | "
		}
		result += name
	}
	return result
}

func bitFlagsExample() {
	fmt.Println("=== 位标志枚举 ===")

	// 用户权限，读写权限
	userPerm := PermRead | PermWrite
	fmt.Printf("User permissions: %d (%s)\n", userPerm, userPerm)

	// 检查权限
	fmt.Printf("Can read? %v\n", hasPermission(userPerm, PermRead))       // true
	fmt.Printf("Can write? %v\n", hasPermission(userPerm, PermWrite))     // true
	fmt.Printf("Can execute? %v\n", hasPermission(userPerm, PermExecute)) // false

	// 添加权限
	userPerm = addPermission(userPerm, PermExecute)
	fmt.Printf("After adding execute: %d (%s)\n", userPerm, userPerm)

	// 移除权限
	userPerm = removePermission(userPerm, PermWrite)
	fmt.Printf("After removing write: %d (%s)\n", userPerm, userPerm)
}

// FileMode 文件权限
type FileMode uint

const (
	ModeRead    FileMode = 1 << iota // 1
	ModeWrite                        // 2
	ModeExecute                      // 4
)

const (
	ModeReadWrite = ModeRead | ModeWrite
	ModeAll       = ModeRead | ModeWrite | ModeExecute
)

func (m FileMode) CanRead() bool {
	return m&ModeRead != 0
}

func (m FileMode) CanWrite() bool {
	return m&ModeWrite != 0
}

func (m FileMode) CanExecute() bool {
	return m&ModeExecute != 0
}

func (m FileMode) String() string {
	result := ""
	if m.CanRead() {
		result += "r"
	} else {
		result += "-"
	}
	if m.CanWrite() {
		result += "w"
	} else {
		result += "-"
	}
	if m.CanExecute() {
		result += "x"
	} else {
		result += "-"
	}
	return result
}

func fileModeExample() {
	fmt.Println("\n=== 文件权限 ===")

	// 读写权限
	mode := ModeReadWrite
	fmt.Printf("Mode: %d (%s)\n", mode, mode)

	// 检查权限
	fmt.Printf("Can read? %v\n", mode.CanRead())       // true
	fmt.Printf("Can write? %v\n", mode.CanWrite())     // true
	fmt.Printf("Can execute? %v\n", mode.CanExecute()) // false
}

// LogLevel 日志级别
type LogLevel uint

const (
	LogDebug   LogLevel = 1 << iota // 1
	LogInfo                         // 2
	LogWarning                      // 4
	LogError                        // 8
)

const LogAll = LogDebug | LogInfo | LogWarning | LogError

func (l LogLevel) String() string {
	switch l {
	case LogDebug:
		return "DEBUG"
	case LogInfo:
		return "INFO"
	case LogWarning:
		return "WARNING"
	case LogError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

func logLevelExample() {
	fmt.Println("\n=== 日志级别 ===")

	level := LogInfo | LogError
	fmt.Printf("Current level: %d\n", level)

	// 检查级别
	if level&LogDebug != 0 {
		fmt.Println("  [DEBUG] Debug message")
	}
	if level&LogInfo != 0 {
		fmt.Println("  [INFO] Info message")
	}
	if level&LogWarning != 0 {
		fmt.Println("  [WARNING] Warning message")
	}
	if level&LogError != 0 {
		fmt.Println("  [ERROR] Error message")
	}
}

func main() {
	bitFlagsExample()
	fileModeExample()
	logLevelExample()
}
