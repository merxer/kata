options = [{:width, 72}, {:style, "light"}, {:style, "print"}]
IO.inspect List.last options

IO.inspect Keyword.get_values options, :style
