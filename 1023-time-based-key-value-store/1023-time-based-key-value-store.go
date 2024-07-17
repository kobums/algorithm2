type data struct {
	time  int
	value string
}

type TimeMap map[string][]data

func Constructor() TimeMap {
	return make(map[string][]data, 1024)
}

func (m TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := m[key]; !ok {
		m[key] = make([]data, 1, 1024)
	}
	m[key] = append(m[key], data{
		time:  timestamp,
		value: value,
	})
}

func (m TimeMap) Get(key string, timestamp int) string {
	d := m[key]
    if len(d) == 0 {
        return ""
    }
    i := sort.Search(len(d), func(i int) bool {
        return timestamp < d[i].time
    })
    i--
	return m[key][i].value

}