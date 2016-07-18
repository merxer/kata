plusOne = Enum.map [1,2,3,4], &(&1 + 1)
plusOne
|> IO.inspect


double = Enum.map [1,2,3,4], &(&1 * &1)
double
|> IO.inspect

less_than_tree = Enum.map [1,2,3,4], &(&1 < 3)
less_than_tree
|> IO.inspect

