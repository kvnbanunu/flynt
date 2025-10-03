import { Get } from "@/lib/api";
import { SS_ENV } from "@/lib/utils";
import React from "react";

export const ExampleComponent: React.FC<{
  title: string;
}> = async (props) => {
  const { title } = props;
  const res = await Get<Models.User[]>(`${SS_ENV.api_url}/api/user`);
  if (res.success) {
    const users: Models.User[] = res.data;

    return (
      <>
        <h1>{title}</h1>
        <ul>
          {users &&
            users.map((result, index) => {
              return (
                <li key={index}>
                  {result.id}: {result.name}
                </li>
              );
            })}
        </ul>
      </>
    );
  } else {
    console.error(`API Error: ${res.error.status_code} - ${res.error.message}`);
    return <div>{res.error.status_code}</div>;
  }
};
