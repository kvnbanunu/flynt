"use client"
import { Get, Post } from "@/lib/api";
import React, { useEffect, useState } from "react";
import { FyreCard } from "./fyrecard";
import { CS_ENV } from "@/lib/utils";
import AddFyre from "./addfyre";

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
      setFyres(res.data.data)
    } else {
      setError(res.error.message);
    }
    setLoading(false);
    setUpdate(false);
  };

  const onAddFyre = () => {
    setUpdate(true)
  }

  useEffect(() => { fetchFyres() }, [update]);

  if (loading) {
    return <div>Loading your Fyres</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div className="p-2 border-2">
      <h1 className="text-xl">Hello {user.name}</h1>
      <div className="mt-1">Your Fyres:</div>
      <ul>
        {fyres.map((fyre) => (
          <li key={fyre.id}>
            <FyreCard fyre={fyre} />
          </li>
        ))}
      </ul>
      <div>
        <AddFyre onSuccessHandler={onAddFyre} user_id={user.id} />
      </div>
    </div>
  );
};
