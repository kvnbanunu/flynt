"use client";

import React from "react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { timezones } from "./Timezones";
import { useIsMobile } from "@/hooks/use-mobile";

interface TimezoneSelectorProps {
  value: string | undefined;
  onValueChange: (value: string) => void;
  placeholder?: string;
}

export const TimezoneSelector: React.FC<TimezoneSelectorProps> = ({value, onValueChange, placeholder}) => {
  const isMobile = useIsMobile();
  return (
    <Select value={value} onValueChange={onValueChange}>
      <SelectTrigger>
        <SelectValue placeholder={placeholder ?? "America/Vancouver"} />
      </SelectTrigger>
      <SelectContent>
        {timezones.map((tz, index) => (
          <SelectItem value={tz.value} key={index}>
            {isMobile ? tz.shortLabel : tz.fullLabel}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
};
