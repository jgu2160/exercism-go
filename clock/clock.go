package clock
import "fmt"

type Clock int

const (
    TestVersion = 1
    minutesPerDay = 1440
    minutesPerHour = 60
)

func New(h,m int) Clock {
    c := Clock((h * minutesPerHour + m) % minutesPerDay)
    if c < 0 {
        c += minutesPerDay
    }
    return c
}

func (c Clock) Add(minutes int) Clock {
    c = (c + Clock(minutes)) % minutesPerDay
    if c < 0 {
        c += minutesPerDay
    }
    return c
}

func (c Clock) String() string {
    hours := c/minutesPerHour
    minutes := c - minutesPerHour * hours
    return fmt.Sprintf("%02d:%02d", c/minutesPerHour, minutes)
}
