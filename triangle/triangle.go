package triangle
import "sort"

type Kind interface {}

const(
    Equ = "equ"
    Iso = "iso"
    Sca = "sca"
    NaT = "nat"
)

func KindFromSides(side1, side2, side3 float64) Kind {
    a := []float64{side1, side2, side3}
    sort.Float64s(a)
    if a[2] >= a[0] + a[1] || a[0] <= 0{
        return NaT
    }

    matches := 0
    for _, test := range a {
        for _, against := range a {
            if test == against {
                matches++
            }
        }
    }

    switch matches {
        case 3: return Sca
        case 5: return Iso
        case 9: return Equ
        default: return NaT
    }
}
