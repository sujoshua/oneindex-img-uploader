package KeyPair

type KeyPair struct {
	Key   string
	Value string
}

type KeyPairs []KeyPair

func NewKeyPairs(n int) KeyPairs {
	return make(KeyPairs, n)
}

func (k KeyPairs) Len() int {
	return len(k)
}

func (k KeyPairs) Less(i, j int) bool {
	return k[i].Key < k[j].Key
}

func (k KeyPairs) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func (k KeyPairs) GetByKey(key string) string {
	for _, v := range k {
		if v.Key == key {
			return v.Value
		}
	}
	return ""
}

// GetByIndex get value by index, if index out of range, return empty string
func (k KeyPairs) GetByIndex(index int) string {
	if index >= k.Len() {
		return ""
	}
	return k[index].Value
}

func (K KeyPairs) Set(index int, value string) {
	if index >= K.Len() {
		return
	}
	K[index].Value = value
}

func (k KeyPairs) Append(keyPair KeyPair) KeyPairs {
	return append(k, keyPair)
}
