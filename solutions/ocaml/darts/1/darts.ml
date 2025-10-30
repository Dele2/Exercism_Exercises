let distance (x : float) (y : float) : float = 
  (x *. x) +. (y *. y)
|> sqrt
;;

let score (x: float) (y: float): int = 
  match distance x y with
  | distance when distance <= 1.0 -> 10
  | distance when distance <= 5.0 -> 5
  | distance when distance <= 10.0 -> 1
  | _ -> 0
;;