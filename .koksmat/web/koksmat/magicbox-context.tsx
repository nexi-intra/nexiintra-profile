"use client"
import { IPublicClientApplication } from "@azure/msal-browser";

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
  user?:User,
  setAccount:(username: string,email:string,image:string)=>void,
  registerAuth : (pca : IPublicClientApplication) => void,
  signIn : (scopes: string[], loginHint: string) => Promise<boolean>,
  signOut : () => void,
  refresh:()=>void

}
export const MagicboxContext = createContext<MagicboxContextType>({
  session: { user: { name: "", email: "", image: "" }, expires: "", roles: [], accessToken: "" }, version: 0, refresh: () => { },
  signIn: function (scopes: string[], loginHint?: string): Promise<boolean> {
    throw new Error("Function not implemented.");
  },
  signOut: function (): void {
    throw new Error("Function not implemented.");
  },
  setAccount: function (username: string, email: string, image: string): void {
    throw new Error("Function not implemented.");
  },
  registerAuth: function (pca: IPublicClientApplication): void {
    throw new Error("Function not implemented.");
  }
});


