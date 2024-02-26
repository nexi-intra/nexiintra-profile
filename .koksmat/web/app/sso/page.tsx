'use client'
 
import { https } from '@/koksmat/httphelper'
import { useSearchParams } from 'next/navigation'
import { useEffect, useState } from 'react'
 
export default function SearchBar() {
  const searchParams = useSearchParams()
 
  const token = searchParams.get('token')
 const [data, setdata] = useState<any>()

  useEffect(() => {
    const load = async () => {

      const result = await https(token??"", "GET", "https://graph.microsoft.com/v1.0/me")
      setdata(result)
    }
    if (token) {
      load()
    }
  
  
  }, [token])
  
  return <pre>{JSON.stringify(data,null,2)}</pre>
}
