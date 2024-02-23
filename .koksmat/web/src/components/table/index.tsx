import { ColumnDef } from "@tanstack/react-table"

import ToSmall from "@/components/tosmall"
import { GenericTableActions } from "./components/GenericTableActions"
import { columns } from "./components/columns"
import { DataTable } from "./components/data-table"
import { DataTableColumnHeader } from "./components/data-table-column-header"
import { GenericItem } from "./data/schema"

export interface GenericTableProps {
  caption?: string
  description?: string
  data: any
  addtionalColumns?: ColumnDef<GenericItem>[]
  actions?: GenericTableActions<GenericItem>
}

export function GenericTable(params: GenericTableProps) {
  const instanceColumns = [...columns]
  if (params.addtionalColumns) {
    
    params.addtionalColumns.forEach((c) =>
    instanceColumns.splice(columns.length - 1, 0, c)
    )
  }
  return (
    <>
      <ToSmall />
      <div className="hidden h-full flex-1 flex-col   md:flex">
        {params.caption && (
          <div className="mb-4 flex items-center  ">
          <div className="text-2xl font-bold">{params.caption}</div>
          <div className="text-md ml-3">{params.description}</div>
          </div>
        )}
          
        <DataTable
          data={params.data}
          columns={instanceColumns}
          actions={params.actions}
        />
      </div>
    </>
  )
}
