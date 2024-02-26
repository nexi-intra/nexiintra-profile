import { cookies } from "next/headers";
import { redirect } from "next/navigation";

export default function RootPage() {

  const data = cookies().has("user") ? JSON.parse(cookies().get("user")?.value as string) : {}
  if (data.country && data.unit) {
    redirect("https://christianiabpos.sharepoint.com/sites/nexiintra-home?country=" + data.country + "&unit=" + data.unit)
  } else[
    redirect("/")
  ]

  return null
}
