"use client";
import { Post } from "@/lib/api";
import React, { ChangeEvent, useState } from "react";

export const FyreCard: React.FC<{ fyre: Models.Fyre }> = async (props) => {
  const [currentFyre, setCurrentFyre] = useState<Models.Fyre>(props.fyre);
  const [isChecked, setIsChecked] = useState<boolean>(false);
  const [streakCount, setStreakCount] = useState<number>(
    props.fyre.streakCount,
  );

  const handleCheckboxChange = async (event: ChangeEvent<HTMLInputElement>) => {
    setIsChecked(event.target.checked);

    let copy: Models.Fyre = currentFyre;

    copy.streakCount = isChecked ? streakCount + 1 : streakCount - 1;

    const res = await Post<Models.Fyre, Models.Fyre>(
      `/api/fyre/${currentFyre.id}`,
      copy,
    );
    if (res.success) {
      setCurrentFyre(res.data.data);
      setStreakCount(res.data.data.streakCount);
    }
  };

  return (
    <div className="mx-8 my-2 p-4">
      <label>
        <input
          type="checkbox"
          checked={isChecked}
          onChange={handleCheckboxChange}
        />
        {currentFyre.title}
      </label>
      <div>Streak Count: {streakCount}</div>
    </div>
  );
};
