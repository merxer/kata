fun1 = fn -> (fn -> "Hello" end) end

other = fun1.()

other.()
|> IO.inspect
