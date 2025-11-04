import React, { ReactNode } from "react";

interface MainContainerProps {
  children?: ReactNode;
}

export const MainContainer: React.FC<MainContainerProps> = ({ children }) => {
  return (
      <div className="flex p-4 w-screen md:w-full h-full place-content-center">
        {children}
      </div>
  );
};
