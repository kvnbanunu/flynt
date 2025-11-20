import React, { ReactNode } from "react";

interface MainContainerProps {
  children?: ReactNode;
}

export const MainContainer: React.FC<MainContainerProps> = ({ children }) => {
  return (
      <div className="flex px-4 pt-4 pb-32 md:pb-4 w-screen md:w-full min-h-full place-content-center">
        {children}
      </div>
  );
};
