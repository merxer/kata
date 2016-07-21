defmodule Outer do

  defmodule Inner do
    IO.puts "Inner module"

    def inner_func do
      IO.puts "Inner function"
    end
  end

  def outer_func do
    IO.puts "Outer function"
    Inner.inner_func
  end
end


Outer.outer_func
Outer.Inner.inner_func
