add_one = &(&1 + 1)

add_one.(44)
|> IO.puts

square = &(&1 * &1)
square.(8)
|> IO.puts

speak = &(IO.puts(&1))
speak.("Hello")
