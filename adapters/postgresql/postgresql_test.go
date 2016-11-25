package postgresql

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReplacePlaceholders(t *testing.T) {
	Convey("Given a SQL string containing placeholders", t, func() {
		sql := "SELECT id, dummy FROM dummies WHERE id > ? AND id < ?"
		Convey("ReplacePlaceholders change all placeholders with $xx", func() {
			sqlWithNewPlaceholders := Adapter.ReplacePlaceholders("?", sql)
			So(sqlWithNewPlaceholders, ShouldEqual, "SELECT id, dummy FROM dummies WHERE id > $1 AND id < $2")
		})
	})
}