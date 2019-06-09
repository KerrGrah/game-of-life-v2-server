import * as React from "react";
import { FC, useCallback, useEffect, useRef, useState } from "react";
import { useBoolean, useNumber } from "react-hanger";
import { GameSetup, LiveCells } from "./types";
import CanvasWrap from "./CanvasWrap";
import { AppWrap, GameOnButton } from "./styles";
import Controls from "./Controls";
import { prepareGameSetup, socketSend } from "./utils";

const location = "localhost:8080/ws";

const url = `${document.location.protocol.replace("http", "ws")}//${location}`;

const INITIAL_GAME_SETUP: GameSetup = {
  width: 150,
  height: 70,
  density: 0.4,
  speed: 1025, // 25 - 1025
  cellSize: 10,
  initiate: true
};

const App: FC = () => {
  const gameOn = useBoolean(false);
  const [gameSetup, setGameSetup] = useState(INITIAL_GAME_SETUP);
  const generations = useNumber(0);
  const [gameState, setGameState] = useState({});

  const socket = useRef<WebSocket>();

  const readSocketData = useCallback(
    ({ data }: MessageEvent) => {
      const parsed = JSON.parse(data) as LiveCells;

      setGameState(parsed);

      generations.increase();
    },
    [generations]
  );

  const start = useCallback(
    () => {
      console.log("start");

      socket.current = new WebSocket(url);
      socket.current.onmessage = readSocketData;
      socket.current.onerror = console.error;
      socket.current.onopen = () => {
        socketSend(prepareGameSetup(gameSetup), socket.current);
        setGameSetup({ ...gameSetup, initiate: false });
      };
      generations.setValue(0);
    },
    [gameSetup, generations, readSocketData]
  );

  const stop = useCallback(() => {
    if (socket.current) {
      socket.current.close();
    }
  }, []);

  useEffect(
    () => {
      if (gameOn.value) start();
      else stop();
    },
    [gameOn.value] // eslint-disable-line react-hooks/exhaustive-deps
  );

  useEffect(
    () => {
      if (gameOn.value) {
        socketSend(prepareGameSetup(gameSetup), socket.current);
      }
    },
    [gameSetup]
  );

  const buttonText = gameOn.value ? "Stop" : "Start";

  return (
    <AppWrap>
      <p> {`generations: ${generations.value}`}</p>
      <CanvasWrap gameSetup={gameSetup} gameState={gameState} />
      <GameOnButton gameOn={gameOn.value} onClick={gameOn.toggle}>
        {buttonText}
      </GameOnButton>
      <Controls
        gameSetup={gameSetup}
        setGameSetup={setGameSetup}
        gameOn={gameOn.value}
      />
    </AppWrap>
  );
};

export default App;
