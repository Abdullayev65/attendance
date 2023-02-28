package models

import (
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int          `bun:",pk,autoincrement" json:"id"`
	Username      string       `bun:",unique,nullzero,notnull" json:"username"`
	Password      string       `bun:",nullzero,notnull" json:"password"`
	FirstName     string       `bun:",nullzero" json:"firstName"`
	LastName      string       `bun:",nullzero" json:"lastName"`
	MiddleName    string       `bun:",nullzero" json:"middleName"`
	DepartmentID  int          `json:"departmentID"`
	Department    *Department  `bun:",rel:belongs-to,join:department_id=id"`
	PositionID    int          `json:"positionID"`
	Position      *Position    `bun:"rel:belongs-to,join:position_id=id"`
	Attendances   []Attendance `bun:"rel:has-many,join:id=user_id"`
}
type Department struct {
	bun.BaseModel `bun:"table:department"`
	ID            int    `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:",nullzero,notnull" json:"name"`
}
type Position struct {
	bun.BaseModel `bun:"table:position"`
	ID            int    `bun:",pk,autoincrement" json:"id"`
	Name          string `bun:",nullzero,notnull" json:"name"`
}
type Attendance struct {
	bun.BaseModel `bun:"table:attendance"`
	ID            int       `bun:",pk,autoincrement" json:"id"`
	Type          int       `bun:",notnull" json:"type"`
	Time          time.Time `bun:",nullzero,notnull" json:"time"`
	UserID        int       `bun:",nullzero,notnull" json:"userID"`
}
