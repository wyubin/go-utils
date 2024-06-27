package serialize

import "time"

const (
	df = "2006-01-02T15:04:05"
)

type DateTime struct {
	time.Time
}

func (d *DateTime) UnmarshalText(text []byte) error {
	t, err := time.Parse(df, string(text))
	if err != nil {
		return err
	}
	*d = DateTime{t}
	return nil
}

func (d *DateTime) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}

func (d *DateTime) String() string {
	return d.Format(df)
}
