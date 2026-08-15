package main

import "fmt"

// WeekDay 自定义类型枚举
type WeekDay int

const (
	Monday WeekDay = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
	weekDays // 7，星期总数
)

// String 方法实现 Stringer 接口
func (w WeekDay) String() string {
	names := [...]string{
		"Monday", "Tuesday", "Wednesday", "Thursday",
		"Friday", "Saturday", "Sunday",
	}
	if w < Monday || w > Sunday {
		return "Unknown"
	}
	return names[w]
}

// IsValid 检查是否是合法的星期
func (w WeekDay) IsValid() bool {
	return w >= Monday && w <= Sunday
}

// FromString 从字符串创建 WeekDay
func FromString(s string) (WeekDay, error) {
	switch s {
	case "Monday":
		return Monday, nil
	case "Tuesday":
		return Tuesday, nil
	case "Wednesday":
		return Wednesday, nil
	case "Thursday":
		return Thursday, nil
	case "Friday":
		return Friday, nil
	case "Saturday":
		return Saturday, nil
	case "Sunday":
		return Sunday, nil
	default:
		return -1, fmt.Errorf("invalid weekday: %s", s)
	}
}

// 工作日检查
func isWorkday(day WeekDay) bool {
	return day != Saturday && day != Sunday
}

// 枚举示例
func enumExample() {
	fmt.Println("=== 类型安全枚举 ===")

	day := Monday
	fmt.Printf("Today is %s (value: %d)\n", day, day)

	// 使用 String 方法
	fmt.Printf("All days:\n")
	for d := Monday; d <= Sunday; d++ {
		fmt.Printf("  %d: %s\n", d, d)
	}

	// IsValid 检查
	invalidDay := WeekDay(100)
	fmt.Printf("Is %d valid? %v\n", invalidDay, invalidDay.IsValid())

	// FromString
	day2, err := FromString("Friday")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("FromString('Friday'): %s (value: %d)\n", day2, day2)
	}

	// 工作日检查
	fmt.Printf("Is Monday a workday? %v\n", isWorkday(Monday))
	fmt.Printf("Is Saturday a workday? %v\n", isWorkday(Saturday))
}

// HTTPStatus 状态码枚举
type HTTPStatus int

const (
	StatusOK HTTPStatus = iota + 200
	StatusCreated
	StatusNoContent
	StatusBadRequest
	StatusUnauthorized
	StatusForbidden
	StatusNotFound
	StatusInternalServerError
)

func (s HTTPStatus) String() string {
	names := [...]string{
		"OK", "Created", "NoContent",
		"BadRequest", "Unauthorized", "Forbidden", "NotFound", "InternalServerError",
	}
	idx := s - StatusOK
	if idx < 0 || int(idx) >= len(names) {
		return "Unknown"
	}
	return names[idx]
}

func httpExample() {
	fmt.Println("\n=== HTTP 状态码枚举 ===")

	status := StatusOK
	fmt.Printf("Status: %s (%d)\n", status, status)

	// 遍历所有状态码
	for s := StatusOK; s <= StatusInternalServerError; s++ {
		fmt.Printf("  %d: %s\n", s, s)
	}
}

// OrderStatus 订单状态枚举
type OrderStatus int

const (
	OrderCreated OrderStatus = iota
	OrderPaid
	OrderShipped
	OrderDelivered
	OrderCancelled
	OrderRefunded
)

func (s OrderStatus) CanTransitionTo(target OrderStatus) bool {
	// 定义状态转换规则
	switch s {
	case OrderCreated:
		return target == OrderPaid || target == OrderCancelled
	case OrderPaid:
		return target == OrderShipped || target == OrderCancelled
	case OrderShipped:
		return target == OrderDelivered || target == OrderCancelled
	case OrderDelivered:
		return false // 已完成，不能再转换
	case OrderCancelled:
		return false // 已取消，不能再转换
	case OrderRefunded:
		return false // 已退款，不能再转换
	default:
		return false
	}
}

func orderExample() {
	fmt.Println("\n=== 订单状态机 ===")

	status := OrderCreated
	fmt.Printf("Initial status: %d\n", status)

	// 状态转换
	if status.CanTransitionTo(OrderPaid) {
		status = OrderPaid
		fmt.Printf("Transitioned to: %d\n", status)
	}

	if status.CanTransitionTo(OrderShipped) {
		status = OrderShipped
		fmt.Printf("Transitioned to: %d\n", status)
	}

	// 尝试非法转换
	if !status.CanTransitionTo(OrderCancelled) {
		fmt.Println("Cannot transition to Cancelled from Delivered")
	}
}

func main() {
	enumExample()
	httpExample()
	orderExample()
}
