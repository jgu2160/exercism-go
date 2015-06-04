package queenattack
import (
    "errors"
)

func CanQueenAttack(white, black string) (attack bool, err error) {

    wf := int(white[0])
    wr := int(white[1])
    bf := int(black[0])
    br := int(black[1])

    if white == black || br >= 57 || wr >= 57 {
        return false, errors.New("Not valid")
    }

    if white[0] == black[0] {
        return true, nil
    } else if white[1] == black[1] {
        return true, nil
    } else if Abs(wr - br) == Abs(wf - bf) {
        return true, nil
    }
    return false, nil
}

func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
