export type GameSetup = {
  width: number;
  height: number;
  cellSize: number;
  density: number;
  speed: number;
};

export type LiveCells = { [key: number]: number[] };

export type CellEntries = [string, number[]][];

export type Dictionary<T> = { [key: string]: T };
