bin = <<3 :: size(2), 5 :: size(4), 1 :: size(2)>>

bin
|> IO.inspect


:io.format("~-8.2b~n", :binary.bin_to_list(bin))
|> IO.inspect

bin
|> byte_size
|> IO.puts
