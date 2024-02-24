

import { getSpAuthToken } from "@/koksmat/officegraph"
import {
  getCountries,
  getNewsCategories,
  getNewsChannels,
  getUnits,
  getValidGuestDomains,
} from "@/app/profile/data/sharepoint"

import { Country, NewsCategory, NewsChannel, Unit } from "./schemas"

const KEY = "profilecache"
export interface ProfileCache {
  date: Date
  key: string
  countries: Country[]
  categories: NewsCategory[]
  units: Unit[]
  channels: NewsChannel[]
}

export async function readProfileData() {
  const token = await getSpAuthToken()
  await getValidGuestDomains(token)

  const countries = (await getCountries(token)) ?? []
  const categories = (await getNewsCategories(token)) ?? []
  const units = (await getUnits(token)) ?? []
  const channels = (await getNewsChannels(token)) ?? []
  const cache: ProfileCache = {
    date: new Date(),
    key: KEY,
    countries,
    categories,
    units,
    channels,
  }
  return cache
}
export async function refreshProfileCache() {
    console.log("Refreshing profile cache")
  //const client = await connect()
  const cache = await readProfileData()


  return cache
}

export async function getProfileCache() {

  
  return {} as any
}
