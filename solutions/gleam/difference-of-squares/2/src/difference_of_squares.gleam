import gleam/io
import gleam/int

fn sum_of_squares_tail(num: Int, acc: Int) -> Int {
  case num {
    num if num == 1 -> acc
    _ -> {
      let new_acc = acc + num * num
      sum_of_squares_tail(num - 1, new_acc)
    }
  }
}

pub fn sum_of_squares(num: Int) -> Int {
  sum_of_squares_tail(num, 1)
}

fn square_of_sum_tail(num: Int, acc: Int) -> Int {
  case num {
    num if num == 1 -> acc
    _ -> {
      let new_acc = acc + num
      square_of_sum_tail(num - 1, new_acc)
    }
  }
}

pub fn square_of_sum(num: Int) -> Int {
  let sum = square_of_sum_tail(num, 1)
  sum * sum
}

pub fn difference(n: Int) -> Int {
  square_of_sum(n) - sum_of_squares(n)
}
