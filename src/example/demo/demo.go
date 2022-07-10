package demo
import ("log")
// 根据长宽获取面积
func GetArea(weight int, height int) int {
    log.Print(weight, height)
    return weight * height
}
