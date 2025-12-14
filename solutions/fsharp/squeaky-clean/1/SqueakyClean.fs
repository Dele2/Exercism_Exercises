module SqueakyClean

open System

let transform (c: char) : string =
    match c with
    | c when Char.IsNumber(c) -> ""
    | c when Char.IsUpper(c) ->  "-" + c.ToString().ToLower()
    | c when c >= 'α' && c <= 'ω' -> "?"
    | '-' -> "_"
    | ' ' -> ""
    | _ -> string c
    
    
let clean (identifier: string): string =
    identifier 
    |> Seq.map transform
    |> String.Concat

