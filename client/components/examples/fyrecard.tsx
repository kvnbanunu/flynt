"use client";
import { Put } from "@/lib/api";
import React, { useState } from "react";
import { CS_ENV } from "@/lib/utils";

export const FyreCard: React.FC<{ fyre: Models.Fyre }> = (props) => {
  const [currentFyre, setCurrentFyre] = useState<Models.Fyre>(props.fyre);
  const [isChecked, setIsChecked] = useState<boolean>(false);
  const [streakCount, setStreakCount] = useState<number>(
    props.fyre.streak_count,
  );

  const handleCheckboxChange = async () => {
    // We need to add another field to fyre to see if it has already been checked off today
    const checked = isChecked ? false : true
    setIsChecked(checked);

    const copy: Models.Fyre = currentFyre;

    copy.streak_count = checked ? streakCount + 1 : streakCount - 1;
    copy.bonfyre_id = -1;

    const res = await Put<Models.Fyre>(
      `${CS_ENV.api_url}/fyre/${currentFyre.id}`,
      copy,
    );
    if (res.success) {
      setCurrentFyre(res.data);
      setStreakCount(res.data.streak_count);
    }
  };

  return (
    <div className="my-2 p-2">
      <label className="text-lg">
        <input
          className="mr-1 cursor-pointer"
          type="checkbox"
          checked={isChecked}
          onChange={handleCheckboxChange}
        />
        {currentFyre.title}
      </label>
      <div>Streak: {streakCount} 🔥</div>
    </div>
  );
};
