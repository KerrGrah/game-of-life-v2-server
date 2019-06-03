import * as React from "react";
import { ChangeEvent, Dispatch, FC, SetStateAction } from "react";
import { GameSetup } from "./types";
import { ControlsContainer } from "./styles";
import Slider from "./Slider";

type Props = {
  gameSetup: GameSetup;
  setGameSetup: Dispatch<SetStateAction<GameSetup>>;
};

const Controls: FC<Props> = props => {
  const {
    setGameSetup,
    gameSetup,
    gameSetup: { width, height, speed, density }
  } = props;

  const handleSlider = ({ target }: ChangeEvent<HTMLInputElement>) => {
    setGameSetup({ ...gameSetup, [target.name]: +target.value });
  };
    console.log('density', density)
    
  return (
    <ControlsContainer>
      <Slider
        value={width}
        name="width"
        range={{ min: 6, max: 200 }}
        handleChange={handleSlider}
      />
      <Slider
        value={height}
        name="height"
        range={{ min: 6, max: 120 }}
        handleChange={handleSlider}
      />
      <Slider
        value={density}
        name="density"
        range={{ min: 0, max: 1 }}
        step={0.01}
        handleChange={handleSlider}
      />
      <Slider
        value={speed}
        name="speed"
        range={{ min: 30, max: 1025 }}
        handleChange={handleSlider}
      />
    </ControlsContainer>
  );
};

export default Controls;
