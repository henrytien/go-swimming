package demo
import "testing"
func TestGetArea(t *testing.T) {
    area := GetArea(40, 50)
    if area != 2000 {
        t.Error("测试失败")
    }
}

func BenchmarkGetArea(t *testing.B) {
    for i := 0; i < t.N; i++ {
        t.Log(i)
        GetArea(40, 50)
    }
}
