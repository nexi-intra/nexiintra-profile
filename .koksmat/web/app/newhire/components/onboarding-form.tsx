"use client"

import { useContext, useEffect, useState } from "react"
import Link from "next/link"
import { zodResolver } from "@hookform/resolvers/zod"
import { is, tr } from "date-fns/locale"

import { useForm } from "react-hook-form"
import * as z from "zod"

import { Result } from "@/koksmat/httphelper"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { toast } from "@/components/ui/use-toast"
import {
  CreateInvitationResult,
  createInvitation,
} from "@/app/profile/actions/onboarding"
import { set } from "date-fns"
import { NewGuest } from "./NewGuest"
import { MagicboxContext } from "@/koksmat/magicbox-context"
import { redirect, useRouter } from "next/navigation"

const profileFormSchema = z.object({
  email: z
    .string({
      required_error: "Please select an email to display.",
    })
    .email(),
})

type ProfileFormValues = z.infer<typeof profileFormSchema>

export function ValidateEmailAccountForm() {
  const magicbox = useContext(MagicboxContext)

  const defaultValues: Partial<ProfileFormValues> = {}

  const form = useForm<ProfileFormValues>({
    resolver: zodResolver(profileFormSchema),
    defaultValues,
    mode: "onChange",
  })

  const [invitationStatus, setInvitationStatus] =
    useState<Result<CreateInvitationResult>>()
  const [signinEnabled, setsigninEnabled] = useState(false)
  const [processing, setProcessing] = useState(false)
  const [isNewGuest, setisNewGuest] = useState(false)
  const [isUnknownDomain, setisUnknownDomain] = useState(false)
  const [guideHTML, setguideHTML] = useState("")
  const router =  useRouter()
  useEffect(() => {
    if (!magicbox || magicbox==undefined) return
    
    if (magicbox.user) {
      if (magicbox.user.email.indexOf("#ext#@") < 0) {
        form.setValue("email", magicbox.user.email ?? "")
      }
    }
  }, [magicbox])

  async function onSubmit(data: ProfileFormValues) {
    setProcessing(true)
    setInvitationStatus(undefined)
    setsigninEnabled(false)
    // toast({
    //   title: "You submitted the following values:",
    //   description: (
    //     <pre className="mt-2 w-[340px] rounded-md bg-slate-950 p-4">
    //       <code className="text-white">{JSON.stringify(data, null, 2)}</code>
    //     </pre>
    //   ),
    // })
    setInvitationStatus(undefined)
    const x = await createInvitation(data)
    setInvitationStatus(x)

    setisNewGuest(x.data?.isNew ?? false)
    setguideHTML(x.data?.guideHTML ?? "")
    if (x.data?.valid) {
      setsigninEnabled(true)
      const signedIn = await magicbox.signIn(
          ["User.Read"],
           form.getValues("email") as string
        )
      if (signedIn) {
        toast({
          title: "You are now signed in",
          description: "You can now continue to the site",
        })
       
      }
      location.href = "/profile/router"
      //router.push("/profile/router")
      
    }

    
  }

  return (
    <div>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <Form {...form}>

          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <div className="flex">
                    <Input placeholder="Enter your email" {...field} autoFocus={true} />
                    <div className="w-4 grow"></div>
                    <Button type="submit" disabled={processing}>Login</Button>
                  </div>
                </FormControl>

                <FormDescription className="max-w-screen-sm">
                  Your email is validated against our user database. If we cannot
                  find yours, and if your are working for a company which have not
                  yet been onboarded we will create an invitation for your
                  account.
                </FormDescription>
                {invitationStatus && invitationStatus.hasError && <div className="rounded-lg bg-red-600 p-5 font-bold text-white">
                  {invitationStatus.errorMessage}

                </div>}
                <FormMessage />
              </FormItem>
            )}
          />


        </Form>
      </form>
      {isNewGuest && <NewGuest continueToSignin={async function () {
        const signedIn = await magicbox.signIn(
          ["User.Read"],
           form.getValues("email") as string
        )
      if (signedIn) {
        toast({
          title: "You are now signed in",
          description: "You can now continue to the site",
        })
       
      }
      location.href = "/profile/router"
      }} cancel={function (): void {
        setisNewGuest(false)
      }} SigninGuideHTML={guideHTML} />}
    </div>
  )
}

export function ValidateNewHireEmailAccountForm() {
  const magicbox = useContext(MagicboxContext)

  const defaultValues: Partial<ProfileFormValues> = {}

  const form = useForm<ProfileFormValues>({
    resolver: zodResolver(profileFormSchema),
    defaultValues,
    mode: "onChange",
  })

  const [invitationStatus, setInvitationStatus] =
    useState<Result<CreateInvitationResult>>()
  const [signinEnabled, setsigninEnabled] = useState(false)
  const [processing, setProcessing] = useState(false)
  const [isNewGuest, setisNewGuest] = useState(false)
  const [isUnknownDomain, setisUnknownDomain] = useState(false)
  const [guideHTML, setguideHTML] = useState("")
  const router =  useRouter()
  useEffect(() => {
    if (!magicbox || magicbox==undefined) return
    
    if (magicbox.user) {
      if (magicbox.user.email.indexOf("#ext#@") < 0) {
        form.setValue("email", magicbox.user.email ?? "")
      }
    }
  }, [magicbox])

  async function onSubmit(data: ProfileFormValues) {
    setProcessing(true)
    setInvitationStatus(undefined)
    setsigninEnabled(false)
    // toast({
    //   title: "You submitted the following values:",
    //   description: (
    //     <pre className="mt-2 w-[340px] rounded-md bg-slate-950 p-4">
    //       <code className="text-white">{JSON.stringify(data, null, 2)}</code>
    //     </pre>
    //   ),
    // })
    setInvitationStatus(undefined)
    const x = await createInvitation(data)
    setInvitationStatus(x)

    setisNewGuest(x.data?.isNew ?? false)
    setguideHTML(x.data?.guideHTML ?? "")
    if (x.data?.valid) {
      setsigninEnabled(true)
      const signedIn = await magicbox.signIn(
          ["User.Read"],
           form.getValues("email") as string
        )
      if (signedIn) {
        toast({
          title: "You are now signed in",
          description: "You can now continue to the site",
        })
       
      }
      location.href = "https://christianiabpos.sharepoint.com/sites/welcome-to-nexi"
      //router.push("/profile/router")
      
    }

    
  }

  return (
    <div>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <Form {...form}>

          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <div className="flex">
                    <Input placeholder="Enter your email" {...field} autoFocus={true} />
                    <div className="w-4 grow"></div>
                    <Button type="submit" disabled={processing}>Login</Button>
                  </div>
                </FormControl>

                <FormDescription className="max-w-screen-sm">
                  Once you have entered your email address, you will be redirected to a login page. 
                  If this is the first time, you will have a few additional steps to complete.
                </FormDescription>
                {invitationStatus && invitationStatus.hasError && <div className="rounded-lg bg-red-600 p-5 font-bold text-white">
                  {invitationStatus.errorMessage}

                </div>}
                <FormMessage />
              </FormItem>
            )}
          />


        </Form>
      </form>
      {isNewGuest && <NewGuest continueToSignin={async function () {
        const signedIn = await magicbox.signIn(
          ["User.Read"],
           form.getValues("email") as string
        )
      if (signedIn) {
        toast({
          title: "You are now signed in",
          description: "You can now continue to the site",
        })
       
      }
      location.href = "/profile/router"
      }} cancel={function (): void {
        setisNewGuest(false)
      }} SigninGuideHTML={guideHTML} />}
    </div>
  )
}
