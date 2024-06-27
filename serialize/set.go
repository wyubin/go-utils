package serialize

import "strings"

type Set map[string]struct{}

func (s *Set) UnmarshalText(text []byte) error {
	*s = Set{}
	for _, v := range strings.Split(string(text), ",") {
		nameItem := strings.TrimSpace(v)
		if nameItem == "" {
			continue
		}
		(*s)[nameItem] = struct{}{}
	}
	return nil
}

func (s *Set) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *Set) String() string {
	keys := make([]string, 0, len(*s))
	for k := range *s {
		keys = append(keys, k)
	}
	return strings.Join(keys, ",")
}

func (s *Set) Add(keys ...string) {
	for _, key := range keys {
		(*s)[key] = struct{}{}
	}
}

func (s *Set) Remove(keys ...string) {
	for _, key := range keys {
		delete(*s, key)
	}
}

func (s *Set) Contains(key string) bool {
	_, ok := (*s)[key]
	return ok
}
