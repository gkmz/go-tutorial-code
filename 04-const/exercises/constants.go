// Package exercises 提供常量章节练习的参考实现。
package exercises

import "fmt"

// HTTPStatus 表示 HTTP 状态码练习中的状态。
type HTTPStatus int

const (
	// StatusOK 表示请求成功。
	StatusOK HTTPStatus = 200
	// StatusCreated 表示资源创建成功。
	StatusCreated HTTPStatus = 201
	// StatusNoContent 表示请求成功但没有响应内容。
	StatusNoContent HTTPStatus = 204
	// StatusBadRequest 表示请求无效。
	StatusBadRequest HTTPStatus = 400
	// StatusUnauthorized 表示缺少有效身份凭据。
	StatusUnauthorized HTTPStatus = 401
	// StatusForbidden 表示服务器拒绝执行请求。
	StatusForbidden HTTPStatus = 403
	// StatusNotFound 表示资源不存在。
	StatusNotFound HTTPStatus = 404
	// StatusInternalServerError 表示服务器内部错误。
	StatusInternalServerError HTTPStatus = 500
)

// String 返回 HTTP 状态码的名称。
func (s HTTPStatus) String() string {
	names := map[HTTPStatus]string{
		StatusOK: "OK", StatusCreated: "Created", StatusNoContent: "NoContent",
		StatusBadRequest: "BadRequest", StatusUnauthorized: "Unauthorized",
		StatusForbidden: "Forbidden", StatusNotFound: "NotFound",
		StatusInternalServerError: "InternalServerError",
	}
	if name, ok := names[s]; ok {
		return name
	}
	return "Unknown"
}

// FileMode 表示可组合的文件权限。
type FileMode uint

const (
	// ModeRead 表示读权限。
	ModeRead FileMode = 1 << iota
	// ModeWrite 表示写权限。
	ModeWrite
	// ModeExecute 表示执行权限。
	ModeExecute
)

// CanRead 判断是否包含读权限。
func (m FileMode) CanRead() bool { return m&ModeRead != 0 }

// CanWrite 判断是否包含写权限。
func (m FileMode) CanWrite() bool { return m&ModeWrite != 0 }

// CanExecute 判断是否包含执行权限。
func (m FileMode) CanExecute() bool { return m&ModeExecute != 0 }

// Color 表示颜色枚举。
type Color int

const (
	// Red 表示红色。
	Red Color = iota
	// Green 表示绿色。
	Green
	// Blue 表示蓝色。
	Blue
)

// String 返回颜色名称。
func (c Color) String() string {
	if !c.IsValid() {
		return "Unknown"
	}
	return [...]string{"Red", "Green", "Blue"}[c]
}

// IsValid 判断颜色值是否有效。
func (c Color) IsValid() bool { return c >= Red && c <= Blue }

// FromString 根据颜色名称返回颜色枚举。
func FromString(value string) (Color, error) {
	for color := Red; color <= Blue; color++ {
		if color.String() == value {
			return color, nil
		}
	}
	return 0, fmt.Errorf("invalid color: %s", value)
}

// OrderStatus 表示订单状态。
type OrderStatus int

const (
	// OrderCreated 表示订单已创建。
	OrderCreated OrderStatus = iota
	// OrderPaid 表示订单已支付。
	OrderPaid
	// OrderShipped 表示订单已发货。
	OrderShipped
	// OrderDelivered 表示订单已送达。
	OrderDelivered
	// OrderCancelled 表示订单已取消。
	OrderCancelled
)

// CanTransitionTo 判断订单是否允许转换到目标状态。
func (s OrderStatus) CanTransitionTo(target OrderStatus) bool {
	switch s {
	case OrderCreated:
		return target == OrderPaid || target == OrderCancelled
	case OrderPaid:
		return target == OrderShipped || target == OrderCancelled
	case OrderShipped:
		return target == OrderDelivered
	default:
		return false
	}
}
