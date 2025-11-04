"use client";
import { Get } from "@/lib/api";
import React, { useCallback, useEffect, useState } from "react";
import { FyreCard } from "./fyrecard";
import AddFyre from "./addfyre";
import { RemoveFyre } from "./removefyre";

export const FyreList: React.FC<{ user: Models.User; onMissedFyres?: (missed: Models.Fyre[]) => void; }> = ({ user, onMissedFyres}) => {
  // const { user } = props;
  const [fyres, setFyres] = useState<Models.Fyre[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [update, setUpdate] = useState<boolean>(false);

  const fetchFyres = useCallback(async () => {
    const res = await Get<Models.Fyre[]>(
      "/fyre",
    );
    if (res.success) {
      setFyres(res.data);
      // Find any missed Fyres
      console.log(res.data) // Check what backend returned
      const missed = res.data.filter(f => f.missed_check === true);
      console.log("Missed Fyres: ", missed);
      if (onMissedFyres) onMissedFyres(missed);
    } else {
      setError(res.error.message);
    }
    setLoading(false);
    setUpdate(false);
  }, [onMissedFyres]);

  const onUpdate = () => {
    setUpdate(true);
  };

  useEffect(() => {
    fetchFyres();
  }, [update, fetchFyres]);

  if (loading) {
    return <div>Loading your Fyres</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div className="p-2 m-2 border-2 rounded-sm">
      <h1 className="text-xl">Hello {user.name}</h1>
      <div className="mt-1">Your Fyres:</div>
      <ul>
        {fyres &&
          fyres.map((fyre) => (
            <li key={fyre.id}>
              <div className="grid grid-cols-2 gap-2 items-center w-full border-2 rounded-sm">
                <FyreCard fyre={fyre} />
                <RemoveFyre onSuccessHandler={onUpdate} fyre_id={fyre.id} />
              </div>
            </li>
          ))}
      </ul>
      <div>
        <AddFyre onSuccessHandler={onUpdate} />
      </div>
    </div>
  );
};
