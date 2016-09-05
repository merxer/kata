defmodule Test do
  def double(n), do: n * 2
end


Test.double(10) |> IO.puts
