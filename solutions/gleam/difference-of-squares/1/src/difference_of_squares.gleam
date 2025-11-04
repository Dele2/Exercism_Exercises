import gleam/io
import gleam/int
import gleam/list

pub fn square_of_sum(n: Int) -> Int {
  list.range(1, n)
  |> list.fold(0, fn(a, b) { a + b }) 
  |> square
}

pub fn sum_of_squares(n: Int) -> Int {
  list.range(1, n)
  |>  list.map(fn(x) { x * x })
  |> list.fold(0, fn(a, b) { a + b })
}

pub fn difference(n: Int) -> Int {
    square_of_sum(n) - sum_of_squares(n)
}

pub fn square(n: Int) -> Int{
  n * n
}
