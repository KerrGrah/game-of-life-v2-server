import * as React from "react";
import { FC, useEffect, useRef } from "react";
import { GameSetup, LiveCells } from "./types";
import Canvas from "./Canvas";
import { makeCanvasHelpers } from "./canvasHelpers";

type Props = { gameSetup: GameSetup; gameState: LiveCells };

const Board: FC<Props> = props => {
  const { gameState, gameSetup } = props;

  const backgroundCanvas = useRef<HTMLCanvasElement>(null);
  const gameCanvas = useRef<HTMLCanvasElement>(null);
  const paint = useRef<(liveCells: LiveCells) => void>();

  useEffect(
    () => {
      const bgCanvas = backgroundCanvas.current;
      const gCanvas = gameCanvas.current;

      const {
        paintBackground,
        setDimensions,
        getContext,
        setGameCanvasStyle,
        makePainter
      } = makeCanvasHelpers(gameSetup);

      if (bgCanvas && gCanvas) {
        setDimensions(bgCanvas);
        setDimensions(gCanvas);

        const bgContext = getContext(bgCanvas);
        const gameContext = getContext(gCanvas);

        if (bgContext && gameContext) {
          paintBackground(bgContext);
          setGameCanvasStyle(gameContext);

          paint.current = makePainter(gameContext);
        }
      }
    },
    [gameSetup]
  );

  useEffect(
    () => {
      const frameId = window.requestAnimationFrame(() => {
        if (paint.current) paint.current(gameState);
      });
      return () => window.cancelAnimationFrame(frameId);
    },
    [gameState]
  );

  return <Canvas setDrawRef={gameCanvas} setBgRef={backgroundCanvas} />;
};

export default Board;
