defmodule Recursive do
  def sum(0), do: 0
  def sum(n), do: n + sum(n-1)
end


Recursive.sum(1)
|> IO.puts

Recursive.sum(3)
|> IO.puts
