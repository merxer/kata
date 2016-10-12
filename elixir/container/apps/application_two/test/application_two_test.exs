defmodule ApplicationTwoTest do
  use ExUnit.Case
  doctest ApplicationTwo

  test "the truth" do
    IO.puts "Running Application Two tests"
    assert 1 + 1 == 2
  end
end
