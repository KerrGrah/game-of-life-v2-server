import { GameSetup, LiveCells } from "./types";
import { theme } from "./theme";

export const makeCanvasHelpers = ({ width, height, cellSize }: GameSetup) => {
  const calcWidth = width * cellSize;
  const calcHeight = height * cellSize;

  return {
    paintBackground: (ctx: CanvasRenderingContext2D) => {
      ctx.fillStyle = theme.bgColor;

      ctx.fillRect(0, 0, calcWidth, calcHeight);

      for (let x = 0; x <= calcWidth; x += cellSize) {
        ctx.moveTo(x, 0);
        ctx.lineTo(x, calcHeight);
      }
      for (let y = 0; y <= calcHeight; y += cellSize) {
        ctx.moveTo(0, y);
        ctx.lineTo(calcWidth, y);
      }

      ctx.lineWidth = 0.3;
      ctx.strokeStyle = theme.bgBorder;
      ctx.stroke();
    },
    setGameCanvasStyle: (ctx: CanvasRenderingContext2D) => {
      ctx.fillStyle = theme.liveCellColor;
    },
    setDimensions: (canvas: HTMLCanvasElement) => {
      canvas.width = calcWidth;
      canvas.height = calcHeight;
    },
    getContext: (canvas: HTMLCanvasElement) => canvas.getContext("2d"),

    makePainter: (gameContext: CanvasRenderingContext2D) => (
      gameState: LiveCells
    ) => {
      gameContext.clearRect(0, 0, calcWidth, calcHeight);

      Object.entries(gameState).forEach(([xKey, yRow]: [string, number[]]) => {
        yRow.forEach(yKey => {
          gameContext.fillRect(
            +xKey * cellSize,
            yKey * cellSize,
            cellSize,
            cellSize
          );
        });
      });
    }
  };
};
