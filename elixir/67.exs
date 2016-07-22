defmodule Times do
  def double(n), do: n * 2
end

defmodule ShortTimes, do: (def double(n), do: n * 2)


Times.double(10)
|> IO.puts


ShortTimes.double(10)
|> IO.puts
