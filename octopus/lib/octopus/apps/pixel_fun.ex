defmodule Octopus.Apps.PixelFun do
  use Octopus.App

  alias Octopus.Canvas
  alias Octopus.Apps.PixelFun.Program

  @width 8 * 10 + 9 * 18
  @height 8

  defmodule State do
    defstruct [:canvas, :program, :source, :easing_interval]
  end

  def name(), do: "Pixel Fun"

  def config_schema() do
    %{
      program: {"Program", :string, %{default: "sin((t-x*0.01)-hypot(x%8-3.5,y-3.5))"}},
      easing_interval: {"Easing Interval", :int, %{default: 50, min: 0, max: 500}}
    }
  end

  def get_config(state) do
    %{program: state.source, easing_interval: state.easing_interval}
  end

  def init(_args) do
    :timer.send_interval((1000 / 60) |> trunc(), :tick)

    canvas = Canvas.new(@width, @height)

    config = config_schema() |> default_config()
    {:ok, program} = config.program |> Program.parse()

    {:ok,
     %State{
       canvas: canvas,
       program: program,
       source: config.program,
       easing_interval: config.easing_interval
     }}
  end

  def handle_config(%{program: program, easing_interval: easing_interval}, %State{} = state) do
    source = program

    program =
      case Program.parse(program) do
        {:ok, program} -> program
        _ -> 0
      end

    {:noreply, %State{state | program: program, source: source, easing_interval: easing_interval}}
  end

  def update_program(pid, program) do
    program =
      case Program.parse(program) do
        {:ok, program} -> program
        _ -> 0
      end

    GenServer.cast(pid, {:update_program, program})
  end

  def handle_cast({:update_program, program}, %State{} = state) do
    {:noreply, %{state | program: program}}
  end

  def handle_info(:tick, %State{} = state) do
    canvas = state |> render()

    canvas
    |> Canvas.to_frame(drop: true)
    |> Map.put(:easing_interval, state.easing_interval)
    |> send_frame()

    {:noreply, %{state | canvas: canvas}}
  end

  defp render(%State{canvas: canvas, program: program} = _state) do
    {seconds, micros} = Time.utc_now() |> Time.to_seconds_after_midnight()
    seconds = seconds + micros / 1_000_000

    for i <- 0..(@width * @height - 1), into: canvas do
      x = rem(i, @width)
      y = div(i, @width)
      {{x, y}, pixels(program, x, y, i, seconds)}
    end
  end

  @default_env %{'pi' => :math.pi(), 'tau' => :math.pi() * 2}

  defp pixels(expr, x, y, i, t) do
    env = [%{'x' => x, 'y' => y, 'i' => i, 't' => t}, @default_env]

    value =
      expr
      |> Program.eval(env)
      |> max(-1.0)
      |> min(1.0)

    interpolate_colors([0x3F, 0xFF, 0x21], [0xFB, 0x48, 0xC4], value)
  end

  defp interpolate_colors([r1, g1, b1], [r2, g2, b2], value) do
    cond do
      value > 0 -> [r1 * value, g1 * value, b1 * value]
      value < 0 -> [r2 * -value, g2 * -value, b2 * -value]
      true -> [0, 0, 0]
    end
    |> Enum.map(&Kernel.trunc/1)
  end
end
