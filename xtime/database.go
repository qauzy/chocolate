package xtime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Scan an interface used by Scan in package database/sql for Scanning value from database to local golang variable.
func (c *Xtime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*c = Xtime{time: value, loc: time.Local}
		return nil
	}
	return fmt.Errorf("can not convert %v to xtime", v)
}

// Value the interface providing the Value method for package database/sql/driver.
func (c Xtime) Value() (driver.Value, error) {
	if c.IsZero() {
		return nil, nil
	}
	return c.time, nil
}
