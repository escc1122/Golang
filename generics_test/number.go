package main

func sum[num int | int64](n1 num, n2 num) num {
	return n1 + n2
}

type Number interface {
	int | int64
}

func sum2[num Number](n1 num, n2 num) num {
	return n1 + n2
}

type Number2 interface {
	~int
}

type intA int

func sum3[num Number2](n1 num, n2 num) num {
	return n1 + n2
}
