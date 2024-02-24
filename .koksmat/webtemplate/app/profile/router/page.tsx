
"use client"

import { MagicboxContext } from "@/koksmat/magicbox-context";
import { redirect } from "next/navigation";
import { useContext, useEffect } from "react";
export const dynamic = 'force-dynamic'

export default async function RedirectToLoggedinUse() {

  const magicboxContext = useContext(MagicboxContext)
  useEffect(() => {
    if (!magicboxContext) {
      return
    }

    if (magicboxContext?.session?.user?.email) {
      redirect("/profile/" + magicboxContext?.session?.user?.email)
    } else { redirect("/profile") }


  }, [magicboxContext])




  return <div>



  </div>
}
