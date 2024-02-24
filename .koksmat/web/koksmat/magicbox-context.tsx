"use client"
import { createContext } from "react";
export interface Session {
  user: User;
  expires: string;
  roles: string[];
  accessToken: string;
}

export interface User {
  name: string;
  email: string;
  image: string;
}


export type  MagicboxContextType= {
  session?:Session,
  version:number,
  signIn : (provider:string,config?:any,params?:any) => void,
  signOut : () => void,
  refresh:()=>void

}
export const MagicboxContext = createContext<MagicboxContextType>({
  session: { user: { name: "", email: "", image: "" }, expires: "", roles: [], accessToken: "" }, version: 0, refresh: () => { },
  signIn: function (provider: string, config: any, params: any): void {
    throw new Error("Function not implemented.");
  },
  signOut: function (): void {
    throw new Error("Function not implemented.");
  }
});


