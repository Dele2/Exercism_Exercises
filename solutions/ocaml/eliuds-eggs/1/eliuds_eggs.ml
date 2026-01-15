let rec decimal_to_binary (num: int) : int list =
  match num with
  | 0 -> []
  | num -> num mod 2 :: decimal_to_binary (num / 2)
;;

let rec count_eggs (lst: int list) : int =
  match lst with
  | [] -> 0
  | (h::t) -> if h == 1 then 1 + count_eggs t
              else count_eggs t
;;

let egg_count (number:int) : int =
  decimal_to_binary number
  |> count_eggs
;;