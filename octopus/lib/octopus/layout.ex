defmodule Octopus.Layout do
  @keys [
    :name,
    :positions,
    :width,
    :height,
    :pixel_size,
    :pixel_margin,
    :image_size,
    :background_image,
    :pixel_image
  ]
  @enforce_keys @keys

  defstruct @keys

  @typedoc """
  Position of a pixel in the image
  """
  @type position :: {integer(), integer()}

  @typedoc """
  Size in pixels
  """
  @type size :: {integer(), integer()}

  @type t :: %__MODULE__{
          name: String.t(),
          positions: list(position()),
          width: integer(),
          height: integer(),
          pixel_size: size(),
          pixel_margin: {integer(), integer(), integer(), integer()},
          image_size: size(),
          background_image: String.t(),
          pixel_image: String.t()
        }

  @callback layout() :: t()

  defmacro __using__(_opts) do
    quote do
      @behaviour Octopus.Layout
    end
  end

  defimpl Jason.Encoder, for: Octopus.Layout do
    def encode(%Octopus.Layout{} = layout, opts) do
      Jason.Encode.map(
        %{
          name: layout.name,
          positions: layout.positions |> Enum.map(&Tuple.to_list/1),
          width: layout.width,
          height: layout.height,
          pixelSize: layout.pixel_size |> Tuple.to_list(),
          pixelMargin: layout.pixel_margin |> Tuple.to_list(),
          imageSize: layout.image_size |> Tuple.to_list(),
          backgroundImage: layout.background_image,
          pixelImage: layout.pixel_image
        },
        opts
      )
    end
  end
end
