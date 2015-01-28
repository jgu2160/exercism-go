package leap

func IsLeapYear(year int) bool{
    four := year % 4 == 0
    hundred := year % 100 == 0
    four_hundred := year % 400 == 0

    if four && !hundred || four_hundred {
      return true
    } else {
      return false
    }
}
