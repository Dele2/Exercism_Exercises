module Hamming

let charList s = s |> Seq.toList

let rec compareLists lst1 lst2 = 
    match lst1 , lst2 with
    | [] , [] -> 0
    | x :: xs , y :: ys -> if x <> y then 
                            1 + compareLists xs ys
                            else
                            compareLists xs ys
    | _ , _ -> 0

let distance (strand1: string) (strand2: string): int option = 
    match strand1 , strand2 with
    | "" , "" -> Some (0)
    | strand1 , strand2 when strand1.Length <> strand2.Length -> None
    | strand1, strand2 -> Some (compareLists (charList strand1) (charList strand2))