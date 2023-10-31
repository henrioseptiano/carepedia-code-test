package app

func (q *Queue) SetRoundRobin(enable bool) {
	q.RoundRobin = enable
	q.LastGenderOut = ""
}
