import React, { ReactNode } from "react";

interface MainContainerProps {
  children?: ReactNode;
}

export const MainContainer: React.FC<MainContainerProps> = ({ children }) => {
  return (
      <div className="flex justify-center items-center px-4 py-4 min-h-screen w-full">
        {children}
      </div>
  );
};
