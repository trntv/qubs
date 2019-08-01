package broker

type Consumer struct {
	tags []string
	link Link
	closed bool
	sent uint64
}

func (s *Consumer) hasTag(tag string) bool {
	for _, t := range s.tags {
		if t == tag {
			return true
		}
	}

	return false
}
