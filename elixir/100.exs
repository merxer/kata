defmodule Solution do
  def main() do
    a = IO.get("") |> String.strip |> String.to_integer
    b = IO.get("") |> String.strip |> String.to_list
    IO.puts b
  end
end

Solution.main()
