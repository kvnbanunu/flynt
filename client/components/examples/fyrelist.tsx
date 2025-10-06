"use client";
import { Get } from "@/lib/api";
import React, { useEffect, useState } from "react";
import { FyreCard } from "./fyrecard";
import { CS_ENV } from "@/lib/utils";
import AddFyre from "./addfyre";
import { RemoveFyre } from "./removefyre";

export const FyreList: React.FC<{ user: Models.User }> = (props) => {
  const { user } = props;
  const [fyres, setFyres] = useState<Models.Fyre[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [update, setUpdate] = useState<boolean>(false);

  const fetchFyres = async () => {
    const res = await Get<Models.Fyre[]>(
      `${CS_ENV.api_url}/api/fyre/user/${user.id}`,
    );
    if (res.success) {
      setFyres(res.data);
    } else {
      setError(res.error.message);
    }
    setLoading(false);
    setUpdate(false);
  };

  const onUpdate = () => {
    setUpdate(true);
  };

  useEffect(() => {
    fetchFyres();
  }, [update]);

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
        <AddFyre onSuccessHandler={onUpdate} user_id={user.id} />
      </div>
    </div>
  );
};
