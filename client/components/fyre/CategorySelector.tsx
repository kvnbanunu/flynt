"use client";
import React from "react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";

interface CategorySelectorProps {
  allCategories: Models.Category[];
  selectedId: number;
  onChange: (id: number) => void;
}

export const CategorySelector: React.FC<CategorySelectorProps> = ({
  allCategories,
  selectedId,
  onChange,
}) => {
  return (
    <Select value={`${selectedId}`} onValueChange={(e) => onChange(Number(e))}>
      <SelectTrigger>
        <SelectValue placeholder="Select a category" />
      </SelectTrigger>
      <SelectContent>
        {allCategories.map((cat) => (
          <SelectItem key={cat.id} value={`${cat.id}`}>
            {cat.name}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
};
