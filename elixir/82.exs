times_2 = fn n -> n * 2 end
apply = fn (fun, value) -> fun.(value) end

apply.(times_2, 6) |> IO.puts
