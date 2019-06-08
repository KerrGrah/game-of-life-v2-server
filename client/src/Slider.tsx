import * as React from "react";
import { ChangeEvent, FC } from "react";
import { capitalize } from "lodash";
import { SliderInput, SliderLabel, SliderWrap } from "./styles";

type Props = {
  name: string;
  value: number;
  range: { min: number; max: number };
  handleChange: (e: ChangeEvent<HTMLInputElement>) => void;
  [key: string]: any;
};
export const Slider: FC<Props> = ({
  name,
  value,
  handleChange,
  range,
  ...rest
}) => {
  return (
    <SliderWrap>
      <SliderLabel>{capitalize(name)}</SliderLabel>
      <SliderInput
        type="range"
        min={range.min}
        max={range.max}
        value={value}
        name={name}
        onChange={handleChange}
        {...rest}
      />
    </SliderWrap>
  );
};

export default Slider;
