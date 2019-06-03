import { Dictionary, GameSetup } from "./types";

export const socketSend = (message: Dictionary<any>, socket?: WebSocket) => {
  if (socket) socket.send(JSON.stringify(message));
};

export const prepareGameSetup = (gameSetup: GameSetup) => {
  const speed = 1030 - gameSetup.speed;
  return { ...gameSetup, speed };
};
