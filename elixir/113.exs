map = %{ name: "Dave", likes: "Programming", where: "Dallas"}

IO.inspect Map.keys map

IO.inspect Map.values map

IO.inspect map[:name]

IO.inspect map.name

map1 = Map.drop map,[:where, :likes]
IO.inspect map1

map2 = Map.put map, :also_likes, "Ruby"
IO.inspect map2

IO.inspect Map.keys map2

{ value, update_map } = Map.pop map2, :also_likes
