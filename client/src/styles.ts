import styled, { css } from "styled-components";
import { ifProp } from "styled-tools";

export const AppWrap = styled.div`
  position: relative;
`;

export const CanvasWrap = styled.div`
  display: flex;
  justify-content: center;
  & > canvas {
    position: absolute;
  }
`;

export const GameOnButton = styled.button<{ gameOn: boolean }>`
  ${ifProp(
    "gameOn",
    css`
      color: red;
    `,
    ""
  )};
`;

export const ControlsContainer = styled.div`
  position: absolute;
  left: 10vw;
  width: 300px;
  z-index: 2;
  opacity: 0.1;
  transition: opacity 0.4s;

  &:hover {
    opacity: 0.7;
  }
`;

export const SliderWrap = styled.div`
  position: relative;
  border: none;
  border-radius: 3px;
  margin: 3px;
  background-color: #222;
  padding: 3px 6px 2px 6px;
  box-shadow: 1px 2px 30px 2px #555;
  transition: transform 0.5s ease-in-out;
  &:hover {
    transform: scale(1.1);
  }

  & > label {
    position: relative;
    bottom: 3.5px;
    color: #c1fcc1;
    font-size: 12px;
    padding-right: 20px;
    font-weight: 200;
  }
`;
