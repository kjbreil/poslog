package poslog

// POSLogs is an array of poslog grouped by store(s) and dayid(s)
type POSLogs struct {
	DayIDs  []string
	Stores  []int
	POSLogs []POSLog
}

func (ps *POSLogs) appendDayID(d string) {
	for _, c := range ps.DayIDs {
		if c == d {
			continue
		} else {
			ps.DayIDs = append(ps.DayIDs, d)
		}
		return
	}
}

func (ps *POSLogs) appendStore(s int) {
	for _, c := range ps.Stores {
		if c == s {
			continue
		} else {
			ps.Stores = append(ps.Stores, s)
		}
		return
	}
}
