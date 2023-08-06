defmodule Octopus.Apps.Supermario.Mario do
  @moduledoc """
  handles the mario logic, still under heavy development
  """
  alias __MODULE__
  alias Octopus.Apps.Supermario.{Game, Level, Matrix}

  @start_position_x 3
  @jump_interval_ms 110_000
  @fall_interval_ms 100_000
  @mario_color [216, 40, 0]

  @type t :: %__MODULE__{
          x_position: integer(),
          y_position: integer(),
          jumps: integer(),
          jumped_at: Time.t() | nil,
          falling_since: Time.t() | nil
        }
  defstruct [
    :x_position,
    :y_position,
    :jumps,
    :jumped_at,
    :falling_since
  ]

  def new(start_position_y) do
    %Mario{
      x_position: @start_position_x,
      y_position: start_position_y,
      jumps: 0,
      jumped_at: nil,
      falling_since: nil
    }
  end

  def draw(pixels, %Mario{} = mario) do
    pixels
    |> Matrix.from_list()
    |> set_mario(mario, @mario_color)
    |> Matrix.to_list()
  end

  def move_left(%Mario{x_position: 0} = mario, _level), do: mario

  def move_left(%Mario{} = mario, _level) do
    %Mario{mario | x_position: mario.x_position - 1}
  end

  def move_right(%Mario{} = mario, _level) do
    %Mario{mario | x_position: mario.x_position + 1}
  end

  def jump(%Mario{jumps: 0} = mario, game) do
    if can_jump?(mario, game) and !can_fall?(mario, game) do
      %Mario{mario | y_position: mario.y_position - 1, jumps: 1, jumped_at: Time.utc_now()}
    else
      mario
    end
  end

  # we can jump a second time in the air
  def jump(%Mario{jumps: 1} = mario, game) do
    if can_jump?(mario, game) do
        %Mario{mario | y_position: mario.y_position - 1, jumps: 2, jumped_at: Time.utc_now()}
      else
        mario
      end
  end

  # no third jump should be possible
  def jump(%Mario{jumps: 2} = mario, _game), do: mario

  # jump a second pixel after a while but only when we are already jumping
  def jump_second_if(
        %Mario{jumps: jumps, y_position: y_position, jumped_at: jumped_at} = mario,
        %Game{} = game
      ) when jumps > 0 do
    # TODO: use another constant, jump_interval_ms is used to prevent a second jump within a short time
    if Time.diff(Time.utc_now(), jumped_at, :microsecond) > @jump_interval_ms do
      new_y_position =
        if can_jump?(mario, game) do
          y_position - 1
        else
          y_position
        end

      %Mario{
        mario
        | y_position: new_y_position,
          jumps: 0,
          jumped_at: nil,
          falling_since: Time.utc_now()
      }
    else
      mario
    end
  end

  # falling_since may be nil, when mario was not jumping before but ran over a hole
  # but then we are falling immediately
  def fall_if(%Mario{falling_since: nil} = mario, game) do
    if can_fall?(mario, game) do
      {true, fall(mario)}
    else
      {false, mario}
    end
  end

  # already falling, check if we can fall further and fall with a delay
  def fall_if(
        %Mario{falling_since: falling_since} = mario,
        %Game{} = game
      ) do
    if Time.diff(Time.utc_now(), falling_since, :microsecond) > @fall_interval_ms do
      if can_fall?(mario, game) do
        # reset falling timestamp, also next fall should be done with a delay
        {true, fall(mario)}
      else
        {false, %Mario{mario | falling_since: nil}}
      end
    else
      {false, mario}
    end
  end

  defp fall(%Mario{y_position: y_position} = mario) do
    %Mario{mario | y_position: y_position + 1, falling_since: Time.utc_now()}
  end

  # very simple implementation wether mario can fall or not. need to provide a pixel matrix
  defp can_fall?(%Mario{y_position: y_position, x_position: x_position}, %Game{
         level: level,
         current_position: current_position
       }) do
    Level.can_fall?(level, x_position + current_position, y_position)
  end

  defp can_jump?(%Mario{y_position: 0}, _), do: false

  defp can_jump?(%Mario{y_position: y_position, x_position: x_position}, %Game{
         level: level,
         current_position: current_position
       }) do
    Level.can_jump?(level, x_position + current_position, y_position)
  end

  def can_move_right?(%Mario{y_position: y_position, x_position: x_position}, %Game{
        level: level,
        current_position: current_position
      }) do
    Level.can_move_right?(level, x_position + current_position, y_position)
  end

  def can_move_left?(%Mario{y_position: y_position, x_position: x_position}, %Game{
        level: level,
        current_position: current_position
      }) do
    Level.can_move_left?(level, x_position + current_position, y_position)
  end

  def start_position_x, do: @start_position_x

  defp set_mario(matrix, mario, mario_color) do
    put_in(matrix[mario.y_position][mario.x_position], mario_color)
  end
end
