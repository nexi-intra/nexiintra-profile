import { Metadata } from "next"

import ToSmall from "@/koksmat/components/tosmall"

import { SiteFooter as LocalSiteFooter } from "./components/site-footer"
import { SiteHeader } from "./components/site-header"

export const metadata: Metadata = {
  title: "intra/profile",
  description: "Configure your profile",
  openGraph: {
  images: [
    {
      url: "https://nexiintra365.blob.core.windows.net/public/nexi-intra-profile.png",
      width: 1200,
      height: 630,
      alt: "Nexi Intra Profile",
    },
  ]
}
}

interface SettingsLayoutProps {
  children: React.ReactNode
}

export default async function SettingsLayout({
  children,
}: SettingsLayoutProps) {
  return (
    <div>

      <div className="relative flex min-h-screen flex-col">
        <SiteHeader />
       {children}
       
        <div className="container relative">
          <LocalSiteFooter />

        
        </div>
      </div>
    </div>
  )
}

