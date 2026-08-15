package main

import "fmt"

// 完整的常量使用示例

// ============ 1. 基本常量 ============

const (
	Pi  = 3.141592653589793
	E   = 2.718281828459045
	Phi = 1.618033988749895 // 黄金比例
)

// ============ 2. 枚举常量 ============

// WeekDay 星期枚举
type WeekDay int

const (
	Monday WeekDay = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (w WeekDay) String() string {
	names := [...]string{
		"Monday", "Tuesday", "Wednesday", "Thursday",
		"Friday", "Saturday", "Sunday",
	}
	return names[w]
}

// HTTPMethod HTTP 方法枚举
type HTTPMethod int

const (
	MethodGet HTTPMethod = iota
	MethodPost
	MethodPut
	MethodPatch
	MethodDelete
	MethodHead
	MethodOptions
)

func (m HTTPMethod) String() string {
	names := [...]string{
		"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
	}
	return names[m]
}

// ============ 3. 位标志常量 ============

// Permission 权限位标志
type Permission uint

const (
	PermRead Permission = 1 << iota
	PermWrite
	PermExecute
)

const (
	PermReadWrite = PermRead | PermWrite
	PermAll       = PermRead | PermWrite | PermExecute
)

func (p Permission) String() string {
	result := ""
	if p&PermRead != 0 {
		result += "r"
	} else {
		result += "-"
	}
	if p&PermWrite != 0 {
		result += "w"
	} else {
		result += "-"
	}
	if p&PermExecute != 0 {
		result += "x"
	} else {
		result += "-"
	}
	return result
}

func hasPermission(userPerm, checkPerm Permission) bool {
	return userPerm&checkPerm == checkPerm
}

// ============ 4. 状态机常量 ============

// OrderStatus 订单状态
type OrderStatus int

const (
	OrderCreated OrderStatus = iota
	OrderPaid
	OrderShipped
	OrderDelivered
	OrderCancelled
	OrderRefunded
)

func (s OrderStatus) String() string {
	names := [...]string{
		"Created", "Paid", "Shipped", "Delivered", "Cancelled", "Refunded",
	}
	return names[s]
}

func (s OrderStatus) CanTransitionTo(target OrderStatus) bool {
	switch s {
	case OrderCreated:
		return target == OrderPaid || target == OrderCancelled
	case OrderPaid:
		return target == OrderShipped || target == OrderCancelled
	case OrderShipped:
		return target == OrderDelivered || target == OrderCancelled
	case OrderDelivered, OrderCancelled, OrderRefunded:
		return false
	default:
		return false
	}
}

// ============ 5. 配置常量 ============

const (
	MaxRetries    = 3
	Timeout       = 30 // 秒
	BufferSize    = 1024
	MaxConnects   = 100
	KeepAliveTime = 60 // 秒
)

// ============ 主函数 ============

func main() {
	fmt.Println("=== Go 常量完整示例 ===\n")

	// 基本常量
	fmt.Println("1. 基本常量:")
	fmt.Printf("   Pi = %.10f\n", Pi)
	fmt.Printf("   E = %.10f\n", E)
	fmt.Printf("   Phi = %.10f\n", Phi)

	// 枚举常量
	fmt.Println("\n2. 枚举常量 (WeekDay):")
	for d := Monday; d <= Sunday; d++ {
		fmt.Printf("   %d: %s\n", d, d)
	}

	// HTTP 方法
	fmt.Println("\n3. HTTP 方法枚举:")
	for m := MethodGet; m <= MethodOptions; m++ {
		fmt.Printf("   %d: %s\n", m, m)
	}

	// 位标志
	fmt.Println("\n4. 位标志 (Permission):")
	userPerm := PermRead | PermWrite
	fmt.Printf("   User permissions: %d (%s)\n", userPerm, userPerm)
	fmt.Printf("   Can read? %v\n", hasPermission(userPerm, PermRead))
	fmt.Printf("   Can write? %v\n", hasPermission(userPerm, PermWrite))
	fmt.Printf("   Can execute? %v\n", hasPermission(userPerm, PermExecute))

	// 状态机
	fmt.Println("\n5. 状态机 (OrderStatus):")
	status := OrderCreated
	fmt.Printf("   Initial: %s\n", status)

	if status.CanTransitionTo(OrderPaid) {
		status = OrderPaid
		fmt.Printf("   -> %s\n", status)
	}
	if status.CanTransitionTo(OrderShipped) {
		status = OrderShipped
		fmt.Printf("   -> %s\n", status)
	}
	if status.CanTransitionTo(OrderDelivered) {
		status = OrderDelivered
		fmt.Printf("   -> %s\n", status)
	}
	fmt.Printf("   Final: %s\n", status)

	// 配置常量
	fmt.Println("\n6. 配置常量:")
	fmt.Printf("   MaxRetries: %d\n", MaxRetries)
	fmt.Printf("   Timeout: %d seconds\n", Timeout)
	fmt.Printf("   BufferSize: %d bytes\n", BufferSize)
	fmt.Printf("   MaxConnects: %d\n", MaxConnects)
	fmt.Printf("   KeepAliveTime: %d seconds\n", KeepAliveTime)

	fmt.Println("\n=== 示例完成 ===")
}
