import * as React from "react";
import { FC, memo, RefObject } from "react";
import { CanvasWrap } from "./styles";

const Canvas: FC<{
  setDrawRef: RefObject<HTMLCanvasElement>;
  setBgRef: RefObject<HTMLCanvasElement>;
}> = ({ setDrawRef, setBgRef }) => (
  <CanvasWrap>
    <canvas ref={setBgRef} />
    <canvas ref={setDrawRef} />
  </CanvasWrap>
);

export default memo(Canvas);
