package utils

var scale float32

func Set_Scale(s float32) {
	*&scale = s
}

func Get_Scale() float32 {
  return scale
}
