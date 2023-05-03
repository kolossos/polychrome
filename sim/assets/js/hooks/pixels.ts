function resize(canvas: HTMLCanvasElement) {
  const dpr = window.devicePixelRatio || 1;
  const rect = canvas.getBoundingClientRect();
  canvas.width = rect.width * dpr;
  canvas.height = rect.height * dpr;
}

interface Layout {
  imageSize: [number, number];
  pixelSize: [number, number];
  positions: [number, number][];
}

export function setup(canvas: HTMLCanvasElement, pixelImageUrl: string) {
  let pixels = new Uint8Array();
  let layout: Layout;

  const textEncoder = new TextEncoder();

  const pixelImage = new Image();
  pixelImage.src = pixelImageUrl;

  const ctx = canvas.getContext("2d");
  if (!ctx) {
    return;
  }

  resize(canvas);

  window.addEventListener("resize", () => resize(canvas));

  this.handleEvent("layout", ({ layout: newLayout }: { layout: Layout }) => {
    layout = newLayout;
  });

  this.handleEvent("pixels", ({ pixels: newPixels }: { pixels: string }) => {
    pixels = textEncoder.encode(newPixels);
  });

  const draw = () => {
    if (!layout) {
      window.requestAnimationFrame(draw);
      return;
    }

    ctx.clearRect(0, 0, canvas.width, canvas.height);

    ctx.save();

    const scale = Math.min(
      canvas.width / layout.imageSize[0],
      canvas.height / layout.imageSize[1]
    );

    const offsetX = canvas.width / 2 - (layout.imageSize[0] / 2) * scale;
    const offsetY = canvas.height / 2 - (layout.imageSize[1] / 2) * scale;

    ctx.translate(offsetX, offsetY);
    ctx.scale(scale, scale);

    layout.positions.forEach(([x, y], i) => {
      const pixel = pixels[i];

      if (pixel !== undefined) {
        ctx.fillStyle = `hsl(40, 80%, ${(95 / 8) * pixel + 5}%)`;
      } else {
        ctx.fillStyle = "hsl(40, 80%, 5%)";
      }

      ctx.fillRect(x, y, layout.pixelSize[0], layout.pixelSize[1]);

      ctx.drawImage(pixelImage, x, y, layout.pixelSize[0], layout.pixelSize[1]);
    });

    ctx.restore();

    window.requestAnimationFrame(draw);
  };

  window.requestAnimationFrame(draw);
}
