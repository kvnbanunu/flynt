import { Get } from "@/lib/api";
import React from "react";

export const ExampleComponent: React.FC<{
    title: string;
}> = async (props) => {
    const { title } = props;
    const users: Models.User[] = await Get("api/users");

    return (
        <>
            <h1>{title}</h1>
            <ul>
                {users?.map((result, index) => {
                    return (
                        <li key={index}>
                            {result.id}: {result.name}
                        </li>
                    );
                })}
            </ul>
        </>
    );
};
