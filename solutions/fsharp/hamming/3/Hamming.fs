module Hamming

let charList s = s |> Seq.toList

let rec compareLists lst1 lst2 acc = 
    match lst1 , lst2 with
    | [] , [] -> acc
    | x :: xs , y :: ys -> if x <> y then 
                            compareLists xs ys (acc + 1)
                            else
                            compareLists xs ys acc

let distance (strand1: string) (strand2: string): int option = 
    match strand1 , strand2 with
    | "" , "" -> Some (0)
    | strand1 , strand2 when strand1.Length <> strand2.Length -> None
    | strand1, strand2 -> Some (compareLists (charList strand1) (charList strand2) 0)