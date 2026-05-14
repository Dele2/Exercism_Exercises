
let string_to_list s = 
  s 
  |> String.to_seq
  |> List.of_seq 
;;

let char_list_to_int_list chars =
  List.map (fun c -> int_of_char c - int_of_char '0') chars
;;

let rec double_every_second_element lst =
  match lst with
  | [] -> []
  | [x] -> [x]
  | (x :: y :: xs) -> x :: y * 2 :: double_every_second_element xs
;;

let rec split_double_digit_ints lst =
  match lst with
  | [] -> []
  | (x :: xs) when x > 9 -> (x / 10) :: (x mod 10) :: split_double_digit_ints xs
  | (x :: xs) -> x :: split_double_digit_ints xs
;;

let sum_list lst = 
  List.fold_left (+) 0 lst
;;

let remove_all_spaces s =
  s 
  |> String.split_on_char ' ' 
  |> String.concat ""
;;

let calculate_luhn s_num =
  s_num
  |> string_to_list
  |> char_list_to_int_list
  |> List.rev
  |> double_every_second_element
  |> split_double_digit_ints
  |> sum_list
;;

let is_all_digits s =
  String.for_all (function '0'..'9' -> true | _ -> false) s

let valid num =
  let s_num = remove_all_spaces num in
  if String.length s_num > 1 && is_all_digits s_num then calculate_luhn s_num mod 10 = 0 else false
;;

