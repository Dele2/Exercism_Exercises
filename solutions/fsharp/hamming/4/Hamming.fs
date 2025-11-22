module Hamming

let rec compareLists str1 str2 acc = 
    match str1 , str2 with
    | "" , "" -> acc
    | _ , _ -> 
        let h = str1.[0]
        let t = str1.Substring(1)
        let hh = str2.[0]
        let tt = str2.Substring(1)
        if h <> hh then 
            compareLists t tt (acc + 1)
        else
            compareLists t tt acc
 

let distance (strand1: string) (strand2: string): int option = 
    match strand1 , strand2 with
    | "" , "" -> Some (0)
    | strand1 , strand2 when strand1.Length <> strand2.Length -> None
    | strand1, strand2 -> Some (compareLists (strand1) (strand2) 0)