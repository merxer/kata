add_n = fn n -> (fn other -> n + other end) end

add_two = add_n.(2)

add_five = add_n.(5)

add_two.(3)
|> IO.inspect

add_five.(7)
|> IO.inspect
