package utils

var (
	scale float32
	fiber int
	stone int
	wood  int
)

func Set_Scale(s float32) {
	*&scale = s
}

func Get_Scale() float32 {
	return scale
}

func Add_Fiber(v int) {
	*&fiber += v
}

func Add_Stone(v int) {
	*&stone += v
}

func Add_Wood(v int) {
	*&wood += v
}

func Get_Fiber() int {
	return fiber
}

func Get_Stone() int {
	return stone
}

func Get_Wood() int {
	return wood
}
