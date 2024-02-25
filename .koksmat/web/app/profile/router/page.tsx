"use client"

import { MagicboxContext } from "@/koksmat/magicbox-context";
import { useContext, useEffect } from "react";

export default  function RedirectToLoggedinUse() {

  const magicbox = useContext(MagicboxContext)
  useEffect(() => {
    if (!magicbox) {
      return
    }

    location.href = "/profile/" + magicbox.user?.email
    

  }, [magicbox])




  return <div>



  </div>
}
