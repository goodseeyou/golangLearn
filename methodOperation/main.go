package main

import "fmt"

type T struct {
    I,J int
}

func (t T) p() {
    fmt.Println(t.I, t.J)
}

func (t *T) pp() {
    fmt.Println(t.I, t.J)
}

func (t *T) ppt() {
    fmt.Printf("%[1]T %[1]v\n", t)
}

func (t T) pt() {
    fmt.Printf("%[1]T %[1]v\n", t)
}


func (T) pT() {
    fmt.Printf("TT\n")
}

/*
// will make redeclare compile error
func (t *T) pt() {
    fmt.Printf("%[1]T %[1]v\n", t)
}*/

func (t *T) Shift4() {
    (*t).I <<= 1
    (*t).J <<= 1
    (*t).I  = (*t).I << 1
    (*t).J = (*t).J << 1
    t.I <<= 1
    t.J <<= 1
    t.I  = t.I << 1
    t.J = t.J << 1
}

func main() {
    T{0,0}.pT()

    var t = T{1,2}
    t.p()
    t.Shift4()
    t.pp()
    t.pt()
    t.ppt()

    var tp = &T{1,2}
    tp.pp()
    tp.Shift4()
    tp.p()
    tp.pt()
    tp.ppt()

    tp = &T{1,2}
    (*tp).p()
    (*tp).Shift4()
    (*tp).p()

    var f = (*T).Shift4
    f(&t)
    t.p()

    // ok
    (&T{3,4}).pp()
    (&T{3,4}).p()

    // error
    //(T{3,4}).pp()
    // ok
    (T{3,4}).p()

    pf := T.p
    pf(T{123,321})
}
