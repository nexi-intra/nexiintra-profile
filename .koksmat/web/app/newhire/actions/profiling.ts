
"use server"

import { cookies } from "next/headers"
import { setMemberships } from "./memberships"


export const saveProfile = async (
  upn: string,
  country: string,
  unit: string,
  membershipsToAdd: string[],
  membershipsToRemove: string[]
) => {
  
  const setMembershipsResuls = await setMemberships(
    upn,
    membershipsToAdd,
    membershipsToRemove
  )


  cookies().set("user", JSON.stringify({ country, unit }), { path: "/" })

  return { href: "https://christianiabpos.sharepoint.com/sites/nexiintra-home?country=" + country +"&unit="+unit, target: "_top", setMembershipsResuls }
}


