package null

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type Null[T any] struct {
	sql.Null[T]
}

func New[T any](v T) *Null[T] {
	n := Null[T]{
		Null: sql.Null[T]{
			V:     v,
			Valid: true,
		},
	}

	return &n
}

func NewFromPtr[T any](v *T) *Null[T] {
	if v == nil {
		return &Null[T]{
			Null: sql.Null[T]{
				Valid: false,
			},
		}
	}

	return &Null[T]{
		Null: sql.Null[T]{
			V:     *v,
			Valid: true,
		},
	}
}

func (n Null[T]) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	v, ok := any(n.V).(driver.Valuer)
	if ok {
		return v.Value()
	}

	return n.V, nil
}

func (n *Null[T]) IsEmpty() bool {
	return !n.Valid
}

func (n *Null[T]) GetValue() *T {
	if !n.Valid {
		return nil
	}

	return &n.V
}

func (n *Null[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(n.V)
}

func (n *Null[T]) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &n.V)

	return err
}
