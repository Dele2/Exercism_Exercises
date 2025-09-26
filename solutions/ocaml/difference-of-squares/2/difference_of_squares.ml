let rec square_of_sum_acc (num: int) (acc: int) : int =
  match num with
  | 1 -> acc
  | num -> square_of_sum_acc (num - 1) (acc + num)
;;

let square_of_sum (num: int) =
  square_of_sum_acc num 1 * square_of_sum_acc num 1
;;

let rec sum_of_squares_acc (num: int) (acc: int) : int =
  match num with
  | 1 -> acc
  | num -> sum_of_squares_acc (num - 1) (acc + num * num)
;;

let sum_of_squares (num: int) =
  sum_of_squares_acc num 1
;;

let difference_of_squares (num: int) : int =
  square_of_sum num - sum_of_squares num
;;

