"use client";
import { Get } from "@/lib/api";
import React, { useEffect, useState } from "react";
import { FyreCard } from "./fyrecard";

export const FyreList: React.FC<{ user: Models.User }> = async (props) => {
  const { user } = props;
  const [fyres, setFyres] = useState<Models.Fyre[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchFyres = async () => {
      const res = await Get<Models.Fyre[]>(`/api/fyre/${user.id}`);
      if (res.success) {
        const fyres: Models.Fyre[] = res.data.data;
        setFyres(fyres);
      } else {
        setError(res.error.message);
      }
      setLoading(false);
    };
    fetchFyres();
  }, []);

  if (loading) {
    return <div>Loading your Fyres</div>;
  }

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div>
      <h1>Hello {user.name}</h1>
      <div>Your Fires</div>
      <ul>
        {fyres.map((fyre) => (
          <li>
            <FyreCard fyre={fyre} />
          </li>
        ))}
      </ul>
    </div>
  );
};
