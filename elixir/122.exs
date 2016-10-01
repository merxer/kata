IO.inspect set1 = Enum.into 1..5, MapSet.new

IO.inspect MapSet.member? set1, 3

IO.inspect set2 = Enum.into 3..8, MapSet.new

IO.inspect MapSet.union set1, set2

IO.inspect MapSet.difference set1, set2

IO.inspect MapSet.difference set2, set1

IO.inspect MapSet.intersection set1, set2
