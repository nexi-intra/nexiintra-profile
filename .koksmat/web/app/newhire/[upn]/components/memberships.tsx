"use client";

import { off } from "process";
import { useContext, useEffect, useState } from "react";
import { Membership, NewsChannel } from "../../data/schemas";
import { Badge } from "@/components/ui/badge";
import { Skeleton } from "@/components/ui/skeleton";
import { TransitiveMemberOf, getMemberOfs } from "../../data/officegraph";
import { MagicboxContext } from "@/koksmat/magicbox-context";



export function MyMemberships() {
  const magicbox = useContext(MagicboxContext);

  const [memberships, setmemberships] = useState<Membership[]>()
  useEffect(() => {
    const load = async () => {

      setmemberships((await getMemberOfs(magicbox.session?.accessToken ?? "")))
    }
   if (magicbox.session?.accessToken) load()
  }, [magicbox.session?.accessToken])


  return <div >
    {!memberships &&
      <div className="items-center">
        {Array.from({ length: 4 }).map((_, i) => {
          return <div key={i}><Skeleton className="mt-3 h-12 w-12 rounded-full" />
            <div className="space-y-2">
              <Skeleton className="h-4 w-[250px]" />
              <Skeleton className="h-4 w-[200px]" />
            </div></div>
        })}
      </div>}
      {memberships && 
      <div>

  {memberships.sort((a,b)=>{
    if (a.groupDisplayName>b.groupDisplayName) return 1
    if (a.groupDisplayName<b.groupDisplayName) return -1
    return 0
  }).map((membership, key) => {
        return <div key={key} >
          {membership.groupDisplayName} [{membership.groupId}]  
        </div>
      })}    
      </div>
      }
  </div>


    ;
}
