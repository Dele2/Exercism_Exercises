module Hamming

let distance (strand1: string) (strand2: string): int  option = 
    match strand1 , strand2 with
    | "" , "" -> Some (0)
    | strand1 , strand2 when strand1.Length <> strand2.Length -> None
    | _ , _ ->
        let lst1 = strand1 |> Seq.toList
        let lst2 = strand2 |> Seq.toList
        let notEq x y= if x <> y then 1 else 0
        let result =List.map2 notEq lst1 lst2
        Some (List.sum result)