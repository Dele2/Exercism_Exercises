

let digit_list num  =
  let rec d_list (acc: int list) (num: int) : int list =
    if num < 10 then num :: acc
    else d_list ((num mod 10) :: acc) (num / 10) in
      d_list [] num
;;

let rec my_pow (x: int) (y: int) : int =
  if y = 0 then 1
  else x * my_pow x (y - 1)
;;

let rec my_sum (lst: int list) : int =
  match lst with
  | [] -> 0
  | hd :: tl -> hd + my_sum tl
;;

let power_digits (lst: int list) : int list =
  let rec p_digits (lst: int list) (exp: int) : int list = 
  match lst with
    | [] -> []
    | hd :: tl -> (my_pow hd exp) :: p_digits tl exp in
      p_digits lst (List.length lst)
;;

let calculate_armstrong (num: int) : int =
  num
  |> digit_list 
  |> power_digits
  |> my_sum
;;

let validate (num: int) : bool =
  calculate_armstrong num = num
;;
