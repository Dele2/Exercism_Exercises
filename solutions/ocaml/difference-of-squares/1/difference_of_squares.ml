let rec sum_of_num (num: int) : int =
  match num with
  | 1 -> 1
  | _ -> num + sum_of_num (num - 1) 
;;

let square_of_sum (num: int) : int =
  (sum_of_num num) * (sum_of_num num)

let rec sum_of_squares (num: int) : int =
  match num with
  | 1 -> 1 
  | _ -> num * num + sum_of_squares (num - 1)
;;

let difference_of_squares (num: int) : int =
  square_of_sum num - sum_of_squares num
;;

