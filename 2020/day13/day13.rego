package day13

part1 := answer {
    [wait, id] := min({x |
        id := constrained_buses[_]
        wait := (ceil(earliest/id)*id)-earliest
        x := [wait, id]
    })
    answer := wait*id
}

ceil(x) = x { round(x) == x } else = round(x+0.5)

constrained_buses := [ to_number(x) | x := split(split(input, "\n")[1], ",")[_]; x != "x" ]

earliest := to_number( split(input, "\n")[0] )

input := big_input

small_input := `939
7,13,x,x,59,x,31,19`

big_input := `1000507
29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,631,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,13,19,x,x,x,23,x,x,x,x,x,x,x,383,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,17`