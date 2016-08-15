fall_velocity = fn(distance) -> :math.sqrt(2 * 9.8 * distance) end

fall_velocity.(20) |> IO.puts

fall_velocity= &(:math.sqrt(2 * 9.8 * &1))

fall_velocity.(20) |> IO.puts
