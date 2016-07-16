a = [1, 2, 3]
IO.inspect a

a = 4
IO.inspect a

4 = a
IO.inspect 4

#[a, b] = [1, 2, 3] #raise MatchError

a = [ [1, 2, 3] ]
IO.inspect a

[ a..5 ] = [ 1..5 ]
IO.inspect a

[a] = [ [1, 2, 3] ]
IO.inspect a

#[ [a] ] = [ [ 1, 2, 3 ] ] #raise MathError
