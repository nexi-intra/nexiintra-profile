"use client"

import { use } from "react"
import { ProfileForm } from "./components/firsttime-profile-form"


import { getProfileCache } from "../data/cache"
export const dynamic = 'force-dynamic'
export default async function SettingsProfilePage() {

  const session = null //await getUserSession()
  const upn ="" //session?.user?.email ?? ""
  const existingProfile : any = {} //await getprofile(upn)
  const data = await getProfileCache()
  const country = existingProfile?.country ?? ""
  const unit = existingProfile?.unit ?? ""
  return (
    <div className="space-y-6">
      <ProfileForm  currentUnit={unit} currentCountry={country} newsCategories={data?.categories??[]} newsChannels={data?.channels ??[]} countries={data?.countries ?? []} units={data?.units ?? []}  />
  
    </div>
  )
  
}
