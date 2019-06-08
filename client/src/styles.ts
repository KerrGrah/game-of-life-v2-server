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
  position: relative;
  cursor: pointer;
  outline: none;
  font-size: 16px;
  background-color: #eee;
  color: #222;
  margin: 4px;
  border: none;
  font-weight: 200;
  transition: all 0.3s ease-in-out;

  &:hover {
    transform: scale(1.05);
    background-color: #111;
  }

  ${ifProp(
    "gameOn",
    css`
      color: red;
    `,
    ""
  )};
`;

export const SliderWrap = styled.div`
  position: relative;
  opacity: 0.1;
  border: none;
  border-radius: 3px;
  margin: 3px;
  background-color: #222;
  padding: 3px 6px 2px 6px;
  box-shadow: 1px 2px 30px 2px #555;
  transition: transform 0.5s ease-in-out, opacity 0.2s;
  &&&:hover {
    transform: scale(1.02);
    opacity: 0.8;
  }
`;

export const SliderLabel = styled.label`
  position: relative;
  bottom: 3.5px;
  color: #eee;
  font-size: 12px;
  padding-right: 20px;
  font-weight: 200;
`;

export const SliderInput = styled.input`
  -webkit-appearance: none;
  background-color: black;
  height: 3px;
  position: relative;
  bottom: 5px;

  &::-webkit-slider-thumb {
    -webkit-appearance: none;
    position: relative;
    bottom: 1px;
    width: 7px;
    height: 14px;
    cursor: pointer;
    -webkit-box-shadow: 0 6px 5px 0 rgba(0, 0, 0, 0.6);
    -moz-box-shadow: 0 6px 5px 0 rgba(0, 0, 0, 0.6);
    box-shadow: 0 6px 5px 0 rgba(0, 0, 0, 0.6);
    border-radius: 20%;
    background-image: -webkit-gradient(
      linear,
      left top,
      left bottom,
      color-stop(0%, #eee),
      color-stop(60%, #aaa),
      color-stop(61%, #000),
      color-stop(100%, #fff)
    );
  }
  &:focus {
    outline: none;
  }
  &:disabled {
    &::-webkit-slider-thumb {
      cursor: not-allowed;
    }
  }
`;

export const ControlsContainer = styled.div`
  position: absolute;
  left: 10vw;
  width: 300px;
  z-index: 2;

  &:hover {
    & ${SliderWrap} {
      opacity: 0.7;
    }
  }
`;
