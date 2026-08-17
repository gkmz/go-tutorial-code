package main

import "fmt"

// 实战建议 1：接口设计原则（接口隔离原则）

// 不支持 不好：接口太大
type DatabaseBad interface {
	Connect() error
	Close() error
	Query(sql string) ([]interface{}, error)
	Insert(table string, data map[string]interface{}) error
	Update(table string, data map[string]interface{}) error
	Delete(table string, id int) error
	BeginTransaction() error
	Commit() error
	Rollback() error
}

// 已完成 好：拆分成小接口
type Connector interface {
	Connect() error
	Close() error
}

type Querier interface {
	Query(sql string) ([]interface{}, error)
}

type Writer interface {
	Insert(table string, data map[string]interface{}) error
	Update(table string, data map[string]interface{}) error
	Delete(table string, id int) error
}

type Transactioner interface {
	BeginTransaction() error
	Commit() error
	Rollback() error
}

// 组合接口
type Database interface {
	Connector
	Querier
	Writer
	Transactioner
}

// 实现示例
type MySQL struct {
	connected bool
}

func (m *MySQL) Connect() error {
	m.connected = true
	fmt.Println("MySQL connected")
	return nil
}

func (m *MySQL) Close() error {
	m.connected = false
	fmt.Println("MySQL closed")
	return nil
}

func (m *MySQL) Query(sql string) ([]interface{}, error) {
	fmt.Println("Executing query:", sql)
	return []interface{}{}, nil
}

func (m *MySQL) Insert(table string, data map[string]interface{}) error {
	fmt.Printf("Inserting into %s: %v\n", table, data)
	return nil
}

func (m *MySQL) Update(table string, data map[string]interface{}) error {
	fmt.Printf("Updating %s: %v\n", table, data)
	return nil
}

func (m *MySQL) Delete(table string, id int) error {
	fmt.Printf("Deleting from %s where id=%d\n", table, id)
	return nil
}

func (m *MySQL) BeginTransaction() error {
	fmt.Println("Transaction started")
	return nil
}

func (m *MySQL) Commit() error {
	fmt.Println("Transaction committed")
	return nil
}

func (m *MySQL) Rollback() error {
	fmt.Println("Transaction rolled back")
	return nil
}

// 使用小接口的好处：可以只依赖需要的功能
func readData(q Querier) {
	q.Query("SELECT * FROM users")
}

func writeData(w Writer) {
	w.Insert("users", map[string]interface{}{"name": "Alice"})
}

func main() {
	fmt.Println("=== 接口设计原则 ===")

	db := &MySQL{}
	db.Connect()

	// 可以传递给只需要部分功能的函数
	readData(db)
	writeData(db)

	db.Close()
}
