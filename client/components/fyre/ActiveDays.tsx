"use client";
import React from "react";
import { Button } from "../ui/button";
import { ButtonGroup } from "../ui/button-group";

interface ActiveDaysProps {
  active: string;
  isOpen: boolean;
  onChange: (index: number) => void;
}

export const ActiveDays: React.FC<ActiveDaysProps> = ({
  active,
  isOpen,
  onChange,
}) => {
  const days = ["S", "M", "T", "W", "T", "F", "S"];

  if (active === "1111111" && !isOpen) {
    return (
      <Button className="w-39 sm:w-46 md:w-60" disabled>
        Everyday
      </Button>
    );
  }

  return (
    <ButtonGroup>
      {days.map((day, index) => (
        <Button
          key={index}
          variant={active[index] === "0" ? "secondary" : "default"}
          className="px-1.5 sm:px-2 md:px-3 border-1"
          disabled={!isOpen}
          onClick={() => onChange(index)}
        >
          {day}
        </Button>
      ))}
    </ButtonGroup>
  );
};
