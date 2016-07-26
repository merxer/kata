ops = fn
      (0,0,_) -> "FizzBuzz"
      (0,_,_) -> "Fizz"
      (_,0,_) -> "Buzz"
      (_,_,n) -> n
end

fizzbuzz = fn(n) -> ops.(rem(n, 3), rem(n, 5), n)  end


for n <- 1..100, do: fizzbuzz.(n) |> IO.puts
