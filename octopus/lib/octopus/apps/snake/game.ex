defmodule Octopus.Apps.Snake.Game do

  defmodule Worm do
    @base_speed 40
    defstruct [:parts,:rem_t, :speed]

    def new() do
      %Worm { parts: [{{3,7},:r}, {{2,7}, :r}, {{1,7}, :r}], rem_t: @base_speed , speed: @base_speed}
    end

    defp move(parts, dir), do: move([], parts, dir)

    defp move(acc,[],_), do: Enum.reverse(acc)
    defp move(acc,[{{x,y}, pdir} | tail], dir) do
      newpos = case pdir do
        :u -> {x,y-1}
        :d -> {x,y+1}
        :l -> {x-1,y}
        :r -> {x+1,y}
      end

      move([{newpos, dir} | acc], tail, pdir)
    end

    def tick(worm, [dir | _]), do: tick(worm, dir)
    def tick(%Worm{ parts: [{_,dir} | _]} = worm, []), do: tick(worm, dir)
    def tick(%Worm{ rem_t: 0} = worm, dir) do
      %Worm{
        worm |
        parts: move(worm.parts, dir),
        rem_t: worm.speed
      } |> IO.inspect()

    end
    def tick(%Worm{parts: [{pos,_} | parttail], rem_t: rem_t} = worm, dir), do: %Worm{ worm | parts: [{pos, dir} | parttail], rem_t: rem_t-1}

    def dead?(%Worm{parts: parts}) do
      parts |> Enum.reduce({MapSet.new(), false}, fn {{x,y} = p, _}, {acc, c} ->
        c = c or x<0 or y<0 or x>=8 or y>=8
        {MapSet.put(acc, p), MapSet.member?(acc, p) or c}
      end ) |> elem(1)
    end

    def positions(%Worm{parts: parts}) do
      parts |> Enum.reduce([], fn {p, _}, acc -> [p | acc] end) |> Enum.reverse()
    end
  end

  defstruct [:worm, :food]
  alias Octopus.Canvas
  alias Octopus.Apps.Snake
  alias Snake.JoyState
  alias Snake.Game

  def new() do
    worm = Worm.new()
    %Game{
      worm: Worm.new(), food: new_food(worm)
    }
  end

  def new_food(%Worm{} = worm) do
    food = {:rand.uniform(8)-1,:rand.uniform(8)-1}
    cond do
      food in Worm.positions(worm) -> new_food(worm)
      true -> food
    end
  end

  def tick(%Game{ food: food} = game, joy) do

    new_worm = game.worm |> Worm.tick(JoyState.direction(joy))

    new_game = case hd(new_worm.parts) do
    {^food, _} ->
      wormy = %Worm{ new_worm | parts: [hd(new_worm.parts) | game.worm.parts] }
      %Game{
        game |
        worm: %Worm{ wormy | speed: (wormy.speed - 1) |> Snake.Util.clamp(10,60)},
        food: new_food(wormy)
      }
    _ -> %Game{
      game |
      worm: new_worm
    }
  end

    cond do
      Worm.dead?(new_game.worm) -> Game.new()
      true -> new_game
    end
  end

  def render_frame(%Game{} = game) do
    gamecanvas =
      Canvas.new(8,8)
      |> Canvas.put_pixel(game.food, {0xff,0xff,0x00})

    gamecanvas =
    game.worm.parts
    |> Enum.reduce(gamecanvas, fn {pos, _dir}, acc -> acc |> Canvas.put_pixel(pos, {0x10,0xff,0x10}) end)

    Canvas.new(60,8) |> Canvas.overlay(gamecanvas, offset: {8*4,0}) |> Canvas.to_frame()
  end
end