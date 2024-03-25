//"use client";

import { redirect } from "next/navigation";
//import { useEffect } from "react";
import { ProfileType } from "./global";

export default function Home() {
  //useEffect(() => {
  const type: ProfileType = process.env.PROFILETYPE as ProfileType;
  switch (type) {
    case "newhire":
      redirect("/newhire");
      break;

    default:
      redirect("/profile");
      break;
  }
  //}, []);

  return <div></div>;
}
