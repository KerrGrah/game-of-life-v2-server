import * as React from "react";
import { SliderWrap } from "./styles";
import { ChangeEvent, FC } from "react";

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
      <label>{name}</label>
      <input
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
