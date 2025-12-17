module GuessingGame

let reply (guess: int): string = 
    match guess with
    | 42 -> "Correct"
    | guess when guess = 41 -> "So close"
    | guess when guess = 43 -> "So close"
    | guess when guess <= 40 -> "Too low"
    | guess when guess >= 44 -> "Too high"
