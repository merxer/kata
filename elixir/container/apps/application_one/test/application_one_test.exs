defmodule ApplicationOneTest do
  use ExUnit.Case
  doctest ApplicationOne

  test "the truth" do
    IO.puts "Running Application On tests"
    assert 1 + 1 == 2
  end
end
