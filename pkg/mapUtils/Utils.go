package mapUtils

type Type interface {
	int | int8 | int16 | int32 | int64 | string | bool
}

func GetMapKey[T Type](data map[T]T) []T {
	result := make([]T, len(data))

	i := 0
	for k, _ := range data {
		result[i] = k
		i++
	}

	return result
}
func GetMapValue[T Type](data map[T]T) []T {
	result := make([]T, len(data))

	i := 0
	for _, v := range data {
		result[i] = v
		i++
	}

	return result
}
