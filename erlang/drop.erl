%% @author Simon St.Laurent <simonstl@simonstl.com> [http://simonstl.com]
%% @doc Functions calculating velocities achieved by objects
%% dropped in a vacuum.
%% @reference from <a href= "http://shop.oreilly.com/product/0636920025818.do" >
%% Introducing Erlang</a>,
%% O'Reilly Media, Inc., 2017.
%% @copyright 2017 by Simon St.Laurent
%% @version 0.1

-module(drop).
-export([fall_velocity/2]).

fall_velocity(earth, Distance) -> math:sqrt(2 * 9.8 * Distance);
fall_velocity(moon, Distance) -> math:sqrt(2 * 1.6 * Distance);
fall_velocity(mars, Distance) -> math:sqrt(2 * 3.71 * Distance).

