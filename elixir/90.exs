defmodule Times do
  def double(n) do
    n * 2
  end
end

Times.double(20) |> IO.puts
