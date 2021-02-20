package main

import "fmt"

// 3    1    2
//  V   Y   Z
//   ^  ^  ^
//    \ | /
//     \|/
//      ------> X - 0

var (
	// размерность
	dimension = 3

	// маска для проверки принадлежности текущей размерности
	dimensionMask = 0

	// размер кубика в боксах
	size = 2
)

// инициализация маски размерности
func init() {
	for ii := 0; ii < dimension; ii++ {
		dimensionMask |= int(directions[ii])
	}
}

// _Direction Вектор или скорее направление
type _Direction int

// возможные направления
const (
	directionX _Direction = 1 << iota
	directionNX
	directionY
	directionNY
	directionZ
	directionNZ
	directionV
	directionNV
)

// вспомогательные маски для проверок принадлежности векторов направлений конкретные координаты
const (
	maskDirectionX = 3 << (2 * iota)
	maskDirectionY
	maskDirectionZ
	maskDirectionV
)

var directions = []_Direction{maskDirectionX, maskDirectionY, maskDirectionZ, maskDirectionV}

const directionLetters = "XYZV"

// String предварительный красивый вывод вектора для консоли
func (d _Direction) String() string {
	if d == 0 {
		return "none"
	}
	result := ""
	for ii := 0; ii < dimension; ii++ {
		if d&maskDirectionX != 0 {
			result += directionLetters[ii : ii+1]
			if d&directionX != 0 {
				result += "+"
			} else {
				result += "-"
			}
		}

		d = d >> 2
	}

	return result
}

// Reverse меняет направление вектора на противоположное
func (d _Direction) Reverse() _Direction {
	for ii := 0; ii < dimension; ii++ {
		if d&directions[ii] != 0 {
			d ^= directions[ii]
		}
	}

	return d
}

// Проверка корректного состаяния вектора направления
func (d _Direction) Check() error {

	// вектор должен быть в текущей размерности
	d &= _Direction(dimensionMask)
	if d == 0 {
		return ErrWrongDimention
	}

	// вектор должен указывать только по одной из координат
	cnt := 0
	for ii := 0; ii < dimension*2; ii++ {
		if d&directionX != 0 {
			cnt++
		}
		d >>= 1
	}

	if cnt != 1 {
		return fmt.Errorf("direction dimention must be 1, %w", ErrWrongDimention)
	}

	return nil
}

// _Plane пока для представления плоскости вращения
// возможно потом появятся и другие назначения
type _Plane [2]_Direction

func (p _Plane) Check() error {

	// сонаправленные векторы не дают однозначной плоскости
	// оба вектора находятся в одной координате и указывают в одном направлении
	//  аля X+ X+
	if p[0]&p[1] != 0 {
		return ErrWrongDirection
	}

	// параллельные разнонаправленные векторы не дают однозначной плоскости
	// оба вектора находятся в одной координате хоть и указывают в ранзых направлениях
	//  аля X+ X-
	xor := p[0] ^ p[1]
	for ii := 0; ii < dimension; ii++ {
		if xor == directions[ii] {
			return ErrWrongDirection
		}
	}

	if err1, err2 := p[0].Check(), p[1].Check(); err1 != nil || err2 != nil {
		return fmt.Errorf("%v: %w", p, ErrWrongDirection)
	}

	return nil
}
