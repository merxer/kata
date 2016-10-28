defmodule Combined do
  def height_to_mps(meters) do
    Convert.mps_to_mph(Drop.fall_velocity(meters))
  end
end
