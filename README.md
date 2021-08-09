# qadrube

не расчитано на работу сразу с несколькими кубами разной размерности или размеров

```go   
    c4 := InitCube(4, 3)    // инициализировать куб
    c4.Print()              // печатает представление куба в консоль
    c4.Move([]int{0, 0, 0, 0}, _Plane{directionX, directionY}) // вращение произвольной грани
    // вращение производится путем указания точки отсчета и заданием плоскости вращения
    // точка отсчета - набор координат по количеству соответствующий размерности пространства и максимальным значением в координате по размеру куба
    // плоскость вращения задается 2мя векторами, направление внутри плоскости определяется порядком векторов
    // например _Plane{directionY, directionX} и _Plane{directionX, directionY} находятся в одной плоскости и являются противонаправленными
    // а _Plane{directionX, directionY} и _Plane(directionNX, directionNY)  в одной плоскости и сонаправлены
```