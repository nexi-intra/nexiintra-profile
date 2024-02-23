"use client";

import { useEffect, useState } from "react";
import { MagicboxContextType, MagicboxContext, Session } from "./magicbox-context";




type Props = {
  children?: React.ReactNode;
};

export const MagicboxProvider = ({ children }: Props) => {

  const [session, setsession] = useState<Session>()
  const [version, setversion] = useState(0)


  
  const magicbox: MagicboxContextType = {
    session, version, refresh: () => {
      setversion(version + 1);
    },
    signIn: function (provider: string, config: any, params: any): void {
      throw new Error("Function not implemented.");
    },
    signOut: function (): void {
      throw new Error("Function not implemented.");
    }
  }
  return <MagicboxContext.Provider value={magicbox}>


    {children}


  </MagicboxContext.Provider>;
};
