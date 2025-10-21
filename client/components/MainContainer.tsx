import React, { ReactNode } from "react";

interface MainContainerProps {
  children?: ReactNode;
}

export const MainContainer: React.FC<MainContainerProps> = ({ children }) => {
  return (
      <div className="flex justify-center items-center px-4 py-4 w-screen md:w-full h-full">
        {children}
      </div>
  );
};
