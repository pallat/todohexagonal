package todo

import (
	"testing"
)

func TestCreateTodoNotAllowSleepTask(t *testing.T) {
	// db, mock, err := sqlmock.New()
	// _ = mock
	// _ = err
	// mock.ExpectQuery("select sqlite_version()").WillReturnRows(mock.NewRows([]string{"version"}).AddRow("3.8.10"))

	// dialector := &sqlite.Dialector{
	// 	DSN:        "sqlmock.db",
	// 	DriverName: "sqlmock",
	// 	Conn:       db,
	// }

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // Slow SQL threshold
	// 		LogLevel:                  logger.Info, // Log level
	// 		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	// 		Colorful:                  true,        // Disable color
	// 	},
	// )

	// gdb, _ := gorm.Open(dialector, &gorm.Config{Logger: newLogger})
	// handler := NewTodoHandler(gdb)

	handler := NewTodoHandler(&TestDB{})
	c := &TestContext{}

	handler.NewTask(c)

	want := "not allowed"

	if want != c.v["error"] {
		t.Errorf("want %s but get %s\n", want, c.v["error"])
	}
}

type TestDB struct{}

func (TestDB) New(*Todo) error {
	return nil
}

type TestContext struct {
	v map[string]interface{}
}

func (TestContext) Bind(v interface{}) error {
	*v.(*Todo) = Todo{
		Title: "sleep",
	}
	return nil
}
func (c *TestContext) JSON(code int, v interface{}) {
	c.v = v.(map[string]interface{})
}
func (TestContext) TransactionID() string {
	return "TestTransactionID"
}
func (TestContext) Audience() string {
	return "Unit Test"
}
