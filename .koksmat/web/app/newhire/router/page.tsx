"use client"
/*---
title: Router
description: Pickup the email of the user and redirect to the profile page of that
 ---

 */
import { MagicboxContext } from "@/koksmat/magicbox-context";
import { useRouter } from "next/navigation";
import { useContext, useEffect } from "react";

export default function RedirectToLoggedinUse() {

  const magicbox = useContext(MagicboxContext)
  const router = useRouter()
  useEffect(() => {
    if (!magicbox) {
      return
    }

    router.push( "/profile/" + magicbox.user?.email)


  }, [magicbox])




  return <div>



  </div>
}
