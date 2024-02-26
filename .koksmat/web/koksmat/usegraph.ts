"use client"

import { useEffect, useState } from "react"

import { run } from "./server"
import { Result, https } from "./httphelper"
import { useMsal } from "@azure/msal-react"
import { set } from "date-fns"

export const version = 1

export function useGraph(
  url: string,
  scopes: string[]


) {
  const [token, settoken] = useState("")
  const [data, setdata] = useState<Result<any>>()
  const [isLoading, setisLoading] = useState(false)
  const [error, seterror] = useState("")
  const [didRun, setdidRun] = useState(false)
  const { accounts, instance,inProgress } = useMsal()
  useEffect(() => {
    const load = async () => {

      const account = accounts[0]

      seterror("")
      try {
        
        const response = await instance.acquireTokenSilent({
          scopes,
          account
        });
        const token = response.accessToken
        const testResponse = await https(response.accessToken, "GET", url)
        setdata(testResponse)
        settoken(token)
        setdidRun(true)
        setisLoading(false)

        return { token, testResponse }

      } catch (error) {
        try {
          const response = await instance.acquireTokenPopup({
            scopes: scopes ?? [],
            account
          });
          const token = response.accessToken
          const testResponse = await https(response.accessToken, "GET", url)
          return { token, testResponse }

        } catch (error) {
          setdidRun(true)
          setisLoading(false)
          setdidRun(true)
          setisLoading(false)

          seterror((error as any).message ?? "Unknown error")
          return null
        }



      }




    
      

    }
    if (accounts && accounts.length > 0 && instance) {
      if (inProgress === "none") {
        setisLoading(true)
        load()
      }
      

    }
  }, [accounts, instance,inProgress])

  return {
    token,
    result:data,
    error,
    isLoading,
  }
}
